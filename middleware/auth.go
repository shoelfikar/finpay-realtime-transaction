package middleware

import (
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/shoelfikar/finpay-realtime-transaction/model"
)

// GenerateJWT is func to generate the token
func GenerateJWT(user *model.User, exp time.Duration) (string, error) {
	var expMins time.Duration = exp
	token := jwt.New(jwt.SigningMethodHS256)
	jwtKey := []byte(os.Getenv("JWT_SECRET"))

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Minute * expMins).Unix()
	claims["user_id"] = user.Id

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Error("error signed token"+ err.Error())
		return "", err
	}
	return tokenString, nil
}

//IsAuthorized is the func for validating the JWT token
func JWTMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	// Check if header is empty or doesn't start with "Bearer "
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		log.Error("Missing or invalid token")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or invalid token"})
	}

	// Extract token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	// If token is invalid
	if err != nil || !token.Valid {
		log.Error("Invalid token")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	// Extract claims (Assuming it's a standard JWT)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Error("Invalid token claims")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
	}

	// Store claims in Fiber locals
	c.Locals("claims", claims)

	// Continue request
	return c.Next()
}