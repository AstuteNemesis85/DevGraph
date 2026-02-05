package code

type SubmitCodeRequest struct {
	Language   string `json:"language" binding:"required"`
	SourceCode string `json:"source_code" binding:"required"`
}
