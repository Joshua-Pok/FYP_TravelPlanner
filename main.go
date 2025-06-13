package main

import (
	"github.com/gin-gonic/gin"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	server := gin.Default()
	InitializeDB()
	server.GET("/plans")
	server.POST("/plans")
	server.Run(":8080")

}
