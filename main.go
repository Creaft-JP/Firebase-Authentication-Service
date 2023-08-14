package main

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"github.com/creaft-jp/firebase-authentication-service/auth"
	"github.com/creaft-jp/firebase-authentication-service/env"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"log"
)

func main() {
	_ = godotenv.Load()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	opt := option.WithCredentialsFile("resource/firebase-credentials.json")
	config := &firebase.Config{ProjectID: env.ProjectId}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	authController := auth.Controller{App: app}
	authController.Route(r.Group(""))

	if err := r.Run(); err != nil {
		panic(err)
	}
}
