package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
	"runtime"
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

func EmptyMap() map[string]string {
	mapp := make(map[string]string)
	return mapp
}

func DbEnvload() (string, string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error in %v is %v ", CurrentFunction(), err)
	}
	DbUser, DbPassword, DbName := os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME")
	return DbUser, DbPassword, DbName
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
