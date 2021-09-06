package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/dao"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/helper"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/pessoa"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// PessoaIndex é um handler/controller que responde a rota '[GET] /pessoas' e retorna StatusOK(200) e uma listagem de pessoas de acordo com o tipo de usuário(admin/comum) caso o TOKEN informado for válido e o usuário associado ao token for cadastrado na API/DB. Caso ocorra algum erro, retorna StatusInternalServerError(500)
func PessoaIndex(w http.ResponseWriter, r *http.Request) {
	db02 := dao.GetDB02()
	defer dao.CloseDB(db02)

	var status = http.StatusInternalServerError // 500

	token, err := helper.GetToken(r, GetMySigningKey())
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioToken, emailToken, admin, err := helper.GetClaims(token)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	var listaPessoas pessoa.IPessoas
	if admin {
		listaPessoas, err = dao.CarregaPessoas02(db02)
		err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
		if err != nil {
			return
		}
	} else {
		listaPessoas, err = dao.CarregaPessoasSimples02(db02)
		err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
		if err != nil {
			return
		}
	}

	p, err := listaPessoas.ProcuraPessoaPorUsuario(usuarioToken)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	err = DefineHeaderRetorno(w, SetHeaderJSON, p.GetEmail() != emailToken, status, errors.New("email de token não confere com email de pessoa"))
	if err != nil {
		return
	}

	status = http.StatusOK // 200
	funcao := "PessoaIndex"
	DefineHeaderRetornoDados(
		w,
		SetHeaderJSON,
		status,
		listaPessoas,
		funcao,
		"Listagem de pessoas",
		"Enviando listagem de pessoas")
}

// PessoaShow é um handler/controller que responde a rota '[GET] /pessoas/{usuario}' e retorna StatusOK(200) e os dados da pessoa(usuário) solicitada caso o TOKEN informado for válido e o usuário associado ao token for cadastrado na API/DB e igual ao usuário da rota. Caso ocorra algum erro, retorna StatusInternalServerError(500)
func PessoaShow(w http.ResponseWriter, r *http.Request) {
	db02 := dao.GetDB02()
	defer dao.CloseDB(db02)

	var status = http.StatusInternalServerError // 500

	vars := mux.Vars(r)
	usuarioRota := vars["usuario"]

	token, err := helper.GetToken(r, GetMySigningKey())
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioToken, _, _, err := helper.GetClaims(token)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	verif := usuarioToken != usuarioRota
	err = DefineHeaderRetorno(w, SetHeaderJSON, verif, status, errors.New("usuário de token diferente do solicitado na rota"))
	if err != nil {
		return
	}

	pessoaEncontrada, err := dao.ProcuraPessoaPorUsuario02(db02, usuarioRota)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	status = http.StatusOK // 200
	funcao := "PessoaShow"
	DefineHeaderRetornoDado(
		w,
		SetHeaderJSON,
		status,
		pessoaEncontrada,
		funcao,
		fmt.Sprintf("Dados de pessoa '%s'", pessoaEncontrada.Usuario),
		fmt.Sprintf("Enviando dados de pessoa '%s'", pessoaEncontrada.Usuario))
}

// PessoaShowAdmin é um handler/controller que responde a rota '[GET] /pessoas/{usuarioAdmin}/{usuario}' e retorna StatusOK(200) e os dados da pessoa(usuário) solicitada caso o TOKEN informado for válido e o usuário administrador associado ao token for cadastrado na API/DB e igual ao usuário admin da rota. Caso não for encontrado o usuário informado no BD, retorna StatusNotFound(404). Para os outros erros, retorna StatusInternalServerError(500)
func PessoaShowAdmin(w http.ResponseWriter, r *http.Request) {
	db02 := dao.GetDB02()
	defer dao.CloseDB(db02)

	var status = http.StatusInternalServerError // 500

	vars := mux.Vars(r)
	usuarioAdmin := vars["usuarioAdmin"]
	usuario := vars["usuario"]

	token, err := helper.GetToken(r, GetMySigningKey())
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioToken, _, admin, err := helper.GetClaims(token)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	verif := usuarioToken != usuarioAdmin
	err = DefineHeaderRetorno(w, SetHeaderJSON, verif, status, errors.New("usuário de token diferente do informado na rota"))
	if err != nil {
		return
	}

	usuarioDB, err := dao.ProcuraPessoaPorUsuario02(db02, usuarioAdmin)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	verif = !(admin && usuarioDB.Administrador)
	err = DefineHeaderRetorno(w, SetHeaderJSON, verif, status, errors.New("somente administradores podem usar essa rota"))
	if err != nil {
		return
	}

	status = http.StatusNotFound // 404
	pessoaEncontrada, err := dao.ProcuraPessoaPorUsuario02(db02, usuario)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	status = http.StatusOK // 200
	funcao := "PessoaShowAdmin"
	DefineHeaderRetornoDado(
		w,
		SetHeaderJSON,
		status,
		pessoaEncontrada,
		funcao,
		fmt.Sprintf("Dados de pessoa '%s'", pessoaEncontrada.Usuario),
		fmt.Sprintf("Enviando dados de pessoa '%s'", pessoaEncontrada.Usuario))
}

