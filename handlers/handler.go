package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"delivery/configs"
	adminController "delivery/controllers/admin"
	"delivery/logger"
	e "delivery/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	httppkg "delivery/pkg/http"
)

type Handler struct {
	cfg             *configs.Configuration
	log             logger.LoggerI
	adminController adminController.AdminController
	validate        *validator.Validate
}

func New(
	cfg *configs.Configuration,
	log logger.LoggerI,
	adminController adminController.AdminController,
	validate *validator.Validate,
) Handler {
	return Handler{
		cfg:             cfg,
		log:             log,
		adminController: adminController,
		validate:        validate,
	}
}

// handleResponse
func (h *Handler) handleResponse(c *gin.Context, status httppkg.Status, data ...interface{}) {
	switch code := status.Code; {
	case code < 300:
		h.log.Info(
			"---Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			// logger.Any("data", data),
		)
	case code < 400:
		h.log.Warn(
			"!!!Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	default:
		h.log.Error(
			"!!!Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	}

	c.JSON(status.Code, httppkg.Response{
		Status:      status.Status,
		Description: status.Description,
		Data:        data,
	})
}

type ResponseForExec struct {
	Info string `json:"info"`
	ID   string `json:"id"`
}

// StatusFromError ...
func StatusFromError(err error) httppkg.Status {
	if err == nil {
		return httppkg.OK
	}

	code, ok := e.ExtractStatusCode(err)
	if !ok || code == http.StatusInternalServerError {
		return httppkg.Status{
			Code:        http.StatusInternalServerError,
			Status:      "INTERNAL_SERVER_ERROR",
			Description: err.Error(),
		}
	} else if code == http.StatusNotFound {
		return httppkg.Status{
			Code:        http.StatusNotFound,
			Status:      "NOT_FOUND",
			Description: err.Error(),
		}
	} else if code == http.StatusBadRequest {
		return httppkg.Status{
			Code:        http.StatusBadRequest,
			Status:      "BAD_REQUEST",
			Description: err.Error(),
		}
	} else if code == http.StatusForbidden {
		return httppkg.Status{
			Code:        http.StatusForbidden,
			Status:      "FORBIDDEN",
			Description: err.Error(),
		}
	} else if code == http.StatusUnauthorized {
		return httppkg.Status{
			Code:        http.StatusUnauthorized,
			Status:      "FORBIDDEN",
			Description: err.Error(),
		}
	} else {
		return httppkg.Status{
			Code:        http.StatusInternalServerError,
			Status:      "INTERNAL_SERVER_ERROR",
			Description: err.Error(),
		}
	}

}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func handleBodyParseError(err error) []ValidationError {
	var validationErrors []ValidationError

	var ute *json.UnmarshalTypeError
	if errors.As(err, &ute) {
		// line, column := findLineAndColumn(body, ute.Offset)
		validationErrors = append(validationErrors, ValidationError{
			Field:   ute.Field,
			Message: fmt.Sprintf("Type mismatch: expected %s but got %s", ute.Type.String(), ute.Value),
		})
	} else if syntaxErr, ok := err.(*json.SyntaxError); ok {
		validationErrors = append(validationErrors, ValidationError{
			Field:   "JSON",
			Message: fmt.Sprintf("Syntax error at offset %d", syntaxErr.Offset),
		})
	} else {
		validationErrors = append(validationErrors, ValidationError{
			Field:   "Unknown",
			Message: err.Error(),
		})
	}

	return validationErrors
}

// func calculateLineAndColumn(data []byte, offset int64) (line, col int) {
// 	// Offset asosida qator va ustun hisoblash
// 	// data, _ := io.ReadAll(body)
// 	line, col = 1, 0
// 	for i, b := range data {
// 		if int64(i) == offset {
// 			break
// 		}
// 		if b == '\n' {
// 			line++
// 			col = 0
// 		} else {
// 			col++
// 		}
// 	}
// 	return line, col
// }

// func findLineAndColumn(jsonBody []byte, offset int64) (line, column int) {
//     if offset > int64(len(jsonBody)) {
//         offset = int64(len(jsonBody)) // Out-of-range ni cheklash
//     }

//     lines := bytes.Split(jsonBody[:offset], []byte("\n"))
//     line = len(lines)
//     column = len(lines[len(lines)-1])
//     return line, column
// }
