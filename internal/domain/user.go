package user

import "errors"

var (
	ErrNameRequired     = errors.New("field 'name' is required")
	ErrLastNameRequired = errors.New("field 'last_name' is required")
	ErrEmailRequired    = errors.New("field 'email' is required")
	ErrPasswordRequired = errors.New("field 'password' is required")
)

const (
	CompanyRole     Role = "company"
	CoordinatorRole Role = "coordinator"
)

type (
	Role string

	User struct {
		ID       int
		Name     string
		LastName string
		Email    string
		Password string
		Role     Role
	}
)

func (u *User) isNameValid() bool {
	return u.Name != ""
}

func (u *User) isLastNameValid() bool {
	return u.LastName != ""
}

func (u *User) isEmailValid() bool {
	return u.Email != ""
}

func (u *User) isPasswordValid() bool {
	return u.Password != ""
}

func (u *User) Validate() error {
	if !u.isNameValid() {
		return ErrNameRequired
	}
	if !u.isLastNameValid() {
		return ErrLastNameRequired
	}
	if !u.isEmailValid() {
		return ErrEmailRequired
	}
	if !u.isPasswordValid() {
		return ErrPasswordRequired
	}
	return nil
}
