requestListTransaction();

$("#button-new-transaction").on("click", function () {
    $("#modal-create-transaction").modal("show");

    $('.error').text("");

    $("#form-create-transaction")[0].reset();

    requestListAccountAndFillSelect()
});

$("#button-create-transaction").on("click", function () {
    if (validationTransation() == true) {
        resquestCreateTransaction();
    }
});

// PREENCHER ELEMENTO

function createTableTransaction(beneficiary_id, beneficiary_name, emitter_id, emitter_name, id, value) {
    var result = document.querySelector('.transaction-table');

    result.innerHTML +=
        `<div class="result filter-preset-1" data-baneficiary-id="${beneficiary_id}" data-baneficiary-name="${beneficiary_name}" data-emitter-id="${emitter_id}" data-emitter-name="${emitter_name}" data-id="${id}" data-value="${value}" style="cursor:context-menu">
            <span class="icon-category text-primary">
                <i class="fa-solid fa-money-bill-transfer"></i>
            </span>
            <span class="description">
                <span class="category">
                    <b>Transferência</b>
                </span>
                <div class="data text-secondary">
                    ${emitter_name} para ${beneficiary_name}
                </div>
                <span class="font-size-14 text-secondary"></span>
            </span>
            <div class="value">
                <b class="text-value">${formatarMoeda(value)}</b> <span class="status mb-1 ms-1"></span>
            </div>
        </div>`
}

// TRATAMENTO DE ERROS

function validationTransation() {
    const ERROR_EMPTY_VALUE = "Informe o valor da tranferência!";
    const ERROR_EMPTY_EMITTER = "Informe a conta emissora!";
    const ERROR_EMPTY_BENEFICIARY = "Informe a conta receptora!";

    const ERROR_CONFLIT_ACCOUNT = "Conta emissora e receptora devem ser diferente!";

    value = $("#create-input-value-transaction").val();
    emitter = $("#create-input-emitter-transaction").val();
    beneficiary = $("#create-input-beneficiary-transaction").val();

    let isValid = true;

    if (value == "" || value == "R$ 0,00") {
        $(".error-msg-value-transaction").text(ERROR_EMPTY_VALUE);
        isValid = false;
    } else {
        $(".error-msg-value-transaction").text("")
    }

    if (emitter == "") {
        $(".error-msg-emitter-transaction").text(ERROR_EMPTY_EMITTER)
        isValid = false;
    } else {
        $(".error-msg-emitter-transaction").text("")
    }

    if (beneficiary == "") {
        $(".error-msg-beneficiary-transaction").text(ERROR_EMPTY_BENEFICIARY)
        isValid = false;

    } else {
        $(".error-msg-beneficiary-transaction").text("")
    }

    if (emitter == beneficiary) {
        $(".error-msg-beneficiary-transaction").text(ERROR_CONFLIT_ACCOUNT)
        isValid = false;
    } else {
        $(".error-msg-beneficiary-transaction").text("")
    }

    return isValid
}

function showModalMessage(backgroundTitle, title, message, code) {
    $(".modal").modal("hide");
    $("#modal-message").modal("show");

    const ERROR_BAD_REQUEST = `Requisição inválida, se o erro persistir contate o suporte!`;
    const ERROR_UNAUTHORIZED = `Seu token de acesso expirou, faça o login novamente!`;
    const ERROR_CONFLIT = `Erro de conflito, registro já existente!`;
    const ERROR_UNPROCESSABLE_ENTITY = `Erro de entidade improcessável, se o erro persisitir contate o suporte!`;
    const ERROR_INTERNAL_SERVER = `Erro interno do servidor, se o erro persistir contate o suporte!`;

    if (code == 400) {
        $(".modal-header").addClass(backgroundTitle);
        $("#modal-message .modal-title").text(title);
        $("#modal-message .message").text(ERROR_BAD_REQUEST);

        $("#modal-message .btn-success").on("click", function () {
            location.reload();
        })
    }
    else if (code == 401 && message == "Unauthorized") {
        $(".modal-header").addClass(backgroundTitle);
        $("#modal-message .modal-title").text(title);
        $("#modal-message .message").text(ERROR_UNAUTHORIZED);

        $("#modal-message .btn-success").on("click", function () {
            location.replace("../login");
        })

    } else if (code == 409) {
        $(".modal-header").addClass(backgroundTitle);
        $("#modal-message .modal-title").text(title);
        $("#modal-message .message").text(ERROR_CONFLIT);

        $("#modal-message .btn-success").on("click", function () {
            location.reload();
        })

    } else if (code == 422) {
        $(".modal-header").addClass(backgroundTitle);
        $("#modal-message .modal-title").text(title);
        $("#modal-message .message").text(ERROR_UNPROCESSABLE_ENTITY);

        $("#modal-message .btn-success").on("click", function () {
            location.reload();
        })
    } else if (code == 500) {
        $(".modal-header").addClass(backgroundTitle);
        $("#modal-message .modal-title").text(title);
        $("#modal-message .message").text(ERROR_INTERNAL_SERVER);

        $("#modal-message .btn-success").on("click", function () {
            location.reload();
        })

    } else {
        $(".modal-header").addClass(backgroundTitle);
        $("#modal-message .modal-title").text(title);
        $("#modal-message .message").text(message);

        $("#modal-message .btn-success").on("click", function () {
            location.reload();
        })
    }
};

// FORMATAÇÃO

function formatValueOniput(input) {
    var value = input.value.replace(/\D/g, '');

    value = (value / 100).toFixed(2);

    value = value.replace(/\B(?=(\d{3})+(?!\d))/g, '.');
    value = value.replace('.', ',');

    input.value = 'R$ ' + value;
}

