<?php include_once("../includes/header.php"); ?>

<?php include_once("../includes/sidebar.php"); ?>

<?php
//  $sql_balance = ("SELECT * FROM `balance` ORDER BY id DESC LIMIT 1");

// $balances = $connection->connection()->query($sql_balance)->fetchAll(PDO::FETCH_ASSOC);

// $sql_expense = ("SELECT * FROM `expense` ORDER BY id DESC LIMIT 1");

// $expenses = $connection->connection()->query($sql_expense)->fetchAll(PDO::FETCH_ASSOC);

?>

<?php
$sql_balance = ("SELECT SUM(`balance`) FROM `account`");

$balances = $connection->connection()->query($sql_balance)->fetchAll(PDO::FETCH_ASSOC);

?>
<div class="container">
    <div class="row ms-4">
        <div class="col-md-12 mt-4">
            <h1>Dashboard</h1>
        </div>
        <div class="col-md-4 mt-3">
            <button class="btn btn-primary text-start" id="card">
                <h4>
                    <label for=""><i class="bi bi-bank"></i> Saldo:</label>
                    <br>
                    <?php foreach ($balances as $balance) {  
                        print_r($balance)
                        ?>
                        <b>R$ <?php echo number_format($balance['SUM(balance)'],2,",",".") ?></b>
                    <?php  } ?>
                </h4>
            </button>
        </div>
        <div class="col-md-4 mt-3">
            <button class="btn btn-success text-start" id="card">
                <h4>
                    <label for=""><i class="bi bi-arrow-up-circle-fill"></i> Receita:</label>
                </h4>
            </button>
        </div>
        <div class="col-md-4 mt-3">
            <a class="btn btn-danger" href="../registration/expense.php" id="card">
                <h4>
                    <label for=""><i class="bi bi-arrow-down-circle-fill"></i> Despesa:  <?php echo date('m/Y')?></label>
                </h4>
            </a>
        </div>
        <hr class="mt-3">
        <div class="col-md-12 text-center mt-1">
            <h3>Transações</h3>
        </div>
        <div class="col-md-12 flow">
        </div>
    </div>
</div>


<?php include_once("../includes/footer.php"); ?>