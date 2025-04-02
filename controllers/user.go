package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/AhmedZeyad/AuthAPI/initializer"
	"github.com/AhmedZeyad/AuthAPI/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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

type SignupBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AddUser(context *gin.Context) {
	var body SignupBody

	// take user email and psee

	if context.BindJSON(&body) != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"msg": "not valid body"})

	}

	// }
	// hashed the pass
	bcryptPassword, err := hash(body.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "filed to hash password"})
		return
	}

	// add user to db

	result := initializer.DB.Exec("INSERT INTO user  VALUES (NULL,?, ?, NOW(), NOW())",
		body.Email, bcryptPassword)

	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to add user"})
		return
	}
	context.JSON(http.StatusAccepted, gin.H{"msg": 1})
}

func Login(context *gin.Context) {
	// take user email and password
	var body SignupBody
	if context.BindJSON(&body) != nil {

		context.JSON(http.StatusBadRequest, gin.H{"msg": "un valid body format"})
		return

	}
	// get user info
	var user models.User
	result := initializer.DB.Raw("Select * from user where email=?", body.Email).Scan(&user)

	if result.Error != nil {
		fmt.Println("Error fetching users:", result.Error)

	}

	// hased the password
	err := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(body.Password))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "un valid password"})
		return
	}

	// gen the jwt

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subject": user.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRETKEY")))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "failed to gen token"})
		fmt.Println(tokenString, err)
		return
	}
	context.SetSameSite(http.SameSiteLaxMode)
	context.SetCookie("token", tokenString, 3600*24*30, "", "", false, true)
	context.JSON(http.StatusOK, gin.H{"token": tokenString})
	// respons

}
func hash(value string) (string, error) {
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(value), 10)
	if err != nil {
		return "", errors.New("filed to hash password")
	}
	return string(bcryptPassword), nil
}
func Validation(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"msg": "you are authenticated"})

}