<?php

namespace App\Models;

use App\Helpers\Formata;
use DateTime;

class Conta {

    public string $nome;
    public string $nomeTipoConta;
    public string $codigo;
    public string $contaPai;
    public string $comentario;
    public Datetime $dataCriacao;
    public Datetime $dataModificacao;
    public bool $estado;

    public function __construct($dados)
    {
        $this->fromJSON($dados);
    }

    public function fromJSON($dados) {
        $this->nome = $dados['nome'];
        $this->nomeTipoConta = $dados['nome_tipo_conta'];
        $this->codigo = $dados['codigo'];
        $this->contaPai = $dados['conta_pai'];
        $this->comentario = $dados['comentario'];
        $this->dataCriacao = Formata::textoParaDatetime($dados['data_criacao']);
        $this->dataModificacao = Formata::textoParaDatetime($dados['data_modificacao']);
        $this->estado = $dados['estado'];
    }

    public function toJSON(): string {
        $json = array(
            'nome' => $this->nome,
            'nome_tipo_conta' => $this->nomeTipoConta,
            'codigo' => $this->codigo,
            'conta_pai' => $this->contaPai,
            'comentario' => $this->comentario,
            'data_criacao' => Formata::DatetimeParaJSON($this->dataCriacao),
            'data_modificacao' => Formata::DatetimeParaJSON($this->dataModificacao),
            'estado' => $this->estado,
        );
        return json_encode($json);
    }
}
