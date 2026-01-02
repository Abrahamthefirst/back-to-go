package webutil

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Abrahamthefirst/back-to-go/internal/dtos"
	"github.com/Abrahamthefirst/back-to-go/internal/entities"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateRequest(c *gin.Context, dto any) error {
	validate := validator.New()

	if err := c.ShouldBindJSON(dto); err != nil {

		slog.WarnContext(c.Request.Context(), "request validation failed",
			"error", err.Error(),
			"path", c.Request.URL.Path,
		)
		c.JSON(http.StatusBadRequest, gin.H{"validation": "all fields required"})
		return err
	}

	if err := validate.Struct(dto); err != nil {

		switch errs := err.(type) {
		case validator.ValidationErrors:
			for _, err := range errs {
				slog.WarnContext(c.Request.Context(), "request validation failed",
					"error", err.Error(),
					"path", c.Request.URL.Path,
				)
				c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%s validation failed", err.Field())})
				return err
			}

		case *validator.InvalidValidationError:
			slog.Error("Validator received invalid input", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal validation error"})
			return err

		default:
			slog.Error("Validator received invalid input", "error", err.Error())

			c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed"})
			return err

		}

	}

	return nil
}

func HandleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, entities.ErrInvalidCredentials):
		c.JSON(http.StatusUnauthorized, dtos.ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusUnauthorized,
		})
	case errors.Is(err, entities.ErrConflict):
		c.JSON(http.StatusConflict, dtos.ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusConflict,
		})

	case errors.Is(err, context.DeadlineExceeded):
		c.JSON(http.StatusRequestTimeout, dtos.ErrorResponse{
			Message:    "The server took too long to respond. Please try again.",
			StatusCode: 408,
		})

	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}
}
