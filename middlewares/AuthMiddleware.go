package middlewares

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		tokenString = tokenString[len("Bearer "):]

		// err := verifyToken(tokenString)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			// Ensure that the token's signing method is HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// Return the secret key used for signing
			return []byte(os.Getenv("JWT_SECRET")), nil // Replace with your actual secret key
		})
		if err != nil {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// You can access the claims here if needed
			userID := uint(claims["id"].(float64))
			c.Set("userID", userID)
			c.Next() // Proceed to the next handler
		} else {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
		}
	}
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return os.Getenv("JWT_SECRET"), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return fmt.Errorf("invalid token claims")
	}
	if claims["exp"] == nil {
		return fmt.Errorf("token does not have expiration time")
	}

	return nil
}
