<?php

namespace App\Models;

use DateTime;

class Lancamento {

    private int $id;
    private string $cpf;
    private string $nomeContaOrigem;
    private DateTime $data;
    private string $numero;
    private string $descricao;
    private string $nomeContaDestino;
    private float $debito;
    private float $credito;

    public function __construct()
    {
        # code...
    }
}
