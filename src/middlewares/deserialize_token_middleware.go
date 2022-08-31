package middlewares

import (
	"github.com/AlfaSakan/my-profile-api.git/src/models"
	"github.com/AlfaSakan/my-profile-api.git/src/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeserializeUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.Request.Header.Get("Authorization")
		if len(accessToken) == 0 {
			c.Next()
			return
		}

		refreshToken := c.Request.Header.Get("X-Refresh")
		if len(refreshToken) == 0 {
			c.Next()
			return
		}

		accessClaim, err := utils.DecodeToken(accessToken)

		if err != nil && len(refreshToken) > 0 {
			newAccessToken, user := utils.ReIssueAccessToken(db, refreshToken)

			if user != nil {
				c.Header("X-Access", newAccessToken)
				c.Set("User", user)
			}

			c.Next()
			return
		}

		data := accessClaim["data"].(map[string]interface{})

		user := &models.User{
			UserId:      data["user_id"].(string),
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
