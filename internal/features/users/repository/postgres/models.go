package users_postgres_repository

import (
	"github.com/zinovev-dm/golang-todoapp/internal/core/domain"
)

type UserModel struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}

func UserDomainsFromModels(userModels []UserModel) []domain.User {
	userDomains := make([]domain.User, len(userModels))

	for i, user := range userModels {
		userDomains[i] = domain.NewUser(user.ID, user.Version, user.FullName, user.PhoneNumber)
	}
	return userDomains
}
