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

                // if (isset($dados)) {
                //     Imprime::console2($dados);
                // }

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
                $contas = $this->filtraContas($contas, $nomeConta);

                $destino = '';
                $tipo = 'debito';

                return view(
                    'Conta.cadastroLancamento',
                    compact(
                        'nomeConta',
                        'mensagem',
                        'tipoMensagem',
                        'cpf',
                        'contas',
                        'destino',
                        'tipo'
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

        $dataFormatada = $this->formataData($data);

        $debito = 0.0;
        $credito = 0.0;
        if ($tipo == 'debito') {
            $debito = floatval($valor);
        } else {
            $credito = floatval($valor);
        }

        $resposta = $http->post(
            "/lancamentos",
            [
                'cpf_pessoa' => $cpf,
                'nome_conta_origem' => $nome_conta_origem,
                'data' => $dataFormatada,
                'numero' => $numero,
                'descricao' => $descricao,
                'nome_conta_destino' => $nome_conta_destino,
                'debito' => $debito,
                'credito' => $credito
            ]
        );

        if ($resposta->successful()) {
            $mensagem = "Lançamento '$descricao' cadastrado com sucesso";
            $tipoMensagem = 'success';

            return redirect()->route(
                'contaEspecifica',
                [
                    'nomeConta' => $nome_conta_origem,
                    'mensagem' => $mensagem,
                    'tipoMensagem' => $tipoMensagem
                ]
            );
        }

        $mensagem = "$resposta";
        $tipoMensagem = 'danger';

        $nomeConta = $nome_conta_origem;
        $contas = $request->session()->get('contas');
        $contas = $this->filtraContas($contas, $nomeConta);

        $destino =  $nome_conta_destino;

        return view(
            'Conta.cadastroLancamento',
            compact(
                'nomeConta',
                'mensagem',
                'tipoMensagem',
                'cpf',
                'contas',
                'cpf',
                'data',
                'numero',
                'descricao',
                'destino',
                'valor',
                'tipo'
            )
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

    /**
     * filtraContas remove do array(associativo) $contas a conta com a chave informada na string $nomeConta
     */
    private function filtraContas(array $contas, string $nomeConta): array {
        return array_filter($contas, function($conta) use(&$nomeConta){
            return $conta !== $nomeConta;
        }, ARRAY_FILTER_USE_KEY);
    }

    private function formataData(string $data): string {
        return "${data}T00:00:00Z";
    }
}
