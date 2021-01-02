package DB

import "golang.org/x/crypto/bcrypt"

type RegisterUser struct {
	Name     string
	Email    string
	Password []byte
}

func NewRegisterUser(name string, email string, password string) (*RegisterUser, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return &RegisterUser{name, email, pass}, err
}

func (user *RegisterUser) GetResgisterData() (string, string) {
	return user.Name, user.Email
}

func (user *RegisterUser) GetName() string {
	return user.Name
}

func (user *RegisterUser) GetEmail() string {
	return user.Email
}

func (user *RegisterUser) GetPassword() []byte {
	return user.Password
}
