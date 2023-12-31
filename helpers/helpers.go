package helpers

import (
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const JWT_SECRET = "kuncirahasia"

func Validation(c *gin.Context, data interface{}) error {
	_, err := govalidator.ValidateStruct(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return err
	}
	return err
}

func EncryptPassword(c *gin.Context, password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed To Encrypt Password"})
		return ""
	}
	return string(hash)
}

func CheckPassword(pass_1 string, pass_2 string) error {
	err := bcrypt.CompareHashAndPassword([]byte(pass_1), []byte(pass_2))
	if err != nil {
		return err
	}
	return err
}

func InitializeToken(userid string) (string, error) {

	expirationTime := time.Now().Add(24 * time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userid,
		"exp": expirationTime.Unix(),
	})

	tokenStr, err := token.SignedString([]byte(JWT_SECRET))
	return tokenStr, err
}

