@extends('layout', ['css' => 'login'])

@section('conteudo')

    <form class="formulario-login text-center" method="post" action="/login">
        @csrf
        <h1 class="titulo-login">Controle pessoal de Finanças</h1>

        <h2 class="mb-3">Informe seu usuário e senha</h2>

        <label for="usuario" class="sr-only">Usuário</label>
        <input type="text" id="usuario" name='usuario' class="form-control" placeholder="Usuário" required autofocus />

        <label for="senha" class="sr-only">Senha</label>
        <input type="password" id="senha" name="senha" class="form-control mt-2 mb-2" placeholder="Senha" required />

        @include('mensagem', ['mensagem' => $mensagem ?? '', 'tipo' => 'danger' ])

        <button class="mt-5 btn btn-primary btn-block" id="botao-logar" type="submit">
            Logar
        </button>
    </form>

@endsection
