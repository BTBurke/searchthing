package main

import "github.com/gin-gonic/gin"
import "github.com/BTBurke/searchthing/search"

func main() {
	router := gin.Default()

	router.GET("/search", search.Search)

	router.Run(":10298")
}
