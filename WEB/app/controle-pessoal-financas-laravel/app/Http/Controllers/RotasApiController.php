<?php

namespace App\Http\Controllers;

use App\Services\RequisicaoHttp;
use Illuminate\Http\Request;

class RotasApiController extends Controller
{
    public function index(RequisicaoHttp $requisicao)
    {
        $resposta = $requisicao->get();
        $dados = $resposta['data'];
        $nomes = array_keys($dados);

        if ($resposta->successful()) {
            return view('rotasApi', [
                'nomes' => $nomes,
                'dados' => $dados
            ]);
        }
        return redirect()->rota('home');
    }
}
