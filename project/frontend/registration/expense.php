<?php include_once("../includes/header.php"); ?>

<?php include_once("../includes/sidebar.php"); ?>

<?php
$sql = ("SELECT * FROM `finance` WHERE `type` = 'EXPENSE' ");

$expenses = $connection->connection()->query($sql)->fetchAll(PDO::FETCH_ASSOC);

?>
<?php
function getCategory($id)
{
    $connection = new Database();

    $sqlCategory = ("SELECT * FROM `category` WHERE `id` = '$id' ");

    $categorys = $connection->connection()->query($sqlCategory)->fetchAll(PDO::FETCH_ASSOC);

    foreach ($categorys as $category) {
        echo $category['name'];
    }
}
?>

<?php
function getAccount($id)
{
    $connection = new Database();

    $sqlAccount = ("SELECT * FROM `account` WHERE `id` = '$id' ");

    $Accounts = $connection->connection()->query($sqlAccount)->fetchAll(PDO::FETCH_ASSOC);

    foreach ($Accounts as $Account) {
        echo $Account['name'];
    }
}
?>

<div class="container">
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
                            <div class="col-md-6">
                                <label class="form-label">Valor:</label>
                                <input class="form-control" type="text" id="value-expense" placeholder="0,00">
                            </div>
                            <div class="col-md-6">
                                <label class="form-label">Data:</label>
                                <input class="form-control" type="date" id="date-expense">
                            </div>
                            <div class="col-md-12 mt-1">
                                <label class="form-label">Descrição:</label>
                                <input class="form-control" type="text" placeholder="Descreva aqui...">
                            </div>
                            <div class="col-md-12 mt-1">
                                <label class="form-label">Categoria:</label>
                                <div class="input-group">
                                    <select class="form-select" id="">
                                        <option value="1">Casa</option>
                                        <option value="1">Serviço</option>
                                        <option value="1">Supermercado</option>
                                    </select>
                                    <button class="btn btn-success" type="button"><i class="fa-solid fa-plus"></i></button>
                                </div>
                            </div>
                        </div>
                    </form>
                </div>
                <div class="modal-footer">
                    <button class="btn btn-danger btn-close-modal-expense" type="button">FECHAR</button>
                    <button class="btn btn-success">SALVAR</button>
                </div>
            </div>
        </div>
    </div>
    <div class="row ms-4">
        <div class="col-md-6 mt-4">
            <h1>Despesas</h1>
        </div>
        <div class="col-md-6 mt-4 text-end">
            <button class="mt-3 btn btn-danger" id="btn-open-modal-expense" type="button">+ NOVA DESPESA</button>
        </div>
        <div class="col-md-12">
            <div class="line-red"></div>
        </div>
        <div class="col-md-12 mt-2">
            <div class="p-3 text-dark" id="expense-total">
                <h4>Total de despesas:</h4>
                <h5 class="text-danger"><b>R$ - 1.788,98</b></h5>
            </div>
        </div>
        <div class="col-md-12 mt-3 text-center">
            <div class="btn btn-danger">
                <button class="btn btn-danger"><</button>
                <button class="btn btn-danger"><b>Novembro</b> 2022</button>
                <button class="btn btn-danger">></button>
            </div>
        </div>
            <div class="col-md-12 mt-3 text-center">
                <div id="div-table">
                    <table class="table table-light table-striped">
                        <thead>
                            <tr>
                                <th>Valor</th>
                                <th>Data</th>
                                <th>Descrição</th>
                                <th>Categoria</th>
                                <th>Conta</th>
                                <th>Ações</th>
                            </tr>
                        </thead>
                        <tbody>
                            <?php foreach ($expenses as $expense) { ?>
                                <tr>
                                    <th class="text-danger"><?php echo "R$ - " . $expense['value'] ?></th>
                                    <th><?php echo date('d/m/Y', strtotime($expense['date'])); ?></th>
                                    <th><?php echo $expense['description'] ?></th>
                                    <th><?php echo getCategory($expense['category_id']) ?></th>
                                    <th><?php echo getAccount($expense['account_id']) ?></th>
                                    <th>
                                        <button class="btn btn-danger" type="button"><i class="fa-solid fa-trash"></i></button>
                                        <button class="btn btn-primary" type="button"><i class="fa-solid fa-pen"></i></button>
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
</div>    

<?php include_once("../includes/footer.php"); ?>