<div class="modal fade" id="modal-delete-account" tabindex="-1" aria-hidden="true">
    <div class="modal-dialog modal-dialog-centered">
        <div class="modal-content">
            <div class="modal-header bg-danger text-light">
                <h5 class="modal-title"><b>EXCLUIR CONTA</b></h5>
                <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
            </div>
            <div class="modal-body">
                <form method="post" id="form-delete-account">
                    <input type="hidden" name="delete_id_account" id="delete-id-account">
                    <div class="row">
                        <div class="col-md-12 mb-2">
                            <span>Tem certeza que deseja excluir a conta:</span>
                        </div>
                        <div class="col-md-12">
                            <b>Conta:</b>
                            <span id="text-name-account"></span>
                        </div>
                        <div class="col-md-12">
                            <b>Saldo:</b>
                            <span id="text-balance-account"></span>
                        </div>
                    </div>
                    <div class="modal-footer mt-3" id="footer-modal-delete-account">
                        <button type="button" class="btn btn-outline-danger" data-bs-dismiss="modal">CANCELAR</button>
                        <button type="button" class="btn btn-success" id="button-confirm-delete-account">SIM</button>
                    </div>
                </form>
            </div>
        </div>
    </div>
</div>