package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/swanhtetaungphyo/burmaGuru/dto"
	"github.com/swanhtetaungphyo/burmaGuru/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"runtime"
	"time"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body"`
}

func NewResponse(statusCode int, message string, data interface{}) Response {
	return Response{
		statusCode,
		message,
		data,
	}
}
func ApiResponse(ctx *fiber.Ctx, statuscode int, message string, data interface{}) error {
	responseBody := NewResponse(statuscode, message, data)
	return ctx.Status(statuscode).JSON(responseBody)
}

func CompareHashAndPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func CurrentFunction() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func LogInit() *os.File {
	file, err := os.OpenFile("../application.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("Failed to open the log file: %s\n", err)
		return nil
	}
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	return file
}

func ToUser(userDTO dto.UserDTO) models.User {
	return models.User{
		UserName:               userDTO.UserName,
		Email:                  userDTO.Email,
		Password:               userDTO.Password,
		ProfilePicture:         userDTO.ProfilePicture,
		IsVerified:             userDTO.IsVerified,
		Interest:               userDTO.Interest,
		Country:                userDTO.Country,
		EmailVerificationToken: userDTO.EmailVerificationToken,
	}
}

func ToUserDTO(user models.User) dto.UserDTO {
	return dto.UserDTO{
		UserName:               user.UserName,
		Email:                  user.Email,
		Password:               user.Password,
		ProfilePicture:         user.ProfilePicture,
		IsVerified:             user.IsVerified,
		Interest:               user.Interest,
		Country:                user.Country,
		EmailVerificationToken: user.EmailVerificationToken,
	}
}

func GenereteTokens(tokenType string, username string, role string, expiration string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	duration, err := time.ParseDuration(expiration)
	if err != nil {
		log.Fatalf("%v", err.Error())
		return "", err
	}

	claims := jwt.MapClaims{
		"exp":  time.Now().Add(duration).Unix(),
		"name": username,
		"role": role,
		"type": tokenType,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
