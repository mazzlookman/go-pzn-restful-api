package web

type UserRegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserLoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID      int              `json:"id"`
	Name    string           `json:"name"`
	Email   string           `json:"email"`
	Avatar  string           `json:"avatar"`
	Token   string           `json:"token"`
	Courses []CourseResponse `json:"courses"`
}
