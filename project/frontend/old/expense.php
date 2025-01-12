<?php include_once("../includes/header.php"); ?>

<?php include_once("../includes/sidebar.php"); ?>

<?php include_once("../includes/request.php"); ?>

<?php
// Consulta das finanças de despesas
$sql = ("SELECT * FROM `finance` WHERE `type` = 'EXPENSE' ORDER BY `date` DESC ");

$expenses = $connection->connection()->query($sql)->fetchAll(PDO::FETCH_ASSOC);


// Consulta das categorias
$sqlCategorys = ("SELECT * FROM `category` WHERE `type` = 'EXPENSE' ");

$categorys = $connection->connection()->query($sqlCategorys)->fetchAll(PDO::FETCH_ASSOC);


$sqlAccount = ("SELECT * FROM `account`");

$accounts = $connection->connection()->query($sqlAccount)->fetchAll(PDO::FETCH_ASSOC);

?>

<!-- MODAL NOVA DESPESA -->
<!-- MODAL NEW EXPENSE -->
<div class="modal fade" id="modal-new-expense">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header bg-danger text-light">
                <h1 class="modal-title fs-5" id="modal-expense-label">NOVA DESPESA</h1>
                <button class="btn-close btn-close-modal-expense" type="button"></button>
            </div>
            <div class="modal-body">
                <form action="" method="post">
                    <div class="row">
                        <input type="hidden" name="type" value="EXPENSE">
                        <div class="col-md-6">
                            <label class="form-label">Situação:</label>
                            <div class="input-group">
                                <select class="form-select" id="" name="status" style="">
                                    <option value="PAID">Paga</option>
                                    <option value="NOT_PAID">Pendente</option>
                                </select>
                            </div>
                        </div>
                        <div class="col-md-6">
                            <label class="form-label">Valor:</label>
                            <input class="form-control" name="value" type="text" id="value-expense" placeholder="0,00">
                        </div>
                        <div class="col-md-6">
                            <label class="form-label">Data:</label>
                            <input class="form-control" name="date" type="date" id="date-expense">
                        </div>
                        <div class="col-md-6 mt-1">
                            <label class="form-label">Descrição:</label>
                            <input class="form-control" name="description" type="text" placeholder="Descreva aqui...">
                        </div>
                        <div class="col-md-6 mt-1">
                            <label class="form-label">Conta:</label>
                            <select class="form-select" id="select-sccount" name="account_id">
                                <?php foreach ($accounts as $account) { ?>
                                    <option value="<?php echo $account["id"] ?>"><?php echo $account["name"] ?></option>
                                <?php } ?>
                            </select>
                        </div>
                        <div class="col-md-6 mt-1">
                            <label class="form-label">Categoria:</label>
                            <div class="input-group">
                                <select class="form-select" id="" name="category_id" style="">
                                    <?php foreach ($categorys as $category) { ?>
                                        <option value="<?php echo $category["id"] ?>"><?php echo $category["name"] ?></option>
                                    <?php } ?>
                                </select>
                                <!-- <button class="btn btn-success" type="button"><i class="fa-solid fa-plus"></i></button> -->
                            </div>
                        </div>
                    </div>
                    <div class="modal-footer">
                        <button class="btn btn-danger btn-close-modal-expense" type="button">FECHAR</button>
                        <button class="btn btn-success" id="btn-save-new-expense" type="submit">SALVAR</button>
                    </div>
                </form>
            </div>
        </div>
    </div>
</div>
<!-- MODAL DE EXCLUSÂO -->
<div class="modal fade" id="modal-delete-expense">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header bg-danger text-light">
                <h5 class="modal-title">EXCLUIR DESPESA</h5>
                <button type="button" class="btn-close btn-close-modal-delete-expense"></button>
            </div>
            <div class="modal-body">
                <h5>
                    < Casa> - < R$ - 1.000,02>
                </h5>
            </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-danger btn-close-modal-delete-expense" data-bs-dismiss="modal">FECHAR</button>
                <button type="button" class="btn btn-success">CONFIRMAR</button>
            </div>
        </div>
    </div>
</div>

