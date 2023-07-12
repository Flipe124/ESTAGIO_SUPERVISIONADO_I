requestListRevenue();

blockInsertDateManual();

let select = document.querySelector("#update-input-category-operation");

$("#button-new-revenue").on("click", function () {
    modalAction("create", "show");
    $(`#create-input-category-operation`).html("");
    $(`#create-input-account-operation`).html("");

    requestListCategoryDefault("create");
    requestListAccount("create");

    var meuInput = document.getElementById('create-input-value-operation');

    formatValueInput(meuInput);
});

$(document).on('click', '.result', function () {
    $("#form-update")[0].reset();
    modalAction("update", "show");

    $(`#update-input-category-operation`).html("");
    $(`#update-input-account-operation`).html("");

    requestListCategoryDefault("update", $(this).data("category-id"));
    requestListAccount("update", $(this).data("account-id"));

    fillModalUpdateRevenue($(this).data("id"), $(this).data("type"), $(this).data("value"), $(this).data("status"), $(this).data("date"), $(this).data("description"), $(this).data("category"), $(this).data("account"));
});

$("#button-delete-revenue").on("click", function () {
    modalAction("update", "hide")
    modalAction("delete", "show")
});

$('#button-cancel-delete').on('click', function () {
    modalAction("update", "show");
});

$("#button-create").on("click", function () {
    if (validationField("create") == true) {
        resquestCreateRevenue();
    }
});

$("#button-update-revenue").on("click", function () {
    if (validationField("update") == true) {
        resquestUpdateRevenue();
    }
});

$("#button-confirm-delete").on("click", function () {
    requestDeleteRevenue();
});

//------------------ Funções ------------------

// Apresentar/esconder modal
function modalAction(modalName, action) {
    $().modal('hide')
    $(".error").text("")
    $("#modal-" + modalName).modal(action)
}

function validationField(operationType) {
    const PAY_STATUS = "OK";

    const ERROR_EMPTY_VALUE = "Informe o valor da receita!";
    const ERROR_EMPTY_STATUS = "Informe o status da receita!";
    const ERROR_EMPTY_DESCRIPTION = "Informe a descrição!";
    const ERROR_EMPTY_DATE = "Informe a data!";
    const ERROR_EMPTY_CATEGORY = "Informe a categoria!";
    const ERROR_EMPTY_ACCOUNT = "Informe a conta!";


    value = $(`#${operationType}-input-value-operation`).val();
    status = $(`#${operationType}-input-status-operation`).val();
    description = $(`#${operationType}-input-description-operation`).val();
    date = $(`#${operationType}-input-date-operation`).val();
    category = $(`#${operationType}-input-category-operation`).val();
    account = $(`#${operationType}-input-account-operation`).val();

    isValid = true;

    if (value == "" || value == "R$ 0,00") {
        $(".error-msg-value-operation").text(ERROR_EMPTY_VALUE);
        isValid = false;

    } else {
        $(".error-msg-value-operation").text();
    }

    if (description == "") {
        $(".error-msg-description-operation").text(ERROR_EMPTY_DESCRIPTION);
        isValid = false;

    } else {
        $(".error-msg-description-operation").text();
    }

    if (date == "") {
        $(".error-msg-date-operation").text(ERROR_EMPTY_DATE);
        isValid = false;

    } else {
        $(".error-msg-date-operation").text();
    }

    if (category == "") {
        $(".error-msg-category-operation").text(ERROR_EMPTY_CATEGORY);
        isValid = false;

    } else {
        $(".error-msg-category-operation").text();
    }

    if(operationType == "create"){
        if (account == "") {
            $(".error-msg-account-operation").text(ERROR_EMPTY_ACCOUNT);
            isValid = false;
    
        } else {
            $(".error-msg-account-operation").text();
        }
    }

    return isValid
}
// ANTES
function generateTableOperation(account_id, category_id, status_id, type_id, id, datetime, description, revenue) {
    var result = document.querySelector('.revenue-table');

    var id = id;
    var type = type_id;
    var value = revenue;
    var statusOp = status_id;
    var date = datetime;
    var data = description;
    var account = account_id;

    var text_type = "text-success";
    var text_type_operation = "";

    var iconCategory = setIconCategory(category_id);

    if (type == '0') {
        text_type_operation = "text-success";
    } else {
        text_type_operation = "text-danger";
    }

    if (status_id != 1) {
        text_type = "text-danger";
    }

    requestNameAccount(account_id, function (accountName) {
        requestNameCategory(category_id, function (categoryName) {
            result.innerHTML +=
                `<div class="result filter-preset-1" data-id="${id}" data-type="${type}" data-value="${value}" data-status="${statusOp}" data-date="${formatData(datetime)}" data-description="${description}" data-category="${categoryName}" data-category-id=${category_id} data-account="${account}" data-account-id=${account_id} >
                <span class="icon-category">
                    ${iconCategory}
                </span>
                <span class="description">
                    <span class="category">
                        <b>${categoryName}</b>
                    </span>
                    <div class="data text-secondary">
                        ${description} |  ${accountName}
                    </div>
                    <span class="font-size-14 text-secondary">${formatData(datetime)}</span>
                </span>
                <div class="value ${text_type_operation}">
                    <b class="text-value">${formatValueMonetary(revenue)}</b> <span class="status mb-1 ms-1"><i class="fas fa-check-circle ${text_type}"></i></span>
                </div>
            </div>`
        });
    });
}

