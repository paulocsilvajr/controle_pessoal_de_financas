<?php

namespace App\Helpers;

use Symfony\Component\Console\Output\ConsoleOutput;

final class Imprime
{
    public static function console(string $mensagem)
    {
        $output = new ConsoleOutput();
        $output->writeln($mensagem);
    }
}
