package main

import (
	"session-redis/internal/router"
)

func main() {
	// Redis()
	r := router.SetupRouter()
	r.Run("localhost:5000")
}



