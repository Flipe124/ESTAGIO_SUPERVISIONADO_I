<?php include_once("../includes/header.php"); ?>

<?php include_once("../includes/sidebar.php"); ?>

<?php
$sql_finance = ("SELECT * FROM `account` ORDER BY `date` DESC");

$accounts = $connection->connection()->query($sql_finance)->fetchAll(PDO::FETCH_ASSOC);
?>

<div class="container">
    <div class="row ms-2">
        <div class="col-md-12 mt-4">
            <h1>Contas</h1>
        </div>
        <div class="col-md-12 mt-3">
            <div class="text-center" id="div-table">
                <table class="table table-light table-striped">
                    <thead>
                        <tr>
                            <th>Conta</th>
                            <th>Saldo</th>
                            <th>Saldo previsto</th>
                            <th>Ações</th>
                        </tr>
                    </thead>
                    <tbody>
                        
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</div>

<?php include_once("../includes/footer.php"); ?>