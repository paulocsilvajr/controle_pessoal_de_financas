@extends('layout', ['css' => 'home'])

@section('cabecalho')
    @include('navegacao')
@endsection

@section('conteudo')
    <div class="container">
        <div class="row margem-navbar-conteudo mb-4">
            <h1 class="text-center" style="margin: auto">Cadastro de Lançamento em <u>{{ ucfirst($nomeConta) }}</u></h1>
        </div>

        <form action="" method="post">
            @csrf

            <input type="text" id="cpf_pessoa" name="cpf_pessoa" value="{{ $cpf ?? 'Enviar CPF para cá' }}" hidden>

            <input type="text" id="nome_conta_origem" name="nome_conta_origem" value="{{ $nomeConta }}" hidden>

            <div class="row">
                <div class="col-sm mb-2">
                    <label for="data">Data</label>
                    <input type="date" id="data" name="data" class="form-control" placeholder="Data" required autofocus />
                </div>
                <div class="col-sm mb-2">
                    <label for="numero">Número</label>
                    <input type="text" id="numero" name="numero" class="form-control" placeholder="Número" required />
                </div>
            </div>

            <div class="row">
                <div class="col mb-2">
                    <label for="descricao">Descrição</label>
                    <input type="text" id="descricao" name="descricao" class="form-control" placeholder="Descrição" required />
                </div>
            </div>

            <div class="row">
                <div class="col mb-2">
                    <label for="nome_conta_destino">Conta</label>
                    <input type="text" id="nome_conta_destino" name="nome_conta_destino" class="form-control" placeholder="Conta" required />
                </div>
            </div>

            <div class="row">
                <div class="col-sm mb-2">
                    <label for="valor">Valor</label>
                    <input type="number" min=".01" step=".01" id="valor" name="valor" class="form-control" placeholder="Valor" required />
                </div>

                <div class="col-sm mb-2">
                    <label for="tipo">Tipo</label>
                    <select name="tipo" id="tipo" class="form-control">
                        <option value="debito">Débito</option>
                        <option value="credito">Crédito</option>
                    </select>
                </div>
            </div>

            @include('mensagem', ['mensagem' => $mensagem ?? '', 'tipo' => $tipoMensagem ])

            <div class="row mt-3">
                <div class="col-sm mb-2">
                    <button class="btn btn-primary btn-block" id="botao-salvar-lancamento" type="submit">
                        Salvar
                    </button>
                </div>
                <div class="col-sm mb-2">
                    <button class="btn btn-secondary btn-block" id="botao-voltar" onclick="history.back()">
                        Voltar
                    </button>
                </div>
            </div>


        </form>
    </div>
@endsection

@section('script')
@endsection
