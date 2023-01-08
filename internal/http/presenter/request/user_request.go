package request

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required,min=3,max=255"`
	Email    string `json:"email" binding:"required,email,max=255"`
	Password string `json:"password" binding:"required,min=8,max=255"`
}

type UpdateUserRequest struct {
	ID       int64  `json:"id"`
	Name     string `json:"name" binding:"required,min=3,max=255"`
	Password string `json:"password" binding:"omitempty,min=8,max=255"`
}
