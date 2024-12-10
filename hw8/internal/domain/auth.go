package domain

type Role string

const (
	NoneRole     Role = "none"
	DirectorRole Role = "admin"
	ExpertRole   Role = "user"
)

type UserInfo struct {
	ID       string
	Username string
	Password string
	Role     Role
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}
