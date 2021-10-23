package Models

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iamrahultanwar/friskco/Config"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("unicorns")

func GetJWT(user *User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["ID"] = user.ID
	claims["user"] = user.Email
	claims["aud"] = "auth.friskco.com"
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role" gorm:"default:user"`
	Drives   []Drive
}

func GetCurrentUser(user *User, userId int) (err error) {
	if err = Config.DB.Find(&user, userId).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *User) (err error) {
	passwordHash, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = passwordHash
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func FindUserByEmail(user *User) (rows int, err error) {
	result := Config.DB.Where("email = ?", user.Email).First(user)
	return int(result.RowsAffected), result.Error
}

func LoginUser(user *User) (token string, err error) {
	password := user.Password
	_, userError := FindUserByEmail(user)

	if userError != nil {
		return "", errors.New("User not found")
	}

	if !checkPasswordHash(password, user.Password) {
		return "", errors.New("password does not match")
	}

	token, jwtError := GetJWT(user)

	return token, jwtError
}

type AuthHeader struct {
	Token string `header:"token"`
}

// authorized middleware
func Authorized() gin.HandlerFunc {

	return func(c *gin.Context) {
		h := AuthHeader{}
		c.ShouldBindHeader(&h)

		if len(h.Token) == 0 {
			c.String(http.StatusUnauthorized, "Token not found")
			c.Abort()
		}

		token, err := jwt.Parse(h.Token, func(token *jwt.Token) (interface{}, error) {

			return mySigningKey, nil
		})

		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			c.Set("userId", claims["ID"])
			c.Set("role", claims["role"])
			c.Set("email", claims["email"])
			c.Next()

		} else {
			c.String(http.StatusUnauthorized, "Invalid Auth")
			c.Abort()
			return
		}

	}
}
