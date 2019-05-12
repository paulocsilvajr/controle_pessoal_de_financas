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

func TestPessoaAlter(t *testing.T) {
	// alterar pessoa(teste20) com usuário administrador(teste01) - 200
	admin := "teste01"
	tokenPessoaAdmin, _ := getToken(admin, "123456")
	usuario := "teste20"
	rota := fmt.Sprintf("/pessoas/%s", usuario)
	res, body, err := put(rota, `{"cpf":"00000002000",  "nome_completo":"Teste 20", "usuario":"teste20", "senha":"123456", "email":"teste2020@email.com"}`, tokenPessoaAdmin)
	if err != nil {
		t.Error(err)
		return
	}

	status := res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// alterar o dados do próprio usuário comum(teste20) - 200
	comum := "teste20"
	tokenPessoaAdmin, _ = getToken(comum, "123456")
	usuario = "teste20"
	rota = fmt.Sprintf("/pessoas/%s", usuario)
	res, body, _ = put(rota, `{"cpf":"00000002000",  "nome_completo":"Teste 20 alterado por ele mesmo", "usuario":"teste20", "senha":"123456", "email":"teste20@email.com"}`, tokenPessoaAdmin)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// alterar pessoa como usuário comum(teste20) diferente dele mesmo(teste30) - 500
	comum = "teste20"
	tokenPessoaAdmin, _ = getToken(comum, "123456")
	usuario = "teste30"
	rota = fmt.Sprintf("/pessoas/%s", usuario)
	res, body, _ = put(rota, `{"cpf":"00000009900",  "nome_completo":"Teste alteração", "usuario":"testeNaoExiste", "senha":"99123456", "email":"teste99@email.com"}`, tokenPessoaAdmin)
	status = res.StatusCode
	if status != 500 {
		t.Error(res, string(body))
	}

	// pessoa(teste99) não pôde ser alterada por não existir no BD - 304
	admin = "teste01"
	tokenPessoaAdmin, _ = getToken(admin, "123456")
	usuario = "teste99"
	rota = fmt.Sprintf("/pessoas/%s", usuario)
	res, body, _ = put(rota, `{"cpf":"00000002020",  "nome_completo":"Teste 20", "usuario":"teste20", "senha":"2020123456", "email":"teste2020@email.com"}`, tokenPessoaAdmin)
	status = res.StatusCode
	if status != 304 {
		t.Error(res, string(body))
	}

	// Dados em JSON não podem ser processados - 422
	admin = "teste01"
	tokenPessoaAdmin, _ = getToken(admin, "123456")
	usuario = "teste20"
	rota = fmt.Sprintf("/pessoas/%s", usuario)
	res, body, _ = put(rota, "", tokenPessoaAdmin)
	status = res.StatusCode
	if status != 422 {
		t.Error(res, string(body))
	}
}

func TestPessoaEstado(t *testing.T) {
	// Ativa pessoa(teste20) com usuário administrador(teste01) - 200
	admin := "teste01"
	tokenPessoaAdmin, _ := getToken(admin, "123456")
	usuario := "teste20"
	rota := fmt.Sprintf("/pessoas/%s/estado", usuario)
	res, body, err := put(rota, `{"estado": true}`, tokenPessoaAdmin)
	if err != nil {
		t.Error(err)
		return
	}

	status := res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// Inativa pessoa(teste20) com usuário administrador(teste01) - 200
	usuario = "teste20"
	rota = fmt.Sprintf("/pessoas/%s/estado", usuario)
	res, body, _ = put(rota, `{"estado": false}`, tokenPessoaAdmin)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// Ativa pessoa(teste01) com usuário seu próprio usuário - 500
	usuario = "teste01"
	rota = fmt.Sprintf("/pessoas/%s/estado", usuario)
	res, body, _ = put(rota, `{"estado": true}`, tokenPessoaAdmin)
	status = res.StatusCode
	if status != 500 {
		t.Error(res, string(body))
	}

	// Erro ao ativa pessoa(teste99) que não exista no DB - 404
	usuario = "teste99"
	rota = fmt.Sprintf("/pessoas/%s/estado", usuario)
	res, body, _ = put(rota, `{"estado": true}`, tokenPessoaAdmin)
	status = res.StatusCode
	if status != 404 {
		t.Error(res, string(body))
	}
}

func TestPessoaAdmin(t *testing.T) {
	// Define como administrador a pessoa(teste20) com usuário administrador(teste01) - 200
	admin := "teste01"
	tokenPessoaAdmin, _ := getToken(admin, "123456")
	usuario := "teste20"
	rota := fmt.Sprintf("/pessoas/%s/admin", usuario)
	res, body, err := put(rota, `{"adminstrador": true}`, tokenPessoaAdmin)
	if err != nil {
		t.Error(err)
		return
	}

	status := res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// Retira privilégio de administrador para a pessoa(teste20) com usuário administrador(teste01) - 200
	usuario = "teste20"
	rota = fmt.Sprintf("/pessoas/%s/admin", usuario)
	res, body, err = put(rota, `{"adminstrador": false}`, tokenPessoaAdmin)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// Retira privilégio de administrador para a pessoa(teste01) com o próprio usuário - 500
	usuario = "teste01"
	rota = fmt.Sprintf("/pessoas/%s/admin", usuario)
	res, body, err = put(rota, `{"adminstrador": false}`, tokenPessoaAdmin)
	status = res.StatusCode
	if status != 500 {
		t.Error(res, string(body))
	}

	// Define como administrador a pessoa(teste01) com usuário comum(teste20) - 500
	comum := "paulo"
	tokenPessoaComum, _ := getToken(comum, "123456")
	usuario = "teste40"
	rota = fmt.Sprintf("/pessoas/%s/admin", usuario)
	res, body, _ = put(rota, `{"adminstrador": true}`, tokenPessoaComum)
	status = res.StatusCode
	if status != 500 {
		t.Error(res, string(body))
	}
}

func TestPessoaRemove(t *testing.T) {
	// remove pessoa1 como usuário comum - 500
	admin := "paulo"
	tokenPessoaAdmin, _ := getToken(admin, "123456")
	usuario := "teste20"
	rota := fmt.Sprintf("/pessoas/%s", usuario)
	res, body, _ := delete(rota, tokenPessoaAdmin)
	status := res.StatusCode
	if status != 500 {
		t.Error(res, string(body))
	}

	// remove pessoa com usuário desta mesma pessoa - 500
	admin = "teste01"
	tokenPessoaAdmin, _ = getToken(admin, "123456")
	usuario = "teste01"
	rota = fmt.Sprintf("/pessoas/%s", usuario)
	res, body, _ = delete(rota, tokenPessoaAdmin)
	status = res.StatusCode
	if status != 500 {
		t.Error(res, string(body))
	}

	// remove pessoa não cadastrada com usuário admin - 500
	admin = "teste01"
	tokenPessoaAdmin, _ = getToken(admin, "123456")
	usuario = "testeABC"
	rota = fmt.Sprintf("/pessoas/%s", usuario)
	res, body, _ = delete(rota, tokenPessoaAdmin)
	status = res.StatusCode
	if status != 500 {
		t.Error(res, string(body))
	}

	// remove pessoa1 como administrador - 200
	admin = "teste01"
	tokenPessoaAdmin, _ = getToken(admin, "123456")
	usuario = "teste20"
	rota = fmt.Sprintf("/pessoas/%s", usuario)
	res, body, err := delete(rota, tokenPessoaAdmin)
	if err != nil {
		t.Error(err)
		return
	}

	status = res.StatusCode
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
