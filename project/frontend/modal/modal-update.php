<div class="modal fade" id="modal-update" tabindex="-1" aria-hidden="true">
    <div class="modal-dialog modal-dialog-centered">
        <div class="modal-content">
            <div class="modal-header bg-primary text-light">
                <h5 class="modal-title"><b>EDITAR RECEITA</b></h5>
                <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
            </div>
            <div class="modal-body">
                <form action="" method="post" name="form_update" id="form-update">
                    <input type="hidden" name="update_id" id="update-id">
                    <input type="hidden" id="update-input-type-operation">
                    <div class="row">
                        <div class="col-md-12 mt-2">
                            <div class="form-floating mb-3">
                                <input type="text" class="form-control" id="update-input-value-operation" placeholder="Valor" oninput="formatValueOniput(this)" maxlength="20">
                                <label for="update-input-value-operation">Valor: <span class="text-danger">*</span></label>
                                <span class="text-danger error error-msg-value-operation"></span>
                            </div>
                            <div class="form-check form-switch mb-3">
                                <input class="form-check-input" type="checkbox" role="switch" id="update-input-status-operation">
                                <label class="form-check-label" for="update-input-status-operation">Recebido</label>
                                <span class="text-danger error error-msg-status-operation"></span>
                            </div>
                            <div class="form-floating mb-3">
                                <input type="date" class="form-control" id="update-input-date-operation" placeholder="Data" maxlength="4">
                                <label for="update-input-date-operation">Data: <span class="text-danger">*</span></label>
                                <span class="text-danger error error-msg-date-operation"></span>
                            </div>
                            <div class="form-floating mb-3">
                                <input type="text" class="form-control" id="update-input-description-operation" value="OK" placeholder="Descrição" maxlength="30">
                                <label for="update-input-description-operation">Descrição: <span class="text-danger">*</span></label>
                                <span class="text-danger error error-msg-description-operation"></span>
                            </div>
                            <div class="form-floating mb-3">
                                <select class="form-select" id="update-input-category-operation" aria-label="Floating">
                                </select>
                                <label for="update-input-category-operation">Categoria: <span class="text-danger">*</span></label>
                                <span class="text-danger error error-msg-category-operation"></span>
                            </div>
                            <div class="form-floating mb-3">
                                <select class="form-select" id="update-input-account-operation" aria-label="Floating">
                                </select>
                                <label for="update-input-account-operation">Conta: <span class="text-danger">*</span></label>
                                <span class="text-danger error error-msg-account-operation"></span>
                            </div>
                        </div>
                        <div class="col-md-12 mt-2">
                            <label class="form-label text-danger">Campos obrigatório *</label>
                        </div>
                    </div>
                    <div class="modal-footer mt-3" id="footer-modal-update">
                        <button type="button" class="btn btn-danger" id="button-delete-revenue" data-id="51">EXCLUIR</button>
                        <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">CANCELAR</button>
                        <button type="button" class="btn btn-success" id="button-update-revenue">SALVAR</button>
                    </div>
                </form>
            </div>
        </div>
    </div>
</div>