function setIconCategory(category) {
    const iconBook = `<i class="fa-solid fa-book" style="color: #964b00"></i>`;
    const iconLeisure = `<i class="fa-solid fa-umbrella-beach text-primary"></i>`;
    const iconFood = `<i class="fa-solid fa-pizza-slice text-warning"></i>`;
    const iconHealth = `<i class="fa-solid fa-house-medical text-danger"></i>`;
    const iconOuther = `<i class="fa-solid fa-question"></i>`;
    const iconError = `<i class="fa-solid fa-bars"></i>`;

    if (category == 1) {
        return iconBook
    } else if (category == 2) {
        return iconLeisure
    } else if (category == 3) {
        return iconFood
    } else if (category == 4) {
        return iconHealth
    } else if (category == 5) {
        return iconOuther
    } else {
        return iconError;
    }
}

function formatarValorReceita(value) {
    var elemento = document.querySelector('.sum-revenue');

    if (typeof value !== 'number') {
        return '';
    }

    var parts = value.toFixed(2).split('.');
    var integerPart = parts[0];
    var decimalPart = parts[1];

    integerPart = integerPart.replace(/\B(?=(\d{3})+(?!\d))/g, '.');

    elemento.textContent = 'R$ ' + integerPart + ',' + decimalPart;
}

function formatValueInput(input) {
    var valor = input.value.replace(/\D/g, '');

    valor = (valor / 100).toFixed(2);

    var partes = valor.split('.');
    var parteInteira = partes[0];
    var parteDecimal = partes[1];

    parteInteira = parteInteira.replace(/\B(?=(\d{3})+(?!\d))/g, '.');

    if (parteDecimal === '00') {
        parteDecimal = '00';
    }

    input.value = 'R$ ' + parteInteira + ',' + parteDecimal;
}

function formatValueOniput(input) {
    var value = input.value.replace(/\D/g, '');

    value = (value / 100).toFixed(2);

    value = value.replace(/\B(?=(\d{3})+(?!\d))/g, '.');
    value = value.replace('.', ',');

    input.value = 'R$ ' + value;
}

function formatValueMonetary(value) {
    if (typeof value !== 'number' || isNaN(value)) {
    }

    return value.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' });
}

