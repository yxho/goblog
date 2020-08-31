package main

import (
	"goblog/model"
	"goblog/routes"
)

// 充满希望
func main() {

	model.InitDb()
	routes.InitRouter()
}
