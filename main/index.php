<?php include_once("../includes/header.php"); ?>

<?php
    $sql_balance = ("SELECT * FROM `balance` WHERE 1 = 1");

    $balance = $connection->connection()->query($sql_balance)->fetchAll(PDO::FETCH_ASSOC);

    ?>

<div class="container">
    <div class="row">
        <div class="col-md-12 text-center">  
            <label for="">Saldo:</label>
            <a href="http://"><?php print_r($balance)?></a>
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