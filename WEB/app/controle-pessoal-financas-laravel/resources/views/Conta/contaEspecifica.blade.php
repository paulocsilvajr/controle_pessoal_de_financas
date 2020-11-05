@extends('layout', ['css' => 'home'])

@section('cabecalho')
    @include('navegacao')
@endsection

@section('conteudo')
    <div class="container">
        <p class="titulo">{{ $nomeConta }}</p>
    </div>
@endsection

@section('script')
@endsection
