package domain

type CommentPageCondition struct {
	ArticleId   int `json:"articleId" form:"articleId"`
	PageSize    int `json:"pageSize" form:"pageSize"`
	CurrentPage int `json:"currentPage" form:"currentPage"`
	AllPageSize int `json:"allPageSize" form:"allPageSize"`
}