// PessoaCreate é um handler/controller que responde a rota '[POST] /pessoas' e retorna StatusCreated(201) e os dados da pessoa criada através das informações informadas via JSON(body) caso o TOKEN informado for válido e o usuário associado ao token for cadastrado na API/DB. Caso ocorra algum erro, retorna StatusInternalServerError(500) ou StatusUnprocessableEntity(422) caso as informações no JSON não corresponderem ao formato {"cpf":"?",  "nome_completo":"?", "usuario":"?", "senha":"?", "email":"?"[, "administrador": ?]}
func PessoaCreate(w http.ResponseWriter, r *http.Request) {
	db02 := dao.GetDB02()
	defer dao.CloseDB(db02)

	var status = http.StatusInternalServerError
	var pessoaFromJSON pessoa.Pessoa
	var novaPessoa *pessoa.Pessoa

	token, err := helper.GetToken(r, GetMySigningKey())
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioToken, _, admin, err := helper.GetClaims(token)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioDB, err := dao.ProcuraPessoaPorUsuario02(db02, usuarioToken)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	verif := !(admin && usuarioDB.Administrador)
	err = DefineHeaderRetorno(w, SetHeaderJSON, verif, status, errors.New("somente administradores podem usar essa rota"))
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, LimitData))
	if err != nil {
		log.Println(err)
	}

	if err := r.Body.Close(); err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(body, &pessoaFromJSON)
	status = http.StatusUnprocessableEntity // 422
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	novaPessoa, err = pessoa.NewPessoa(
		pessoaFromJSON.Cpf,
		pessoaFromJSON.NomeCompleto,
		pessoaFromJSON.Usuario,
		pessoaFromJSON.Senha,
		pessoaFromJSON.Email)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}
	adicionaPessoa := dao.AdicionaPessoa02
	if pessoaFromJSON.Administrador {
		adicionaPessoa = dao.AdicionaPessoaAdmin02
	}

	p, err := adicionaPessoa(db02, novaPessoa)
	status = http.StatusInternalServerError // 500
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	status = http.StatusCreated // 201
	funcao := "PessoaCreate"
	DefineHeaderRetornoDado(
		w,
		SetHeaderJSON,
		status,
		p,
		funcao,
		fmt.Sprintf("Dados de pessoa '%s'", p.Usuario),
		fmt.Sprintf("Enviando dados de pessoa '%s'", p.Usuario))
}

// PessoaRemove é um handler/controller que responde a rota '[DELETE] /pessoas/{usuario}' e retorna StatusOK(200) e uma mensagem de confirmação caso o TOKEN informado for válido, o usuário associado ao token for cadastrado na API/DB e seja um administrador, que o usuário informado na rota seja diferente ao do token e seja cadastrado no BD. Caso ocorra algum erro, retorna StatusInternalServerError(500) ou StatusNotFound(404) caso não encontre o registro para remoção
func PessoaRemove(w http.ResponseWriter, r *http.Request) {
	db02 := dao.GetDB02()
	defer dao.CloseDB(db02)

	var status = http.StatusInternalServerError // 500

	vars := mux.Vars(r)
	usuarioRemocao := vars["usuario"]

	token, err := helper.GetToken(r, GetMySigningKey())
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioToken, _, admin, err := helper.GetClaims(token)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioDB, err := dao.ProcuraPessoaPorUsuario02(db02, usuarioToken)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	verif := !(admin && usuarioDB.Administrador)
	err = DefineHeaderRetorno(w, SetHeaderJSON, verif, status, errors.New("somente administradores podem usar essa rota"))
	if err != nil {
		return
	}

	verif = usuarioToken == usuarioRemocao
	err = DefineHeaderRetorno(w, SetHeaderJSON, verif, status, fmt.Errorf("o usuário %s não pode remover a si mesmo", usuarioToken))
	if err != nil {
		return
	}

	status = http.StatusNotFound // 404
	err = dao.RemovePessoaPorUsuario02(db02, usuarioRemocao)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	status = http.StatusOK // 200
	funcao := "PessoaRemove"
	DefineHeaderRetornoDado(
		w,
		SetHeaderJSON,
		status,
		usuarioRemocao,
		funcao,
		fmt.Sprintf("Removido pessoa '%s'", usuarioRemocao),
		fmt.Sprintf("Enviando resposta de remoção de pessoa '%s'", usuarioRemocao))

}

