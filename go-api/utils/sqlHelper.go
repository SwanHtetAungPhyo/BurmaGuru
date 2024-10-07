package utils

import (
	"bufio"
	"os"
	"strings"
)


type SqlQuaries struct {
	GetAllUsers string 
	GetUserByID string 
	CreateUser string 
	UpdateUser string 
	DeleteUser string 
}


func LoadSqlQuaries(filepath string) (*SqlQuaries, error){
	file, err := os.Open(filepath)
	if err != nil {
		return nil , err
	}

	defer file.Close()

	commandArray :=  &SqlQuaries{}
	scanner := bufio.NewScanner(file)

	var currentCommand string 
	var commandBuilder strings.Builder

	for scanner.Scan(){
		line := scanner.Text()
		if strings.HasPrefix(line,"--"){
			if currentCommand  != ""{
				assignQuary(currentCommand,commandBuilder.String(), commandArray)
				commandBuilder.Reset()
			}
			currentCommand = strings.TrimPrefix(line,"--")
		}else{
			commandBuilder.WriteString(line +"\n")
		}
	}
	if currentCommand != ""{
		assignQuary(currentCommand,commandBuilder.String(),commandArray)
	}

	if err := scanner.Err(); err != nil{
		return nil , err
	}
	return commandArray, nil 
}

func assignQuary(command, sqlText string,quary *SqlQuaries){
	switch command{
	case "SQL command for retrieving all users":
		quary.GetAllUsers = sqlText
	case "SQL command for retrieving a user by ID":
		quary.GetUserByID = sqlText
	case "SQL command for creating a new user":
		quary.CreateUser = sqlText
	case "SQL command for updating a user":
		quary.UpdateUser = sqlText
	case "SQL command for deleteing a user":
		quary.DeleteUser = sqlText
	}
}