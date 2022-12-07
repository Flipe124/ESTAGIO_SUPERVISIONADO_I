<?php

// REQUEST LIST

$ch = curl_init('http://localhost:8008/list?table=finance');

curl_setopt_array($ch, [

    CURLOPT_CUSTOMREQUEST => 'GET',

    // Permite obter o resultado
    CURLOPT_RETURNTRANSFER => 1,
]);

$expensess = json_decode(curl_exec($ch), true);
curl_close($ch);

// REQUEST SAVE

$account_id = isset($_POST['account_id']) && $_POST['account_id'] != '';
$category_id = isset($_POST['category_id']) && $_POST['category_id'] != '';
$value = isset($_POST['value']) && $_POST['value'] != '';
$type = isset($_POST['type']) && $_POST['type'] != '';
$description = isset($_POST['description']) && $_POST['description'] != '';
$date = isset($_POST['date']) && $_POST['date'] != '';


var_dump($value);
print_r($value);
echo $value;

$saveExpense = [
    'account_id' => $account_id,
    'category_id' => $category_id,
    'value' => $value,
    'type' => $type,
    'description' => $description,
    'date' => $date
];

$urlRequestSave = curl_init('http://localhost:8008/save?table=finance');

curl_setopt_array($urlRequestSave, [

    CURLOPT_CUSTOMREQUEST => 'POST',

    // Permite obter o resultado
    CURLOPT_RETURNTRANSFER => 1,

    CURLOPT_POSTFIELDS => $saveExpense

]);

$expenseSave = json_decode(curl_exec($urlRequestSave), true);
curl_close($urlRequestSave);

