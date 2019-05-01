package controller

import (
	"fmt"
	"testing"
)

func TestPessoaIndex(t *testing.T) {
	tokenPessoaAdmin, _ := getToken("teste01", "123456")
	rota := "/pessoas"
	res, body, err := get(rota, tokenPessoaAdmin)
	if err != nil {
		t.Error(err)
		return
	}

	status := res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	tokenPessoaComum, _ := getToken("paulo", "123456")
	res, body, _ = get(rota, tokenPessoaComum)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}
}

func TestPessoaShow(t *testing.T) {
	// retorno de dados de pessoa admin - 200
	usuario := "teste01"
	tokenPessoaAdmin, _ := getToken(usuario, "123456")
	rota := fmt.Sprintf("/pessoas/%s", usuario)
	res, body, err := get(rota, tokenPessoaAdmin)
	if err != nil {
		t.Error(err)
		return
	}

	status := res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// retorno de dados de pessoa comum - 200
	usuario = "paulo"
	tokenPessoaComum, _ := getToken(usuario, "123456")
	rota = fmt.Sprintf("/pessoas/%s", usuario)
	res, body, _ = get(rota, tokenPessoaComum)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// usuário token diferente do usuário da rota - 500
	usuario = "joao02"
	tokenPessoaAdmin, _ = getToken("teste01", "123456")
	rota = fmt.Sprintf("/pessoas/%s", usuario)
	res, body, _ = get(rota, tokenPessoaAdmin)
	status = res.StatusCode
	if status != 500 {
		t.Error(res, string(body))
	}
}

func TestPessoaShowAdmin(t *testing.T) {
	// retorno de dados de pessoa comum com token de admin - 200
	admin := "teste01"
	usuario := "joao02"
	tokenPessoaAdmin, _ := getToken(admin, "123456")
	rota := fmt.Sprintf("/pessoas/%s/%s", admin, usuario)
	res, body, err := get(rota, tokenPessoaAdmin)
	if err != nil {
		t.Error(err)
		return
	}

	status := res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// retorno de dados de pessoa comum com token comum - 500
	usuario = "paulo"
	tokenPessoaComun, _ := getToken(usuario, "123456")
	rota = fmt.Sprintf("/pessoas/%s/%s", usuario, usuario)
	res, body, _ = get(rota, tokenPessoaComun)
	status = res.StatusCode
	if status != 500 {
		t.Error(res, string(body))
	}

	// rota somente para administradores - 500
	usuario = "paulo"
	usuario2 := "paulo2"
	tokenPessoaComun, _ = getToken(usuario, "123456")
	rota = fmt.Sprintf("/pessoas/%s/%s", usuario, usuario2)
	res, body, _ = get(rota, tokenPessoaComun)
	status = res.StatusCode
	if status != 500 {
		t.Error(res, string(body))
	}
}
