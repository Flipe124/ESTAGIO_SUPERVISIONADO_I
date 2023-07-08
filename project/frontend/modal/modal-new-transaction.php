<div class="modal fade" id="modal-create-transaction" tabindex="-1" aria-hidden="true">
    <div class="modal-dialog modal-dialog-centered">
        <div class="modal-content">
            <div class="modal-header bg-success text-light">
                <h5 class="modal-title"><b>NOVA TRANSFÊRENCIA</b></h5>
                <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
            </div>
            <div class="modal-body">
                <form action="" method="post" name="form_create_transaction" id="form-create-transaction">
                    <input type="hidden" name="create_id_transaction" id="create-id-transaction">
                    <div class="row">
                        <div class="col-md-12">
                            <div class="form-floating mb-3">
                                <input type="text" class="form-control" id="create-input-value-transaction" placeholder="Saldo atual" oninput="formatValueOniput(this)" maxlength="16" value="R$ 0,00">
                                <label for="input-value-transaction">Valor da tranferência: <span class="text-danger">*</span></label>
                                <span class="text-danger error error-msg-value-transaction"></span>
                            </div>
                            <div class="form-floating mb-3">
                                <select class="form-select" id="create-input-emitter-transaction" aria-label="Floating">
                                </select>
                                <label for="create-input-emitter-transaction">Conta emissora: <span class="text-danger">*</span></label>
                                <span class="text-danger error error-msg-emitter-transaction"></span>
                            </div>
                            <div class="form-floating mb-3">
                                <select class="form-select" id="create-input-beneficiary-transaction" aria-label="Floating">
                                </select>
                                <label for="create-input-beneficiary-transaction">Conta receptora: <span class="text-danger">*</span></label>
                                <span class="text-danger error error-msg-beneficiary-transaction"></span>
                            </div>
                        </div>
                        <div class="col-md-12 mt-2">
                            <label class="form-label text-danger">Campos obrigatório *</label>
                        </div>
                    </div>
                    <div class="modal-footer mt-3" id="footer-modal-create-transaction">
                        <button type="button" class="btn btn-outline-danger" data-bs-dismiss="modal">CANCELAR</button>
                        <button type="button" class="btn btn-success" id="button-create-transaction">SALVAR</button>
                    </div>
                </form>
            </div>
        </div>
    </div>
</div>