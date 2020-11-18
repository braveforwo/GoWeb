package domain

type UploadResponseModel struct {
	Success int         `json:"success"`
	Message string      `json:"message"`
	Url     string      `json:"url"`
	Data    interface{} `json:"data"`
}
