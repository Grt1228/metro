package req

type MetroListReq struct {
	Page int    `json:"page" binding:"required"`
	Size int    `json:"size" binding:"required"`
	Name string `json:"name"`
}
