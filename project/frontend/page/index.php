<?php include_once("../includes/header.php"); ?>

<?php include_once("../includes/sidebar.php"); ?>

<?php
//  $sql_balance = ("SELECT * FROM `balance` ORDER BY id DESC LIMIT 1");

// $balances = $connection->connection()->query($sql_balance)->fetchAll(PDO::FETCH_ASSOC);

// $sql_expense = ("SELECT * FROM `expense` ORDER BY id DESC LIMIT 1");

// $expenses = $connection->connection()->query($sql_expense)->fetchAll(PDO::FETCH_ASSOC);

$sql_finance = ("SELECT * FROM `finance` ORDER BY `date` DESC");

$finances = $connection->connection()->query($sql_finance)->fetchAll(PDO::FETCH_ASSOC);

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
                        <b>R$ <?php echo number_format($balance['SUM(balance)'], 2, ",", ".") ?></b>
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
                    <label for=""><i class="bi bi-arrow-down-circle-fill"></i> Despesa: <?php echo date('m/Y') ?></label>
                </h4>
            </a>
        </div>
        <hr class="mt-3">
        <div class="col-md-12 text-center mt-1">
            <h3>Transações</h3>
        </div>
        <div class="col-md-12">
            <div class="text-center" id="div-table">
                <table class="table table-light table-striped">
                    <thead>
                        <tr>
                            <th>Situação</th>
                            <th>Valor</th>
                            <th>Data</th>
                            <th>Descrição</th>
                            <th>Categoria</th>
                            <th>Conta</th>
                            <th>Ações</th>
                        </tr>
                    </thead>
                    <tbody>
                        <?php foreach ($finances as $finance) { ?>
                            <tr>
                                <th><?php echo $finance["status"]; ?></th>
                                <?php if($finance["type"] == "EXPENSE") { ?>
                                    <td class="text-danger"><?php echo "R$ - ". number_format($finance['value'],2,",","."); ?></td>
                                <?php } else { ?>
                                    <td class="text-success"><?php echo "R$ ". number_format($finance['value'],2,",","."); ?></td>
                                <?php } ?> 
                                <td><?php echo date('d/m/Y', strtotime($finance['date'])); ?></td>
                                <td><?php echo $finance["description"]; ?></td>
                                <td><?php echo getCategory($finance["category_id"]); ?></td>
                                <td><?php echo getAccount($finance["account_id"]); ?></td>
                                <td>
                                    <button class="btn btn-danger btn-delete-expense" type="button"><i class="fa-solid fa-trash"></i></button>
                                    <button class="btn btn-primary btn-edit-expense" type="button"><i class="fa-solid fa-pen"></i></button>
                                </td>
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