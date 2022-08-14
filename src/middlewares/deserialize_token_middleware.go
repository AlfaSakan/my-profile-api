package middlewares

import (
	"myProfileApi/src/models"
	"myProfileApi/src/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeserializeUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := strconv.Unquote(c.Request.Header.Get("Authorization"))
		if err != nil {
			c.Next()
			return
		}

		refreshToken, err := strconv.Unquote(c.Request.Header.Get("x-refresh"))
		if err != nil {
			c.Next()
			return
		}

		accessClaim, err := utils.DecodeToken(accessToken)

		if err != nil && len(refreshToken) > 0 {
			newAccessToken, user := utils.ReIssueAccessToken(db, refreshToken)

			if user != nil {
				c.Header("x-access", newAccessToken)
				c.Set("User", user)
			}

			c.Next()
			return
		}

		data := accessClaim["data"].(map[string]interface{})

		user := &models.User{
			UserId:      uint(data["user_id"].(float64)),
			CountryCode: data["country_code"].(string),
			PhoneNumber: data["phone_number"].(string),
			Name:        data["name"].(string),
			ImageUrl:    data["image_url"].(string),
			Status:      data["status"].(string),
			CreatedAt:   utils.ConvertFloat64ToInt64(data["created_at"].(float64)),
			UpdatedAt:   utils.ConvertFloat64ToInt64(data["updated_at"].(float64)),
		}

		c.Set("User", user)
		c.Next()
	}
}
