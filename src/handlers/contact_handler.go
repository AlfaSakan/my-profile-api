package handlers

import (
	"net/http"

	"github.com/AlfaSakan/my-profile-api.git/src/models"
	"github.com/AlfaSakan/my-profile-api.git/src/schemas"
	"github.com/AlfaSakan/my-profile-api.git/src/services"
	"github.com/AlfaSakan/my-profile-api.git/src/utils"

	"github.com/gin-gonic/gin"
)

type ContactHandler struct {
	contactService services.IContactService
}

func NewContactHandler(contactService services.IContactService) *ContactHandler {
	return &ContactHandler{contactService}
}

func (ch *ContactHandler) GetContacts(c *gin.Context) {
	user, _ := c.Get("User")
	userId := user.(*models.User).UserId

	response := &schemas.Response{}

	contacts, err := ch.contactService.GetContactsByUserId(userId)
	if err != nil {
		utils.ResponseBadRequest(c, response, err)
		return
	}

	response.Data = contacts
	response.Message = "OK"
	response.Status = http.StatusOK
	c.JSON(response.Status, response)
}

func (ch *ContactHandler) PatchContacts(c *gin.Context) {
	response := &schemas.Response{}

	contact := &schemas.ContactRequest{}

	user, _ := c.Get("User")
	userId := user.(*models.User).UserId
	contact.UserId = userId

	err := c.ShouldBindJSON(contact)
	if err != nil {
		utils.ResponseBadRequest(c, response, err)
		return
	}

	err = ch.contactService.BlokirContact(contact)
	if err != nil {
		utils.ResponseBadRequest(c, response, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "OK"
	response.Data = contact
	c.JSON(response.Status, response)
}

func (ch *ContactHandler) PostContacts(c *gin.Context) {
	response := &schemas.Response{}

	contact := &schemas.ContactRequest{}

	user, _ := c.Get("User")
	userId := user.(*models.User).UserId
	contact.UserId = userId

	err := c.ShouldBindJSON(contact)
	if err != nil {
		utils.ResponseBadRequest(c, response, err)
		return
	}

	err = ch.contactService.CreateContact(contact)
	if err != nil {
		utils.ResponseBadRequest(c, response, err)
		return
	}

	response.Status = http.StatusCreated
	response.Message = "Created"
	response.Data = contact
	c.JSON(response.Status, response)
}
