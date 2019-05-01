package controller

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

const URLBaseTest string = "https://localhost:8085/"

var (
	TokenTest string
)

func TestLogin(t *testing.T) {
	// url := URLBaseTest + "login/teste01"

	// payload := strings.NewReader(`{"usuario":"teste01",  "senha":"123456"}`)

	// req, _ := http.NewRequest("POST", url, payload)

	// req.Header.Add("Content-Type", "application/json")

	// // Desativa segurança para não barrar o certificado auto-assinado
	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{
	// 		InsecureSkipVerify: true,
	// 	},
	// }
	// client := &http.Client{Transport: tr}
	// res, _ := client.Do(req)

	// body, _ := ioutil.ReadAll(res.Body)

	// if res.StatusCode == 200 {
	// 	t.Error(res)
	// 	t.Error(string(body))
	// }

	// res.Body.Close()

	usuario := "teste01"
	senha := "123456"
	jsonPost := fmt.Sprintf(`{"usuario":"%s",  "senha":"%s"}`, usuario, senha)

	res, body := post("login/teste01", jsonPost, "")

	if res.StatusCode == 200 {
		retorno := ReturnTokenJson{}
		json.Unmarshal(body, &retorno)
		TokenTest = retorno.Token
	} else {
		t.Error(res, body)
	}
}

func post(urlRota, json, token string) (*http.Response, []byte) {
	return requisicao("POST", urlRota, json, token)
}

func requisicao(tipo, urlRota, json, token string) (*http.Response, []byte) {
	url := URLBaseTest + urlRota

	payload := strings.NewReader(json)
	if tipo == "GET" || tipo == "DELETE" {
		payload = nil
	}

	req, _ := http.NewRequest(tipo, url, payload)

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
	res, _ := client.Do(req)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	return res, body
}
