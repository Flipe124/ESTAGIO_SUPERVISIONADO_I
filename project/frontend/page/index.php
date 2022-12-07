<?php include_once("../includes/header.php"); ?>

<?php include_once("../includes/sidebar.php"); ?>

<?php include_once("../includes/request.php"); ?>

<?php
//  $sql_balance = ("SELECT * FROM `balance` ORDER BY id DESC LIMIT 1");

// $balances = $connection->connection()->query($sql_balance)->fetchAll(PDO::FETCH_ASSOC);

// $sql_expense = ("SELECT * FROM `expense` ORDER BY id DESC LIMIT 1");

// $expenses = $connection->connection()->query($sql_expense)->fetchAll(PDO::FETCH_ASSOC);

$sql_finance = ("SELECT * FROM `finance` ORDER BY `date` DESC");

$finances = $connection->connection()->query($sql_finance)->fetchAll(PDO::FETCH_ASSOC);

?>

<?php
// $sql_balance = ("SELECT * SUM(`balance`) FROM `account`");

// $balances = $connection->connection()->query($sql_balance)->fetchAll(PDO::FETCH_ASSOC);

?>
<div class="container">
    <div class="row ms-4">
        <div class="col-md-12 mt-4">
            <h1>Dashboard</h1>
        </div>
        <div class="col-md-4 mt-3">
            <a class="btn btn-primary" href="../page/account.php" id="card">
                <h4>
                    <label for=""><i class="bi bi-bank"></i> Saldo:</label>
                    <br>
                    <b>R$ <?php  echo number_format($resultSumBalance, 2, ",", ".") ?></b>
                    
                </h4>
            </a>
        </div>
        <div class="col-md-4 mt-3">
            <a class="btn btn-success" href="../page/revenue.php"  id="card">
                <h4>
                    <label for=""><i class="bi bi-arrow-up-circle-fill"></i> Receita:</label>
                    <br>
                    <b>R$ <?php  echo number_format($resultSumRevenue, 2, ",", ".") ?></b>
                </h4>
            </a>
        </div>
        <div class="col-md-4 mt-3">
            <a class="btn btn-danger" href="../page/expense.php" id="card">
                <h4>
                    <label for=""><i class="bi bi-arrow-down-circle-fill"></i> Despesa: <?php echo date('m/Y') ?></label>
                    <br>
                    <b>R$ <?php  echo number_format($resultSumExpense, 2, ",", ".") ?></b>
                </h4>
            </a>
        </div>
        <hr class="mt-3">
        <div class="col-md-12 text-center mt-1">
            <h3>Pendente</h3>
        </div>
        <div class="col-md-12">
            
        </div>
        <div class="col-md-12 flow">
        </div>
    </div>
</div>


<?php include_once("../includes/footer.php"); ?>