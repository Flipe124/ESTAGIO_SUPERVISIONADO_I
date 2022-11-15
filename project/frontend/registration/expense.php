<?php include_once("../includes/header.php"); ?>

<?php include_once("../includes/sidebar.php"); ?>

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
                                <select class="form-select" id="">
                                    <option value="1">Casa</option>
                                    <option value="1">Serviço</option>
                                    <option value="1">Supermercado</option>
                                </select>
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
            <button class="mt-3 btn btn-success" id="btn-open-modal-expense" type="button">+ NOVA DESPESA</button>
        </div>
        <div class="col-md-12">
            <div class="line-red"></div>
        </div>
        <div class="flow"></div>
    </div>
</div>

<?php include_once("../includes/footer.php"); ?>