<?php

$data = file_get_contents("input.txt");

printf("Part 1: %s\n", part1($data));
printf("Part 2: %s\n", part2($data));

function part1(string $data): int{
    $result = 0;
    $data = parseInput($data);

    $re = '/mul\((\d{1,3})\,(\d{1,3})\)/m';

    preg_match_all($re, $data, $matches, PREG_SET_ORDER, 0);

    foreach ($matches as $match) {
        $result += intval($match[1]) * intval($match[2]);
    }

    return $result;
}

function part2(string $data): int{
    $result = 0;
    $data = parseInput($data);

    $re = '/mul\((\d{1,3})\,(\d{1,3})\)|do\(\)|don\'t\(\)/m';

    preg_match_all($re, $data, $matches, PREG_SET_ORDER, 0);
    $do = true;
    foreach ($matches as $match) {

        if ($match[0] == "do()"){
            $do = true;
            continue;
        } else if ($match[0] == "don't()"){
            $do = false;
            continue;
        }

        if ($do && ($match[1])) {
            $result += intval($match[1]) * intval($match[2]);
        }
        
    }

    return $result;
}

function parseInput(string $data):string{
    return $data;
}



