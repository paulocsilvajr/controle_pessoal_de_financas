<?php

namespace App\Http\Controllers;

use App\Helpers\Imprime;
use App\Services\RequisicaoHttp;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;

class EntrarController extends Controller
{
    public function index(Request $request)
    {
        $this->apiOnLine($request);

        $mensagem = $request->session()->get('mensagem');
        $tipoMensagem = $request->session()->get('tipoMensagem');
        $estadoApi = $request->session()->get('estadoApi');
        return view(
            'login',
            compact(
                'mensagem',
                'tipoMensagem',
                'estadoApi',
            )
        );
    }

    public function entrar(Request $request, RequisicaoHttp $http)
    {
        $usuario = $request->usuario;
        $senha = $request->senha;

        $resposta = $http->postWithoutToken(
            "/login/{$usuario}",
            [
                'usuario' => $usuario,
                'senha' => $senha
            ]
        );

        if ($resposta->successful()) {
            $this->definirSessaoUsuario($request, $usuario, $senha, $resposta['token']);
            $this->logar($request);

            return redirect()->route('home');
        }

        $msg = 'Usuário ou senha inválida';
        if ($resposta->serverError()) {
            $msg = "Problema interno do servidor";
        }

        $this->deslogar($request);
        $this->removerSessaoUsuario($request);
        $request->session()->flash('mensagem', $msg);
        $request->session()->flash('tipoMensagem', 'danger');

        return redirect()->route('login');
    }

    public function sair(Request $request)
    {
        $this->deslogar($request);
        return redirect()->route('login');
    }

    private function apiOnLine(Request $request)
    {
        $http = new RequisicaoHttp($request);
        try {
            $resposta = $http->getWithoutToken('/API');

            if ($resposta->successful()) {
                $request->session()->put('estadoApi', true);
                Imprime::console(">>> API ONLINE <<<");
                return;
            }
            $request->session()->put('estadoApi', false);
            Imprime::console(">>> API OFFLINE <<<");
        } catch (\Throwable $th) {
            $request->session()->put('estadoApi', false);
            Imprime::console(">>> API OFFLINE <<<");
        }
    }

    private function logar(Request $request)
    {
        $this->definirChaveLogadoSessao($request, true);
    }

    private function deslogar(Request $request)
    {
        $this->definirChaveLogadoSessao($request, false);
    }

    private function definirChaveLogadoSessao(Request $request, bool $logado)
    {
        $request->session()->put('logado', $logado);
    }

    private function definirSessaoUsuario(Request $request, string $usuario, string $senha, string $token)
    {
        $request->session()->put('usuario', $usuario);
        $request->session()->put('senha', $senha);
        $request->session()->put('token', $token);
    }

    private function removerSessaoUsuario(Request $request)
    {
        $request->session()->remove('usuario');
        $request->session()->remove('senha');
        $request->session()->remove('token');
    }
}