function blockInsertDateManual() {
    document.getElementById("update-input-date-operation").addEventListener("keydown", function (event) {
        event.preventDefault();
    });
    document.getElementById("create-input-date-operation").addEventListener("keydown", function (event) {
        event.preventDefault();
    });
}

function fillModalUpdateRevenue(id, type, value, status, date, description, category, account) {
    $("#update-id").val(id);
    $("#update-input-type-operation").val(type);
    $("#update-input-value-operation").val(formatValueMonetary(value));
    $("#update-input-date-operation").val(converterFormatoData(date));
    $("#update-input-description-operation").val(description);
    $("#update-input-category-operation").val(category);
    $("#update-input-account-operation").val(account);

    if (status == 1) {
        $("#update-input-status-operation").prop("checked", true);
    } else {
        $("#update-input-status-operation").prop("checked", false);
    }

    var meuInput = document.getElementById('update-input-value-operation');

    formatValueInput(meuInput);

    selecionarCheckbox(status)

    $("#delete-id").val(id)

}

function converterFormatoData(date) {
    var partes = date.split('/');
    var dia = partes[0];
    var mes = partes[1];
    var ano = partes[2];
    var dataFormatada = ano + '-' + mes + '-' + dia;

    return dataFormatada;
}

function selecionarCheckbox(value) {
    var checkbox = document.getElementById('update-input-status-operation');
    if (value === 'OK' && checkbox) {
        checkbox.checked = true;
    }
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

function disabledButton(button, disabled) {
    button.text("Carregando...");
    button.prop("disabled", disabled);
}

function fillSelectCategory(id, name, icon, form) {
    categoria =
        `
        <option value="${id}">${name}</option>
        `

    $(`#${form}-input-category-operation`).append(categoria);
}

function fillSelectAccount(id, bank, balance, form) {
    conta =
        `
        <option value="${id}">${bank}</option>
        `

    $(`#${form}-input-account-operation`).append(conta);
}

function formatData(data) {
    if (typeof data !== 'string' || !/\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}-\d{2}:\d{2}/.test(data)) {
        return '';
    }

    var dateObj = new Date(data);

    var dia = dateObj.getDate();
    var mes = dateObj.getMonth() + 1;
    var ano = dateObj.getFullYear();

    var dataFormatada = padZero(dia) + '/' + padZero(mes) + '/' + ano;

    return dataFormatada;
}

function padZero(numero) {
    return numero < 10 ? '0' + numero : numero;
}

function checkStatus(status) {
    if (status == "on") {
        return 1
    }
    return 0
}

function verificarCheckboxAtivo(checkboxId) {
    var checkbox = document.getElementById(checkboxId);

    if (checkbox.checked) {
        return 1; // Está marcado (ativo)
    } else {
        return 0; // Não está marcado (inativo)
    }
}

function formatarData(data) {
    var dataFormatada = data + " 23:59:59";
    return dataFormatada;
}

function formatarValor(valor) {
    var valorFormatado = valor.replace(/[^0-9]/g, "");
    valorFormatado = parseFloat((parseFloat(valorFormatado) / 100).toFixed(2));
    return valorFormatado;
}

// REQUEST

