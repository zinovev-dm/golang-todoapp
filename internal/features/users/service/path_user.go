package users_service

import (
	"context"
	"fmt"

	"github.com/zinovev-dm/golang-todoapp/internal/core/domain"
)

func (s *UsersService) PathUser(ctx context.Context, id int, path domain.UserPath) (domain.User, error) {
	user, err := s.usersRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user from repository: %w", err)
	}
	if err := user.ApplyPath(path); err != nil {
		return domain.User{}, fmt.Errorf("apply path to user: %w", err)
	}
	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("validate user domain: %w", err)
	}
	patchedUser, err := s.usersRepository.PathUser(ctx, id, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("path user : %w", err)
	}
	return patchedUser, nil
}
