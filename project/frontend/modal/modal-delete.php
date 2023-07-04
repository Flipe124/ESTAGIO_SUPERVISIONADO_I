<div class="modal fade" id="modal-delete" tabindex="-1" aria-hidden="true">
    <div class="modal-dialog modal-dialog-centered">
        <div class="modal-content">
            <div class="modal-header bg-danger text-light">
                <h5 class="modal-title"><b>EXCLUIR RECEITA</b></h5>
                <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
            </div>
            <div class="modal-body">
                <form method="post" id="form-delete">
                    <input type="hidden" name="delete_id" id="delete-id">
                    <div class="row">
                        <div class="col-md-12">
                            Tem certeza que deseja excluir está receita?
                        </div>
                    </div>
                    <div class="modal-footer mt-3" id="footer-modal-delete">
                        <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">CANCELAR</button>
                        <button type="button" class="btn btn-success" id="button-confirm-delete">SIM</button>
                    </div>
                </form>
            </div>
        </div>
    </div>
</div>