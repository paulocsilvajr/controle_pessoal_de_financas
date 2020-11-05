<?php

namespace App\Http\Controllers;

use App\Services\RequisicaoHttp;
use App\Services\Token;
use Illuminate\Http\Request;

class ContaController extends Controller
{
    public function index(Request $request, RequisicaoHttp $http, Token $token)
    {
        if ($token->valido()) {
            $resposta = $http->get('/contas');

            if ($resposta->successful()) {
                $dados = $resposta['data'];

                return view(
                    'Conta.conta',
                    compact(
                        'dados',
                    )
                );
            }

            return redirect()->route('home');
        } else {
            return redirect()->route('login');
        }
    }

    public function contaEspecifica(Request $request, RequisicaoHttp $http, Token $token, string $nomeConta)
    {
        if ($token->valido()) {
            return view(
                'Conta.contaEspecifica',
                compact(
                    'nomeConta'
                )
            );
            return redirect()->route('home');
        } else {
            return redirect()->route('login');
        }
    }
}
