<?php

// REQUEST SOMA DE RECEITA

$urlSumRevenue = "http://localhost:8008/get";
$curlSumRevenue = curl_init($urlSumRevenue);

$getSumRevenue = array("optional" => "revenue");

$payloadSumRevenue = json_encode($getSumRevenue);

curl_setopt($curlSumRevenue, CURLOPT_CUSTOMREQUEST, 'GET');
curl_setopt($curlSumRevenue, CURLOPT_POSTFIELDS, $payloadSumRevenue);
curl_setopt($curlSumRevenue, CURLOPT_HTTPHEADER, array('Content-Type: application/json'));
curl_setopt($curlSumRevenue, CURLOPT_RETURNTRANSFER, true);

$resultSumRevenue = curl_exec($curlSumRevenue);

curl_close($curlSumRevenue);


// REQUEST SOMA DE DESPESAS

$urlSumExpense = "http://localhost:8008/get";
$curlSumExpense = curl_init($urlSumExpense);

$getSumExpense = array("optional" => "expense");

$payloadSumExpense = json_encode($getSumExpense);

curl_setopt($curlSumExpense, CURLOPT_CUSTOMREQUEST, 'GET');
curl_setopt($curlSumExpense, CURLOPT_POSTFIELDS, $payloadSumExpense);
curl_setopt($curlSumExpense, CURLOPT_HTTPHEADER, array('Content-Type: application/json'));
curl_setopt($curlSumExpense, CURLOPT_RETURNTRANSFER, true);

$resultSumExpense = curl_exec($curlSumExpense);

curl_close($curlSumExpense);

// REQUEST SOMA DE SALDO

$urlSumBalance = "http://localhost:8008/get";
$curlSumBalance = curl_init($urlSumBalance);

$getSumBalance = array("optional" => "balance");

$payloadSumBalance = json_encode($getSumBalance);

curl_setopt($curlSumBalance, CURLOPT_CUSTOMREQUEST, 'GET');
curl_setopt($curlSumBalance, CURLOPT_POSTFIELDS, $payloadSumBalance);
curl_setopt($curlSumBalance, CURLOPT_HTTPHEADER, array('Content-Type: application/json'));
curl_setopt($curlSumBalance, CURLOPT_RETURNTRANSFER, true);

$resultSumBalance = curl_exec($curlSumBalance);

curl_close($curlSumBalance);


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
$url = "http://localhost:8008/create";
$curl = curl_init($url);

$account_id  = isset($_POST['account_id'])  && $_POST['account_id']  != '' ? $_POST['account_id']  : null;
$category_id = isset($_POST['category_id']) && $_POST['category_id'] != '' ? $_POST['category_id'] : null;
$value       = isset($_POST['value'])       && $_POST['value']       != '' ? $_POST['value']       : null;
$type        = isset($_POST['type'])        && $_POST['type']        != '' ? "'" . $_POST['type'] . "'"  : null;
$status      = isset($_POST['status'])      && $_POST['status']      != '' ? "'" . $_POST['status'] . "'"      : null;
$description = isset($_POST['description']) && $_POST['description'] != '' ? "'" . $_POST['description'] . "'": null;
$date        = isset($_POST['date'])        && $_POST['date']        != '' ? "'" . $_POST['date'] . "'"      : null;

// $description = 'strval($description)';

$gambiarra = "";

// $gambiarra = $gambiarra . $description;

// var_dump($description);

$data = array(
    "name" => "finance",
    "columns" => array(
        "account_id",
        "category_id",
        "value",
        "type",
        "status",
        "description",
        "date",
    ),
    "rows" => array(
        array(
            "row" => array(
                $account_id,
                $category_id,
                $value,
                $type,
                $status,
                $description,
                $date,
            ),
        ),
    ),
);

$payload = json_encode($data);

curl_setopt($curl, CURLOPT_POSTFIELDS, $payload);
curl_setopt($curl, CURLOPT_HTTPHEADER, array('Content-Type: application/json'));
curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);

$result = curl_exec($curl);

curl_close($curl);



// echo json_encode($arraySaveExpense);

// $JsonSaveNewExpense = json_encode($data);

// $urlRequestSave = curl_init('http://localhost:8008/create?table=finance');


// curl_setopt_array($urlRequestSave, [

//     CURLOPT_CUSTOMREQUEST => 'POST',

//     // Permite obter o resultado
//     CURLOPT_RETURNTRANSFER => 1,

//     CURLOPT_POSTFIELDS => $JsonSaveNewExpense

// ]);

// $expenseSave = json_decode(curl_exec($urlRequestSave), true);

// curl_close($urlRequestSave);
