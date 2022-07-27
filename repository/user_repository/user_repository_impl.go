package user_repository

import (
	"backend-onboarding2/model/entity"
	"log"

	"github.com/jinzhu/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB}
}

func (repository *UserRepositoryImpl) AddUsers(user *entity.User) {
	db := repository.DB
	err := db.Exec("UPSERT INTO users (ID, PERSONAL_NUMBER, PASSWORD, EMAIL, NAME, ROLE_ID, IS_ACTIVE) VALUES (?, ?, ?, ?, ?, ?, ?)", user.Id, user.PersonalNumber, user.Password, user.Email, user.Name, user.RoleID, user.IsActive).Error

	if err != nil {
		log.Printf("Error create %v\n", err)
	} else {
		log.Print("Successfully create user")
	}
}
