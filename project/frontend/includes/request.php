<?php

// REQUEST LIST

// $ch = curl_init('http://localhost:8008/list?table=finance');

// curl_setopt_array($ch, [

//     CURLOPT_CUSTOMREQUEST => 'GET',

//     // Permite obter o resultado
//     CURLOPT_RETURNTRANSFER => 1,
// ]);

// $expensess = json_decode(curl_exec($ch), true);
// curl_close($ch);

// REQUEST SAVE

$account_id = isset($_POST['account_id']) && $_POST['account_id'] != ''  ? $_POST['account_id'] : null;
$category_id = isset($_POST['category_id']) && $_POST['category_id'] != '' ? $_POST['category_id'] : null;
$value = isset($_POST['value']) && $_POST['value'] != '' ? $_POST['value'] : null;
$type = isset($_POST['type']) && $_POST['type'] != '' ? $_POST['type'] : null;
$description = isset($_POST['description']) && $_POST['description'] != '' ? $_POST['description'] : null;
$date = isset($_POST['date']) && $_POST['date'] != '' ? $_POST['date'] : null;

$saveExpense = [
    // 'account_id' => $account_id,
    // 'category_id' => $category_id,
    // 'value' => $value,
    // 'type' => $type,
    // 'description' => $description,
    // 'date' => $date

    'account_id' => 1,
    'category_id' => 1,
    'value' => 1100,
    'type' => "EXPENSE",
    'description' => "TESTEEE",
    'date' => "08-10-2002"
];


$JsonSaveExpense = json_encode($saveExpense);

// print_r($saveExpense);

$urlRequestSave = curl_init('http://localhost:8008/create?table=finance');

curl_setopt_array($urlRequestSave, [

    CURLOPT_CUSTOMREQUEST => 'POST',

    // Permite obter o resultado
    CURLOPT_RETURNTRANSFER => 1,

    CURLOPT_POSTFIELDS => $JsonSaveExpense

]);

echo $JsonSaveExpense;

$expenseSave = json_decode(curl_exec($urlRequestSave), true);

curl_close($urlRequestSave);

