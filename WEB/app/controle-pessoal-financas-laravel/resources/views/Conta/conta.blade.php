@extends('layout', ['css' => 'home'])

@section('cabecalho')
    @include('navegacao')
@endsection

@section('conteudo')
    <div class="container">
        <h1 class="text-center mt-3">Contas</h1>

        {{-- impressão de variável com HTML --}}
        {{-- {!! $lista !!} --}}

        <?php
            function imprime(array $contas, string $nomeAnterior, string &$texto)
            {
                $texto .= '<ul class="nav-item" style="list-style: none;">';
                foreach ($contas as $conta) {
                    if ($conta['conta_pai'] == $nomeAnterior) {
                        $texto .= '<li class="nav-item">';
                        $texto .= '<a href="conta/' . $conta['nome'] . '" class="nav-link">' . ucfirst($conta['nome']) . '</a>';
                        imprime($contas, $conta['nome'], $texto);
                        $texto .= '</li>';
                    }
                }
                $texto .= '</ul>';
            }
        ?>

        {{-- impressão de contas recursiva --}}
        <ul class="nav flex-column mt-3">
            @foreach ($dados as $conta)
                @if (empty($conta['conta_pai']))
                    <li class="nav-item">
                    <a href="conta/{{ $conta['nome'] }}" class="nav-link"><strong>{{ ucfirst($conta['nome']) }}</strong></a>
                        <?php
                            $texto = '';
                            imprime($dados, $conta['nome'], $texto);
                            echo $texto;
                        ?>
                    </li>
                @endif
            @endforeach
        </ul>

    </div>
@endsection

@section('script')
@endsection
