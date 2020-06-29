<?php

namespace App\Services;

use App\Helpers\Imprime;
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
                Imprime::console(">>> TOKEN válido <<<");
                return true;
            } else if ($resposta->clientError()) {
                Imprime::console(">>> Renovando TOKEN <<<");
                $usuario = $this->request->session()->get('usuario');
                $senha = $this->request->session()->put('senha');

                $resposta = $this->renovar($http, $usuario, $senha);
                if ($resposta->successful()) {
                    $this->request->session()->put('token', $resposta['token']);
                    return true;
                }
            }
        } catch (\Throwable $th) {
            Imprime::console(">>> NÃO foi obtido resposta do servidor/API para a rota '/token' <<<");
        }

        return false;
    }

    private function renovar(RequisicaoHttp $http, string $usuario, string $senha): Response
    {
        return $http->post(
            "/login/{$usuario}",
            [
                'usuario' => $usuario,
                'senha' => $senha
            ]
        );
    }
}
