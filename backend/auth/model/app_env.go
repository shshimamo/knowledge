package model

import "os"

type AppEnv string

const (
	Production AppEnv = "production"
	Test       AppEnv = "test"
	Develop    AppEnv = "develop"
)

func NewAppEnv() AppEnv {
	switch os.Getenv("APP_ENV") {
	case "production":
		return Production
	case "test":
		return Test
	case "develop":
		return Develop
	default:
		return Develop
	}
}
