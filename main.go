package main

import (
	"session-redis/internal/router"
)

func main() {
	r := router.SetupRouter()
	r.Run("localhost:5000")
}



