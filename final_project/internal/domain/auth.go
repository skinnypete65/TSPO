package domain

type Role string

const (
	NoneRole  Role = "none"
	AdminRole Role = "admin"
	UserRole  Role = "user"
)

type UserInfo struct {
	ID       string `gorm:"column:user_id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:hash_pass"`
	Role     Role   `gorm:"column:user_role"`
}

func (UserInfo) TableName() string {
	return "users"
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}
