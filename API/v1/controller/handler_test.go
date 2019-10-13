package controller

import (
	"controle_pessoal_de_financas/API/v1/dao"
	"controle_pessoal_de_financas/API/v1/model/pessoa"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

// const URLBaseTest string = "https://064ce4ef.ngrok.io" // execute o script run.sh e posteriormente o script run_ngrok.sh para receber o ip externo(público) disponibilizado pelo NGROK e substitua a URLBaseTest pelo host atualizado
const URLBaseTest string = "https://localhost:8085" // local

var (
	TokenTest                      string
	testTokenAdmin, testTokenComum string
)

func init() {
	// Criação no BD das pessoas para o teste em handlers
	db := dao.GetDB()

	p1 := pessoa.New("11111111111", "teste 01", "teste01", "123456", "teste01@gmail.com")
	p2 := pessoa.New("12345678910", "Paulo C Silva Jr", "paulo", "123456", "pauluscave@gmail.com")
	p3 := pessoa.New("33333333333", "João Alcântara", "joao02", "123456", "joaoa@gmail.com")

	pessoas := pessoa.Pessoas{p1, p2, p3}

	for _, p := range pessoas {
		dao.AdicionaPessoa(db, p)
	}
	dao.SetAdministrador(db, p1.Cpf, true)
	dao.InativaPessoa(db, p3.Cpf)

	// tokens(Administrador e usuário comum) para teste
	testTokenAdmin, _ = getToken("teste01", "123456")
	testTokenComum, _ = getToken("paulo", "123456")
}

func TestLogin(t *testing.T) {
	// status OK - 200
	usuario := "teste01"
	senha := "123456"
	rota := fmt.Sprintf("/login/%s", usuario)
	jsonPost := fmt.Sprintf(`{"usuario":"%s", "senha":"%s"}`, usuario, senha)

	res, body, err := post(rota, jsonPost, "")
	if err != nil {
		t.Error(err)
		return
	}

	status := res.StatusCode
	if status == 200 {
		retorno := ReturnTokenJSON{}
		json.Unmarshal(body, &retorno)
		TokenTest = retorno.Token
	} else {
		t.Error(res, string(body))
	}

	// não pode processar o json - 422
	res, body, _ = post("/login/teste01", "", "")
	status = res.StatusCode
	if status != 422 {
		t.Error(res, string(body))
	}

	// não foi encontrado um registro com o usuário - 406
	usuario = "teste99"
	senha = "123"
	rota = fmt.Sprintf("/login/%s", usuario)
	jsonPost = fmt.Sprintf(`{"usuario":"%s", "senha":"%s"}`, usuario, senha)
	res, body, _ = post(rota, jsonPost, "")
	status = res.StatusCode
	if status != 406 {
		t.Error(res, string(body))
	}

	// usuário inativo - 404
	usuario = "joao02"
	senha = "123456"
	rota = fmt.Sprintf("/login/%s", usuario)
	jsonPost = fmt.Sprintf(`{"usuario":"%s",  "senha":"%s"}`, usuario, senha)
	res, body, _ = post(rota, jsonPost, "")
	status = res.StatusCode
	if status != 404 {
		t.Error(res, string(body))
	}

	// usuário ou senha inválida - 406
	usuario = "teste01"
	senha = "12345"
	rota = fmt.Sprintf("/login/%s", usuario)
	jsonPost = fmt.Sprintf(`{"usuario":"%s",  "senha":"%s"}`, usuario, senha)
	res, body, _ = post(rota, jsonPost, "")
	status = res.StatusCode
	if status != 406 {
		t.Error(res, string(body))
	}
}

func TestTokenValido(t *testing.T) {
	// token válido
	rota := "/token"
	res, body, err := get(rota, TokenTest)
	if err != nil {
		t.Error(err)
		return
	}

	status := res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// token inválido
	tokenTestInvalido := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZW1haWwiOiJ0ZXN0ZTAxQGdtYWlsLmNvbSIsImV4cCI6MTU1NjcxNjM3MiwidXN1YXJpbyI6InRlc3RlMDEifQ.pT6ljh_Ykw3FIEop9CTTuXlI1R1QpsplRBwa-LujiH8"
	res, body, _ = get(rota, tokenTestInvalido)
	status = res.StatusCode
	if status != 401 {
		t.Error(res, string(body))
	}
}

// TestIndex verifica se a rota /, que retorna as rotas cadastradas, está tendo um StatusCode diferente de 200
func TestIndex(t *testing.T) {
	rota := "/"
	res, body, err := get(rota, TokenTest)
	if err != nil {
		t.Error(err)
		return
	}

	status := res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}
}

func post(urlRota, json, token string) (*http.Response, []byte, error) {
	return requisicao("POST", urlRota, json, token)
}

func get(urlRota, token string) (*http.Response, []byte, error) {
	return requisicao("GET", urlRota, "", token)
}

func delete(urlRota, token string) (*http.Response, []byte, error) {
	return requisicao("DELETE", urlRota, "", token)
}

func put(urlRota, json, token string) (*http.Response, []byte, error) {
	return requisicao("PUT", urlRota, json, token)
}

func requisicao(tipo, urlRota, json, token string) (*http.Response, []byte, error) {
	url := URLBaseTest + urlRota

	payload := strings.NewReader(json)
	var req *http.Request
	if len(json) != 0 {
		req, _ = http.NewRequest(tipo, url, payload)
	} else {
		req, _ = http.NewRequest(tipo, url, nil)
	}

	if tipo == "POST" || tipo == "PUT" {
		req.Header.Add("Content-Type", "application/json")
	}

	if len(token) > 0 {
		req.Header.Add(
			"Authorization",
			fmt.Sprintf("Bearer %s", token),
		)
	}

	// Desativa segurança para não barrar o certificado auto-assinado
	// fonte: https://stackoverflow.com/questions/12122159/how-to-do-a-https-request-with-bad-certificate
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	return res, body, nil
}

func getToken(usuario, senha string) (string, error) {
	rota := fmt.Sprintf("/login/%s", usuario)
	jsonPost := fmt.Sprintf(`{"usuario":"%s", "senha":"%s"}`, usuario, senha)

	res, body, err := post(rota, jsonPost, "")
	if err != nil {
		return "", err
	}

	status := res.StatusCode
	if status == 200 {
		retorno := ReturnTokenJSON{}
		json.Unmarshal(body, &retorno)
		TokenTest = retorno.Token

		return TokenTest, nil
	}

	return "", fmt.Errorf("%v %s", res, string(body))
}
