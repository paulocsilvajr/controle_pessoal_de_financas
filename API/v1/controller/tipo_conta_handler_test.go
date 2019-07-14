package controller

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTipoContaCreate(t *testing.T) {
	// criar TipoConta como administrador - 201 ou 500(chave duplicada)
	rota := fmt.Sprintf("/tipos_conta")
	res, body, err := post(rota, `{"nome":"tipo conta teste 01",  "descricao_debito":"saída", "descricao_credito":"entrada"}`, testTokenAdmin)
	if err != nil {
		t.Error(err)
		return
	}

	status := res.StatusCode
	if !(status == 201 || status == 500) {
		t.Error(res, string(body))
	}

	res, body, err = post(rota, `{"nome":"tipo conta teste 02",  "descricao_debito":"saída", "descricao_credito":"entrada"}`, testTokenAdmin)
	status = res.StatusCode
	if !(status == 201 || status == 500) {
		t.Error(res, string(body))
	}

	res, body, err = post(rota, `{"nome":"tipo conta teste 03",  "descricao_debito":"saída", "descricao_credito":"entrada"}`, testTokenAdmin)
	status = res.StatusCode
	if !(status == 201 || status == 500) {
		t.Error(res, string(body))
	}

	// criar TipoConta como usuário comum - 201 ou 500(chave duplicada)
	res, body, err = post(rota, `{"nome":"base",  "descricao_debito":"-", "descricao_credito":"+"}`, testTokenComum)
	status = res.StatusCode
	if !(status == 201 || status == 500) {
		t.Error(res, string(body))
	}

	// chave duplicada na inclusão de tipo conta como admin - 500
	res, body, err = post(rota, `{"nome":"base",  "descricao_debito":"-", "descricao_credito":"+"}`, testTokenAdmin)
	status = res.StatusCode
	if status != 500 {
		t.Error(res, string(body))
	}

	// erro ao criar tipo conta com json inválido - 422
	res, body, err = post(rota, `"nome":"tipo conta teste 04",  "descricao_debito":"saída", "descricao_credito":"entrada"`, testTokenAdmin)
	status = res.StatusCode
	if status != 422 {
		t.Error(res, string(body))
	}
}

func TestTipoContaEstado(t *testing.T) {
	// Inativa tipo conta(tipo conta teste 01) com usuário administrador - 200
	tipoConta := "tipo conta teste 01"
	rota := fmt.Sprintf("/tipos_conta/%s/estado", tipoConta)
	res, body, err := put(rota, `{"estado": false}`, testTokenAdmin)
	if err != nil {
		t.Error(err)
		return
	}

	status := res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// Inativar tipo conta(base) como usuário comun - 200
	tipoConta = "base"
	rota = fmt.Sprintf("/tipos_conta/%s/estado", tipoConta)
	res, body, err = put(rota, `{"estado": false}`, testTokenComum)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// Ativar tipo conta(tipo conta teste 01) como usuário comum - 500
	tipoConta = "tipo conta teste 01"
	rota = fmt.Sprintf("/tipos_conta/%s/estado", tipoConta)
	res, body, err = put(rota, `{"estado": true}`, testTokenComum)
	status = res.StatusCode
	if status != 500 {
		t.Error(res, string(body))
	}

	// Ativar tipo conta(tipo conta teste 01) como administrador - 200
	tipoConta = "tipo conta teste 01"
	rota = fmt.Sprintf("/tipos_conta/%s/estado", tipoConta)
	res, body, err = put(rota, `{"estado": true}`, testTokenAdmin)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}
}

func TestTipoContaIndex(t *testing.T) {
	var retorno ReturnData
	var quantAdmin, quantComum int

	// retorno de dados dos tipos de conta cadastradas usando um token como administrador - 200
	rota := "/tipos_conta"
	res, body, err := get(rota, testTokenAdmin)
	if err != nil {
		t.Error(err)
		return
	}

	status := res.StatusCode
	json.Unmarshal(body, &retorno)
	quantAdmin = retorno.Count

	if status != 200 {
		t.Error(res, string(body))
	}

	// retorno de dados de tipos de conta usando um token de um usuário comum - 200
	res, body, _ = get(rota, testTokenComum)
	status = res.StatusCode
	json.Unmarshal(body, &retorno)
	quantComum = retorno.Count

	if status != 200 {
		t.Error(res, string(body))
	}

	diferenca := quantAdmin - quantComum
	if diferenca < 1 {
		t.Error("Diferença entre a quantidade de registros na busca de Tipo de Conta como adminitrador está menor do que a quantidade de registros como usuário comum")
	}
}

