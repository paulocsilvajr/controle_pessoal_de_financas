@extends('layout', ['css' => 'home'])

@section('cabecalho')
    @include('navegacao')
@endsection

@section('conteudo')
    <div class="container">
        <h1 class="text-center mt-3">Contas</h1>

        <ul class="list-group mt-3">
            @foreach ($dados as $conta)
                @if (empty($conta['conta_pai']))
                    <li class="list-group-item list-group-item-action">
                        <a href="#">{{ ucfirst($conta['nome']) }}</a>
                    </li>
                @endif
            @endforeach
        </ul>
@endsection

@section('script')
@endsection
