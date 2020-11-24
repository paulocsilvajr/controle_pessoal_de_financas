<?php

namespace App\Http\Controllers;

use App\Helpers\Imprime;
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
            $resposta = $http->get("/lancamentos_conta/$nomeConta");

            if ($resposta->successful()) {
                if ($resposta['count'] == 0) {
                    Imprime::console(">>> Sem registro de Lançamentos para a conta '$nomeConta' <<<");
                }

                $dados = $resposta['data'];

                return view(
                    'Conta.contaEspecifica',
                    compact(
                        'nomeConta',
                        'dados',
                    )
                );
            }

            return redirect()->route('home');
        } else {
            return redirect()->route('login');
        }
    }

    public function carregaCadastroLancamento(Request $request, RequisicaoHttp $http, Token $token, string $nomeConta) {
        if ($token->valido()) {
            $mensagem = '';
            $tipoMensagem = '';

            return view(
                'Conta.cadastroLancamento',
                compact(
                    'nomeConta',
                    'mensagem',
                    'tipoMensagem',
                )
            );
        } else {
            return redirect()->route('login');
        }
    }

    public function cadastraLancamento(Request $request, RequisicaoHttp $http, Token $token) {
        $cpf = $request->cpf_pessoa;
        $nome_conta_origem = $request->nome_conta_origem;
        $data = $request->data;
        $numero = $request->numero;
        $descricao = $request->descricao;
        $nome_conta_destino = $request->nome_conta_destino;
        $valor = $request->valor;
        $tipo = $request->tipo;

        Imprime::console("<<<
CPF: $cpf
Origem: $nome_conta_origem
Data: $data
Número: $numero
Descrição: $descricao
Destino: $nome_conta_destino
Valor: $valor
Tipo: $tipo
>>>");

        return redirect()->route(
            'contaEspecifica',
            ['nomeConta' => $nome_conta_origem]
        );
    }
}
