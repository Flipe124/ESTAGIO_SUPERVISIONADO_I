<?php include_once("../includes/header.php"); ?>

<?php
$sql_balance = ("SELECT * FROM `balance` ORDER BY id DESC LIMIT 1");

$balances = $connection->connection()->query($sql_balance)->fetchAll(PDO::FETCH_ASSOC);

?>

<div class="container">
    <div class="row">
        <div class="col-md-12 text-center bg-primary" id="balance">
            <b>
                <label for="">Saldo:</label>
                <?php foreach ($balances as $balance) {  ?>
                    <a href="http://">R$ <?php echo $balance["balance"] ?></a>
                <?php  } ?>
            </b>

        </div>
        <div class="col-md-6 text-center">
            <label for="">Ganhos:</label>
        </div>
        <div class="col-md-6 text-center">
            <label for="">Despesas:</label>
        </div>

    </div>
</div>


<?php include_once("../includes/footer.php"); ?>