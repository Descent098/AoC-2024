<?php
$data = file_get_contents("input.txt");

printf("%s\n", part1($data));
printf("%s\n", part2($data));

function parse_input_text(string $input): array {
    $lines = explode("\r\n", $input);

    $leftList = [];
    $rightList = [];

    foreach ($lines as $line) {
        $temp = explode("   ", $line);

        array_push($leftList, intval( $temp[0]));
        array_push($rightList, intval($temp[1]));
    }

    sort($leftList);
    sort($rightList);
    return [$leftList, $rightList];
}

function part1(string $input){
    $temp = parse_input_text($input);

    $leftList = $temp[0];
    $rightList = $temp[1];

    $difference = 0;
    for ($i=0; $i < count($leftList); $i++) { 
        $difference += abs($leftList[$i]-$rightList[$i]);
    }
    return $difference;
}

function part2(string $input){
    $temp = parse_input_text($input);

    $leftList = $temp[0];
    $rightList = $temp[1];

    $similarity = 0;
    for ($i=0; $i < count($leftList); $i++) { 
        $appearances = 0;
        for ($j=0; $j < count($rightList) ; $j++) { 
            

            if ($leftList[$i] == $rightList[$j]){
                $appearances += 1;

            } else if ($leftList[$i] < $rightList[$j]){
                break;
            }
        }
        $similarity += $leftList[$i] * $appearances;
    }
    return $similarity;

}


