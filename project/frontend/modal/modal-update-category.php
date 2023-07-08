<div class="modal fade" id="modal-update-category" tabindex="-1" aria-hidden="true">
    <div class="modal-dialog modal-dialog-centered">
        <div class="modal-content">
            <div class="modal-header bg-primary text-light">
                <h5 class="modal-title"><b>EDITAR CONTA</b></h5>
                <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
            </div>
            <div class="modal-body">
                <form action="" method="post" name="form_update_category" id="form-update-category">
                    <input type="hidden" name="update_id_category" id="update-input-id-category">
                    <div class="row">
                        <div class="col-md-12">
                            <div class="form-floating mb-3">
                                <input type="text" class="form-control" id="update-input-name-category" placeholder="Nome" maxlength="30">
                                <label for="input-name-category">Nome: <span class="text-danger">*</span></label>
                                <span class="text-danger error error-name-category"></span>
                            </div>
                        </div>
                        <!-- <div class="col-md-12">
                            <div class="form-floating mb-3">
                                <input type="text" class="form-control" id="update-input-balance-category" placeholder="Saldo atual" oninput="formatValue(this)" maxlength="16">
                                <label for="input-balance-category">Saldo atual: <span class="text-danger">*</span></label>
                                <span class="text-danger error error-balance-category"></span>
                            </div>
                        </div> -->
                        <div class="col-md-12 mt-2">
                            <label class="form-label text-danger">Campos obrigatório *</label>
                        </div>
                    </div>
                    <div class="modal-footer mt-3" id="footer-modal-update-category">
                        <button type="button" class="btn btn-outline-danger" data-bs-dismiss="modal">CANCELAR</button>
                        <button type="button" class="btn btn-success" id="button-update-category">SALVAR</button>
                    </div>
                </form>
            </div>
        </div>
    </div>
</div>