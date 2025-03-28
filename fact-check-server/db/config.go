package db

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func LoadConfig() Config {
	defaultDBHost := "localhost"
	defaultDBPort := 5432
	defaultDBUser := "postgres"
	defaultDBPassword := "1234"
	defaultDBName := "postgres"

	Host := os.Getenv("DB_HOST")

	if Host == "" {
		log.Println("DB_HOST environment variable not set, cascading to default: " + defaultDBHost)
		Host = defaultDBHost
	}

	var Port int
	inputPort := os.Getenv("DB_PORT")

	if inputPort == "" {
		log.Println("DB_PORT environment variable not set, cascading to default:", defaultDBPort)
		Port = defaultDBPort
	} else {
		var err error
		Port, err = strconv.Atoi(inputPort)
		if err != nil {
			log.Println("DB_PORT environment variable not set, cascading to default: " + inputPort)
			Port = defaultDBPort
		}
	}

	User := os.Getenv("DB_USER")

	if User == "" {
		log.Println("DB_USER environment variable not set, cascading to default: " + defaultDBUser)
		User = defaultDBUser
	}

	Password := os.Getenv("DB_PASSWORD")

	if Password == "" {
		log.Println("DB_PASSWORD environment variable not set, cascading to default: " + defaultDBPassword)
		Password = defaultDBPassword
	}

	DBName := os.Getenv("DB_NAME")

	if DBName == "" {
		log.Println("DB_NAME environment variable not set, cascading to default: " + defaultDBName)
		DBName = defaultDBName
	}

	config := Config{
		Host:     Host,
		Port:     Port,
		User:     User,
		Password: Password,
		DBName:   DBName,
	}
	return config
}