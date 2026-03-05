// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/mustafa4341/database"
// )

// func main() {
// 	client := database.Connect()
// 	if client == nil {
// 		log.Fatal("Failed to connect to MongoDB")
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	err := client.Ping(ctx, nil)
// 	if err != nil {
// 		log.Fatal("Failed to ping MongoDB")
// 	}
// 	fmt.Println("Successfully connected and pinged MongoDB!")

// 	col := database.OpenCollection("movies", client)
// 	if col == nil {
// 		log.Fatal("Failed to open collection")
// 	}
// 	fmt.Println("Successfully opened collection:", col.Name())
// }
