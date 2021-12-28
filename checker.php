<?php
$duplicateIds = [];
$data = [];
$total_data = [];
$duplicateCount = 0;
$files_path = './log/';
$files = [
    'Golang-Publisher-A.log' => [
        'Golang-Publisher-A.log' => 0,
        'Golang-Publisher-B.log' => 0,
        'Golang-Publisher-C.log' => 0,
        'Golang-Publisher-D.log' => 0,
    ],
    'Golang-Publisher-B.log' => [
        'Golang-Publisher-A.log' => 0,
        'Golang-Publisher-B.log' => 0,
        'Golang-Publisher-C.log' => 0,
        'Golang-Publisher-D.log' => 0,
    ],
    'Golang-Publisher-C.log' => [
        'Golang-Publisher-A.log' => 0,
        'Golang-Publisher-B.log' => 0,
        'Golang-Publisher-C.log' => 0,
        'Golang-Publisher-D.log' => 0,
    ],
    'Golang-Publisher-D.log' => [
        'Golang-Publisher-A.log' => 0,
        'Golang-Publisher-B.log' => 0,
        'Golang-Publisher-C.log' => 0,
        'Golang-Publisher-D.log' => 0,
    ]
];

echo "Starting...\n";
$key = "";

foreach($files as $file => $dupAmount) {

    $f_content = file($files_path.$file);

    echo "Get file: ".$file." lines " . count($f_content) . " \n";

    foreach ($f_content as $line) {
        
        $line = trim($line);
        if(empty($line)) {
            continue 2;
        }

        
        $json = json_decode($line);

        $key = $json->appendix->event_id;

        $total_data[] = $key . ' ' . json_encode($json->data) . " " . json_encode($json->appendix);
        
        if(isset($data[$key])) {
            // $duplicateIds[] = json_encode($json->data);
            // $duplicateIds[] = $data[$key];
            $duplicateCount++;
            //$files[$file][$data[$key]['file']]++;
        }

        if(!isset($data[$key])) {
            $data[$key] = [
                'files' => [$file],
                'data' => [json_encode($json->data) . ' - ' . $json->appendix->event_id]
            ];
        } else {
            $data[$key]['files'][] = $file;
            $data[$key]['data'][] = json_encode($json->data) . ' - ' . $json->appendix->event_id;
        }
        
    }

}

sort($duplicateIds);
sort($total_data);

file_put_contents($files_path . 'Total.log', implode("\n", $total_data));

//print_r($files);

foreach($data as $key => $d) {
    echo "Key: " . $key . " Files: " . count($d['files']) . " Data: " . count($d['data']) . "\n";   
    foreach($d['data'] as $line) {
        echo "    " . $line . "\n";
    }
}

echo "Unique: " . count($data) . "\n";
echo "Duplicate: " . $duplicateCount . "\n";
echo "Each file duplication: \n";

//print_r($duplicateIds);
