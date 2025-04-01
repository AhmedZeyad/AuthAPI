package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/AhmedZeyad/AuthAPI/initializer"
	"github.com/AhmedZeyad/AuthAPI/models"
	"github.com/gin-gonic/gin"
)

func GetUser() ([]models.User, error) {
	// var result User
	var users []models.User

	result := initializer.DB.Raw("SELECT * FROM user").Scan(&users)
	if result.Error != nil {
		fmt.Println("Error fetching users:", result.Error)
		return users, errors.New("not foun user")
	}

	// Print all users
	for _, user := range users {
		fmt.Printf("ID: %d, Email: %s, CreatedAt: %s\n",
			user.ID,
			deref(user.Email, "N/A"), // Handle nil values
			user.CreatedAt.Format("2006-01-02 15:04:05"))
	}
	return users, nil
}
func deref(s *string, defaultValue string) string {
	if s == nil {
		return defaultValue
	}
	return *s
}
type SignupBody struct{
	Email    string `json:"email"`
    Password string `json:"password"`
} 
func AddUser(context *gin.Context){
	var body SignupBody 

	// take user email and psee 

if context.BindJSON(&body)!=nil{
	context.JSON(
		http.StatusBadRequest,
		gin.H{"msg":"not valid body",})

}

// }
	// hashed the pass


// add user to db
	
initializer.DB.Raw("INSERT INTO USER VALUSE(%s,%s,%s,%s)",body.Email,body.Password,time.DateTime,time.DateTime)
context.JSON(http.StatusAccepted,gin.H{"msg":1})}
