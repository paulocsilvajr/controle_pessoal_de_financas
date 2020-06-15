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

            return redirect()->route('home');
        }

        $msg = 'Usuário ou senha inválida';
        if ($resposta->serverError()) {
            $msg = "Problema interno do servidor";
        }

        $request->session()->flash('mensagem', $msg);

        return redirect()->route('login');
    }
}
