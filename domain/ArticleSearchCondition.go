package domain

type ArticleSearchCondition struct {
	SearchContext string `json:"searchContext" form:"searchContext"`
	CurrentPage   int    `json:"currentPage" form:"currentPage"`
	PageSize      int    `json:"pageSize" form:"pageSize"`
}
