package entities

const (
    ROLE_ADMIN = "admin"
    ROLE_USER  = "user"
)

type UserEntity struct {
    BaseEntity

    Username    string `json:"username"`
    Email       string `json:"email"`
    DisplayName string `json:"display_name"`
    Password    string `json:"-"`
    Avatar      string `json:"avatar"`
    Role        string `json:"role"`
}
