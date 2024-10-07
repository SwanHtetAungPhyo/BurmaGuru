package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassowrd(password string) (string, error){
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		return "", err
	}
	return string(hashedPassword), nil
}

func LogInit() *os.File {
	file, err := os.OpenFile("./application.log",  os.O_CREATE| os.O_APPEND | os.O_WRONLY, 0666)
	if err != nil{
		fmt.Printf("Failed to open the log file: %s\n", err)
		return nil
	}
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	return file
}

func EmailTokenGenerator(length int) string {
	const charset ="abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	tokenByte := make([]byte,length)
	for i := range tokenByte {
		tokenByte[i] = charset[rand.Intn(len(charset))]
	}
	return string(tokenByte)
}

func SendEmail(to []string, subject, body string) error {
	smtpHost := "smtp.gmail.com" 
	smtpPort := "587"
	username := "swanhtet102002@gmail.com"
	password := "Swanhtet12@"

	auth := smtp.PlainAuth("",username, password, smtpHost)

	header := make(map[string]string)
	header["From"] = username
	header["To"] = strings.Join(to,",")
	header["Subject"] = subject
	header["MIME-version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"UTF-8\""
	message := ""
	for key, value := range header {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	message += "\r\n" + body
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, username, to, []byte(message))
	if err != nil {
		return err
	}
	return nil

}

func  LoadSqlFile(filename string)  (string, error){
	content , err := ioutil.ReadFile(filepath.Join("sql",filename))
	if err != nil {
		return "", err
	}
	return string(content), nil 
}
