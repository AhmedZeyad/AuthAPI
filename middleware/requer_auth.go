package middleware

import (
	// "fmt"
	"net/http"
	"os"
	"time"

	"github.com/AhmedZeyad/AuthAPI/initializer"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CheckAuth(context *gin.Context) {
	println("im form middleware")
	// git token from cookie
	tokenString,err:=context.Cookie("token")
if err!=nil{
	context.JSON(http.StatusUnauthorized,gin.H{"msg":"Unautherized user"})
	context.Abort()
	return 
}



	// decode the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRETKEY")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil && token == nil {
		context.JSON(http.StatusUnauthorized,gin.H{"msg":"Unautherized user2"})
		context.Abort()
		return 

	}
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// check token exp
		if float64(time.Now().Unix()) > claims["exp"].(float64){
			context.JSON(http.StatusUnauthorized,gin.H{"msg":"Unautherized user3"})
			context.Abort()
			return

			

		}
		// chek if user exist
		result :=initializer.DB.Exec("SELECT * FROM user WHERE ID=?",claims["subject"])
		if result.Error != nil {
			context.JSON(http.StatusUnauthorized,gin.H{"msg":"Unautherized user4"})
			context.Abort()
			return 

		}
	// continue
	context.Next()

	} else {
		context.JSON(http.StatusUnauthorized,gin.H{"msg":"Unautherized user5"})
		context.Abort()
		return 
	}
	
}
