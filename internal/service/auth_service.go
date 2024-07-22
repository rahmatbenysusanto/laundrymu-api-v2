package service

import (
	"crypto/rsa"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"io/ioutil"
	"time"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func init() {
	privateKeyBytes, err := ioutil.ReadFile("private.pem")
	if err != nil {
		log.Fatalf("Error reading private key file: %v", err)
	}
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		log.Fatalf("Error parsing private key: %v", err)
	}

	// Load public key
	publicKeyBytes, err := ioutil.ReadFile("public.pem")
	if err != nil {
		log.Fatalf("Error reading public key file: %v", err)
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		log.Fatalf("Error parsing public key: %v", err)
	}
}

type UserClaims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func CreateToken(name, email, phone, role string) (string, error) {
	claims := UserClaims{
		Name:  name,
		Email: email,
		Phone: phone,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 336)),
			Issuer:    "laundrymu-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(privateKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && !token.Valid {
		log.Fatalf("Invalid token")
	}

	return nil
}
