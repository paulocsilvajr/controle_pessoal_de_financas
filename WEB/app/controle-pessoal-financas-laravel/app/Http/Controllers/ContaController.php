<?php

namespace App\Http\Controllers;

use App\Services\RequisicaoHttp;
use Illuminate\Http\Request;

class ContaController extends Controller
{
    public function index(Request $request, RequisicaoHttp $http)
    {
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
    }
}
