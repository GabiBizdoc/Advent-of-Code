package main

import (
	"aoc/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	server.Start()
}
