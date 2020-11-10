@extends('layout', ['css' => 'home'])

@section('cabecalho')
    @include('navegacao')
@endsection

@section('conteudo')
    <div class="container">
        <div class="row margem-navbar-conteudo">
            <h1 class="text-center" style="margin: auto">Cadastro de <u>{{ ucfirst($nomeConta) }}</u></h1>
        </div>
    </div>
@endsection

@section('script')
@endsection
