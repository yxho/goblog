package main

import (
	"goblog/model"
	"goblog/routes"
)

func main() {

	model.InitDb()
	routes.InitRouter()
}
