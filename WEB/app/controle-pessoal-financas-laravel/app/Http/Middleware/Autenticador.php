<?php

namespace App\Http\Middleware;

use App\Helpers\Imprime;
use App\Helpers\LogPersonalizado;
use Closure;
// use Symfony\Component\Console\Output\ConsoleOutput;

class Autenticador
{
    /**
     * Handle an incoming request.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \Closure  $next
     * @return mixed
     */
    public function handle($request, Closure $next)
    {
        $logado = $request->session()->get('logado');
        $estaLogado = $logado === true;

        // $output = new ConsoleOutput();
        // $output->writeln(">>> Está logado: $estaLogado <<<");

        if (!$estaLogado) {
            $rotaDestino = "login";
            LogPersonalizado::redirecionamento($request, $rotaDestino);

            return redirect()->route('login');
        }
        return $next($request);
    }
}
