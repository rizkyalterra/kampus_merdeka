package middlewares

import (
	"fmt"
	"kampus_merdeka/constants"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JwtClaims struct {
	UserId int `json:"userId"`
	jwt.StandardClaims
}

func GenerateTokenJWT(id int) (string, error) {
	claims := JwtClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 1).Unix(),
		},
	}

	fmt.Println(claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtToken, err := token.SignedString([]byte(constants.SECRET_JWT))

	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func GetClaimsUserId(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64)
	return int(userId)
}
