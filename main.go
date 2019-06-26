package main

import (
	"crud-go-redis/api"
)

const (
	apiPort = 3000
)

func main() {
	api.Connect(apiPort)
}
