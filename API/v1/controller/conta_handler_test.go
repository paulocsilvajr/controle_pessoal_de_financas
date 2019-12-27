package controller

import (
	"fmt"
	"testing"
)

func TestContaCreate(t *testing.T) {
	// criar Conta como administrador - 201 ou 500(chave duplicada)
	post("/tipos_conta", `{"nome":"tipo conta teste 01",  "descricao_debito":"saída", "descricao_credito":"entrada"}`, testTokenAdmin)

	rota := fmt.Sprintf("/contas")
	res, body, err := post(rota, `{"nome":"conta teste 01",  "nome_tipo_conta":"tipo conta teste 01", "codigo":"000000001", "conta_pai": "", "comentario": "Teste de inclusão 01 de Conta via POST"}`, testTokenAdmin)
	if err != nil {
		t.Error(err)
		return
	}

	status := res.StatusCode
	if !(status == 201 || status == 500) {
		t.Error(res, string(body))
	}

	// criar conta com conta pai a conta "conta teste 01" - 201 ou 500(chave duplicada)
	res, body, err = post(rota, `{"nome":"conta teste 02",  "nome_tipo_conta":"tipo conta teste 01", "codigo":"000000002", "conta_pai": "conta teste 01", "comentario": "Teste de inclusão 02 de Conta via POST"}`, testTokenAdmin)
	status = res.StatusCode
	if !(status == 201 || status == 500) {
		t.Error(res, string(body))
	}

	// criar Conta sem código, conta pai e comentário - 201 ou 500(chave duplicada)
	res, body, err = post(rota, `{"nome":"conta teste 03",  "nome_tipo_conta":"tipo conta teste 01", "codigo":"", "conta_pai": "", "comentario": ""}`, testTokenAdmin)
	status = res.StatusCode
	if !(status == 201 || status == 500) {
		t.Error(res, string(body))
	}

	// criar Conta como usuário comum com conta pai "conta teste 03" - 201 ou 500(chave duplicada)
	res, body, err = post(rota, `{"nome":"conta teste 04",  "nome_tipo_conta":"tipo conta teste 01", "codigo":"000000004", "conta_pai": "conta teste 03", "comentario": "Teste de inclusão 04 de Conta via POST"}`, testTokenComum)
	status = res.StatusCode
	if !(status == 201 || status == 500) {
		t.Error(res, string(body))
	}

	// chave duplicada na inclusão de tipo conta como admin - 500
	res, body, err = post(rota, `{"nome":"conta teste 01",  "nome_tipo_conta":"tipo conta teste 01", "codigo":"000000001", "conta_pai": "", "comentario": "Teste de inclusão 01 de Conta via POST"}`, testTokenAdmin)
	status = res.StatusCode
	if status != 500 {
		t.Error(res, string(body))
	}

	// criar conta com campo obrigatório vazio("nome_tipo_conta": "") - 422
	res, body, err = post(rota, `{"nome":"conta teste 01",  "nome_tipo_conta":"", "codigo":"000000001", "conta_pai": "", "comentario": "Teste de inclusão 01 de Conta via POST"}`, testTokenAdmin)
	status = res.StatusCode
	if status != 422 {
		t.Error(res, string(body))
	}

	// erro ao criar conta com json inválido - 422
	res, body, err = post(rota, `{"nome":"conta teste 05", "codigo":"000000001", "comentario": "Teste de inclusão 05 de Conta via POST"}`, testTokenAdmin)
	status = res.StatusCode
	if status != 422 {
		t.Error(res, string(body))
	}
}

func TestContaEstado(t *testing.T) {
	// Inativa conta(conta teste 01) com usuário administrador - 200
	conta := "conta teste 01"
	rota := fmt.Sprintf("/contas/%s/estado", conta)
	res, body, err := put(rota, `{"estado": false}`, testTokenAdmin)
	if err != nil {
		t.Error(err)
		return
	}

	status := res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// Inativar conta(conta teste 04) como usuário comun - 200
	conta = "conta teste 04"
	rota = fmt.Sprintf("/contas/%s/estado", conta)
	res, body, err = put(rota, `{"estado": false}`, testTokenComum)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// Ativar conta(conta teste 01) como usuário comum - 500
	conta = "conta teste 01"
	rota = fmt.Sprintf("/contas/%s/estado", conta)
	res, body, err = put(rota, `{"estado": true}`, testTokenComum)
	status = res.StatusCode
	if status != 500 {
		t.Error(res, string(body))
	}

	// Ativar conta(conta teste 01) como administrador - 200
	rota = fmt.Sprintf("/contas/%s/estado", conta)
	res, body, err = put(rota, `{"estado": true}`, testTokenAdmin)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}
}

