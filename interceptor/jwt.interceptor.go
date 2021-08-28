package interceptor

import (
	"fmt"
	"main/model"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gin-gonic/gin"
)

var secretKey = "Br6v9UKqwRChmnqMdcfq3nBd8JBxsmVj4evEaQE2FDRZEQ39MYnZ5wMuFBcvBwkQ3emcjh5F3M2haZK59CuVcPFXc8ZndwNXvEDBhJmTNCXfnQvuQujFSAPVFxgUZv5ksLSNhMvc6S9A9T2Zhcck2LUaWLgjVdf7VNzusqPQWckrmxvvtSWzf9TgrLvBRUpB6bRsBLvh7DGRr7mKEFWjcBR3j6WszJeYjwbEjxjLvdKLmRPw68JzBUYZWYja3yYhHctXpH5MxuVd9BXLVduQ8VjA3yVvcbKPfaL5MACWE78ve9TtFhfUHRB5t7WCCLkL8amVYABzKDnesJue8e3DEg3nLe7GAqZ2qCQdLruywZDkJtSWGFB6dECduzJ6HcN9stcrKeHdqBpTAaj894PSKeKjUKKrcPZehZJD5KSTGx2VV2QFA8QnhKPycFuzuCYPLcQsHtxgvNumuf6Xy6LhR6muwhV5KGuNDUfSC3VKKE33vyyBWZ2APtfcfwmsL2pZxbggRmqR9uZVGPeZ3XEENBdkvqdxZm9eFZReuqkK4xHtVvLdqr8XeDGpt4Mq5skGK7WPcd7Dp7GLvFEvHswHpuEUj7B85PChKKgfkUUUF6JZArSmmUEj7qKJgwAuej3bwqStzLexauc89zWzgmCAznH6futqLgAScYVs4tnZRcFzrzxGRGSany4jUzAN5DCLen5EQadHewyZc3XDsMnJYWjML5pXcpDvLajC4HudSsCNnyX9X5uUD3zQzPcGp4fxLBuPs3HubtPmxntUmXbJMd5ewu3wmzRrGMVGukDBGfkaSNsujvEs8Zc63FEDw6HcveRTw2pXuFFY4AVYep5PMZrQYWaz54nZ3GKxTdZjTLY3j6yjn9VZptXXvkGm9ZWdwFJNaMc2kqV5yEWxMFcgeLz2Eh2Kz2uKZtjd4dRb7NZBKZJF4x4E6L3Mhnn8W53sJ4xcH7b4teXXfGJEcnr2McJkFNbeMBgFYEvZ54fbrHa3FTvLyfEhDKvptGCZyjAB"

func JwtSign(payload model.Score) string {
	atClaims := jwt.MapClaims{}

	// Payload begin
	atClaims["id"] = payload.ID
	atClaims["student_id"] = payload.StudentID
	atClaims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	// Payload end

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, _ := at.SignedString([]byte(secretKey))
	return token

}

func JwtVerify(c *gin.Context) {
	tokenString := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	// fmt.Println(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// fmt.Println(claims)

		studentID := fmt.Sprintf("%v", claims["student_id"])
		c.Set("jwt_student_id", studentID)

		c.Next()
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "nok", "message": "invalid token", "error": err})
		c.Abort()
	}
}
