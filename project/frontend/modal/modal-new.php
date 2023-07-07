<div class="modal fade" id="modal-create" tabindex="-1" aria-hidden="true">
    <div class="modal-dialog modal-dialog-centered">
        <div class="modal-content">
            <div class="modal-header bg-success text-light">
                <h5 class="modal-title"><b>NOVA RECEITA</b></h5>
                <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
            </div>
            <div class="modal-body">
                <form action="" method="post" name="form_create" id="form-create">
                    <input type="hidden" name="create_id" id="create-id">
                    <div class="row">
                        <div class="col-md-12 mt-2">
                            <div class="form-floating mb-3">
                                <input type="text" class="form-control" id="create-input-value-operation" placeholder="Valor" oninput="formatValue(this)" maxlength="20">
                                <label for="create-input-value-operation">Valor: <span class="text-danger">*</span></label>
                                <span class="text-danger error error-msg-value-operation"></span>
                            </div>
                            <div class="form-check form-switch form-check-reverse mb-3">
                                <label class="form-check-label" for="create-input-status-operation" value="OK">Recebido</label>
                                <input class="form-check-input" type="checkbox" role="switch" id="create-input-status-operation" checked>
                                <span class="text-danger error error-msg-status-operation"></span>
                            </div>
                            <div class="form-floating mb-3">
                                <input type="date" class="form-control" id="create-input-date-operation" placeholder="Data" maxlength="4">
                                <label for="create-input-date-operation">Data: <span class="text-danger">*</span></label>
                                <span class="text-danger error error-msg-date-operation"></span>
                            </div>
                            <div class="form-floating mb-3">
                                <input type="text" class="form-control" id="create-input-description-operation" placeholder="Descrição" maxlength="30">
                                <label for="create-input-description-operation">Descrição: <span class="text-danger">*</span></label>
                                <span class="text-danger error error-msg-description-operation"></span>
                            </div>
                            <div class="form-floating mb-3">
                                <select class="form-select" id="create-input-category-operation" aria-label="Floating">
                                    <option value="Salário">Salário</option>
                                    <option value="Investimento">Investimento</option>
                                    <option value="Presente">Presente</option>
                                </select>
                                <label for="create-input-category-operation">Categoria: <span class="text-danger">*</span></label>
                                <span class="text-danger error error-msg-category-operation"></span>
                            </div>
                            <div class="form-floating mb-3">
                                <select class="form-select" id="create-input-account-operation" aria-label="Floating">
                                    <option value="1">Nubank</option>
                                    <option value="2">Inter</option>
                                    <option value="3">Bradesco</option>
                                </select>
                                <label for="create-input-account-operation">Conta: <span class="text-danger">*</span></label>
                                <span class="text-danger error error-msg-account-operation"></span>
                            </div>
                        </div>
                        <div class="col-md-12 mt-2">
                            <label class="form-label text-danger">Campos obrigatório *</label>
                        </div>
                    </div>
                    <div class="modal-footer mt-3" id="footer-modal-create">
                        <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">CANCELAR</button>
                        <button type="button" class="btn btn-success" id="button-create">SALVAR</button>
                    </div>
                </form>
            </div>
        </div>
    </div>
</div>