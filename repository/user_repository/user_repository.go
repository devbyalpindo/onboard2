package user_repository

import "backend-onboarding2/model/entity"

type UserRepository interface {
	AddUsers(*entity.User)
	//UpdateUsers(entity.User, string) (string, error)
	//DeleteUsers(string) error
}
