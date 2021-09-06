package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/helper"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/logger"
)

// https://play.golang.org/p/6dX5SMdVtr

const diretorioLog = "config"
const arquivoLog = "config.json"

type rota struct {
	Tipo, Rota, Descricao, Retorno, Documentacao string
}

type rotas map[string]rota

func (r rotas) Len() int {
	return len(r)
}

func (r rotas) DefineDocumentacao(rota, doc string) {
	var temp = r[rota]
	temp.Documentacao = doc
	r[rota] = temp
}

// Rotas é uma variável que guarda todas as rotas definidas da API. Usa o tipo interno rota{Tipo, Rota, Descricao} em cada elemento desse hashMap rotas map[string]rota
var Rotas = rotas{
	"API": rota{
		"GET",
		"/API",
		"",
		`{"status":int, "message":"string", "count":int, "data": {"API OnLine":"string"}}`,
		"",
	},
	"Index": rota{
		"GET",
		"/",
		"",
		`{"status":int, "message":"string", "count":int, "data":{"CONTROLLER_1": {"Tipo":"VERBO_HTML|string", "Rota":"/ROTA", "Descricao":"string", "Documentacao":"string"}, "CONTROLLER_2": {...}, ... } }`,
		"",
	},
	"Login": rota{
		"POST",
		"/login/{usuario}",
		`Body: {"usuario":"string", "senha":"string"}`,
		`{"status":int, "message":"string", "token":"string"}`,
		"",
	},
	"TokenValido": rota{
		"GET",
		"/token",
		"",
		`{"status":int, "message":"string"}`,
		"",
	},
	"PessoaIndex": rota{
		"GET",
		"/pessoas",
		"",
		`{"status":int, "message":"string", "count":int, "data": [ {"cpf":"string", "nome_completo":"string", "usuario":"string", "senha":"string", "email":"string", "data_criacao":"YYYY-MM-DDTHH:mm:SS.nnn...-nn:nn", "data_modificacao":"YYYY-MM-DDTHH:mm:SS.nnn...-nn:nn", "estado":boolean, "administrador":boolean}, {"cpf":"string", ... }, ... ] } ou {"status":int, "message":"string","count":int, "data": [ {"usuario":"string", "email":"", "data_criacao":"YYYY-MM-DDTHH:mm:SS.nnn...-nn:nn", "data_modificacao":"YYYY-MM-DDTHH:mm:SS.nnn...-nn:nn"}, {"usuario": ... } ] }`,
		"",
	},
	"PessoaShow": rota{
		"GET",
		"/pessoas/{usuario}",
		"",
		`{"status":int, "message":"string", "count":int, "data": {"cpf":"string", "nome_completo":"string", "usuario":"string", "senha":"string", "email":"string", "data_criacao":"YYYY-MM-DDTHH:mm:SS.nnn...-nn:nn", "data_modificacao":"YYYY-MM-DDTHH:mm:SS.nnn...-nn:nn", "estado":boolean, "administrador":boolean} }`,
		"",
	},
	"PessoaShowAdmin": rota{
		"GET",
		"/pessoas/{usuarioAdmin}/{usuario}",
		"",
		`{"status":int, "message":"string", "count":int, "data": {"cpf":"string", "nome_completo":"string", "usuario":"string", "senha":"string", "email":"string", "data_criacao":"YYYY-MM-DDTHH:mm:SS.nnn...-nn:nn", "data_modificacao":"YYYY-MM-DDTHH:mm:SS.nnn...-nn:nn", "estado":boolean, "administrador":boolean} }`,
		"",
	},
	"PessoaCreate": rota{
		"POST",
		"/pessoas",
		`Body: {"cpf":"string",  "nome_completo":"string", "usuario":"string", "senha":"string", "email":"string"[, "administrador": boolean]}`,
		`{"status":int, "message":"string", "count":int, "data": {"cpf":"string", "nome_completo":"string", "usuario":"string", "senha":"string", "email":"string", "data_criacao":"YYYY-MM-DDTHH:mm:SS.nnn...-nn:nn", "data_modificacao":"YYYY-MM-DDTHH:mm:SS.nnn...-nn:nn", "estado":boolean, "administrador":boolean}}`,
		"",
	},
	"PessoaRemove": rota{
		"DELETE",
		"/pessoas/{usuario}",
		"",
		`{"status":int, "message":"string", "count":int, "data":"string"}`,
		"",
	},
	"PessoaAlter": rota{
		"PUT",
		"/pessoas/{usuario}",
		`Body: {"cpf":"string",  "nome_completo":"string", "usuario":"string", "senha":"string", "email":"string"}`,
		`{"status":int, "message":"string", "count":int, "data": {"cpf":"string", " nome_completo":"string", "usuario":"string", "senha":"string", "email":"string", "data_criacao":"yyyy-mm-ddthh:mm:ss.nnn...-nn:nn", "data_modificacao":"YYYY-MM-DDTHH:mm:SS.nnn...-nn:nn", "estado":boolean, "administrador":boolean} }`,
		"",
	},
	"PessoaEstado": rota{
		"PUT",
		"/pessoas/{usuario}/estado",
		`Body: {"estado": boolean}`,
		`{"status":int, "message":"string", "count":int, "data": {"cpf":"string", "nome_completo":"string", "usuario":"string", "senha":"string", "email":"string", "data_criacao":"2021-03-02T09:41:12.800977-03:00", "data_modificacao":"2021-03-03T09:34:51.286881051-03:00", "estado":boolean, "administrador":boolean} }`,
		"",
	},
	"PessoaAdmin": rota{
		"PUT",
		"/pessoas/{usuario}/admin",
		`Body: {"administrador": boolean}`,
		`{"status":int, "message":"string", "count":int, "data": {"cpf":"string", "nome_completo":"string", "usuario":"string", "senha":"string", "email":"string", "data_criacao":"YYYY-MM-DDTHH:mm:SS.nnn...-nn:nn", "data_modificacao":"YYYY-MM-DDTHH:mm:SS.nnn...-nn:nn", "estado":boolean, "administrador":boolean} }`,
		"",
	},
	"TipoContaIndex": rota{
		"GET",
		"/tipos_conta",
		"",
		"",
		"",
	},
	"TipoContaShow": rota{
		"GET",
		"/tipos_conta/{tipoConta}",
		"",
		"",
		"",
	},
	"TipoContaCreate": rota{
		"POST",
		"/tipos_conta",
		`Body: {"nome":"?",  "descricao_debito":"?", "descricao_credito":"?"}`,
		"",
		"",
	},
	"TipoContaRemove": rota{
		"DELETE",
		"/tipos_conta/{tipoConta}",
		"",
		"",
		"",
	},
	"TipoContaAlter": rota{
		"PUT",
		"/tipos_conta/{tipoConta}",
		`Body: {["nome":"?",]  "descricao_debito":"?", "descricao_credito":"?"}`,
		"",
		"",
	},
	"TipoContaEstado": rota{
		"PUT",
		"/tipos_conta/{tipoConta}/estado",
		`Body: {"estado": ?}`,
		"",
		"",
	},
	"ContaIndex": rota{
		"GET",
		"/contas",
		"",
		"",
		"",
	},
	"ContaShow": rota{
		"GET",
		"/contas/{conta}",
		"",
		"",
		"",
	},
	"ContaCreate": rota{
		"POST",
		"/contas",
		`Body: {"nome":"?",  "nome_tipo_conta":"?", "codigo":"?", "conta_pai":"?", "comentario":"?"}`,
		"",
		"",
	},
	"ContaRemove": rota{
		"DELETE",
		"/contas/{conta}",
		"",
		"",
		"",
	},
	"ContaAlter": rota{
		"PUT",
		"/contas/{conta}",
		`Body: {["nome":"?",]  "nome_tipo_conta":"?", "codigo":"?", "conta_pai":"?", "comentario":"?"}`,
		"",
		"",
	},
	"ContaEstado": rota{
		"PUT",
		"/contas/{conta}/estado",
		`Body: {"estado": ?}`,
		"",
		"",
	},
	"LancamentoIndex": rota{
		"GET",
		"/lancamentos",
		"",
		"",
		"",
	},
	"LancamentoShow": rota{
		"GET",
		"/lancamentos/{lancamento}",
		"",
		"",
		"",
	},
	"LancamentoCreate": rota{
		"POST",
		"/lancamentos",
		`Body: {"cpf_pessoa":"?", "nome_conta_origem":"?", "data":"?", "numero":"?", "descricao":"?", "nome_conta_destino":"?", "debito":?, "credito":?}`,
		"",
		"",
	},
	"LancamentoRemove": rota{
		"DELETE",
		"/lancamentos/{lancamento}",
		"",
		"",
		"",
	},
	"LancamentoAlter": rota{
		"PUT",
		"/lancamentos/{lancamento}/{origen}/{destino}",
		`Body: {["cpf_pessoa":"?",] "nome_conta_origem":"?", "data":"?", "numero":"?", "descricao":"?", "nome_conta_destino":"?", "debito":?, "credito":?}`,
		"",
		"",
	},
	"LancamentoEstado": rota{
		"PUT",
		"/lancamentos/{lancamento}/estado",
		`Body: {"estado": ?}`,
		"",
		"",
	},
	"LancamentoPorConta": rota{
		"GET",
		"/lancamentos_conta/{conta}",
		"",
		"",
		"",
	},
}

