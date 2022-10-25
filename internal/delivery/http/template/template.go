package template

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type CreateTemplateRequest struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

type TemplateResponse struct {
	ID         uuid.UUID `json:"id" binding:"required,uuid" example:"00000000-0000-0000-0000-000000000000" format:"uuid"`
	Path       string    `json:"path" binding:"required" example:"/storage/file"`
	CreatedAt  time.Time `json:"createdAt"  binding:"required"`
	ModifiedAt time.Time `json:"modifiedAt"  binding:"required"`
}
