<?php include_once("../includes/header.php"); ?>

<?php include_once("../includes/sidebar.php"); ?>

<div class="container">
    <!-- MODAL NOVA DESPESA -->
    <!-- MODAL NEW EXPENSE -->
    <div class="modal fade" id="modal-new-expense">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <h1 class="modal-title fs-5" id="modal-expense-label">Nova despesa</h1>
                    <button class="btn-close" type="button" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
                <div class="modal-body">

                </div>
                <div class="modal-footer">
                    <button class="btn btn-danger" type="button" data-bs-dismiss="modal">Fechar</button>
                    <button class="btn btn-success">Salvar</button>
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