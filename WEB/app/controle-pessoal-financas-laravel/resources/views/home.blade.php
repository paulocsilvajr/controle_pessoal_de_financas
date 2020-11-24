@extends('layout', ['css' => 'home'])

@section('cabecalho')
    @include('navegacao')
@endsection

@section('conteudo')
    <div class="container">
        <div class="row">
            <div class="col">
                <h1 class="mensagem-inicial text-center">Bem vindo {{ $usuario ?? '' }}</h1>
            </div>
        </div>
        <div class="row">
            <div class="col">
                <p class="text-center mt-3">Use a barra de navegação para acessar as funcionalidades do sistema</p>
            </div>
        </div>

        @if (env('APP_DEBUG', false))
            <div class="row">
                <div class="col">
                    <div class="card informacoes-sessao" style="width: 20rem;">
                        <div class="card-body">
                            <h5 class="card-title">Dados em sessão</h5>
                            <p class="card-text">usuário: {{ $usuario }}</p>
                            <p class="card-text">senha: {{ $senha }}</p>
                            <p class="card-text">token: {{ $tokenParcial }}</p>
                            <p class="card-text">logado: {{ $logado }}</p>
                        </div>
                    </div>
                </div>
            </div>
        @endif

        <div class="row">
            <div class="col">
                <div id="app">
                    <exemplo></exemplo>
                </div>
            </div>
        </div>
    </div>

    {{-- <h1 class="mensagem-inicial text-center">Bem vindo {{ $usuario ?? '' }}</h1> --}}
    {{-- <p class="text-center mt-3">Use a barra de navegação para acessar as funcionalidades do sistema</p> --}}

    {{-- @if (env('APP_DEBUG', false))
        <div class="card informacoes-sessao" style="width: 20rem;">
            <div class="card-body">
                <h5 class="card-title">Dados em sessão</h5>
                <p class="card-text">usuário: {{ $usuario }}</p>
                <p class="card-text">senha: {{ $senha }}</p>
                <p class="card-text">token: {{ $tokenParcial }}</p>
                <p class="card-text">logado: {{ $logado }}</p>
            </div>
        </div>
    @endif --}}

    {{-- <div id="app">
        <exemplo></exemplo>
    </div> --}}

@endsection

@section('script')
    <script src="{{ asset('js/app.js') }}"></script>
@endsection