func TestContaShow(t *testing.T) {
	// retorno de dados de conta como administrador - 200
	conta := "conta teste 01"
	rota := fmt.Sprintf("/contas/%s", conta)
	res, body, err := get(rota, testTokenAdmin)
	if err != nil {
		t.Error(err)
		return
	}

	status := res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// retorno de dados de conta como usuário comum - 200
	rota = fmt.Sprintf("/contas/%s", conta)
	res, body, _ = get(rota, testTokenComum)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// consulta de conta inativa como administrador - 200
	conta = "conta teste 04"
	rota = fmt.Sprintf("/contas/%s", conta)
	res, body, _ = get(rota, testTokenAdmin)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// consulta de conta inativa como usuário comum - 500
	rota = fmt.Sprintf("/contas/%s", conta)
	res, body, _ = get(rota, testTokenComum)
	status = res.StatusCode
	if status != 500 {
		t.Error(res, string(body))
	}
}

func TestContaAlter(t *testing.T) {
	// alterar conta(conta teste 01) como administrador - 200
	conta := "conta teste 01"
	rota := fmt.Sprintf("/contas/%s", conta)
	res, body, err := put(rota, `{"nome":"conta teste 06",  "nome_tipo_conta":"tipo conta teste 01", "codigo":"000000006", "conta_pai":"", "comentario":"Teste de alteração 01 de Conta via PUT"}`, testTokenAdmin)
	if err != nil {
		t.Error(err)
		return
	}

	status := res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	conta = "conta teste 06"
	rota = fmt.Sprintf("/contas/%s", conta)
	res, body, err = put(rota, `{"nome":"conta teste 01",  "nome_tipo_conta":"tipo conta teste 01", "codigo":"000000001", "conta_pai": "", "comentario": "Teste de alteração 02 de Conta via PUT"}`, testTokenAdmin)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// alter conta como usuário comum - 200
	conta = "conta teste 02"
	rota = fmt.Sprintf("/contas/%s", conta)
	res, body, err = put(rota, `{"nome":"conta teste 02",  "nome_tipo_conta":"tipo conta teste 01", "codigo":"000000007", "conta_pai": "conta teste 01", "comentario": "Teste de alteração 03 de Conta via PUT"}`, testTokenComum)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// alter conta como administrador sem informar em JSON o campo nome - 200
	conta = "conta teste 01"
	rota = fmt.Sprintf("/contas/%s", conta)
	res, body, err = put(rota, `{"nome_tipo_conta":"tipo conta teste 01", "codigo":"000000001", "conta_pai": "", "comentario": "Teste de alteração 04 de Conta via PUT"}`, testTokenAdmin)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// conta não pôde ser alterada por não existir no BD - 304
	conta = "conta teste 05"
	rota = fmt.Sprintf("/contas/%s", conta)
	res, body, err = put(rota, `{"nome":"conta teste 06",  "nome_tipo_conta":"tipo conta teste 01", "codigo":"000000008", "conta_pai": "", "comentario": "Teste de alteração 05 de Conta via PUT"}`, testTokenAdmin)
	status = res.StatusCode
	if status != 304 {
		t.Error(res, string(body))
	}

	// Dados em JSON não podem ser processados - 422
	conta = "conta teste 01"
	rota = fmt.Sprintf("/contas/%s", conta)
	res, body, err = put(rota, `"nome":"conta teste 01",  "nome_tipo_conta":"tipo conta teste 01", "codigo":"000000009", "conta_pai": "", "comentario": "Teste de alteração 06 de Conta via PUT"`, testTokenAdmin)
	status = res.StatusCode
	if status != 422 {
		t.Error(res, string(body))
	}
}

func TestContaRemove(t *testing.T) {
	// remove conta como usuário comum - 500
	conta := "conta teste 01"
	rota := fmt.Sprintf("/contas/%s", conta)
	res, body, _ := delete(rota, testTokenComum)
	status := res.StatusCode
	if status != 500 {
		t.Error(res, string(body))
	}

	// remove conta que não existe como administrador - 500
	conta = "conta teste 05"
	rota = fmt.Sprintf("/contas/%s", conta)
	res, body, _ = delete(rota, testTokenAdmin)
	status = res.StatusCode
	if status != 500 {
		t.Error(res, string(body))
	}

	// remove conta como administrador
	conta = "conta teste 01"
	rota = fmt.Sprintf("/contas/%s", conta)
	res, body, _ = delete(rota, testTokenAdmin)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// "conta teste 02" é removido automaticamente quando removido "conta teste 01", FOREIGN KEY com DELETE CASCADE para o campo "conta_pai"

	conta = "conta teste 03"
	rota = fmt.Sprintf("/contas/%s", conta)
	res, body, _ = delete(rota, testTokenAdmin)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// "conta teste 04" é removido automaticamente quando removido "conta teste 03", FOREIGN KEY com DELETE CASCADE para o campo "conta_pai"

	// removido TipoConta "tipo conta teste 01"
	tipoConta := "tipo conta teste 01"
	rota = fmt.Sprintf("/tipos_conta/%s", tipoConta)
	res, body, _ = delete(rota, testTokenAdmin)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}
}
