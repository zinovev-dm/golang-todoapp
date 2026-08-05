package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/zinovev-dm/golang-todoapp/internal/core/errors"
)

type User struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}

func NewUser(
	id int,
	version int,
	fullName string,
	phoneNumber *string,
) User {
	return User{
		ID:          id,
		Version:     version,
		FullName:    fullName,
		PhoneNumber: phoneNumber,
	}
}

func NewUserUninitialized(fullName string, phoneNumber *string) User {
	return NewUser(UninitializedID, UninitializedVersion, fullName, phoneNumber)
}

func (u *User) Validate() error {
	fullNameLength := len([]rune(u.FullName))
	if fullNameLength < 3 || fullNameLength > 100 {
		return fmt.Errorf("invalid FullName length: %d %w", fullNameLength, core_errors.ErrInvalidArgument)
	}

	if u.PhoneNumber != nil {
		phoneNumberLength := len([]rune(*u.PhoneNumber))
		if phoneNumberLength < 10 || phoneNumberLength > 15 {
			return fmt.Errorf("invalid PhoneNumber length: %d %w", phoneNumberLength, core_errors.ErrInvalidArgument)
		}

		re := regexp.MustCompile(`^\+[0-9]+$`)

		if !re.MatchString(*u.PhoneNumber) {
			return fmt.Errorf("invalid PhoneNumber format: %w", core_errors.ErrInvalidArgument)
		}
	}
	return nil
}

type UserPath struct {
	FullName    Nullable[string]
	PhoneNumber Nullable[string]
}

func (p *UserPath) Validate() error {
	if p.FullName.Set && p.FullName.Value == nil {
		return fmt.Errorf("FullName can't be path to NULL: %w", core_errors.ErrInvalidArgument)
	}
	return nil
}

func (u *User) ApplyPath(path UserPath) error {
	if err := path.Validate(); err != nil {
		return fmt.Errorf("validate user path: %w", err)
	}

	tmp := *u

	if path.FullName.Set {
		tmp.FullName = *path.FullName.Value
	}
	if path.PhoneNumber.Set {
		tmp.PhoneNumber = path.PhoneNumber.Value
	}

	if err := tmp.Validate(); err != nil {
		fmt.Errorf("validate patched user: %w", err)
	}

	*u = tmp

	return nil
}
