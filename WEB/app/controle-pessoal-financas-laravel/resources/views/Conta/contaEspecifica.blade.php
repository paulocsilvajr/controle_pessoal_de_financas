@extends('layout', ['css' => 'home'])

@section('cabecalho')
    @include('navegacao')
@endsection

@section('conteudo')
    <div class="container">
        <div class="row margem-navbar-conteudo">
            <div class="col-sm-6 col-md-10 titulo-conta-sm">
                <h1 class="text-center" style="margin: auto">{{ ucfirst($nomeConta) }}</h1>
            </div>
            <div class="col-sm-6 col-md-2 d-flex align-items-center">
                <button class="btn btn-primary btn-block">
                    Novo
                    <i class="fa fa-plus" aria-hidden="true"></i>
                </button>
            </div>
        </div>
        <p>Continua...</p>
    </div>
@endsection

@section('script')
@endsection
