package types

type UserServiceReq struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15" example:"chan"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=16" example:"123456"`
}

type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

type UserResp struct {
	ID       uint   `json:"id" form:"id" example:"1"`
	Username string `json:"username" form:"user_name" example:"chan"`
	CreateAt int64  `json:"createAt" form:"create_at"`
}
