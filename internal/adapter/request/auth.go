package request

type RegisterRequest struct {
	Name     string `json:"name" binding:"required" example:"Victor Alarcon"`
	Username string `json:"username" binding:"required" example:"vicalar"`
	Email    string `json:"email" binding:"required,email" example:"vicalar@gmail.com"`
	Password string `json:"password" binding:"required,min=8,max=32,alphanum" example:"Prueba123"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"vicalar@gmail.com"`
	Password string `json:"password" binding:"required,min=8,max=32,alphanum" example:"Prueba123"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" binding:"required" example:"Victor Alarcon"`
	Username string `json:"username" binding:"required" example:"vicalar"`
	Email    string `json:"email" binding:"required,email" example:"vicalar@gmail.com"`
	Password string `json:"password" binding:"required,min=8,max=32,alphanum" example:"Prueba123"`
}
