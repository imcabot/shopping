package user

//创建用户请求结构体
type CreatUserRequest struct {
	Username  string `json:"username"`
	password  string `json:"password"`
	password2 string `json:"password2"`
}

//创建用户响应
type CreatUserResponse struct {
	Username string `json:"username"`
}

//登录请求
type LoginRequest struct {
	Username string `json:"username"`
	password string `json:"password"`
}

//登录响应
type LoginResponse struct {
	Username string `json:"username"`
	UserID   uint   `json:"UserID"`
	Token    string `json:"token"`
}
