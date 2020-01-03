package controller

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
)

var (
	idLanc01, idLanc02, idLanc03, idLanc04, idLanc05, idLanc06, idLanc07 int
)

func TestLancamentoCreate(t *testing.T) {
	// adicionado pessoa e conta a partir de funções de teste já criadas
	TestPessoaCreate(t)
	TestContaCreate(t)

	// criar Lancamentos(idLanc01, idLanc02, idLanc03) como administrador - 201
	rota := fmt.Sprintf("/lancamentos")
	res, body, err := post(rota, `{"cpf_pessoa":"00000002000",  "nome_conta_origem":"conta teste 01", "data":"2019-12-31T12:30:00.000Z", "numero":"A1", "descricao":"Lancamento teste 01", "nome_conta_destino":"conta teste 02", "debito":200.0, "credito":0.0}`, testTokenAdmin)
	if err != nil {
		t.Error(err)
		return
	}

	status := res.StatusCode
	if status != 201 {
		t.Error(res, string(body))
	} else {
		idLanc01 = getID(body)
	}

	res, body, err = post(rota, `{"cpf_pessoa":"00000003000",  "nome_conta_origem":"conta teste 02", "data":"2020-01-01T09:20:00.000Z", "numero":"B1", "descricao":"Lancamento teste 02", "nome_conta_destino":"conta teste 01", "debito":0.0, "credito":500.25}`, testTokenAdmin)
	status = res.StatusCode
	if status != 201 {
		t.Error(res, string(body))
	} else {
		idLanc02 = getID(body)
	}

	res, body, err = post(rota, `{"cpf_pessoa":"00000002000",  "nome_conta_origem":"conta teste 02", "data":"2020-01-02T08:30:00.000Z", "numero":"B1", "descricao":"Lancamento teste 03", "nome_conta_destino":"conta teste 03", "debito":253, "credito":13}`, testTokenAdmin)
	status = res.StatusCode
	if status != 201 {
		t.Error(res, string(body))
	} else {
		idLanc03 = getID(body)
	}

	// criar Lancamento como usuário comum - 201
	res, body, err = post(rota, `{"cpf_pessoa":"00000002000",  "nome_conta_origem":"conta teste 02", "data":"2020-01-02T08:30:00.000Z", "numero":"B1", "descricao":"Lancamento teste 03", "nome_conta_destino":"conta teste 03", "debito":253, "credito":13}`, testTokenComum)
	status = res.StatusCode
	if status != 201 {
		t.Error(res, string(body))
	} else {
		idLanc04 = getID(body)
	}

	// criar Lancamento com JSON inválido - 422
	res, body, err = post(rota, `"cpf_pessoa":"00000002000",  "nome_conta_origem":"conta teste 02", "data":"2020-01-02T08:30:00.000Z", "numero":"B1", "descricao":"Lancamento teste 03", "nome_conta_destino":"conta teste 03"`, testTokenAdmin)
	status = res.StatusCode
	if status != 422 {
		t.Error(res, string(body))
	}

	res, body, err = post(rota, `{"nome_conta_origem":"conta teste 02", "data":"2020-01-02T08:30:00.000Z", "descricao":"Lancamento teste 03", "nome_conta_destino":"conta teste 03", "debito":253, "credito":13}`, testTokenAdmin)
	status = res.StatusCode
	if status != 422 {
		t.Error(res, string(body))
	}

	// criar Lancamentos com os cpf do token admin e comum
	res, body, err = post(rota, `{"cpf_pessoa":"11111111111",  "nome_conta_origem":"conta teste 02", "data":"2020-01-01T09:20:00.000Z", "numero":"B1", "descricao":"Lancamento teste 02", "nome_conta_destino":"conta teste 01", "debito":0.0, "credito":500}`, testTokenAdmin)
	status = res.StatusCode
	if status != 201 {
		t.Error(res, string(body))
	} else {
		idLanc05 = getID(body)
	}

	res, body, err = post(rota, `{"cpf_pessoa":"11111111111",  "nome_conta_origem":"conta teste 02", "data":"2020-01-01T09:20:00.000Z", "numero":"B1", "descricao":"Lancamento teste 02", "nome_conta_destino":"conta teste 01", "debito":400.0, "credito":40}`, testTokenAdmin)
	status = res.StatusCode
	if status != 201 {
		t.Error(res, string(body))
	} else {
		idLanc06 = getID(body)
	}

	res, body, err = post(rota, `{"cpf_pessoa":"12345678910",  "nome_conta_origem":"conta teste 02", "data":"2020-01-02T08:30:00.000Z", "numero":"B1", "descricao":"Lancamento teste 03", "nome_conta_destino":"conta teste 03", "debito":200, "credito":10}`, testTokenComum)
	status = res.StatusCode
	if status != 201 {
		t.Error(res, string(body))
	} else {
		idLanc07 = getID(body)
	}
}

