package main

import (
	"geo.org/gin-members/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	controllers.InitMemberRoutes(router)

	//pprof.Register(router)

	router.Run(":8080")
}
