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

                $nomesCompletos = $this->geraListaContasCompleto($dados);

                $request->session()->put('contas', $nomesCompletos);

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

            $usuario = $request->session()->get('usuario');
            $resposta = $http->get("/pessoas/$usuario");

            if ($resposta->successful()) {
                $cpf = $resposta['data']['cpf'];
                $contas = $request->session()->get('contas');

                return view(
                    'Conta.cadastroLancamento',
                    compact(
                        'nomeConta',
                        'mensagem',
                        'tipoMensagem',
                        'cpf',
                        'contas',
                    )
                );
            }
            return redirect()->route('home');
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

    private function geraListaContasCompleto(array $contas): array{
        $lista = array();

        foreach ($contas as $conta) {
            if (empty($conta['conta_pai'])) {
                $nomeCompleto = $conta['nome'];
                $lista[$conta['nome']] = $conta['nome'];
            } else {
                $nomeCompleto = $conta['conta_pai'] . '>' . $conta['nome'];

                // foreach ($contas as $conta2) {
                //     if (empty($conta2['conta_pai'])) {
                //         continue;
                //     } else if ($conta2['nome'] == $conta['conta_pai']) {
                //         $nomeCompleto = $conta2['conta_pai'] . '>' . $nomeCompleto;
                //     }
                // }
                $this->geraNomeCompletoR($contas, $conta['conta_pai'], $nomeCompleto);

                $lista[$conta['nome']] = $nomeCompleto;
            }
        }
        return $lista;
    }

    private function geraNomeCompletoR(array $contas, string $contaPaiAnterior, string &$nomeCompleto) {
        foreach ($contas as $conta) {
            if (empty($conta['conta_pai'])) {
                continue;
            } else if ($conta['nome'] == $contaPaiAnterior) {
                $nomeCompleto = $conta['conta_pai'] . '>' . $nomeCompleto;

                $this->geraNomeCompletoR($contas, $conta['conta_pai'], $nomeCompleto);
            }
        }
    }
}
