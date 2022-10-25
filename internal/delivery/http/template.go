package http

import (
	"context"
	templateStruct "emailservice/internal/delivery/http/template"
	"io/ioutil"
	"mime"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	defaultRequestTimeout = time.Second * 15
)

// CreateTemplate
// @Summary Добавляет шаблон письма.
// @Description Добавляет шаблон письма.
// @Tags templates
// @Accept  multipart/form-data
// @Produce json
// @Param 	file  		formData 	file true										"Файл шаблона"
// @Success 201			{object}    templateStruct.TemplateResponse					"Структура шаблона"
// @Failure 400 		{object}    ErrorResponse									"Bad Request"
// @Failure 500 	    {object} 	ErrorResponse									"Internal Server Error"
// @Router /template/ [post]
func (d *Delivery) CreateTemplate(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	template := templateStruct.CreateTemplateRequest{}
	if err := c.ShouldBind(&template); err != nil {
		d.logger.Info("invalid request params %s", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{Msg: "invalid request params"})
		return
	}

	mediaType, _, err := mime.ParseMediaType(template.File.Header.Get("Content-Type"))
	if err != nil {
		d.logger.Error("error parsing mime type %s", err.Error())
		c.JSON(http.StatusInternalServerError, ErrorResponse{Msg: "error processing file"})
		return
	}

	if mediaType != "text/html" {
		d.logger.Info("invalid file type")
		c.JSON(http.StatusBadRequest, ErrorResponse{Msg: "invalid file type"})
		return
	}

	file, err := template.File.Open()
	if err != nil {
		d.logger.Error("error openning file: %s", err.Error())
		c.JSON(http.StatusInternalServerError, ErrorResponse{Msg: "error processing file"})
		return
	}
	defer file.Close()

	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		d.logger.Error("error reading file: %s", err.Error())
		c.JSON(http.StatusInternalServerError, ErrorResponse{Msg: "error processing file"})
		return
	}

	result, err := d.useCase.CreateTemplate(ctx, fileContent)
	if err != nil {
		d.logger.Error("error creating template: %s", err.Error())
		c.JSON(http.StatusInternalServerError, ErrorResponse{Msg: "error creating template"})
		return
	}

	c.JSON(http.StatusCreated, templateStruct.TemplateResponse{ID: result.ID(), Path: result.Path().String(), CreatedAt: result.CreatedAt(), ModifiedAt: result.ModifiedAt()})
}
