<?php include_once("../includes/header.php"); ?>

<?php include_once("../includes/sidebar.php"); ?>

<?php include_once("../includes/request.php"); ?>

<?php
// Consulta das finanças de despesas
$sql = ("SELECT * FROM `finance` WHERE `type` = 'REVENUE' ORDER BY `date` DESC ");

$revenues = $connection->connection()->query($sql)->fetchAll(PDO::FETCH_ASSOC);


// Consulta das categorias
$sqlCategorys = ("SELECT * FROM `category` WHERE `type` = 'revenue' ");

$categorys = $connection->connection()->query($sqlCategorys)->fetchAll(PDO::FETCH_ASSOC);


$sqlAccount = ("SELECT * FROM `account`");

$accounts = $connection->connection()->query($sqlAccount)->fetchAll(PDO::FETCH_ASSOC);

?>

<!-- MODAL NOVA DESPESA -->
<!-- MODAL NEW revenue -->
<div class="modal fade" id="modal-new-revenue">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header bg-success text-light">
                <h1 class="modal-title fs-5" id="modal-revenue-label">NOVA RECEITA</h1>
                <button class="btn-close btn-close-modal-revenue" type="button"></button>
            </div>
            <div class="modal-body">
                <form action="" method="post">
                    <div class="row">
                        <input type="hidden" name="type" value="revenue">
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
                            <input class="form-control" name="value" type="text" id="value-revenue" placeholder="0,00">
                        </div>
                        <div class="col-md-6">
                            <label class="form-label">Data:</label>
                            <input class="form-control" name="date" type="date" id="date-revenue">
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
                        <button class="btn btn-danger btn-close-modal-revenue" type="button">FECHAR</button>
                        <button class="btn btn-success" id="btn-save-new-revenue" type="submit">SALVAR</button>
                    </div>
                </form>
            </div>
        </div>
    </div>
</div>
<!-- MODAL DE EXCLUSÂO -->
<div class="modal fade" id="modal-delete-revenue">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header bg-success text-light">
                <h5 class="modal-title">EXCLUIR RECEITA</h5>
                <button type="button" class="btn-close btn-close-modal-delete-revenue"></button>
            </div>
            <div class="modal-body">
                <h5>
                    < Casa> - < R$ - 1.000,02>
                </h5>
            </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-danger btn-close-modal-delete-revenue" data-bs-dismiss="modal">FECHAR</button>
                <button type="button" class="btn btn-success">CONFIRMAR</button>
            </div>
        </div>
    </div>
</div>

<div class="container">
    <div class="row ms-4">
        <?php
        // foreach ($revenuess as $revenue) {
        //     echo  $revenue['value'] . "<br>";
        // }
        ?>
        <div class="col-md-6 mt-4">
            <h1>Receitas</h1>
        </div>
        <div class="col-md-6 mt-4 text-end">
            <button class="mt-3 btn btn-success" id="btn-open-modal-revenue" type="button">+ NOVA DESPESA</button>
        </div>
        <div class="col-md-12">
            <div class="line-green"></div>
        </div>
        <div class="col-md-6 mt-2">
            <div class="p-3 text-dark value-total">
                <h4>Despesas pagas:</h4>
                <h5 class="text-success"><b>R$ <?php echo getSum('REVENUE', 'PAID') ?></b></h5>
            </div>
        </div>
        <div class="col-md-6 mt-2">
            <div class="p-3 text-dark value-total-peding">
                <h4>Despesas pendente:</h4>
                <h5 class="text-success"><b>R$ <?php echo getSum('REVENUE', 'NOT_PAID') ?></b></h5>
            </div>
        </div>
        <div class="col-md-12 mt-3 text-center">
            <div class="">
                <button class="btn btn-success"><i class="fa-solid fa-arrow-left"></i></button>
                <button class="btn btn-success"><b>Novembro</b> 2022</button>
                <button class="btn btn-success"><i class="fa-solid fa-arrow-right"></i></button>
            </div>
        </div>
        <div class="col-md-12 mt-3">
            <div id="div-table">
                <table class="table table-light table-striped">
                    <thead>
                        <tr class="">
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
                        <?php foreach ($revenues as $revenue) { ?>
                            <tr>
                                <?php if ($revenue["status"] == "PAID") { ?>
                                    <th class="text-center"><i class="text-success fa-solid fa-circle-check"></i></th>
                                <?php } else { ?>
                                    <th class="text-center"><i class="text-danger fa-solid fa-circle-xmark"></i></th>
                                <?php } ?>
                                <th class="text-success text-end"><?php echo "R$ " . number_format($revenue['value'], 2, ",", "."); ?></th>
                                <th class="text-center"><?php echo date('d/m/Y', strtotime($revenue['date'])); ?></th>
                                <th><?php echo $revenue['description'] ?></th>
                                <th><?php echo getCategory($revenue['category_id']) ?></th>
                                <th><?php echo getAccount($revenue['account_id']) ?></th>
                                <th class="text-center">
                                    <button class="btn btn-danger btn-delete-revenue" type="button" data-id="<?php echo $revenue['id'] ?>" data-value="<?php echo $revenue['value'] ?>" data-description="<?php echo $revenue['description'] ?>"><i class="fa-solid fa-trash"></i></button>
                                    <button class="btn btn-primary btn-update-revenue" id="BTN" type="button" data-id="<?php echo $revenue['id'] ?>" data-status="<?php echo $revenue['status'] ?>" data-value="<?php echo $revenue['id'] ?>" data-date="<?php echo $revenue['date'] ?>" data-description="<?php echo $revenue['description'] ?>" data-category="<?php echo getCategory($revenue['category_id']) ?>" data-account="<?php echo getAccount($revenue['account_id']) ?>"><i class="fa-solid fa-pen"></i></button>
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
<div class="modal fade" id="modal-update-revenue">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header bg-danger text-light">
                <h1 class="modal-title fs-5" id="modal-revenue-label">EDITAR DESPESA</h1>
                <button class="btn-close btn-close-modal-revenue" type="button"></button>
            </div>
            <div class="modal-body">
                <form action="" method="post">
                    <div class="row">
                        <input type="hidden" name="type" value="revenue">
                        <div class="col-md-6">
                            <label class="form-label">Situação:</label>
                            <div class="input-group">
                                <select class="form-select" id="status" name="status">
                                    <option value="PAID">Recebido</option>
                                    <option value="NOT_PAID">Pendente</option>
                                </select>
                            </div>
                        </div>
                        <div class="col-md-6">
                            <label class="form-label">Valor:</label>
                            <input class="form-control" name="value" type="text" id="value-revenue" placeholder="0,00">
                        </div>
                        <div class="col-md-6">
                            <label class="form-label">Data:</label>
                            <input class="form-control" name="date" type="date" id="date-revenue">
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
                        <button class="btn btn-danger btn-close-modal-update-revenue" type="button">FECHAR</button>
                        <button class="btn btn-success" id="btn-save-new-revenue" type="submit">SALVAR</button>
                    </div>
                </form>
            </div>
        </div>
    </div>
</div>

<?php include_once("../includes/footer.php"); ?>