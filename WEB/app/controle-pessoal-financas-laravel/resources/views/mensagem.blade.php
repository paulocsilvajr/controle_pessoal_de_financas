@if(!empty($mensagem))
<div class="alert alert-{{ $tipo }}">
    {{ $mensagem }}
</div>
@endif
