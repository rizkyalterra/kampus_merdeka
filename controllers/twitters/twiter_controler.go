package twitters

import (
	"fmt"
	"kampus_merdeka/configs"
	"kampus_merdeka/models/response"

	"kampus_merdeka/models/twiters"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetTwitterController(c echo.Context) error {

	twitters := []twiters.Tweet{}

	result := configs.DB.Preload("User").Find(&twitters)

	if result.Error != nil {
		fmt.Println(result.Error)
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, response.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error ketika input mendapatkan data twitter dari DB",
				Data:    nil,
			})
		}
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil mendapatkan data twiter",
		Data:    twitters,
	})
}