// Configuracoes é a representação de um hashMap das configurações da API
type Configuracoes map[string]string

// AbrirConfiguracoes carrega as configurações no retorno da função do tipo Configuracoes. Se o arquivo não existir, cria o arquivo de configuração com os dados padrões definidos na função interna criarConfigPadrao
func AbrirConfiguracoes() Configuracoes {
	decodeFile, err := os.Open(getArquivoLog())
	if err != nil {
		err = criarConfigPadrao()
		if err != nil {
			panic(fmt.Errorf("Erro(1) ao abrir arquivo de configurações[%s]", err))

		}

		decodeFile, err = os.Open(getArquivoLog())
		if err != nil {
			panic(fmt.Errorf("Erro(2) ao abrir arquivo de configurações[%s]", err))
		}
	}
	defer decodeFile.Close()

	decoder := json.NewDecoder(decodeFile)

	configuracoes := make(Configuracoes)

	decoder.Decode(&configuracoes)

	return configuracoes
}

// AbrirConfiguracoesParaTestes retorna um map com as configurações de teste, mas adiciona ao nome do Banco o sufixo "_teste", para ser usado exclusivamente nos testes Test...
func AbrirConfiguracoesParaTestes() Configuracoes {
	config := AbrirConfiguracoes()
	sufixo := "_teste"

	config["DBnome"] = fmt.Sprintf("%s%s", config["DBnome"], sufixo)

	return config
}

func getArquivoLog() string {
	dirBase, _ := helper.GetDiretorioAbs()
	dirBaseLog := path.Join(dirBase, diretorioLog)
	helper.CriarDiretorioAbs(diretorioLog)

	return fmt.Sprintf("%s/%s", dirBaseLog, arquivoLog)
}

func criarConfigPadrao() error {
	arq := getArquivoLog()

	configuracoes := make(Configuracoes)
	configuracoes["porta"] = "8085"
	configuracoes["host"] = "localhost"
	configuracoes["duracao_token"] = "3600"
	configuracoes["protocolo"] = "https"
	configuracoes["DB"] = "postgres"
	configuracoes["DBhost"] = "localhost"
	configuracoes["DBporta"] = "15432"
	configuracoes["DBnome"] = "controle_pessoal_financas_2"
	configuracoes["DBusuario"] = "postgres"
	configuracoes["DBsenha"] = "Postgres2019!"

	encodeFile, err := os.Create(arq)
	defer encodeFile.Close()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(encodeFile)

	if err := encoder.Encode(configuracoes); err != nil {
		return err
	}

	logger.GeraLogFS(fmt.Sprintf("Criado arquivo de config '%s'", arq), time.Now())

	return nil
}