// PessoaAlter é um handler/controller que responde a rota '[PUT] /pessoas/{usuario}' e retorna StatusOK(200) e uma mensagem de confirmação com os dados da pessoa alterada caso o TOKEN informado for válido, o usuário associado ao token for cadastrado na API/DB e o usuário informado na rota existir. Somente usuários administradores podem alterar qualquer usuário. Um usuário comum somente pode alterar a si mesmo. Caso ocorra algum erro, retorna StatusInternalServerError(500) ou StatusUnprocessableEntity(422), caso o JSON não seguir o formato {"cpf":"?",  "nome_completo":"?", "usuario":"?", "senha":"?", "email":"?"} ou StatusNotModified(304) caso ocorra algum erro na alteração do BD
func PessoaAlter(w http.ResponseWriter, r *http.Request) {
	db02 := dao.GetDB02()
	defer dao.CloseDB(db02)

	var status = http.StatusInternalServerError // 500
	var pessoaFromJSON pessoa.Pessoa

	vars := mux.Vars(r)
	usuarioAlteracao := vars["usuario"]

	token, err := helper.GetToken(r, GetMySigningKey())
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioToken, _, admin, err := helper.GetClaims(token)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioDB, err := dao.ProcuraPessoaPorUsuario02(db02, usuarioToken)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	verif := !(admin && usuarioDB.Administrador || usuarioToken == usuarioAlteracao)
	err = DefineHeaderRetorno(w, SetHeaderJSON, verif, status, errors.New("somente administradores ou o próprio usuário pode alterar seus dados"))
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, LimitData))
	if err != nil {
		log.Println(err)
	}

	if err := r.Body.Close(); err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(body, &pessoaFromJSON)
	status = http.StatusUnprocessableEntity // 422
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioParaAlterar, err := dao.ProcuraPessoaPorUsuario02(db02, usuarioAlteracao)
	status = http.StatusInternalServerError // 500
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	status = http.StatusUnprocessableEntity // 422
	verificacaoCPF := usuarioParaAlterar.Cpf != pessoaFromJSON.Cpf
	err = DefineHeaderRetorno(w, SetHeaderJSON, verificacaoCPF, status, errors.New("CPF informado em JSON diferente do cadastrado no Banco de dados"))
	if err != nil {
		return
	}

	p, err := dao.AlteraPessoaPorUsuario02(
		db02,
		usuarioAlteracao,
		&pessoaFromJSON)
	status = http.StatusNotModified // 304
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	status = http.StatusOK // 200
	funcao := "PessoaAlter"
	DefineHeaderRetornoDado(
		w,
		SetHeaderJSON,
		status,
		p,
		funcao,
		fmt.Sprintf("Novos dados de pessoa '%s'", p.Usuario),
		fmt.Sprintf("Enviando novos dados de pessoa '%s'", p.Usuario))
}

