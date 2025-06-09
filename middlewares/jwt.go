package middlewares

import (
	"english-ai-go/repositories"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(email string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Minute)

	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

func AuthToken(repo repositories.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lấy token từ cookie
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token not found!",
			})
			c.Abort()
			return
		}

		// Parse và kiểm tra token
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid or expired token!",
			})
			c.Abort()
			return
		}

		user, err := repo.FindByEmail(claims.Email)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "User not found!",
			})
			c.Abort()
			return
		}

		// Lưu thông tin user vào context
		c.Set("user", user)

		c.Next()
	}
}
