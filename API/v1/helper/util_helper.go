package helper

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	jwtreq "github.com/dgrijalva/jwt-go/request"
)

func GetSenhaSha256(senha string) string {
	senhaSha256 := sha256.Sum256([]byte(senha))

	dst := make([]byte, hex.EncodedLen(len(senhaSha256)))
	hex.Encode(dst, senhaSha256[:])

	return string(dst[:len(dst)])
}

func GetToken(r *http.Request, secret_key []byte) (token *jwt.Token, err error) {
	tokenString, err := jwtreq.HeaderExtractor{"Authorization"}.ExtractToken(r)
	if err != nil {
		return
	}
	tokenString = strings.Split(tokenString, " ")[1]

	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secret_key, nil
	})

	return
}

// GetClaims obtém variáveis incluídas no token e verifica se ele é válido. error != nil caso token inválido.
func GetClaims(token *jwt.Token) (usuario string, email string, admin bool, err error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuario = claims["usuario"].(string)
		email = claims["email"].(string)
		admin = claims["admin"].(bool)
		return
	} else {
		err = errors.New("Token vazio/inválido, sem claims")
		return
	}

}

func SetClaims(token *jwt.Token, duracaoSegundos time.Duration, usuario, email string, admin bool) (claims jwt.MapClaims) {
	claims = token.Claims.(jwt.MapClaims)
	claims["usuario"] = usuario
	claims["email"] = email
	claims["admin"] = admin
	claims["exp"] = time.Now().Add(time.Second * duracaoSegundos).Unix()

	return
}

func CriarDiretorioSeNaoExistir(nomeDiretorio string) (err error) {
	if _, err = os.Stat(nomeDiretorio); os.IsNotExist(err) {
		err = os.MkdirAll(nomeDiretorio, os.ModePerm)
		if err != nil {
			return
		}
	}
	return
}

func GetLocalIP() string {
	// Fonte: https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