<div class="container">
    <div class="row ms-4">
        <?php
        // foreach ($expensess as $expense) {
        //     echo  $expense['value'] . "<br>";
        // }
        ?>
        <div class="col-md-6 mt-4">
            <h1>Despesas</h1>
        </div>
        <div class="col-md-6 mt-4 text-end">
            <button class="mt-3 btn btn-danger" id="btn-open-modal-expense" type="button">+ NOVA DESPESA</button>
        </div>
        <div class="col-md-12">
            <div class="line-red"></div>
        </div>
        <div class="col-md-6 mt-2">
            <div class="p-3 text-dark value-total">
                <h4>Despesas pagas:</h4>
                <h5 class="text-danger"><b>R$ - <?php echo getSum('EXPENSE', 'PAID') ?></b></h5>
            </div>
        </div>
        <div class="col-md-6 mt-2">
            <div class="p-3 text-dark value-total-peding">
                <h4>Despesas pendente:</h4>
                <h5 class="text-danger"><b>R$ - <?php echo getSum('EXPENSE', 'NOT_PAID') ?></b></h5>
            </div>
        </div>
        <div class="col-md-12 mt-3 text-center">
            <div class="">
                <button class="btn btn-danger"><i class="fa-solid fa-arrow-left"></i></button>
                <button class="btn btn-danger"><b>Novembro</b> 2022</button>
                <button class="btn btn-danger"><i class="fa-solid fa-arrow-right"></i></button>
            </div>
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
                        <?php foreach ($expenses as $expense) { ?>
                            <tr>
                                <?php if ($expense["status"] == "PAID") { ?>
                                    <th class="text-center"><i class="text-success fa-solid fa-circle-check"></i></th>
                                <?php } else { ?>
                                    <th class="text-center"><i class="text-danger fa-solid fa-circle-xmark"></i></th>
                                <?php } ?>
                                <th class="text-danger text-end"><?php echo "R$ - " . number_format($expense['value'], 2, ",", "."); ?></th>
                                <th class="text-center"><?php echo date('d/m/Y', strtotime($expense['date'])); ?></th>
                                <th><?php echo $expense['description'] ?></th>
                                <th><?php echo getCategory($expense['category_id']) ?></th>
                                <th><?php echo getAccount($expense['account_id']) ?></th>
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
    </div>
    <div class="flow"></div>
</div>

<!-- MODAL DE EDIÇÂO DE DESPESA -->
<div class="modal fade" id="modal-update-expense">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header bg-danger text-light">
                <h1 class="modal-title fs-5" id="modal-expense-label">EDITAR DESPESA</h1>
                <button class="btn-close btn-close-modal-expense" type="button"></button>
            </div>
            <div class="modal-body">
                <form action="" method="post">
                    <div class="row">
                        <input type="hidden" name="type" value="EXPENSE">
                        <div class="col-md-6">
                            <label class="form-label">Situação:</label>
                            <div class="input-group">
                                <select class="form-select" id="status" name="status">
                                    <option value="PAID">Paga</option>
                                    <option value="NOT_PAID">Pendente</option>
                                </select>
                            </div>
                        </div>
                        <div class="col-md-6">
                            <label class="form-label">Valor:</label>
                            <input class="form-control" name="value" type="text" id="value-expense" placeholder="0,00">
                        </div>
                        <div class="col-md-6">
                            <label class="form-label">Data:</label>
                            <input class="form-control" name="date" type="date" id="date-expense">
                        </div>
                        <div class="col-md-6 mt-1">
                            <label class="form-label">Descrição:</label>
                            <input class="form-control" name="description" type="text" placeholder="Descreva aqui...">
                        </div>
                        <div class="col-md-6 mt-1">
                            <label class="form-label">Conta:</label>
                            <select class="form-select" id="select-sccount" name="account_id">
                                <?php foreach ($accounts as $account) { ?>
                                    <option value="<?php echo $account["id"] ?>"><?php echo $account["name"] ?></option>
                                <?php } ?>
                            </select>
                        </div>
                        <div class="col-md-6 mt-1">
                            <label class="form-label">Categoria:</label>
                            <div class="input-group">
                                <select class="form-select" id="" name="category_id" style="">
                                    <?php foreach ($categorys as $category) { ?>
                                        <option value="<?php echo $category["id"] ?>"><?php echo $category["name"] ?></option>
                                    <?php } ?>
                                </select>
                            </div>
                        </div>
                    </div>
                    <div class="modal-footer">
                        <button class="btn btn-danger btn-close-modal-update-expense" type="button">FECHAR</button>
                        <button class="btn btn-success" id="btn-save-new-expense" type="submit">SALVAR</button>
                    </div>
                </form>
            </div>
        </div>
    </div>
</div>

<?php include_once("../includes/footer.php"); ?>