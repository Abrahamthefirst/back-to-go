package webutil

import (
	"fmt"
	"log/slog"
	"net/http"

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
