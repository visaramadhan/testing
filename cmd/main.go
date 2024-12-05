package main

import (
	"log"
	"project/infra"
	"project/routes"
)

func main() {
	ctx, err := infra.NewServiceContext()
	if err != nil {
		log.Fatal("can't init service context %w", err)
	}

	r := routes.NewRoutes(*ctx)

	if err = r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
