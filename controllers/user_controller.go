package controllers

import (
	"errors"
	"kampus_merdeka/configs"
	"kampus_merdeka/middlewares"
	"kampus_merdeka/models/response"
	"kampus_merdeka/models/users"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterController(c echo.Context) error {
	var userRegister users.UserRegister
	c.Bind(&userRegister)

	// validasi
	if userRegister.Name == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Nama masih kosong",
			Data:    nil,
		})
	}

	// etc

	var userDB users.User
	userDB.Name = userRegister.Name

	// userDB.Password = helpers.Hash(userRegister.Password)
	userDB.Address = userRegister.Address
	userDB.Email = userRegister.Email

	result := configs.DB.Create(&userDB)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error ketika input data user ke DB",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil register",
		Data:    userDB,
	})
}

func LoginController(c echo.Context) error {
	userLogin := users.UserLogin{}
	c.Bind(&userLogin)

	user := users.User{}

	result := configs.DB.First(&user, "email = ? AND password = ?",
		userLogin.Email, userLogin.Password)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusForbidden, response.BaseResponse{
				Code:    http.StatusForbidden,
				Message: "User tidak ditemukan atau password tidak sesuai",
				Data:    nil,
			})
		} else {
			return c.JSON(http.StatusInternalServerError, response.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Ada keselahan di server",
				Data:    nil,
			})
		}

	}

	token, err := middlewares.GenerateTokenJWT(user.Id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Ada keselahan di server",
			Data:    nil,
		})
	}

	userResponse := users.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Address:   user.Address,
		Token:     token,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil login",
		Data:    userResponse,
	})
}

func DetailUserController(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Gagal konversi userId",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil",
		Data:    users.User{Id: userId},
	})
}

func GetUserController(c echo.Context) error {

	users := []users.User{}

	result := configs.DB.Find(&users)

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, response.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error ketika input mendapatkan data user dari DB",
				Data:    nil,
			})
		}
	}

	userId, _ := middlewares.GetClaimsUserId(c)
	/// olah

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    userId,
		Message: "Berhasil mendapatkan data user",
		Data:    users,
	})
}
