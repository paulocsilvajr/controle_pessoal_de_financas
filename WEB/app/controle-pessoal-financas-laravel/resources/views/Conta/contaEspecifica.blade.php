@extends('layout', ['css' => 'home'])

@section('cabecalho')
    @include('navegacao')
@endsection

@section('conteudo')
    <div class="container">
        <div class="row margem-navbar-conteudo mb-2">
            <div class="col-sm-6 col-md-9 titulo-conta-sm">
                <h1 class="text-center" style="margin: auto">{{ ucfirst($nomeConta) }}</h1>
            </div>
            <div class="col-sm-6 col-md-3 d-flex align-items-center">
            <button class="btn btn-primary btn-block" onclick="location.href = '/conta/{{ $nomeConta }}/cadastroLancamento';">
                    Novo Lançamento
                    <i class="fa fa-plus" aria-hidden="true"></i>
                </button>
            </div>
        </div>

        @if($dados == null)
            @php
                $mensagem = 'Sem lançamentos para a conta \'' . ucfirst($nomeConta) . '\'';
            @endphp
            @include('mensagem', ['mensagem' => $mensagem, 'tipo' => 'danger' ])
        @else
            <div class="table-responsive">

                <table class="table table-striped table-hover mt-3">
                    <thead class="thead-dark">
                        <tr>
                            <th scope="col">#</th>
                            <th scope="col">Data</th>
                            <th scope="col">Número</th>
                            <th scope="col">Descrição</th>
                            <th scope="col">Conta</th>
                            <th scope="col">Débito</th>
                            <th scope="col">Crédito</th>
                        </tr>
                    </thead>
                    <tbody>
                        @foreach ($dados as $dado)
                            <tr>
                                <th scope="row">{{ $dado['id'] }}</th>
                                <td>{{ App\Helpers\Formata::textoParaDataBrasil($dado['data']) }}</td>
                                <td>{{ $dado['numero'] }}</td>
                                <td><strong>{{ $dado['descricao'] }}</strong></td>
                                @if ($dado['nome_conta_origem'] == $nomeConta)
                                    <td>{{ ucfirst($dado['nome_conta_destino']) }}</td>
                                @else
                                    <td>{{ ucfirst($dado['nome_conta_origem']) }}</td>
                                @endif
                                <td class="alinhamento-numeros-tabela">{{ App\Helpers\Formata::valorParaMonetarioBrasil($dado['debito']) }}</td>
                                <td class="alinhamento-numeros-tabela">{{ App\Helpers\Formata::valorParaMonetarioBrasil($dado['credito']) }}</td>
                            </tr>
                        @endforeach
                    </tbody>
                </table>

            </div>
        @endif

    </div>
@endsection

@section('script')
@endsection
