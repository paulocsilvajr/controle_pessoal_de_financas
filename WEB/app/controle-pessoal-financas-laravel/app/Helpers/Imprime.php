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

    public static function console2(array $array)
    {
        // Fonte: https://stackoverflow.com/questions/48970080/errorexception-array-to-string-conversion-in-php
        Imprime::console(str_replace("'", "\'", json_encode($array)));
    }

    public static function console3($mensagem)
    {
        $output = new ConsoleOutput();
        $output->writeln($mensagem);
    }
}
