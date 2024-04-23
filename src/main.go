package main

import "backend/src/initializers"

func main() {
	initializers.LoadEnv()
	initializers.StartServer()
}
