package entities

type ResponLogin struct {
	UserId  string
	Session string
}

type ResponRegister struct {
	Status string
}

type ResponGetSession struct {
	UserSession string
	RememberMe  bool
}