func TestLancamentoEstado(t *testing.T) {
	// Inativa lancamento(idLanc01) com usuário administrador - 200
	rota := fmt.Sprintf("/lancamentos/%d/estado", idLanc01)
	res, body, err := put(rota, `{"estado": false}`, testTokenAdmin)
	if err != nil {
		t.Error(err)
		return
	}

	status := res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// Inativar lancamento(idLanc02) como usuário comun - 200
	rota = fmt.Sprintf("/lancamentos/%d/estado", idLanc02)
	res, body, err = put(rota, `{"estado": false}`, testTokenComum)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// Ativar lancamento(idLanc02) como usuário comun - 500
	rota = fmt.Sprintf("/lancamentos/%d/estado", idLanc02)
	res, body, err = put(rota, `{"estado": true}`, testTokenComum)
	status = res.StatusCode
	if status != 500 {
		t.Error(res, string(body))
	}

	// Ativar lancamento(idLanc01) como usuário administrador - 200
	rota = fmt.Sprintf("/lancamentos/%d/estado", idLanc01)
	res, body, err = put(rota, `{"estado": true}`, testTokenAdmin)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}
}

func TestLancamentoIndex(t *testing.T) {
	var retorno ReturnData
	var quantAdmin, quantComum int

	// retorno de dados dos lancamento cadastradas usando um token como administrador - 200
	rota := "/lancamentos"
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

	// retorno de dados de lancamento usando um token de um usuário comum - 200
	res, body, _ = get(rota, testTokenComum)
	status = res.StatusCode
	json.Unmarshal(body, &retorno)
	quantComum = retorno.Count

	if status != 200 {
		t.Error(res, string(body))
	}

	// deve haver 2 lancamentos pertencentes ao usuário Admin e 1 ao usuário comum, tendo como diferença 1. Se for 0 ou menor(negativo), ocorreu algum problema
	diferenca := quantAdmin - quantComum
	if diferenca < 1 {
		t.Error("Diferença entre a quantidade de registros na busca de Lancamento como adminitrador está menor do que a quantidade de registros como usuário comum", diferenca, "=", quantAdmin, "-", quantComum)
	}
}

func TestLancamentoPorConta(t *testing.T) {
	var retorno ReturnData
	var quantAdmin, quantComum int

	// retorno de dados dos lancamento cadastradas usando um token como administrador e conta específica - 200
	conta := "conta teste 01"
	rota := fmt.Sprintf("/lancamentos_conta/%s", conta)
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

	// retorno de dados de lancamento usando um token de um usuário comum e conta específica - 200
	res, body, _ = get(rota, testTokenComum)
	status = res.StatusCode
	json.Unmarshal(body, &retorno)
	quantComum = retorno.Count

	if status != 200 {
		t.Error(res, string(body))
	}

	// deve haver 2 lancamentos pertencentes ao usuário Admin e 1 ao usuário comum, tendo como diferença 1. Se for 0 ou menor(negativo), ocorreu algum problema
	diferenca := quantAdmin - quantComum
	if diferenca < 1 {
		t.Error("Diferença entre a quantidade de registros na busca de Lancamento como adminitrador está menor do que a quantidade de registros como usuário comum", diferenca, "=", quantAdmin, "-", quantComum)
	}
}

