<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class PrincipalController extends Controller
{
    public function index(Request $request)
    {
        $usuario = $request->session()->get('usuario');
        $senha = $request->session()->get('senha');
        $token = $request->session()->get('token');
        $logado = $request->session()->get('logado');

        $tokenParcial = substr($token ?? '', 0, 10) . "..." . substr($token ?? '', -10);

        // enviando e exibindo senha somente para teste
        return view('home', compact('usuario', 'senha', 'tokenParcial', 'logado'));
    }
}
