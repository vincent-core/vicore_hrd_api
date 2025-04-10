package rest

import (
	"net/http"
	"os"
	"strings"
	"time"
	"vicore_hrd/pkg/helper"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	jwtMiddleware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Middleware interface {
	GenerateTokenUser(Email string, UserPeson string, KDBagian string) (map[string]string, error)
}

type jwtMidleware struct {
}

func NewMiddleware() *jwtMidleware {
	return &jwtMidleware{}
}

func CORSMiddleware() gin.HandlerFunc {
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
		if c.Request.Method == "OPTIONS" {
			c.Writer.Write([]byte("allowed"))
			return
		}

		c.Next()
	}
}

func GenerateTokenUser(Email string, UserPeson string, KdBagian string, UserID string) (map[string]string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = Email
	claims["person"] = UserPeson
	claims["kd_bagian"] = KdBagian
	claims["user_id"] = UserID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return nil, err
	}

	// REFRESH TOKEN NOT USE
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["email"] = Email
	rtClaims["person"] = UserPeson
	rtClaims["kd_bagian"] = KdBagian
	rtClaims["exp"] = time.Now().Add(time.Hour * 360).Unix()
	rtClaims["user_id"] = UserID

	resf := ""
	resf, err = refreshToken.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"token":         t,
		"refresh_token": resf,
	}, nil
}

func JWTProtected() func(*fiber.Ctx) error {
	config := jwtMiddleware.Config{
		SigningKey:     []byte(os.Getenv("JWT_SECRET_KEY")),
		ErrorHandler:   jwtError,
		SuccessHandler: jwtSuccess,
	}
	return jwtMiddleware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		response := helper.APIResponseFailure("Token Failed", http.StatusUnauthorized)
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}
	return c.Status(fiber.StatusUnauthorized).JSON(helper.APIResponseFailure("Token invalid", http.StatusUnauthorized))
}

func jwtSuccess(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if !strings.Contains(authHeader, "Bearer") {
		response := helper.APIResponseFailure("Unauthorized", http.StatusUnauthorized)
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

	tokenString := ""

	arrayToken := strings.Split(authHeader, " ")
	if len(arrayToken) == 2 {
		tokenString = arrayToken[1]
	}

	data, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		response := helper.APIResponseFailure("Token Expired", http.StatusUnauthorized)
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

	claim, ok := data.Claims.(jwt.MapClaims)

	if !ok || !data.Valid {
		response := helper.APIResponseFailure("Token Expired", http.StatusUnauthorized)
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

	person := string(claim["person"].(string))
	userID := string(claim["user_id"].(string))

	c.Locals("person", person)
	c.Locals("userID", userID)

	return c.Next()
}
