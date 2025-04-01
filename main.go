package main

import (
	"net/http"
	"github.com/AhmedZeyad/AuthAPI/initializer"
	"github.com/gin-gonic/gin"
	"github.com/AhmedZeyad/AuthAPI/controllers"
)
func init(){
	initializer.LoadEnvVariables()
	initializer.ConnectDB()
	initializer.SyncDB()
}
func main() {

	println("Hello, World!")
	router:=gin.Default()
// set the end point
router.GET("/test",func(ctx *gin.Context) {
	users,err:=controllers.GetUser()
	ctx.JSON(
http.StatusOK,
	gin.H{"msg":"API is runing",
	"err":err,
	"data":users,
})

})
router.POST("/signup",controllers.AddUser)

	// run the api
	router.Run(":9090")
}
