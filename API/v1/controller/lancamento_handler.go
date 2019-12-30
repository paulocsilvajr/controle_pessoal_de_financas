package controller

import (
	"net/http"
	"time"
)

type lancamentoPersJSON struct {
	ID               int       `json:"id"`
	CpfPessoa        string    `json:"cpf_pessoa"`
	NomeContaOrigem  string    `json:"nome_conta_origem"`
	Data             time.Time `json:"data"`
	Numero           string    `json:"numero"`
	Descricao        string    `json:"descricao"`
	NomeContaDestino string    `json:"nome_conta_destino"`
	Debito           float64   `json:"debito"`
	Credito          float64   `json:"credito"`
	DataCriacao      time.Time `json:"data_criacao"`
	DataModificacao  time.Time `json:"data_modificacao"`
	Estado           time.Time `json:"estado"`
}

func LancamentoIndex(w http.ResponseWriter, r *http.Request) {
	// var status = http.StatusInternalServerError // 500

	// token, err := helper.GetToken(r, GetMySigningKey())
	// err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	// if err != nil {
	// 	return
	// }

	// usuarioToken, _, admin, err := helper.GetClaims(token)
	// err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	// if err != nil {
	// 	return
	// }

	// _, err = dao.ProcuraPessoaPorUsuario(db, usuarioToken)
	// err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	// if err != nil {
	// 	return
	// }

}
