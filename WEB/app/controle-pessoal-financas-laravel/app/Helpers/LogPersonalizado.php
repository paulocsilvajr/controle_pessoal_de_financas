<?php

namespace App\Helpers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Log;

final class LogPersonalizado
{
    public static function info($mensagem)
    {
        Log::info("$mensagem");
    }

    public static function error($mensagem)
    {
        Log::error("$mensagem");
    }

    public static function redirecionamento(Request $request, string $destino)
    {
        $origem = $request->url();
        Log:info("Redirecionamento de rota '$origem' para rota '$destino'");
    }
}
