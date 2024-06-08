package main

import (
	"backend/db"
	"backend/src/initializers"
)

func main() {
	initializers.LoadEnv()
	db.ConnectDB()
	initializers.StartApp()
}
