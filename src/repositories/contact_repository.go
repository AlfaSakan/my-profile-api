package repositories

import (
	"github.com/AlfaSakan/my-profile-api.git/src/models"

	"gorm.io/gorm"
)

type IContactRepository interface {
	FindContacts(userId string, contacts *[]models.Contact) error
	CreateContact(contact *models.Contact) error
	BlokirContact(contact *models.Contact) error
}

type ContactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) *ContactRepository {
	return &ContactRepository{
		db,
	}
}

func (r *ContactRepository) FindContacts(userId string, contacts *[]models.Contact) error {
	return r.db.Where(&models.Contact{UserId: userId}).Find(contacts).Error
}

func (r *ContactRepository) CreateContact(contact *models.Contact) error {
	return r.db.Create(contact).Error
}

func (r *ContactRepository) BlokirContact(contact *models.Contact) error {
	return r.db.Where(contact).Update("status", "blokir").Error
}
