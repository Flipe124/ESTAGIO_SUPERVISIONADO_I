<?php include_once("../includes/header.php"); ?>

<?php include_once("../includes/sidebar.php"); ?>

<?php include_once("../includes/request.php"); ?>

<?php
//  $sql_balance = ("SELECT * FROM `balance` ORDER BY id DESC LIMIT 1");

// $balances = $connection->connection()->query($sql_balance)->fetchAll(PDO::FETCH_ASSOC);

// $sql_expense = ("SELECT * FROM `expense` ORDER BY id DESC LIMIT 1");

// $expenses = $connection->connection()->query($sql_expense)->fetchAll(PDO::FETCH_ASSOC);

$sqlFinanceNotPaid = ("SELECT * FROM `finance` WHERE `status` = 'NOT_PAID' ORDER BY `date` DESC");

$financesNotPaid = $connection->connection()->query($sqlFinanceNotPaid)->fetchAll(PDO::FETCH_ASSOC);

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
            <a class="btn btn-<?php echo verifyStatusBalance($resultSumBalance)?>" href="../page/account.php" id="card">
                <h4>
                    <label for=""><i class="bi bi-bank"></i> Saldo:</label>
                    <br>
                    <?php if ($resultSumBalance < 0) { ?>
                        <b class="text-danger">R$ <?php echo number_format($resultSumBalance, 2, ",", ".") ?></b>
                    <?php } else { ?>
                        <b>R$ <?php echo number_format($resultSumBalance, 2, ",", ".") ?></b>
                    <?php } ?>

                </h4>
            </a>
        </div>
        <div class="col-md-4 mt-3">
            <a class="btn btn-success" href="../page/revenue.php" id="card">
                <h4>
                    <label for=""><i class="bi bi-arrow-up-circle-fill"></i> Receita:</label>
                    <br>
                    <b>R$ <?php echo number_format($resultSumRevenue, 2, ",", ".") ?></b>
                </h4>
            </a>
        </div>
        <div class="col-md-4 mt-3">
            <a class="btn btn-danger" href="../page/expense.php" id="card">
                <h4>
                    <label for=""><i class="bi bi-arrow-down-circle-fill"></i> Despesa:</label>
                    <br>
                    <b>R$ - <?php echo number_format($resultSumExpense, 2, ",", ".") ?></b>
                </h4>
            </a>
        </div>
        <hr class="mt-3">
        <div class="col-md-12 text-center mt-1">
            <h3>Pendente</h3>
        </div>
        <div class="col-md-12">
            <div class="text-center" id="div-table">
                <table class="table table-light table-striped">
                    <thead>
                        <tr>
                            <th class="text-center">Situação</th>
                            <th class="text-center">Valor</th>
                            <th class="text-center">Data</th>
                            <th>Descrição</th>
                            <th>Categoria</th>
                            <th>Conta</th>
                            <th class="text-center">Ações</th>
                        </tr>
                    </thead>
                    <tbody>
                        <?php foreach ($financesNotPaid as $financeNotPaid) { ?>
                            <tr>
                                <?php if ($financeNotPaid["status"] == "PAID") { ?>
                                    <th class="text-center"><i class="text-success fa-solid fa-circle-check"></i></th>
                                <?php } else { ?>
                                    <th class="text-center"><i class="text-danger fa-solid fa-circle-xmark"></i></th>
                                <?php } ?>
                                <th class="text-danger text-end"><?php echo "R$ - " . number_format($financeNotPaid['value'], 2, ",", "."); ?></th>
                                <th class="text-center"><?php echo date('d/m/Y', strtotime($financeNotPaid['date'])); ?></th>
                                <th><?php echo $financeNotPaid['description'] ?></th>
                                <th><?php echo getCategory($financeNotPaid['category_id']) ?></th>
                                <th><?php echo getAccount($financeNotPaid['account_id']) ?></th>
                                <th class="text-center">
                                    <button class="btn btn-danger btn-delete-expense" type="button" data-id="<?php echo $expense['id'] ?>" data-value="<?php echo $expense['value'] ?>" data-description="<?php echo $expense['description'] ?>"><i class="fa-solid fa-trash"></i></button>
                                    <button class="btn btn-primary btn-update-expense" id="BTN" type="button" data-id="<?php echo $expense['id'] ?>" data-status="<?php echo $expense['status'] ?>" data-value="<?php echo $expense['id'] ?>" data-date="<?php echo $expense['date'] ?>" data-description="<?php echo $expense['description'] ?>" data-category="<?php echo getCategory($expense['category_id']) ?>" data-account="<?php echo getAccount($expense['account_id']) ?>"><i class="fa-solid fa-pen"></i></button>
                                </th>
                            </tr>
                        <?php } ?>
                    </tbody>
                </table>
            </div>
        </div>
        <div class="col-md-12 flow">
        </div>
    </div>
</div>


<?php include_once("../includes/footer.php"); ?>