func TestTipoContaShow(t *testing.T) {
	// retorno de dados de tipo conta como administrador - 200
	tipoConta := "tipo conta teste 01"
	rota := fmt.Sprintf("/tipos_conta/%s", tipoConta)
	res, body, err := get(rota, testTokenAdmin)
	if err != nil {
		t.Error(err)
		return
	}

	status := res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// retorno de dados de tipo conta como usuário comum - 200
	rota = fmt.Sprintf("/tipos_conta/%s", tipoConta)
	res, body, _ = get(rota, testTokenComum)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// consulta de tipo conta inativo como administrador - 200
	tipoConta = "base"
	rota = fmt.Sprintf("/tipos_conta/%s", tipoConta)
	res, body, _ = get(rota, testTokenAdmin)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// consulta de tipo conta inativo como usuário comum - 500
	rota = fmt.Sprintf("/tipos_conta/%s", tipoConta)
	res, body, _ = get(rota, testTokenComum)
	status = res.StatusCode
	if status != 500 {
		t.Error(res, string(body))
	}
}

func TestTipoContaAlter(t *testing.T) {
	// alterar tipo conta(tipo conta teste 01) como administrador - 200
	tipoConta := "tipo conta teste 01"
	rota := fmt.Sprintf("/tipos_conta/%s", tipoConta)
	res, body, err := put(rota, `{"nome":"tipo conta teste 04",  "descricao_debito":"<", "descricao_credito":">"}`, testTokenAdmin)
	if err != nil {
		t.Error(err)
		return
	}

	status := res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	tipoConta = "tipo conta teste 04"
	rota = fmt.Sprintf("/tipos_conta/%s", tipoConta)
	res, body, err = put(rota, `{"nome":"tipo conta teste 01",  "descricao_debito":"diminuir", "descricao_credito":"aumentar"}`, testTokenAdmin)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// alter tipo conta como usuário comum - 200
	tipoConta = "tipo conta teste 02"
	rota = fmt.Sprintf("/tipos_conta/%s", tipoConta)
	res, body, err = put(rota, `{"nome":"tipo conta teste 02",  "descricao_debito":"-", "descricao_credito":"+"}`, testTokenComum)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// alter tipo conta como administrador sem informar em JSON o campo nome - 200
	tipoConta = "tipo conta teste 01"
	rota = fmt.Sprintf("/tipos_conta/%s", tipoConta)
	res, body, err = put(rota, `{"descricao_debito":"-", "descricao_credito":"+"}`, testTokenAdmin)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// tipo conta não pôde ser alterada por não existir no BD - 304
	tipoConta = "tipo conta teste 05"
	rota = fmt.Sprintf("/tipos_conta/%s", tipoConta)
	res, body, err = put(rota, `{"descricao_debito":"menos", "descricao_credito":"mais"}`, testTokenAdmin)
	status = res.StatusCode
	if status != 304 {
		t.Error(res, string(body))
	}

	// Dados em JSON não podem ser processados - 422
	tipoConta = "tipo conta teste 01"
	rota = fmt.Sprintf("/tipos_conta/%s", tipoConta)
	res, body, err = put(rota, `"descricao_debito":"-", "descricao_credito":"+"`, testTokenAdmin)
	status = res.StatusCode
	if status != 422 {
		t.Error(res, string(body))
	}
}

func TestTipoContaRemove(t *testing.T) {
	// remove tipo conta como usuário comum - 500
	tipoConta := "tipo conta teste 01"
	rota := fmt.Sprintf("/tipos_conta/%s", tipoConta)
	res, body, _ := delete(rota, testTokenComum)
	status := res.StatusCode
	if status != 500 {
		t.Error(res, string(body))
	}

	// remove tipo conta que não existe como administrador - 500
	tipoConta = "tipo conta teste 05"
	rota = fmt.Sprintf("/tipos_conta/%s", tipoConta)
	res, body, _ = delete(rota, testTokenAdmin)
	status = res.StatusCode
	if status != 500 {
		t.Error(res, string(body))
	}

	// remove tipos conta como administrador
	tipoConta = "tipo conta teste 01"
	rota = fmt.Sprintf("/tipos_conta/%s", tipoConta)
	res, body, _ = delete(rota, testTokenAdmin)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	tipoConta = "tipo conta teste 02"
	rota = fmt.Sprintf("/tipos_conta/%s", tipoConta)
	res, body, _ = delete(rota, testTokenAdmin)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	tipoConta = "tipo conta teste 03"
	rota = fmt.Sprintf("/tipos_conta/%s", tipoConta)
	res, body, _ = delete(rota, testTokenAdmin)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	tipoConta = "base"
	rota = fmt.Sprintf("/tipos_conta/%s", tipoConta)
	res, body, _ = delete(rota, testTokenAdmin)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}
}
