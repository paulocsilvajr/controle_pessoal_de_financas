<?php

namespace App\Models;

use App\Helpers\Formata;
use DateTime;

class Lancamento {

    public int $id;
    public string $cpfPessoa;
    public string $nomeContaOrigem;
    public DateTime $data;
    public string $numero;
    public string $descricao;
    public string $nomeContaDestino;
    public float $debito;
    public float $credito;
    public Datetime $dataCriacao;
    public Datetime $dataModificacao;
    public bool $estado;

    public function __construct($dados)
    {
        $this->fromJSON($dados);
    }

    public function fromJSON($dados) {
        $this->id = $dados['id'];
        $this->cpfPessoa = $dados['cpf_pessoa'];
        $this->nomeContaOrigem = $dados['nome_conta_origem'];
        $this->data = Formata::textoParaDatetime($dados['data']);
        $this->numero = $dados['numero'];
        $this->descricao = $dados['descricao'];
        $this->nomeContaDestino = $dados['nome_conta_destino'];
        $this->debito = $dados['debito'];
        $this->credito = $dados['credito'];
        $this->dataCriacao = Formata::textoParaDatetime($dados['data_criacao']);
        $this->dataModificacao = Formata::textoParaDatetime($dados['data_modificacao']);
        $this->estado = $dados['estado'];
    }

    public function toJSON(): string {
        $json = array(
            "cpf_pessoa" => $this->cpfPessoa,
            "nome_conta_origem" => $this->nomeContaOrigem,
            "data" => Formata::DatetimeParaJson($this->data),
            "numero" => $this->numero,
            "descricao" => $this->descricao,
            "nome_conta_destino" => $this->nomeContaDestino,
            "debito" => $this->debito,
            "credito" => $this->credito
        );
        return json_encode($json);
    }
}