function resquestCreateRevenue() {
    // disabledButton($('#button-create'), true);
    var accessToken = sessionStorage.getItem('accessToken');
    var objeto = JSON.parse(accessToken);
    token = objeto.token;

    var value = formatarValor($("#create-input-value-operation").val())
    var account_id = parseInt($("#create-input-account-operation").val());
    var category_id = parseInt($("#create-input-category-operation").val());
    var date = formatarData($("#create-input-date-operation").val());
    var description = $("#create-input-description-operation").val();
    var status_id = verificarCheckboxAtivo("create-input-status-operation") // checkStatus($("#create-input-status-operation").val())
    var type_id = 0;

    var connect_success = true;

    var xhr = new XMLHttpRequest();

    xhr.open('POST', 'http://localhost:9999/api/v0/finance/');// ALTERAR

    xhr.setRequestHeader('Token', `Bearer ${token}`);

    xhr.onload = function () {
        if (xhr.status === 200 || xhr.status === 201 || xhr.status === 204) {
            // disabledButton($('#button-create'), false);

            showModalMessage("bg-success", "NOVA RECEITA", `Receita cadastrada com sucesso!`, 0);

        } else {
            // disabledButton($('#button-create'), false);

            connect_success = false;

            var objMessage;

            if (xhr.responseText) {
                objMessage = JSON.parse(xhr.responseText);
                var code = objMessage.code;
                var msg = objMessage.error;

                showModalMessage("bg-danger", "ERROR", msg, code);

            } else {
                showModalMessage("bg-danger", "ERROR", "Ocorreu um erro desconhecido.", "");
            }

            return connect_success
        }
    };

    var data = {
        "account_id": account_id,
        "category_id": category_id,
        "date_time": date,
        "description": description,
        "status_code": status_id,
        "type_code": type_id,
        "value": value
    }

    var json = JSON.stringify(data);

    xhr.send(json);

    return connect_success
};

function resquestUpdateRevenue() {
    // disabledButton($('#button-update-revenue'), true);

    var accessToken = sessionStorage.getItem('accessToken');
    var objeto = JSON.parse(accessToken);
    token = objeto.token;

    console.log(token)

    var id = parseInt($("#update-id").val());
    var value = formatarValor($("#update-input-value-operation").val());
    // var account_id = parseInt($("#update-input-account-operation").val());
    var category_id = parseInt($("#update-input-category-operation").val());
    var date = formatarData($("#update-input-date-operation").val());
    var status_id = verificarCheckboxAtivo("update-input-status-operation");
    var description = $("#update-input-description-operation").val();
    var type_id = 0;

    var connect_success = true;

    var xhr = new XMLHttpRequest();

    xhr.open('PATCH', `http://localhost:9999/api/v0/finance/${id}`);// ALTERAR

    xhr.setRequestHeader('Token', `Bearer ${token}`);

    xhr.onload = function () {
        if (xhr.status === 200 || xhr.status === 201 || xhr.status === 204) {
            // disabledButton($('#button-update-revenue'), false);

            showModalMessage("bg-success", "EDITAR RECEITA", `Receita editada com sucesso!`, 0);

        } else {
            // disabledButton($('#button-update-revenue'), false);

            connect_success = false;
            var objMessage;

            if (xhr.responseText) {
                objMessage = JSON.parse(xhr.responseText);
                var code = objMessage.code;
                var msg = objMessage.error;

                showModalMessage("bg-danger", "ERROR", msg, code);

            } else {
                showModalMessage("bg-danger", "ERROR", "Ocorreu um erro desconhecido.", "");
            }

            return connect_success
        }
    };

    var data = {
        // "account_id": account_id,
        "category_id": category_id,
        "date_time": date,
        "description": description,
        "status_code": status_id,
        "type_code": type_id,
        "value": value
    }

    var json = JSON.stringify(data);

    xhr.send(json);

    return connect_success
};

function requestDeleteRevenue() {
    disabledButton($('#button-confirm-delete'), true);

    var accessToken = sessionStorage.getItem('accessToken');
    var objeto = JSON.parse(accessToken);
    token = objeto.token;

    var id = $('#delete-id').val();

    var connect_success = true;

    var xhr = new XMLHttpRequest();

    xhr.open('DELETE', `http://localhost:9999/api/v0/finance/${id}`);

    xhr.setRequestHeader('Token', `Bearer ${token}`);

    xhr.onload = function () {
        if (xhr.status === 200 || xhr.status === 204) {
            disabledButton($('#button-confirm-delete'), false);

            showModalMessage("bg-success", "EXCLUIR RECEITA", `Receita excluida com sucesso!`, 0);

        } else {
            disabledButton($('#button-confirm-delete'), false);

            connect_success = false;

            var objMessage = JSON.parse(xhr.responseText);

            var code = objMessage.code;
            var msg = objMessage.error;

            showModalMessage("bg-danger", "ERROR", msg, code);

            return connect_success
        }
    };

    var data = {
        "id": id,
    }

    var json = JSON.stringify(data);

    xhr.send(json);

    return connect_success
};

