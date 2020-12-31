package DB

type LoginUser struct {
	email    string
	password string
}

func NewLoginUser(email string, password string) *LoginUser {

	return &LoginUser{email, password}
}

func (user *LoginUser) GetLoginData() string {
	return user.email
}

func (user *LoginUser) GetEmail() string {
	return user.email
}

func (user *LoginUser) GetPassword() string {
	return user.password
}