function formatarMoeda(valor) {
    var formatter = new Intl.NumberFormat('pt-BR', {
        style: 'currency',
        currency: 'BRL'
    });

    return formatter.format(valor);
}

// REQUEST 

function resquestCreateTransaction() {
    // disabledButton($('#button-create'), true);

    var accessToken = sessionStorage.getItem('accessToken');
    var objeto = JSON.parse(accessToken);

    token = objeto.token;

    var value = $("#create-input-value-transaction").val();

    value = value.replace(/[^0-9]/g, "");

    value = parseFloat((parseFloat(value) / 100).toFixed(2));

    var beneficiary_id = parseInt($("#create-input-beneficiary-transaction").val());
    var emitter_id = parseInt($("#create-input-emitter-transaction").val());

    var connect_success = true;

    var xhr = new XMLHttpRequest();

    xhr.open('POST', 'http://localhost:9999/api/v0/transaction/');

    xhr.setRequestHeader('Token', `Bearer ${token}`);

    xhr.onload = function () {
        if (xhr.status === 200 || xhr.status === 201) {
            // disabledButton($('#button-create'), false);

            showModalMessage("bg-success", "NOVA TRANSFERÊNCIA", `Transferência cadastrada com sucesso!`, 0);

        } else {
            // disabledButton($('#button-create'), false);

            connect_success = false;

            var objMessage = JSON.parse(xhr.responseText);

            var code = objMessage.code;
            var msg = objMessage.error;

            showModalMessage("bg-danger", "ERROR", msg, code);

            return connect_success
        }
    };

    var data = {
        "beneficiary_id": beneficiary_id,
        "emitter_id": emitter_id,
        "value": value
    }

    var json = JSON.stringify(data);

    xhr.send(json);

    return connect_success
};

function requestListTransaction() {
    var accessToken = sessionStorage.getItem('accessToken');
    var objeto = JSON.parse(accessToken);
    token = objeto.token;

    var connect_success = true;

    var xhr = new XMLHttpRequest();

    xhr.open('GET', 'http://localhost:9999/api/v0/transaction/');

    xhr.setRequestHeader('Token', `Bearer ${token}`);

    xhr.onload = function () {
        if (xhr.status === 200) {
            $(".text-empty-transaction").text("");
            var resposta = JSON.parse(xhr.responseText);

            for (var i = 0; i < resposta.length; i++) {
                var id_emitter = resposta[i].emitter_id;
                var id_beneficiary = resposta[i].beneficiary_id;
                var id = resposta[i].id;
                var value = resposta[i].value;
                
                requestNameAccount(id_emitter, function (accountNameEmitter) {
                    requestNameAccount(id_beneficiary, function (accountNameBenificiary) {
                        createTableTransaction(id_beneficiary, accountNameBenificiary, id_emitter, accountNameEmitter, id, value);
                    });
                });
            
            }

        } else if (xhr.status === 204) {
            $(".text-empty-transaction").text("Sem transferências realizadas!");

        } else {
            connect_success = false;

            var objMessage = JSON.parse(xhr.responseText);

            var code = objMessage.code;
            var msg = objMessage.error;

            showModalMessage("bg-danger", "ERRO", msg, code);

            return connect_success
        }
    };

    xhr.send();
};

function requestListAccountAndFillSelect() {
    var accessToken = sessionStorage.getItem('accessToken');
    var objeto = JSON.parse(accessToken);
    token = objeto.token;

    var connect_success = true;

    var xhr = new XMLHttpRequest();

    xhr.open('GET', 'http://localhost:9999/api/v0/account/');

    xhr.setRequestHeader('Token', `Bearer ${token}`);

    xhr.onload = function () {
        if (xhr.status === 200) {
            var resposta = JSON.parse(xhr.responseText);

            var result = document.querySelector('#create-input-emitter-transaction');

            var resultBenificiary = document.querySelector('#create-input-beneficiary-transaction');

            result.innerHTML = "";
            resultBenificiary.innerHTML = "";

            for (var i = 0; i < resposta.length; i++) {
                selected = "";

                result.innerHTML +=
                    `
                    <option value="${resposta[i].id}">${resposta[i].name}</option>
                    `

                if (i == 1) {
                    selected = "selected";
                }

                resultBenificiary.innerHTML +=
                    `
                    <option value="${resposta[i].id}" ${selected}>${resposta[i].name}</option>
                    `

            }

        } else if (xhr.status === 204) {
            console.log("Sem contas registradas!");

        } else {
            connect_success = false;

            var objMessage = JSON.parse(xhr.responseText);

            var code = objMessage.code;
            var msg = objMessage.error;

            showModalMessage("bg-danger", "ERRO", msg, code);

            return connect_success
        }
    };

    xhr.send();
};

function requestNameAccount(id, callback) {
    var accessToken = sessionStorage.getItem('accessToken');
    var objeto = JSON.parse(accessToken);
    token = objeto.token;

    var connect_success = true;

    var xhr = new XMLHttpRequest();

    xhr.open('GET', `http://localhost:9999/api/v0/account/${id}`);

    xhr.setRequestHeader('Token', `Bearer ${token}`);

    xhr.onload = function () {
        if (xhr.status === 200 || xhr.status === 201) {
            var resposta = JSON.parse(xhr.responseText);
            var name = resposta.name;

            callback(name);

        } else if (xhr.status === 204) {
            console.log("vazio");
            callback(null);

        } else {
            connect_success = false;
            callback(connect_success);
        }
    };

    xhr.send();
}