func TestLancamentoShow(t *testing.T) {
	// retorno de dados de lancamento como administrador - 200
	rota := fmt.Sprintf("/lancamentos/%d", idLanc01)
	res, body, err := get(rota, testTokenAdmin)
	if err != nil {
		t.Error(err)
		return
	}

	status := res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// retorno de dados de lancamento como usuário comum - 200
	rota = fmt.Sprintf("/lancamentos/%d", idLanc01)
	res, body, _ = get(rota, testTokenComum)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// retorno de dados de lancamento inativo como usuário administrador - 200
	rota = fmt.Sprintf("/lancamentos/%d", idLanc02)
	res, body, _ = get(rota, testTokenAdmin)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// retorno de dados de lancamento inativo como usuário comum - 500
	rota = fmt.Sprintf("/lancamentos/%d", idLanc02)
	res, body, _ = get(rota, testTokenComum)
	status = res.StatusCode
	if status != 500 {
		t.Error(res, string(body))
	}

}

func TestLancamentoRemove(t *testing.T) {
	// Ao remover um lancamento, todos os seus detalhes são removido automaticamente

	// remover lancamento como usuário comum - 500
	id := strconv.Itoa(idLanc04)
	rota := fmt.Sprintf("/lancamentos/%s", id)
	res, body, _ := delete(rota, testTokenComum)
	status := res.StatusCode
	if status != 500 {
		t.Error(res, string(body))
	}

	// remove lancamentos como administrador - 200
	id = strconv.Itoa(idLanc01)
	rota = fmt.Sprintf("/lancamentos/%s", id)
	res, body, err := delete(rota, testTokenAdmin)
	if err != nil {
		t.Error(err, res, string(body))
	}
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	id = strconv.Itoa(idLanc02)
	rota = fmt.Sprintf("/lancamentos/%s", id)
	res, body, _ = delete(rota, testTokenAdmin)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	id = strconv.Itoa(idLanc03)
	rota = fmt.Sprintf("/lancamentos/%s", id)
	res, body, _ = delete(rota, testTokenAdmin)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	id = strconv.Itoa(idLanc04)
	rota = fmt.Sprintf("/lancamentos/%s", id)
	res, body, _ = delete(rota, testTokenAdmin)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	id = strconv.Itoa(idLanc05)
	rota = fmt.Sprintf("/lancamentos/%s", id)
	res, body, _ = delete(rota, testTokenAdmin)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	id = strconv.Itoa(idLanc06)
	rota = fmt.Sprintf("/lancamentos/%s", id)
	res, body, _ = delete(rota, testTokenAdmin)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	id = strconv.Itoa(idLanc07)
	rota = fmt.Sprintf("/lancamentos/%s", id)
	res, body, _ = delete(rota, testTokenAdmin)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	// remover lancamento inexistente, ID = 0 ou "abc" - 500
	rota = fmt.Sprintf("/lancamentos/%s", "0")
	res, body, _ = delete(rota, testTokenAdmin)
	status = res.StatusCode
	if status != 500 {
		t.Error(res, string(body))
	}

	rota = fmt.Sprintf("/lancamentos/%s", "abc")
	res, body, _ = delete(rota, testTokenAdmin)
	status = res.StatusCode
	if status != 500 {
		t.Error(res, string(body))
	}

	// removido contas e pessoas criadas em TesteLancamentoCreate
	TestContaRemove(t)
	TestPessoaRemove(t)
}

func getID(jsonEmBytes []byte) int {
	var jsonRetornado *ReturnData
	json.Unmarshal(jsonEmBytes, &jsonRetornado)
	d, _ := jsonRetornado.Data.([]interface{})
	m, _ := d[0].(map[string]interface{})
	id, _ := m["ID"].(float64)
	return int(id)
}
