<?php

namespace App\Helpers;

use DateTime;
use UnexpectedValueException;

final class Formata
{
    /**
     * valorParaMonetarioBrasil converte um número(float ou integer) em uma string no formato monetário brasileiro(vírgula separando decimais) SEM o símbolo de reais(R$)
     */
    public static function valorParaMonetarioBrasil(float $valor): string
    {
        return number_format($valor, 2, ',', '.');
    }

    /**
     * textoParaDataBrasil formata a data no formato texto YYYY-MM-DDTHH:mm:SSZ obtido via JSON da API para o formato brasileiro DD/MM/YYYY
     */
    public static function textoParaDataBrasil(string $dataCompleta): string
    {
        $data = explode('T', $dataCompleta); // separado a data
        $data = explode('-', $data[0]);
        $dia = intval($data[2]);
        $mes = intval($data[1]);
        $ano = intval($data[0]);
        return sprintf('%02d/%02d/%04d', $dia, $mes, $ano);
    }

    public static function textoParaDatetime(string $data): DateTime
    {
        $data = substr($data, 0, 19);
        $data = str_replace("T", " ", $data);
        $dataConvertida = DateTime::createFromFormat('Y-m-d H:i:s', $data);
        if (!$dataConvertida) {
            throw new UnexpectedValueException("Não foi possível converter a data informada[$data]");
        }
        return $dataConvertida;
    }

    public static function DatetimeParaJSON(Datetime $data): string {
        return $data->format('Y-m-d\TH:i:s.000\Z');
    }
}
