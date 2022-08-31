package utils

import (
	"fmt"
	"time"

	"github.com/AlfaSakan/my-profile-api.git/src/models"

	"github.com/golang-jwt/jwt/v4"
)

type CustomClaim struct {
	*models.User
	SessionId string `json:"session_id"`
}

func GenerateToken(data *CustomClaim, expireTime int64) (string, error) {
	privateKey := ViperEnvVariable("PRIVATEKEY")

	claims := jwt.MapClaims{
		"iss": "issuer",
		"exp": expireTime,
		"data": map[string]interface{}{
			"name":         data.Name,
			"phone_number": data.PhoneNumber,
			"country_code": data.CountryCode,
			"image_url":    data.ImageUrl,
			"user_id":      data.UserId,
			"status":       data.Status,
			"created_at":   data.CreatedAt,
			"updated_at":   data.UpdatedAt,
			"session_id":   data.SessionId,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString([]byte(privateKey))

	if err != nil {
		return "", err
	}

	return token, nil
}

func DecodeToken(token string) (jwt.MapClaims, error) {
	privateKey := ViperEnvVariable("PRIVATEKEY")

	decode, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(privateKey), t.Claims.Valid()
	})

	if err != nil {
		return nil, err
	}

	claims := decode.Claims.(jwt.MapClaims)

	expiredMilliSecond := ConvertFloat64ToInt64(claims["exp"].(float64))

	isExpired := expiredMilliSecond-time.Now().UnixMilli() <= 0

	if isExpired {
		return claims, fmt.Errorf("token is expired")
	}

	return claims, err
}
