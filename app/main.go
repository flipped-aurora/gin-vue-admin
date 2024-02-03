package main

import "app/router"

func main() {
	//r := gin.Default()
	//r.Run(":8081")
	r := router.GetRouter()
	r.Run(":8082")
	return
}
