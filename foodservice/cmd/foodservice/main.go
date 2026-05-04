package main

import (
	"context"
	"log"

	"github.com/boginskiy/FoodMoment/foodservice/internal/app"
)

func main() {
	ctx := context.Background()

	application, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init application: %s", err.Error())
	}

	err = application.Run(ctx)
	if err != nil {
		log.Fatalf("failed to run application: %s", err.Error())
	}
}