// PessoaEstado é um handler/controller que responde a rota '[PUT] /pessoas/{usuario}/estado' e retorna StatusOK(200) e uma mensagem de confirmação com os dados da pessoa alterada caso o TOKEN informado for válido, o usuário associado ao token for cadastrado na API/DB e o usuário informado na rota existir. Somente usuários administradores podem alterar o estado de usuários, mas não pode alterar o próprio estado. Caso ocorra algum erro, retorna StatusInternalServerError(500), StatusUnprocessableEntity(422), caso o JSON não seguir o formato {"estado": ?}, StatusNotModified(304) caso ocorra algum erro na alteração do BD ou StatusNotFound(404) caso o usuário informado na rota não existir
func PessoaEstado(w http.ResponseWriter, r *http.Request) {
	db02 := dao.GetDB02()
	defer dao.CloseDB(db02)

	var status = http.StatusInternalServerError // 500
	var estadoPessoa estado

	vars := mux.Vars(r)
	usuarioAlteracao := vars["usuario"]

	token, err := helper.GetToken(r, GetMySigningKey())
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioToken, _, admin, err := helper.GetClaims(token)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioDB, err := dao.ProcuraPessoaPorUsuario02(db02, usuarioToken)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	verif := !(admin && usuarioDB.Administrador && usuarioToken != usuarioAlteracao)
	err = DefineHeaderRetorno(w, SetHeaderJSON, verif, status, errors.New("somente administradores podem alterar o estado de pessoas que sejam diferentes do próprio administrador"))
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, LimitData))
	if err != nil {
		log.Println(err)
	}

	if err := r.Body.Close(); err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(body, &estadoPessoa)
	status = http.StatusUnprocessableEntity // 422
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioDBAlteracao, err := dao.ProcuraPessoaPorUsuario02(db02, usuarioAlteracao)
	status = http.StatusNotFound // 404
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	var alteraEstado func(*gorm.DB, string) (*pessoa.Pessoa, error)
	if estadoPessoa.Estado {
		alteraEstado = dao.AtivaPessoa02
	} else {
		alteraEstado = dao.InativaPessoa02
	}
	p, err := alteraEstado(db02, usuarioDBAlteracao.Cpf)
	status = http.StatusNotModified // 304
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	status = http.StatusOK // 200
	funcao := "PessoaEstado"
	DefineHeaderRetornoDado(
		w,
		SetHeaderJSON,
		status,
		p,
		funcao,
		fmt.Sprintf("Novos dados de pessoa '%s'", p.Usuario),
		fmt.Sprintf("Enviando novos dados de pessoa '%s'", p.Usuario))
}

// PessoaAdmin é um handler/controller que responde a rota '[PUT] /pessoas/{usuario}/admin' e retorna StatusOK(200) e uma mensagem de confirmação com os dados da pessoa alterada caso o TOKEN informado for válido, o usuário associado ao token for cadastrado na API/DB e o usuário informado na rota existir. Somente usuários administradores podem redefinir usuários como administrador, mas não pode alterar a sí mesmo. Caso ocorra algum erro, retorna StatusInternalServerError(500), StatusUnprocessableEntity(422), caso o JSON não seguir o formato {"adminstrador": ?}, StatusNotModified(304) caso ocorra algum erro na alteração do BD ou StatusNotFound(404) caso o usuário informado na rota não existir
func PessoaAdmin(w http.ResponseWriter, r *http.Request) {
	db02 := dao.GetDB02()
	defer dao.CloseDB(db02)

	var status = http.StatusInternalServerError // 500
	type administrador struct {
		Administrador bool `json:"administrador"`
	}
	var adminPessoa administrador

	vars := mux.Vars(r)
	usuarioAlteracao := vars["usuario"]

	token, err := helper.GetToken(r, GetMySigningKey())
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioToken, _, admin, err := helper.GetClaims(token)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioDB, err := dao.ProcuraPessoaPorUsuario02(db02, usuarioToken)
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	verif := !(admin && usuarioDB.Administrador && usuarioToken != usuarioAlteracao)
	err = DefineHeaderRetorno(w, SetHeaderJSON, verif, status, errors.New("somente administradores podem redefinir como administrador pessoas que sejam diferentes do próprio administrador"))
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, LimitData))
	if err != nil {
		log.Println(err)
	}

	if err := r.Body.Close(); err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(body, &adminPessoa)
	status = http.StatusUnprocessableEntity // 422
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	usuarioDBAlteracao, err := dao.ProcuraPessoaPorUsuario02(db02, usuarioAlteracao)
	status = http.StatusNotFound // 404
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	p, err := dao.SetAdministrador02(db02, usuarioDBAlteracao.Cpf, adminPessoa.Administrador)
	status = http.StatusNotModified // 304
	err = DefineHeaderRetorno(w, SetHeaderJSON, err != nil, status, err)
	if err != nil {
		return
	}

	status = http.StatusOK // 200
	funcao := "PessoaAdmin"
	DefineHeaderRetornoDado(
		w,
		SetHeaderJSON,
		status,
		p,
		funcao,
		fmt.Sprintf("Novos dados de pessoa '%s'", p.Usuario),
		fmt.Sprintf("Enviando novos dados de pessoa '%s'", p.Usuario))
}
