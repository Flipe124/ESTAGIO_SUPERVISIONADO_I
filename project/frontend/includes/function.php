<?php

//

function verifyStatusBalance($balance){
    if($balance < 0){
        return "warning";
    } else {
        return "primary";
    }
}

// SOMANDO O SALDO DAS CONTAS
function getBalance()
{
    $connection = new Database();

    $sqlBalance = ("SELECT SUM(`balance`) FROM `account`");

    $balances = $connection->connection()->query($sqlBalance)->fetchAll(PDO::FETCH_ASSOC);

    foreach ($balances as $balance) {
        // FORMATAÇÂO COM "." e ","
        return number_format($balance['SUM(balance)'],2,",",".");
    }


}
// Somando despesas/receitas pagas/pendente
function getSum($type, $status)
{
    $connection = new Database();

    $sqlExpense = ("SELECT SUM(value) FROM `finance` WHERE `type` = '$type' AND `status` = '$status'");

    $expense = $connection->connection()->query($sqlExpense)->fetchAll(PDO::FETCH_ASSOC);

    foreach ($expense as $expenseValue) {
        // FORMATAÇÂO COM "." e ","
        return number_format($expenseValue['SUM(value)'],2,",",".");
    }
}

// Pegando o ID da conta e apresentando o NOME
function getAccount($id)
{
    $connection = new Database();

    $sqlAccount = ("SELECT * FROM `account` WHERE `id` = '$id' ");

    $Accounts = $connection->connection()->query($sqlAccount)->fetchAll(PDO::FETCH_ASSOC);

    foreach ($Accounts as $Account) {
        return $Account['name'];
    }
}

// Pegando o ID da categoria e apresentando o NOME
function getCategory($id)
{
    $connection = new Database();

    $sqlCategory = ("SELECT * FROM `category` WHERE `id` = '$id' ");

    $categorys = $connection->connection()->query($sqlCategory)->fetchAll(PDO::FETCH_ASSOC);

    foreach ($categorys as $category) {
        return $category['name'];
    }
}