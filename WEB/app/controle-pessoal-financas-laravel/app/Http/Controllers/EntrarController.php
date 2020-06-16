<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;

class EntrarController extends Controller
{
    public function index(Request $request)
    {
        $mensagem = $request->session()->get('mensagem');
        return view('login', compact('mensagem'));
    }

    public function entrar(Request $request)
    {
        $usuario = $request->usuario;
        $senha = $request->senha;

        $verificarCertificadoSSL = false;
        $resposta = Http::withOptions([
            'verify' => $verificarCertificadoSSL
        ])->withHeaders([
            'Content-Type' => 'application:json'
        ])->post("https://localhost:8085/login/{$usuario}", [
            'usuario' => $usuario,
            'senha' => $senha
        ]);

        if ($resposta->successful()) {
            $request->session()->put('usuario', $usuario);
            $request->session()->put('senha', $senha);
            $request->session()->put('token', $resposta['token']);
            $this->logar($request);

            return redirect()->route('home');
        }

        $msg = 'Usuário ou senha inválida';
        if ($resposta->serverError()) {
            $msg = "Problema interno do servidor";
        }

        $this->deslogar($request);
        $request->session()->flash('mensagem', $msg);

        return redirect()->route('login');
    }

    public function sair(Request $request)
    {
        $this->deslogar($request);
        return redirect()->route('login');
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
}
