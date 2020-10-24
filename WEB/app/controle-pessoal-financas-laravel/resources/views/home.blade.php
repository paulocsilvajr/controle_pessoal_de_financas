@extends('layout', ['css' => 'home'])

@section('cabecalho')
    @include('navegacao')
@endsection

@section('conteudo')
    <h1 class="mensagem-inicial text-center">Bem vindo {{ $usuario ?? '' }}</h1>
    <p class="text-center mt-3">Use a barra de navegação para acessar as funcionalidades do sistema</p>

    <div class="card informacoes-sessao" style="width: 20rem;">
        <div class="card-body">
            <h5 class="card-title">Dados em sessão</h5>
            <p class="card-text">usuário: {{ $usuario }}</p>
            <p class="card-text">senha: {{ $senha }}</p>
            <p class="card-text">token: {{ $tokenParcial }}</p>
            <p class="card-text">logado: {{ $logado }}</p>
        </div>
    </div>
@endsection
