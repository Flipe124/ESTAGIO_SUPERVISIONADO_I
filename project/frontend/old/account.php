<?php include_once("../includes/header.php"); ?>

<?php include_once("../includes/sidebar.php"); ?>

<?php
$sql_account = ("SELECT * FROM `account` ORDER BY `name` DESC");

$accounts = $connection->connection()->query($sql_account)->fetchAll(PDO::FETCH_ASSOC);
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
                            <!-- <th>Saldo previsto</th> -->
                            <th>Ações</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <?php foreach ($accounts as $Account) { ?>
                                <th><?php echo $Account['name'] ?></th>
                                <th class="text-success">R$ <?php echo number_format($Account['balance'], 2, ",", ".") ?></th>
                                <th>
                                    <button class="btn btn-danger btn-delete-expense" type="button"><i class="fa-solid fa-trash"></i></button>
                                    <button class="btn btn-primary btn-update-expense" id="BTN" type="button"><i class="fa-solid fa-pen"></i></button>
                                </th>
                        </tr>
                    <?php } ?>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</div>

<?php include_once("../includes/footer.php"); ?>