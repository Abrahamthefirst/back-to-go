package middleware

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/Abrahamthefirst/back-to-go/internal/entities"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")

		slog.Info(header)

		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(401, gin.H{"error": "Authorization header is required"})
			return
		}

		tokenString := strings.TrimPrefix(header, "Bearer ")
		token, err := jwt.ParseWithClaims(tokenString, &entities.AccesTokenClaims{}, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil
		})

		if err != nil {
			slog.Error("Parse error: " + err.Error())
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
			return
		}

		if claims, ok := token.Claims.(*entities.AccesTokenClaims); ok && token.Valid {
			slog.Info("User authenticated", "id", claims.UserID)
			c.Set("userID", claims.UserID)
		}

		c.Next()
	}
}
