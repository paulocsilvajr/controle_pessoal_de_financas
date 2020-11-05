@extends('layout', ['css' => 'login'])

@section('conteudo')

    <form class="formulario-login text-center" method="post" action="/login">
        @csrf
        <h1 class="titulo-login">Controle pessoal de Finanças</h1>

        <h2 class="mb-3">Informe seu usuário e senha</h2>

        <label for="usuario" class="sr-only">Usuário</label>
        <input type="text" id="usuario" name='usuario' class="form-control" placeholder="Usuário" required autofocus />

        <label for="senha" class="sr-only">Senha</label>
        <input type="password" id="senha" name="senha" class="form-control mt-2 mb-3" placeholder="Senha" required />

        @include('mensagem', ['mensagem' => $mensagem ?? '', 'tipo' => $tipoMensagem ])

        @include('mensagemApi', ['estadoApi' => $estadoApi ?? false])

        <button class="mt-3 btn btn-primary btn-block" id="botao-logar" type="submit" {{ $estadoApi === true ? '' : 'disabled' }}>
            Logar
        </button>
    </form>

    <div id="app">
        <desenvolvedor></desenvolvedor>
    </div>

@endsection

@section('script')
    <script src="{{ asset('js/app.js') }}"></script>
@endsection
