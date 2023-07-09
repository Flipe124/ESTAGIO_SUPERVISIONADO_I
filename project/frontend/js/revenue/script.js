requestListRevenue();

blockInsertDateManual();

let select = document.querySelector("#update-input-category-operation");

console.log(select)

$("#btn-open-modal-revenue").on("click", function () {
    modalAction("new-revenue", "show")
})

$("#button-new-revenue").on("click", function () {
    modalAction("create", "show");
    $(`#create-input-category-operation`).html("");
    $(`#create-input-account-operation`).html("");

    requestListCategory("create");
    requestListAccount("create");

    var meuInput = document.getElementById('create-input-value-operation');

    formatValueInput(meuInput);
});

$(document).on('click', '.result', function () {
    console.log(".result")
    $("#form-update")[0].reset();
    modalAction("update", "show");

    $(`#update-input-category-operation`).html("");
    $(`#update-input-account-operation`).html("");

    requestListCategory("update");
    requestListAccount("update");

    fillModalUpdateRevenue($(this).data("id"), $(this).data("type"), $(this).data("value"), $(this).data("status"), $(this).data("date"), $(this).data("description"), $(this).data("category"), $(this).data("account"));
});


$("#button-delete-revenue").on("click", function () {
    console.log($(this).data("id"))
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
    // resquestDeleteRevenue();
    showModalMessage("bg-success", "EXCLUIR RECEITA", `Receita excluida com sucesso!`, 0);

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

    if (account == "") {
        $(".error-msg-account-operation").text(ERROR_EMPTY_ACCOUNT);
        isValid = false;

    } else {
        $(".error-msg-account-operation").text();
    }

    return isValid
}

function generateTableOperation(account_id, category_id, status_id, type_id, id, datetime, description, revenue) {
    var result = document.querySelector('.revenue-table');

    type = type_id;
    value = revenue;
    statusOp = status_id;
    date = datetime;
    data = description;
    account = account_id;

    text_type = "";
    text_type_operation = "";

    // Chama a função requestNameCategory e passa uma função de callback para receber o nome da categoria
    requestNameCategory(category_id, function (categoryName) {
        iconCategory = setIconCategory(2);

        if (type == '0') {
            text_type_operation = "text-success";
        } else {
            text_type_operation = "text-danger";
        }

        if (statusOp == 1) {
            text_type = "text-success";
        } else {
            text_type = "text-danger";
        }

        result.innerHTML +=
            `<div class="result filter-preset-1" data-id="${id}" data-type="${type}" data-value="${value}" data-status="${statusOp}" data-date="${formatData(date)}" data-description="${data}" data-category="${categoryName}" data-account="${account}" >
                <span class="icon-category">
                    ${iconCategory}
                </span>
                <span class="description">
                    <span class="category">
                        <b>${categoryName}</b>
                    </span>
                    <div class="data text-secondary">
                        ${data} |  ${account}
                    </div>
                    <span class="font-size-14 text-secondary">${formatData(date)}</span>
                </span>
                <div class="value ${text_type_operation}">
                    <b class="text-value">${formatValueMonetary(value)}</b> <span class="status mb-1 ms-1"><i class="fas fa-check-circle ${text_type}"></i></span>
                </div>
            </div>`

        sumRevenueAndExpenseAndFormat();
    });
}


function setIconCategory(category) {
    const iconSalary = `<i class="fab fa-sellcast text-success"></i>`;
    const iconCoins = `<i class="fas fa-coins text-warning"></i>`;
    const iconError = `<i class="fas fa-times text-danger"></i>`;

    if (category == 1) {
        return iconSalary;

    } else if (category == 2) {
        return iconCoins;

    } else {
        return iconError;
    }

}

function sumRevenueAndExpenseAndFormat(valorFloat) {
    var sumRevenue = 0;
    var sumExpense = 0;

    $('.result').each(function () {
        var value = parseFloat($(this).attr('data-value'));
        var type = $(this).attr('data-type');
        if (!isNaN(value)) {
            if (type === "0") {
                sumRevenue += value;
            } else if (type === "1") {
                sumExpense += value;
            }
        }
    });

    if (!isNaN(valorFloat)) {
        if (sumRevenue > sumExpense) {
            sumRevenue += valorFloat;
        } else {
            sumExpense += valorFloat;
        }
    }

    var sumRevenueFormatted = sumRevenue.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' });
    var sumExpenseFormatted = sumExpense.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' });

    var divSumRevenue = document.querySelector('.sum-revenue');
    var divSumExpense = document.querySelector('.sum-expense');

    if (divSumRevenue) {
        divSumRevenue.innerHTML = sumRevenueFormatted;
    }

    if (divSumExpense) {
        divSumExpense.innerHTML = sumExpenseFormatted;
    }
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

function formatValue(valor) {
    var valor = valor.replace(/\D/g, '');

    var valor = (valor / 100).toFixed(2);

    var partes = valor.split('.');
    var parteInteira = partes[0];
    var parteDecimal = partes[1];

    parteInteira = parteInteira.replace(/\B(?=(\d{3})+(?!\d))/g, '.');

    if (parteDecimal === '00') {
        parteDecimal = '00';
    }

    return 'R$ ' + parteInteira + ',' + parteDecimal;
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
    $("#update-input-value-operation").val(value);
    // $("#update-input-status-operation").val(status);
    $("#update-input-date-operation").val(converterFormatoData(date));
    $("#update-input-description-operation").val(description);
    $("#update-input-category-operation").val(category);
    $("#update-input-account-operation").val(account);

    if (status == 1) {
        $("#update-input-status-operation").prop("checked", true);
    } else {
        $("#update-input-status-operation").prop("checked", false);
    }

    selectCategory = document.getElementById('update-input-category-operation');
    selectAccount = document.getElementById('update-input-account-operation');

    // selecionarOpcaoPorValor(selectCategory, category);
    // selecionarOpcaoPorValor(selectAccount, account);

    selecionarOpcaoPorValor(category)

    var select = $('#update-input-category-operation'); // Selecionar o elemento select pelo seu id
    var valorString = category.toString(); // Converter o valor para string

    select.val(valorString); // Definir o valor do select como o valor fornecid

    var meuInput = document.getElementById('update-input-value-operation');

    formatValueInput(meuInput);

    selecionarCheckbox(status)

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

    // code = 400

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

function formatDataReverse(data) {
    if (typeof data !== 'string' || !/\d{4}-\d{2}-\d{2}/.test(data)) {
        return '';
    }

    var dataAtual = new Date().toLocaleDateString('pt-BR').split('/');
    var diaAtual = dataAtual[0];
    var mesAtual = dataAtual[1];
    var anoAtual = dataAtual[2];

    var partes = data.split('-');
    var ano = partes[0];
    var mes = partes[1];
    var dia = partes[2];

    if (ano > anoAtual || (ano == anoAtual && mes > mesAtual) || (ano == anoAtual && mes == mesAtual && dia > diaAtual)) {
        return '';
    }

    var horaAtual = new Date().toLocaleTimeString('pt-BR', { hour12: false });

    var dataFormatada = ano + '-' + mes + '-' + dia + ' ' + horaAtual;

    return dataFormatada;
}

function checkStatus(status) {
    if (status == "on") {
        return 1
    }
    return 0
}

function selecionarOpcaoPorValor(valor) {
    var select = $('#update-input-category-operation');
    var valorString = valor.toString();

    select.val(valorString);

    var opcaoEncontrada = select.find('option[value="' + valorString + '"]');
    // if (opcaoEncontrada.length > 0) {
    //     console.log("Valor encontrado: " + opcaoEncontrada.text());
    // } else {
    //     console.log("Valor não encontrado.");
    // }
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
    var status_id = checkStatus($("#create-input-status-operation").val())
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
        "datetime": date,
        "description": description,
        "status_code": status_id,
        "type_code": type_id,
        "value": value
    }

    console.log(data);

    var json = JSON.stringify(data);

    xhr.send(json);

    return connect_success
};

function resquestUpdateRevenue() {
    // disabledButton($('#button-update-revenue'), true);

    var accessToken = sessionStorage.getItem('accessToken');
    var objeto = JSON.parse(accessToken);
    token = objeto.token;

    id = parseInt($("#update-id").val());

    console

    var value = $("#update-input-value-operation").val();

    value = value.replace(/[^0-9]/g, "");

    value = parseFloat((parseFloat(value) / 100).toFixed(2));

    var status_id = $("#update-input-status-operation").val();
    // var date = $("#update-input-date-operation").val();
    var date = formatDataReverse($("#update-input-date-operation").val());
    var description = $("#update-input-description-operation").val();
    var category_id = $("#update-input-category-operation").val();
    var account_id = $("#update-input-account-operation").val();

    var type_id = 0;

    var connect_success = true;


    console.log("---------------")
    console.log(account_id);
    console.log(category_id);
    console.log(date);
    console.log(description);
    console.log(status_id);
    console.log(type_id);
    console.log(value);

    var xhr = new XMLHttpRequest();

    xhr.open('PATCH', `http://localhost:9999/api/v0/user/${id}`);// ALTERAR

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
        "account_id": account_id,
        "category_id": category_id,
        "datetime": date,
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

    xhr.open('DELETE', `http://localhost:9999/api/v0/user/${id}`); // ALTERAR

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
            console.log("Status OK");
            var resposta = JSON.parse(xhr.responseText);

            for (var i = 0; i < resposta.length; i++) {
                if (resposta[i].type_code == 0) {// 0 entrada
                    generateTableOperation(resposta[i].account_id, resposta[i].category_id, resposta[i].status_code, resposta[i].type_code, resposta[i].id, resposta[i].datetime, resposta[i].description, resposta[i].value)
                }
            }

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

function requestListCategory(form) {
    var accessToken = sessionStorage.getItem('accessToken');
    var objeto = JSON.parse(accessToken);
    token = objeto.token;

    console.log(token); // Remover na versão final

    var connect_success = true;

    var xhr = new XMLHttpRequest();

    xhr.open('GET', 'http://localhost:9999/api/v0/category/');

    xhr.setRequestHeader('Token', `Bearer ${token}`);

    xhr.onload = function () {
        if (xhr.status === 200) {
            var resposta = JSON.parse(xhr.responseText);

            for (var i = 0; i < resposta.length; i++) {
                fillSelectCategory(resposta[i].id, resposta[i].name, resposta[i].icon, form)
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

function requestListAccount(form) {
    var accessToken = sessionStorage.getItem('accessToken');
    var objeto = JSON.parse(accessToken);
    token = objeto.token;

    console.log(token); // Remover na versão final

    var connect_success = true;

    var xhr = new XMLHttpRequest();

    xhr.open('GET', 'http://localhost:9999/api/v0/account/');

    xhr.setRequestHeader('Token', `Bearer ${token}`);

    xhr.onload = function () {
        if (xhr.status === 200) {
            var resposta = JSON.parse(xhr.responseText);

            for (var i = 0; i < resposta.length; i++) {
                fillSelectAccount(resposta[i].id, resposta[i].name, resposta[i].balance, form)
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

    xhr.open('GET', `http://localhost:9999/api/v0/category/${id}`);

    xhr.setRequestHeader('Token', `Bearer ${token}`);

    xhr.onload = function () {
        if (xhr.status === 200 || xhr.status === 201 || xhr.status === 204) {
            var resposta = JSON.parse(xhr.responseText);
            var name = resposta.name;

            callback(name); // Chama o callback com o valor desejado

        } else if (xhr.status === 204) {
            console.log("vazio");
            callback(null); // Chama o callback com valor nulo

        } else {
            connect_success = false;
            callback(connect_success); // Chama o callback com o valor de connect_success
        }
    };

    xhr.send();
}



function requestNameAccount(id) {

    console.log("ID: " + id)

    var accessToken = sessionStorage.getItem('accessToken');
    var objeto = JSON.parse(accessToken);
    token = objeto.token;

    console.log(token); // Remover na versão final

    var connect_success = true;

    var xhr = new XMLHttpRequest();

    xhr.open('GET', `http://localhost:9999/api/v0/account/${id}`);

    xhr.setRequestHeader('Token', `Bearer ${token}`);

    xhr.onload = function () {
        if (xhr.status === 200 || xhr.status === 201) {
            var resposta = JSON.parse(xhr.responseText);
            var name = resposta.name

            for (var i = 0; i < resposta.length; i++) {
                // fillSelectCategory(resposta[i].id, resposta[i].name, resposta[i].icon, form)


            }
            return name

        } else if (xhr.status === 204) {
            console.log("vazio");

        } else {
            connect_success = false;

            // var objMessage = JSON.parse(xhr.responseText);

            // var code = objMessage.code;
            // var msg = objMessage.error;

            // showModalMessage("bg-danger", "ERRO", msg, code);

            return connect_success
        }
    };

    xhr.send();
}