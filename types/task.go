package types

type CreateTaskReq struct {
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=1000"`
	Status  int    `form:"status" json:"status"`
}

type ListTaskReq struct {
	CurrentPage int `form:"current_page" json:"currentPage"`
	PageSize    int `form:"page_size" json:"pageSize"`
}

type UpdateListTaskReq struct {
	ID      uint   `form:"id" json:"id"`
	Content string `form:"content" json:"content" binding:"max=1000"`
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Status  int    `form:"status" json:"status"`
}

type TaskResp struct {
	ID        uint   `json:"id"`
	Content   string `json:"content"`
	Status    int    `json:"status"`
	Title     string `json:"title"`
	StartTime int64  `json:"startTime"`
	//View      uint64 `json:"view"`
	EndTime  int64 `json:"endTime"`
	CreateAt int64 `json:"createAt"`
}

type DeleteTaskReq struct {
	ID uint `form:"id" json:"id"`
}
