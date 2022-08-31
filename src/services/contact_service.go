package services

import (
	"github.com/AlfaSakan/my-profile-api.git/src/models"
	"github.com/AlfaSakan/my-profile-api.git/src/repositories"
	"github.com/AlfaSakan/my-profile-api.git/src/schemas"
)

type IContactService interface {
	BlokirContact(contact *schemas.ContactRequest) error
	CreateContact(contact *schemas.ContactRequest) error
	GetContactsByUserId(userId string) (*[]models.Contact, error)
}

type ContactService struct {
	contactRepository repositories.IContactRepository
}

func NewContactService(contactRepository repositories.IContactRepository) *ContactService {
	return &ContactService{contactRepository}
}

func (s *ContactService) GetContactsByUserId(userId string) (*[]models.Contact, error) {
	contacts := &[]models.Contact{}
	return contacts, s.contactRepository.FindContacts(userId, contacts)
}

func (s *ContactService) CreateContact(contactRequest *schemas.ContactRequest) error {
	contact := &models.Contact{
		UserId:   contactRequest.UserId,
		FriendId: contactRequest.FriendId,
		Status:   contactRequest.Status,
	}

	return s.contactRepository.CreateContact(contact)
}

func (s *ContactService) BlokirContact(contactRequest *schemas.ContactRequest) error {
	contact := &models.Contact{
		Status: contactRequest.Status,
	}

	return s.contactRepository.BlokirContact(contact)
}
