package routes

import (
	"github.com/AlfaSakan/my-profile-api.git/src/handlers"
	"github.com/AlfaSakan/my-profile-api.git/src/middlewares"

	"github.com/gin-gonic/gin"
)

const CONTACT_ROUTE = "/contact"

func ContactRoute(r *gin.RouterGroup, ch *handlers.ContactHandler) {
	r.GET(CONTACT_ROUTE, middlewares.RequireUser(), ch.GetContacts)

	r.POST(CONTACT_ROUTE, middlewares.RequireUser(), ch.PostContacts)

	r.PATCH(CONTACT_ROUTE, middlewares.RequireUser(), ch.PatchContacts)
}
