package auth

type Authenticator interface {
	GenerateToken() (string, error)
	ValidateToken(token string) (string, error)
}
