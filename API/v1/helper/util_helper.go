package helper

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"time"

	// jwt "github.com/dgrijalva/jwt-go"  // original
	// jwtreq "github.com/dgrijalva/jwt-go/request"  // repositório original
	jwt "github.com/form3tech-oss/jwt-go"            // fork
	jwtreq "github.com/form3tech-oss/jwt-go/request" //fork
)

// GetSenhaSha256 retorna uma string hasheada(sha256) da senha informada no parâmetro senha
func GetSenhaSha256(senha string) string {
	senhaSha256 := sha256.Sum256([]byte(senha))

	dst := make([]byte, hex.EncodedLen(len(senhaSha256)))
	hex.Encode(dst, senhaSha256[:])

	return string(dst[:len(dst)])
}

// GetToken obtem o token da requisição http informada ao parâmetro r através da chave secreta contida no parâmetro secret_key. Se ocorrer um erro, retorna uma como token NIL e um erro
func GetToken(r *http.Request, secretKey []byte) (token *jwt.Token, err error) {
	tokenString, err := jwtreq.HeaderExtractor{"Authorization"}.ExtractToken(r)
	if err != nil {
		return
	}
	tokenString = strings.Split(tokenString, " ")[1]

	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
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
	}
	err = errors.New("Token vazio/inválido, sem claims")
	return

}

// SetClaims define os Clains ao token informado no primeiro parâmetro de acordo com os valores passados nos outros parâmetros(duracaoSegundos, usuario, email, admin)
func SetClaims(token *jwt.Token, duracaoSegundos time.Duration, usuario, email string, admin bool) (claims jwt.MapClaims) {
	claims = token.Claims.(jwt.MapClaims)
	claims["usuario"] = usuario
	claims["email"] = email
	claims["admin"] = admin
	claims["exp"] = time.Now().Add(time.Second * duracaoSegundos).Unix()

	return
}

// CriarDiretorioSeNaoExistir cria o diretório informado no parâmetro nomeDiretorio se ele não existir
func CriarDiretorioSeNaoExistir(nomeDiretorio string) (err error) {
	if _, err = os.Stat(nomeDiretorio); os.IsNotExist(err) {
		err = os.MkdirAll(nomeDiretorio, os.ModePerm)
		if err != nil {
			return
		}
	}
	return
}

// GetLocalIP retorna uma string contendo o endereço IP local do PC. Usado para exibir o endereço na inicialização da API. Fonte: https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go
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

// FormatarPorta recebe uma string representando a porta e deixa no formato :porta
func FormatarPorta(porta string) string {
	return fmt.Sprintf(":%s", porta)
}

// GetDiretorioAbs retorna uma string como diretório absoluto de executável. Se erro != nil, retorna um erro. Fonte: https://stackoverflow.com/questions/18537257/how-to-get-the-directory-of-the-currently-running-file
func GetDiretorioAbs() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}

// CriarDiretorioAbs cria o diretório informado em parâmetro string dirLOG
func CriarDiretorioAbs(dirLOG string) (diretorio string) {
	dirBase, _ := GetDiretorioAbs()
	dirBaseLog := path.Join(dirBase, dirLOG)
	err := CriarDiretorioSeNaoExistir(dirBaseLog)
	if err != nil {
		return ""
	}

	return dirBaseLog
}

// GetEstado retorna uma string representado o estado boleano informado em um texto amigável para usuários
func GetEstado(estado bool) string {
	estadoEmString := "ativo"
	if !estado {
		estadoEmString = "inativo"
	}

	return estadoEmString
}

// Fonte de funções abaixo: https://stackoverflow.com/questions/54558527/how-to-get-func-documentation-in-golang

// FuncPathAndName Get the name and path of a func
func FuncPathAndName(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

// FuncName Get the name of a func (with package path)
func FuncName(f interface{}) string {
	splitFuncName := strings.Split(FuncPathAndName(f), ".")
	return splitFuncName[len(splitFuncName)-1]
}

// FuncDescription Get description of a func
func FuncDescription(f interface{}) string {
	fileName, _ := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).FileLine(0)
	funcName := FuncName(f)
	fset := token.NewFileSet()

	// Parse src
	parsedAst, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	pkg := &ast.Package{
		Name:  "Any",
		Files: make(map[string]*ast.File),
	}
	pkg.Files[fileName] = parsedAst

	importPath, _ := filepath.Abs("/")
	myDoc := doc.New(pkg, importPath, doc.AllDecls)
	for _, theFunc := range myDoc.Funcs {
		if theFunc.Name == funcName {
			return theFunc.Doc
		}
	}
	return ""
}
