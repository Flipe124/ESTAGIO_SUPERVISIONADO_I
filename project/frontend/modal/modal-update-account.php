<div class="modal fade" id="modal-update-account" tabindex="-1" aria-hidden="true">
    <div class="modal-dialog modal-dialog-centered">
        <div class="modal-content">
            <div class="modal-header bg-primary text-light">
                <h5 class="modal-title"><b>EDITAR CONTA</b></h5>
                <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
            </div>
            <div class="modal-body">
                <form action="" method="post" name="form_update_account" id="form-update-account">
                    <input type="hidden" name="update_id_account" id="update-id-account">
                    <div class="row">
                        <div class="col-md-12">
                            <div class="form-floating mb-3">
                                <input type="text" class="form-control" id="input-name-account" placeholder="Nome" maxlength="30">
                                <label for="input-name-account">Nome: <span class="text-danger">*</span></label>
                            </div>
                        </div>
                        <div class="col-md-12">
                            <div class="form-floating mb-3">
                                <input type="text" class="form-control" id="input-balance-account" placeholder="Saldo atual" oninput="formatValue(this)" maxlength="20">
                                <label for="input-balance-account">Saldo atual: <span class="text-danger">*</span></label>
                            </div>
                        </div>
                        <div class="col-md-12 mt-2">
                            <label class="form-label text-danger">Campos obrigatório *</label>
                        </div>
                    </div>
                    <div class="modal-footer mt-3" id="footer-modal-update-account">
                        <button type="button" class="btn btn-outline-danger" data-bs-dismiss="modal">CANCELAR</button>
                        <button type="button" class="btn btn-success" id="button-update-account">SALVAR</button>
                    </div>
                </form>
            </div>
        </div>
    </div>
</div>