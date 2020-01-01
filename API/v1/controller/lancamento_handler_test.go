package controller

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
)

var (
	idLanc01 int
)

func TestLancamentoCreate(t *testing.T) {
	TestPessoaCreate(t)
	TestContaCreate(t)

	// criar Lancamento como administrador - 201 ou 500(chave duplicada)
	rota := fmt.Sprintf("/lancamentos")
	res, body, err := post(rota, `{"cpf_pessoa":"00000002000",  "nome_conta_origem":"conta teste 01", "data":"2019-12-31T12:30:00.000Z", "numero":"A1", "descricao":"Lancamento teste 01", "nome_conta_destino":"conta teste 02", "debito":200.0, "credito":0.0}`, testTokenAdmin)
	if err != nil {
		t.Error(err)
		return
	}

	status := res.StatusCode
	if !(status == 201 || status == 500) {
		t.Error(res, string(body))
	} else {
		idLanc01 = getID(body)
	}
}

func TestLancamentoRemove(t *testing.T) {
	// Ao remover um lancamento, todos os seus detalhes são removido automaticamente

	// remove conta como administrador
	id := strconv.Itoa(idLanc01)
	rota := fmt.Sprintf("/lancamentos/%s", id)
	res, body, _ := delete(rota, testTokenAdmin)
	status := res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

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
