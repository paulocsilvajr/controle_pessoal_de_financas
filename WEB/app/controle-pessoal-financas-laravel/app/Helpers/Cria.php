<?php

namespace App\Helpers;

use App\Models\Conta;

final class Cria
{
    public static function arrayContas(array $dados): array
    {
        return Cria::arrayTipado($dados, Conta::class);
    }

    private static function arrayTipado($dados, $construtor): array
    {
        $lista = array();
        foreach ($dados as $dado) {
            $lista[] = new $construtor($dado);
        }
        return $lista;
    }
}
