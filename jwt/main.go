package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

/**
HMACSHA256(
  base64UrlEncode(header) + "." +
  base64UrlEncode(payload),
  secret)

head.payload.signature
*/

type User struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

const secret = "hello"

var (
	header = map[string]string{
		"type":      "JWT",
		"algorithm": "HS256",
	}
)

func main() {
	router := gin.Default()

	router.StaticFile("/", "./index.html")

	router.POST("/login", handleLogin)

	router.GET("/user", handleUser)

	err := router.Run(":7000")

	if err != nil {
		log.Fatal(err)
	}
}

func handleLogin(context *gin.Context) {
	var user *User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	} else {
		authorization, err := jwtSignature(user)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}
		context.Header("Authorization", authorization)
		context.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}

func handleUser(context *gin.Context) {
	authorization := context.GetHeader("Authorization")
	if authorization == "" {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "no login",
		})
	} else {

	}
}

func jwtSignature(user *User) (string, error) {
	header, err := headerString()
	if err != nil {
		return "", err
	}
	payload, err := payloadString(user)
	if err != nil {
		return "", err
	}
	signature, err := computeSignature(header, payload)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s.%s.%s", header, payload, signature), nil
}

func computeSignature(header string, payload string) (string, error) {
	h := hmac.New(sha256.New, []byte(secret))
	message := fmt.Sprintf("%s.%s", header, payload)
	_, err := io.WriteString(h, message)
	if err != nil {
		return "", err
	}
	signature := h.Sum(nil)
	return string(signature), nil
}

func headerString() (string, error) {
	buffer, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(buffer), nil
}

func payloadString(user *User) (string, error) {
	buffer, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(buffer), nil
}
