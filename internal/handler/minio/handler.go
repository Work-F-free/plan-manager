package minio

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"seatPlanner/internal/common/dto"
	"seatPlanner/internal/common/errors"
	"seatPlanner/internal/common/responses"
	"seatPlanner/pkg/minio"
	"seatPlanner/pkg/minio/helpers"
)

type Handler struct {
	minioService minio.Client
}

func NewMinioHandler(minioService minio.Client) *Handler {
	return &Handler{
		minioService: minioService,
	}
}

func (h *Handler) CreateOne(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrorResponse{
			Status:  http.StatusBadRequest,
			Error:   "No file is received",
			Details: err,
		})
		return
	}

	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "Unable to open the file",
			Details: err,
		})
		return
	}
	defer f.Close()

	fileBytes, err := io.ReadAll(f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "Unable to read the file",
			Details: err,
		})
		return
	}

	fileData := helpers.FileDataType{
		FileName: file.Filename,
		Data:     fileBytes,
	}

	link, err := h.minioService.CreateOne(fileData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "Unable to save the file",
			Details: err,
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Status:  http.StatusOK,
		Message: "File uploaded successfully",
		Data:    link,
	})
}

func (h *Handler) CreateMany(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrorResponse{
			Status:  http.StatusBadRequest,
			Error:   "Invalid form",
			Details: err,
		})
		return
	}

	files := form.File["files"]
	if files == nil {
		c.JSON(http.StatusBadRequest, errors.ErrorResponse{
			Status:  http.StatusBadRequest,
			Error:   "No files are received",
			Details: err,
		})
		return
	}

	data := make(map[string]helpers.FileDataType)

	for _, file := range files {
		f, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, errors.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Error:   "Unable to open the file",
				Details: err,
			})
			return
		}
		defer f.Close()

		fileBytes, err := io.ReadAll(f)
		if err != nil {
			c.JSON(http.StatusInternalServerError, errors.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Error:   "Unable to read the file",
				Details: err,
			})
			return
		}

		data[file.Filename] = helpers.FileDataType{
			FileName: file.Filename,
			Data:     fileBytes,
		}
	}

	links, err := h.minioService.CreateMany(data)
	if err != nil {
		fmt.Printf("err: %+v\n ", err.Error())
		c.JSON(http.StatusInternalServerError, errors.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "Unable to save the files",
			Details: err,
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Files uploaded successfully",
		Data:    links,
	})
}

func (h *Handler) GetOne(c *gin.Context) {
	objectID := c.Param("objectID")

	link, err := h.minioService.GetOne(objectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "Enable to get the object",
			Details: err,
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Status:  http.StatusOK,
		Message: "File received successfully",
		Data:    link,
	})
}

func (h *Handler) GetMany(c *gin.Context) {
	var objectIDs dto.ObjectIdsDto

	if err := c.ShouldBindJSON(&objectIDs); err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrorResponse{
			Status:  http.StatusBadRequest,
			Error:   "Invalid request body",
			Details: err,
		})
		return
	}

	links, err := h.minioService.GetMany(objectIDs.ObjectIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "Enable to get many objects",
			Details: err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Files received successfully",
		"data":    links,
	})
}

func (h *Handler) DeleteOne(c *gin.Context) {
	objectID := c.Param("objectID")

	if err := h.minioService.DeleteOne(objectID); err != nil {
		c.JSON(http.StatusInternalServerError, errors.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "Cannot delete the object",
			Details: err,
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Status:  http.StatusOK,
		Message: "File deleted successfully",
	})
}

func (h *Handler) DeleteMany(c *gin.Context) {
	var objectIDs dto.ObjectIdsDto

	if err := c.BindJSON(&objectIDs); err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrorResponse{
			Status:  http.StatusBadRequest,
			Error:   "Invalid request body",
			Details: err,
		})
		return
	}

	if err := h.minioService.DeleteMany(objectIDs.ObjectIDs); err != nil {
		c.JSON(http.StatusInternalServerError, errors.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   "Cannot delete many objects",
			Details: err,
		})
		return
	}

	// Возвращаем успешный ответ, если объекты успешно удалены
	c.JSON(http.StatusOK, responses.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Files deleted successfully",
	})
}
