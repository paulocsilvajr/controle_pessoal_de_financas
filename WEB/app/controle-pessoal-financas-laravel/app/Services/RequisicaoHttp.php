<?php

namespace App\Services;

use Illuminate\Http\Client\Response;
use Illuminate\Support\Facades\Http;

class RequisicaoHttp
{

    public $verificarCertificadoSSL = false;
    private $requisicao;
    private $rotaBase;

    public function __construct()
    {
        $this->requisicao = Http::withOptions([
            'verify' => $this->verificarCertificadoSSL
        ]);
        $this->rotaBase = "https://localhost:8085";
    }

    public function setRotaBase(string $rota)
    {
        if (!empty($rota)) {
            $this->rotaBase = $rota;
        }

        return $this;
    }

    public function getRotaBase(): string
    {
        return $this->rotaBase;
    }

    public function post(string $rota, array $body): Response
    {
        $headers = [
            'Content-Type' => 'application:json'
        ];
        return $this->requisicao
            ->withHeaders($headers)
            ->post($this->rotaBase . $rota, $body);
    }
}
