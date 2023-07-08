<div class="modal fade" id="modal-create-category" tabindex="-1" aria-hidden="true">
    <div class="modal-dialog modal-dialog-centered">
        <div class="modal-content">
            <div class="modal-header bg-success text-light">
                <h5 class="modal-title"><b>NOVA CATEGORIA</b></h5>
                <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
            </div>
            <div class="modal-body">
                <form action="" method="post" name="form_create_category" id="form-create-category">
                    <input type="hidden" name="create_id_account" id="create-id-category">
                    <div class="row">
                        <div class="col-md-12">
                            <div class="form-floating mb-3">
                                <input type="text" class="form-control" id="create-input-name-category" placeholder="Nome" maxlength="30">
                                <label for="input-name-category">Nome: <span class="text-danger">*</span></label>
                                <span class="text-danger error error-name-category"></span>
                            </div>
                        </div>
                        <!-- <div class="col-md-12">
                            <div class="form-floating mb-3">
                                <input type="text" class="form-control" id="create-input-balance-account" placeholder="Saldo atual" oninput="formatValueOniput(this)" maxlength="16">
                                <label for="input-balance-account">Saldo atual: <span class="text-danger">*</span></label>
                                <span class="text-danger error error-balance-account"></span>
                            </div>
                        </div> -->
                        <div class="col-md-12 mt-2">
                            <label class="form-label text-danger">Campos obrigatório *</label>
                        </div>
                    </div>
                    <div class="modal-footer mt-3" id="footer-modal-create-category">
                        <button type="button" class="btn btn-outline-danger" data-bs-dismiss="modal">CANCELAR</button>
                        <button type="button" class="btn btn-success" id="button-create-category">SALVAR</button>
                    </div>
                </form>
            </div>
        </div>
    </div>
</div>