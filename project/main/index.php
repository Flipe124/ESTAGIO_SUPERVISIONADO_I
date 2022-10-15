<?php include_once("../includes/header.php"); ?>

<?php
$sql_balance = ("SELECT * FROM `balance` ORDER BY id DESC LIMIT 1");

$balances = $connection->connection()->query($sql_balance)->fetchAll(PDO::FETCH_ASSOC);

$sql_expense = ("SELECT * FROM `expense` ORDER BY id DESC LIMIT 1");

$expenses = $connection->connection()->query($sql_expense)->fetchAll(PDO::FETCH_ASSOC);

?>

<div class="container">
    <div class="row justify-content-center">
        <div class="btn btn-primary col-md-10 text-center mt-2" id="balance">
            <h2>
                <label for=""><i class="bi bi-bank"></i> Saldo:</label>
                <?php foreach ($balances as $balance) {  ?>
                    <a href="http://">R$ <?php echo $balance["balance"] ?></a>
                <?php  } ?>
            </h2>
        </div>
        <div class="btn btn-success col-md-5 text-start gain-and-expense mt-2">
            <h3>
                <label for=""><i class="bi bi-arrow-up-circle-fill"></i> Ganhos:</label>
            </h3>
        </div>
        <div class="btn btn-danger col-md-5 text-start gain-and-expense mt-2">
            <h3>
                <label for=""><i class="bi bi-arrow-down-circle-fill"></i> Despesas:</label>
            </h3>
        </div>
        <hr class="mt-3">
        <div class="col-md-12 text-center mt-1">
            <h3>Transações</h3>
        </div>
    </div>


</div>


<?php include_once("../includes/footer.php"); ?>