function requestListRevenue() {
    var accessToken = sessionStorage.getItem('accessToken');
    var objeto = JSON.parse(accessToken);
    token = objeto.token;

    var connect_success = true;

    var xhr = new XMLHttpRequest();

    xhr.open('GET', 'http://localhost:9999/api/v0/finance/');

    xhr.setRequestHeader('Token', `Bearer ${token}`);

    xhr.onload = function () {
        if (xhr.status === 200) {
            var resposta = JSON.parse(xhr.responseText);
            var somaReceita = 0.00;

            for (var i = 0; i < resposta.length; i++) {
                if (resposta[i].type_code == 0) {// 0 entrada
                    generateTableOperation(resposta[i].account_id, resposta[i].category_id, resposta[i].status_code, resposta[i].type_code, resposta[i].id, resposta[i].date_time, resposta[i].description, resposta[i].value)
                }
                if (resposta[i].type_code == 0 && resposta[i].status_code == 1) {
                    somaReceita += parseFloat(resposta[i].value.toFixed(2));
                }
            }
            formatarValorReceita(somaReceita);

        } else if (xhr.status === 204) {
            console.log("Sem receitas registradas!");

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

function requestListCategoryDefault(form, category) {
    var accessToken = sessionStorage.getItem('accessToken');
    var objeto = JSON.parse(accessToken);
    token = objeto.token;

    var connect_success = true;

    var xhr = new XMLHttpRequest();

    xhr.open('GET', 'http://localhost:9999/api/v0/category/default/');

    xhr.setRequestHeader('Token', `Bearer ${token}`);

    xhr.onload = function () {
        if (xhr.status === 200) {
            var resposta = JSON.parse(xhr.responseText);

            for (var i = 0; i < resposta.length; i++) {
                fillSelectCategory(resposta[i].id, resposta[i].name, resposta[i].icon, form)
                $("#update-input-category-operation").val(category);
            }

            requestListCategory(form, category);

        } else if (xhr.status === 204) {
            console.log("Sem categorias registradas!");

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

function requestListCategory(form, category) {
    var accessToken = sessionStorage.getItem('accessToken');
    var objeto = JSON.parse(accessToken);
    token = objeto.token;

    var connect_success = true;

    var xhr = new XMLHttpRequest();

    xhr.open('GET', 'http://localhost:9999/api/v0/category/');

    xhr.setRequestHeader('Token', `Bearer ${token}`);

    xhr.onload = function () {
        if (xhr.status === 200) {
            var resposta = JSON.parse(xhr.responseText);

            for (var i = 0; i < resposta.length; i++) {
                fillSelectCategory(resposta[i].id, resposta[i].name, resposta[i].icon, form)
                $("#update-input-category-operation").val(category);
            }

        } else if (xhr.status === 204) {
            console.log("Sem categorias registradas!");

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

function requestListAccount(form, account) {
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

            for (var i = 0; i < resposta.length; i++) {
                fillSelectAccount(resposta[i].id, resposta[i].name, resposta[i].balance, form)
                $("#update-input-account-operation").val(account);
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

function requestNameCategory(id, callback) {
        var accessToken = sessionStorage.getItem('accessToken');
        var objeto = JSON.parse(accessToken);
        token = objeto.token;

        var connect_success = true;

        var xhr = new XMLHttpRequest();

        if(id >= 6) {
            xhr.open('GET', `http://localhost:9999/api/v0/category/${id}`);
        } else {
            xhr.open('GET', `http://localhost:9999/api/v0/category/default/${id}`);
        }

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