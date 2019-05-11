package controller

import (
	"fmt"
	"testing"
)

func TestPessoaIndex(t *testing.T) {
	// retorno de dados das pessoas cadastradas como administrador - 200
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

	// retorno de dados das pessoas cadastradas como usuário comum - 200
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

	// retorno de caso pessoa informada não seja cadastrada no BD - 404
	admin = "teste01"
	usuario = "joao03"
	tokenPessoaAdmin, _ = getToken(admin, "123456")
	rota = fmt.Sprintf("/pessoas/%s/%s", admin, usuario)
	res, body, _ = get(rota, tokenPessoaAdmin)
	status = res.StatusCode
	if status != 404 {
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

func TestPessoaCreate(t *testing.T) {
	// criar pessoa como administrador - 201
	admin := "teste01"
	tokenPessoaAdmin, _ := getToken(admin, "123456")
	rota := fmt.Sprintf("/pessoas")
	res, body, err := post(rota, `{"cpf":"00000002000",  "nome_completo":"Teste 20", "usuario":"teste20", "senha":"20123456", "email":"teste20@email.com"}`, tokenPessoaAdmin)
	if err != nil {
		t.Error(err)
		return
	}

	status := res.StatusCode
	if status != 201 {
		t.Error(res, string(body))
	}

	// chave duplicada na inclusão de pessoa como admin - 500
	res, body, _ = post(rota, `{"cpf":"00000002000",  "nome_completo":"Teste 20", "usuario":"teste20", "senha":"20123456", "email":"teste20@email.com"}`, tokenPessoaAdmin)
	status = res.StatusCode
	if status != 500 {
		t.Error(res, string(body))
	}

	// criar pessoa2 como administrador - 201
	rota = fmt.Sprintf("/pessoas")
	res, body, _ = post(rota, `{"cpf":"00000003000",  "nome_completo":"Teste 30", "usuario":"teste30", "senha":"30123456", "email":"teste30@email.com"}`, tokenPessoaAdmin)
	status = res.StatusCode
	if status != 201 {
		t.Error(res, string(body))
	}

	// criar pessoa3 como administrador - 201
	rota = fmt.Sprintf("/pessoas")
	res, body, _ = post(rota, `{"cpf":"00000004000",  "nome_completo":"Teste 40", "usuario":"teste40", "senha":"40123456", "email":"teste40@email.com"}`, tokenPessoaAdmin)
	status = res.StatusCode
	if status != 201 {
		t.Error(res, string(body))
	}

	// criar pessoa como usuário comum - 500
	admin = "paulo"
	tokenPessoaAdmin, _ = getToken(admin, "123456")
	rota = fmt.Sprintf("/pessoas")
	res, body, _ = post(rota, `{"cpf":"00000005000",  "nome_completo":"Teste 50", "usuario":"teste50", "senha":"50123456", "email":"teste50@email.com"}`, tokenPessoaAdmin)
	status = res.StatusCode
	if status != 500 {
		t.Error(res, string(body))
	}

	// erro ao criar pessoa com json inválido - 422
	admin = "teste01"
	tokenPessoaAdmin, _ = getToken(admin, "123456")
	rota = fmt.Sprintf("/pessoas")
	res, body, _ = post(rota, "", tokenPessoaAdmin)
	status = res.StatusCode
	if status != 422 {
		t.Error(res, string(body))
	}
}

func TestPessoaRemove(t *testing.T) {
	// remove pessoa1 como administrador - 200
	admin := "teste01"
	tokenPessoaAdmin, _ := getToken(admin, "123456")
	usuario := "teste20"
	rota := fmt.Sprintf("/pessoas/%s", usuario)
	res, body, err := delete(rota, tokenPessoaAdmin)
	if err != nil {
		t.Error(err)
		return
	}

	status := res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// remove pessoa2 como administrador - 200
	usuario = "teste30"
	rota = fmt.Sprintf("/pessoas/%s", usuario)
	res, body, _ = delete(rota, tokenPessoaAdmin)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// remove pessoa3 como administrador - 200
	usuario = "teste40"
	rota = fmt.Sprintf("/pessoas/%s", usuario)
	res, body, _ = delete(rota, tokenPessoaAdmin)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}
}
