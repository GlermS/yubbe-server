package interfaces

type IRegisterUser interface {
	GetResgisterData() (string, string)
	GetName() string
	GetEmail() string
	GetPassword() []byte
}

type IDatabase interface {
	GetDatabaseData() (string, string)
	Register(IRegisterUser) error
	Verify(ILoginUser) error
}

type ILoginUser interface {
	GetLoginData() string
	GetEmail() string
	GetPassword() string
}
