package middlewares

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gin-gonic/gin"
)

// VerifyJwt verify jwt is valid
var VerifyJwt gin.HandlerFunc = func(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	// Verifica se o header existe e se tem o formato correto (Bearer <token>)
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(401, gin.H{
			"Data": "Token inexistente ou formato inválido",
		})
		c.Abort()
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) < 2 {
		c.JSON(401, gin.H{
			"Data": "Token mal formatado",
		})
		c.Abort()
		return
	}

	HeaderToken := parts[1]
	hash := jwt.NewHS256([]byte("secret"))
	payload := jwt.Payload{}

	token := []byte(HeaderToken)
	hd, err := jwt.Verify(token, hash, &payload)

	if err != nil {
		c.JSON(401, gin.H{
			"Data": "Token Invalido",
		})
		c.Abort()
		return
	} else {
		tokenInt, err := strconv.Atoi(hd.KeyID)
		if err != nil {
			fmt.Println(err)
		}
		tokenUint := uint(tokenInt)

		c.Set("id", tokenUint)
		c.Next()
	}
}

// GenerateJwt return JWT
func GenerateJwt(ID uint) string {
	hs := jwt.NewHS256([]byte("secret"))
	now := time.Now()
	pl := jwt.Payload{
		Issuer:   "gbrlsnchs",
		Subject:  "login",
		IssuedAt: jwt.NumericDate(now),
	}
	stringID := strconv.Itoa(int(ID))
	token, err := jwt.Sign(pl, hs, jwt.ContentType("JWT"), jwt.KeyID(stringID))
	if err != nil {
		log.Fatal(err)
	}
	return string(token)
}