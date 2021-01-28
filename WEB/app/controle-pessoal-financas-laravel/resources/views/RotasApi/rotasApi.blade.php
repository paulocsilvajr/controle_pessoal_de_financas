@extends('layout', ['css' => 'home'])

@section('cabecalho')
    @include('navegacao')
@endsection

@section('conteudo')
    <div class="container-fluid">
        <h1 class="text-center margem-navbar-conteudo">Rotas da API</h1>

        <p>{{ $quant }} rotas cadastradas</p>

        <table class="table table-striped table-hover table-responsive mt-3">
            <thead class="thead-dark">
                <tr>
                    <th scope="col">#</th>
                    <th scope="col">Nome</th>
                    <th scope="col">Tipo</th>
                    <th scope="col">Rota</th>
                    <th scope="col">Descrição</th>
                    <th scope="col">Retorno</th>
                    <th scope="col">Documentação</th>
                </tr>
            </thead>
            <tbody>
                @foreach ($nomes as $cont => $nome)
                    <tr>
                        <th scope="row">{{ $cont + 1 }}</th>
                        <td>{{ $nome }}</td>
                        <td>{{ $dados[$nome]['Tipo'] }}</td>
                        <td><strong>{{ $dados[$nome]['Rota'] }}</strong></td>
                        <td>{{ $dados[$nome]['Descricao'] }}</td>
                        <td>{{ $dados[$nome]['Retorno'] }}</td>
                        <td>{{ $dados[$nome]['Documentacao'] }}</td>
                    </tr>
                @endforeach
            </tbody>
        </table>
    </div>
@endsection
