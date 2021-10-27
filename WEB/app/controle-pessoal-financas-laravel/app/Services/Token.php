<?php

namespace App\Services;

use App\Helpers\Imprime;
use App\Helpers\LogPersonalizado;
use Illuminate\Http\Client\Response;
use Illuminate\Http\Request;

final class Token
{
    public function __construct(Request $request)
    {
        $this->request = $request;
    }

    public function valido(): bool
    {
        $http = new RequisicaoHttp($this->request);

        try {
            $resposta = $http->get('/token');

            if ($resposta->successful()) {
                LogPersonalizado::info("TOKEN válido");
                return true;
            } else if ($resposta->clientError()) {
                LogPersonalizado::info("Renovando TOKEN");
                $usuario = $this->request->session()->get('usuario');
                $senha = $this->request->session()->put('senha');

                $resposta = $this->renovar($http, $usuario, $senha);
                if ($resposta->successful()) {
                    $this->request->session()->put('token', $resposta['token']);
                    return true;
                }
            }
        } catch (\Throwable $th) {
            LogPersonalizado::info("NÃO foi obtido resposta do servidor/API para a rota '/token'");
        }

        return false;
    }

    private function renovar(RequisicaoHttp $http, string $usuario, string $senha): Response
    {
        return $http->postWithoutToken(
            "/login/{$usuario}",
            [
                'usuario' => $usuario,
                'senha' => $senha
            ]
        );
    }
}
