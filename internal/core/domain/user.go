package domain

type User struct {
	ID        uint64 `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
