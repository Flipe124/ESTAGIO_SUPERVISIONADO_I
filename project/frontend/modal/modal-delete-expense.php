<div class="modal fade" id="modal-delete" data-bs-backdrop="static" data-bs-keyboard="false" tabindex="-1" aria-labelledby="staticBackdropLabel" aria-hidden="true">
    <div class="modal-dialog modal-dialog-centered">
        <div class="modal-content">
            <div class="modal-header bg-danger text-light">
                <h5 class="modal-title"><b>EXCLUIR DESPESA</b></h5>
            </div>
            <div class="modal-body">
                <form method="post" id="form-delete">
                    <input type="hidden" name="delete_id" id="delete-id">
                    <div class="row">
                        <div class="col-md-12">
                            Tem certeza que deseja excluir está despesa?
                        </div>
                    </div>
                    <div class="modal-footer mt-3" id="footer-modal-delete">
                        <button type="button" class="btn btn-secondary" id="button-cancel-delete" data-bs-dismiss="modal">CANCELAR</button>
                        <button type="button" class="btn btn-success" id="button-confirm-delete">SIM</button>
                    </div>
                </form>
            </div>
        </div>
    </div>
</div>