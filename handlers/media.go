package handlers

import (
	"fmt"
	"go-core/config"
	Middlewares "go-core/middlewares"
	"go-core/services"
	"go-core/types"
	Utils "go-core/utils"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type MediaHandler interface {
	UploadFile(c *gin.Context)
	MediaList(c *gin.Context)
}

type mediaHandler struct {
	mediaService services.MediaService
}

func NewMediaHandler(service services.MediaService) MediaHandler {
	return &mediaHandler{mediaService: service}
}

// Media godoc
// @Summary      	Upload media
// @Description  	Upload media
// @Tags         	Media
// @accept        multipart/form-data
// @Param				 	files formData file true "Upload"
// @Router       	/api/v1/admin/media [post]
// @Security			Bearer
func (h *mediaHandler) UploadFile(c *gin.Context) {
	var request types.UploadFileRequest

	if err := c.ShouldBind(&request); err != nil {
		c.Error(Middlewares.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	userId, _ := c.Get("UserID")

	sec := time.Now().Unix()
	path := fmt.Sprintf("public/uploads/%d/%d", userId, sec)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, os.FileMode(0755)); err != nil {
			c.Error(err)
			return
		}
	}

	files := request.Files
	var filesUploaded []string
	for _, file := range files {
		path := path + "/" + filepath.Base(file.Filename)
		log.Println(path, file.Size, file.Header.Get("Content-Type"))
		if err := c.SaveUploadedFile(file, path); err != nil {
			c.Error(err)
			return
		}

		fileMime := file.Header.Get("Content-Type")
		fileMimeType := strings.Split(fileMime, "/")[0]
		fileType := config.FILE
		if fileMimeType == "image" {
			fileType = config.IMAGE
		}

		request := types.StoreMediaRequest{
			FileName: file.Filename,
			FileSize: file.Size,
			FileMime: fileMime,
			FileType: fileType,
			UserID:   userId.(uint64),
			Path:     path,
		}

		_, err := h.mediaService.StoreMedia(&request)
		if err != nil {
			c.Error(err)
			return
		}

		filesUploaded = append(filesUploaded, path)
	}

	InitContext(c).ResponseEntity(&BaseOutput{
		Data:    filesUploaded,
		Message: fmt.Sprintf("Uploaded sucessfully %d files", len(files)),
	})
}

// Media godoc
// @Summary      	Get media list
// @Description  	Get media list
// @Tags         	Media
// @Param					Query query types.FilterMediaRequest true "BaseRequest"
// @Success				200 {object} handlers.BaseOutput{data=types.MediaDTO} "Media"
// @Router       	/api/v1/admin/media [get]
// @Security			Bearer
func (h *mediaHandler) MediaList(c *gin.Context) {
	UserID, _ := c.Get("UserID")
	request := types.FilterMediaRequest{
		UserID: UserID.(uint64),
	}
	page := c.DefaultQuery("page", "1")
	size := c.DefaultQuery("size", "20")
	request.Page = Utils.ParseInt(page)
	request.Size = Utils.ParseInt(size)
	responses, err := h.mediaService.MediaList(&request)
	if err != nil {
		c.Error(err)
		return
	}

	InitContext(c).ResponseEntity(&BaseOutput{
		Data: responses,
	})
}
