<?php include_once("../includes/header.php"); ?>

<?php include_once("../includes/sidebar.php"); ?>

<?php
$sql_balance = ("SELECT * FROM `balance` ORDER BY id DESC LIMIT 1");

$balances = $connection->connection()->query($sql_balance)->fetchAll(PDO::FETCH_ASSOC);

$sql_expense = ("SELECT * FROM `expense` ORDER BY id DESC LIMIT 1");

$expenses = $connection->connection()->query($sql_expense)->fetchAll(PDO::FETCH_ASSOC);

?>
<div class="container">
    <div class="row">
        <div class="col-md-4 mt-5">
            <button class="btn btn-primary text-start" id="card">
                <h4>
                    <label for=""><i class="bi bi-bank"></i> Saldo:</label>
                    <br>
                    <?php foreach ($balances as $balance) {  ?>
                        <b href="http://">R$ <?php echo $balance["balance"] ?></b>
                    <?php  } ?>
                </h4>         
            </button>
        </div>
        <div class="col-md-4 mt-5">
            <button class="btn btn-success text-start" id="card">
                <h4>
                    <label for=""><i class="bi bi-arrow-up-circle-fill"></i> Ganhos:</label>
                </h4>
            </button>
        </div>
        <div class="col-md-4 mt-5">
            <button class="btn btn-danger text-start" id="card">
                <h4>
                    <label for=""><i class="bi bi-arrow-down-circle-fill"></i> Despesas:</label>
                </h4>
            </button>
        </div>
        <hr class="mt-3">
        <div class="col-md-12 text-center mt-1">
            <h3>Transações</h3>
        </div>
    </div>
</div>


<?php include_once("../includes/footer.php"); ?>