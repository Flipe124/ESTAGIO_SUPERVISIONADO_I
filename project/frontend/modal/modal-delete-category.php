<div class="modal fade" id="modal-delete-category" tabindex="-1" aria-hidden="true">
    <div class="modal-dialog modal-dialog-centered">
        <div class="modal-content">
            <div class="modal-header bg-danger text-light">
                <h5 class="modal-title"><b>EXCLUIR CATEGORIA</b></h5>
                <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
            </div>
            <div class="modal-body">
                <form method="post" id="form-delete-category">
                    <input type="hidden" name="delete_id_category" id="delete-id-category">
                    <div class="row">
                        <div class="col-md-12 mb-2">
                            <span>Tem certeza que deseja excluir a categoria:</span>
                        </div>
                        <div class="col-md-12">
                            <b>Conta:</b>
                            <span id="text-name-category"></span>
                        </div>
                        <!-- <div class="col-md-12">
                            <b>Ícone:</b>
                            <span id="text-icon"></span>
                        </div> -->
                    </div>
                    <div class="modal-footer mt-3" id="footer-modal-delete-category">
                        <button type="button" class="btn btn-outline-danger" data-bs-dismiss="modal">CANCELAR</button>
                        <button type="button" class="btn btn-success" id="button-confirm-delete-category">SIM</button>
                    </div>
                </form>
            </div>
        </div>
    </div>
</div>