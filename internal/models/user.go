package models

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"` // O "-" impede que o password seja incluído no JSON
}

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
