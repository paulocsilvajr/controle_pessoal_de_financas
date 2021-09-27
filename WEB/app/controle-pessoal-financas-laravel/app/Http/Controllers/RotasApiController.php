<?php

namespace App\Http\Controllers;

use App\Services\RequisicaoHttp;
use App\Services\Token;
use Illuminate\Http\Request;

class RotasApiController extends Controller
{
    public function index(RequisicaoHttp $requisicao, Token $token)
    {
        if ($token->valido()) {
            $resposta = $requisicao->get();
            $dados = $resposta['data'];
            $quant = $resposta['count'];
            $nomes = array_keys($dados);

            if ($resposta->successful()) {
                return view('RotasApi.rotasApi', [
                    'nomes' => $nomes,
                    'dados' => $dados,
                    'quant' => $quant
                ]);
            }
            return redirect()->route('home');
        } else {
            return redirect()->route('login');
        }
    }
}
