package env

import (
	"os"
	"log"
	"fmt"
	"path/filepath"
	
	"github.com/joho/godotenv"
)

const (
	APP_ENV = iota
	PORT
	DB_DRIVER
	DB_HOST
	DB_USER
	DB_PASSWORD
	DB_NAME
	DB_PORT
	JWT_SECRET
	AWS_REGION
	S3_BUCKET_NAME
)

var keys = []string{
	"APP_ENV",
	"PORT",
	"DB_DRIVER",
	"DB_HOST",
	"DB_USER",
	"DB_PASSWORD",
	"DB_NAME",
	"DB_PORT",
	"JWT_SECRET",
	"AWS_REGION",
	"S3_BUCKET_NAME",
}

var envs = map[string]string{
	"TEST":        ".env.test",
	"DEV":         ".env.dev",
	"PRODUCTION":  ".env",
}


func LoadEnv() {
	_, isPresent := os.LookupEnv(keys[APP_ENV])
	if !isPresent {
		err := os.Setenv(keys[APP_ENV], "DEV")
		if err != nil {
			log.Fatalf("error occured when setting APP_ENV %v", err)
		}
	}
	envFile, isValidEnv := envs[GetEnv(APP_ENV)]
	if !isValidEnv {
		log.Printf("invalid APP_ENV: %s\n", GetEnv(APP_ENV))
	} else {
		path := envFile
		if GetEnv(APP_ENV) == "TEST" {
			pwd, err := os.Getwd()
			if err != nil {
				panic(err)
			}
			path = filepath.Join(pwd, "../" + envFile)
		}
		err := godotenv.Load(path)
		if err != nil {
			log.Printf("%s file not found\n", envFile)
		}
	}
	log.Printf("Environment: %s\n", GetEnv(APP_ENV))
	for idx := range keys {
		key := keys[idx]
		_, isPresent := os.LookupEnv(key)
		if !isPresent {
			panic(fmt.Sprintf("missing environment variable: %s", key))
		}
	}
}

func GetEnv(key int) string {
	if key < 0 || key >= len(keys) {
		panic("Invalid env key")
	}

	return os.Getenv(keys[key])
}
