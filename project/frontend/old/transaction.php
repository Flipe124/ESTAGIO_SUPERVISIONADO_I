<?php include_once("../includes/header.php"); ?>

<?php include_once("../includes/sidebar.php"); ?>

<?php
$sql_finance = ("SELECT * FROM `finance` ORDER BY `date` DESC");

$finances = $connection->connection()->query($sql_finance)->fetchAll(PDO::FETCH_ASSOC);
?>

<div class="container">
    <div class="row ms-2">
        <div class="col-md-12 mt-4">
            <h1>Transações</h1>
        </div>
        <div class="col-md-12 text-center">
            <button class="btn btn-warning"><i class="fa-solid fa-arrow-left"></i></button>
            <button class="btn btn-warning"><b>Novembro</b> 2022</button>
            <button class="btn btn-warning"><i class="fa-solid fa-arrow-right"></i></button>
        </div>
        <div class="col-md-12 mt-3">
            <div id="div-table">
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
                        <?php foreach ($finances as $finance) { ?>
                            <tr>
                                <?php if ($finance["status"] == "PAID") { ?>
                                    <th class="text-center"><i class="text-success fa-solid fa-circle-check"></i></th>
                                <?php } else { ?>
                                    <th class="text-center"> <i class="text-danger fa-solid fa-circle-xmark"></i></th>
                                <?php } ?>
                                <?php if ($finance["type"] == "EXPENSE") { ?>
                                    <td class="text-danger text-end"><?php echo "R$ - " . number_format($finance['value'], 2, ",", "."); ?></td>
                                <?php } else { ?>
                                    <td class="text-success text-end"><?php echo "R$ " . number_format($finance['value'], 2, ",", "."); ?></td>
                                <?php } ?>
                                <td class="text-center"><?php echo date('d/m/Y', strtotime($finance['date'])); ?></td>
                                <td><?php echo $finance["description"]; ?></td>
                                <td><?php echo getCategory($finance["category_id"]); ?></td>
                                <td><?php echo getAccount($finance["account_id"]); ?></td>
                                <td class="text-center">
                                    <button class="btn btn-danger btn-delete-expense" type="button"><i class="fa-solid fa-trash"></i></button>
                                    <button class="btn btn-primary btn-edit-expense" type="button"><i class="fa-solid fa-pen"></i></button>
                                </td>
                            </tr>
                        <?php } ?>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</div>

<?php include_once("../includes/footer.php"); ?>