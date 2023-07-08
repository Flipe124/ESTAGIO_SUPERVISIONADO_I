generateTableOperation();
generateTableOperation();

requestListRevenue();

blockInsertDateManual();

// sumRevenueAndFormated();

$("#btn-open-modal-revenue").on("click", function () {
    modalAction("new-revenue", "show")
})

$("#button-new-revenue").on("click", function () {
    modalAction("create", "show");
    requestListCategory();
    requestListAccount();

    var meuInput = document.getElementById('create-input-value-operation');

    formatValueInput(meuInput);
});

$(".result").on("click", function () {
    $("#form-update")[0].reset();
    modalAction("update", "show");
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
        showModalMessage("bg-success", "NOVA RECEITA", `Receita cadastrada com sucesso!`, 0);
        // resquestCreateRevenue();
    }
});

$("#button-update-revenue").on("click", function () {
    if (validationField("update") == true) {
        showModalMessage("bg-success", "EDITAR RECEITA", `Receita editada com sucesso!`, 0);
        // resquestCreateRevenue();
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

    // if (status != "" || status != PAY_STATUS) {
    //     $(".error-msg-status-operation").text(ERROR_EMPTY_STATUS);
    //     isValid = false;

    // } else {
    //     $(".error-msg-status-operation").text();
    // }

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

    // id = 1;
    type = type_id;
    value = revenue;
    statusOp = status_id;
    date = datetime;
    data = description;
    categoryName = category_id;
    account = account_id;

    console.log("----------------")
    console.log(id)
    console.log(value)
    console.log(type)
    console.log(statusOp)
    console.log(date)
    console.log(data)
    console.log(categoryName)
    console.log(account)

    text_type = "";
    text_type_operation = "";

    iconCategory = setIconCategory(2);

    if (type == "receita") {
        text_type_operation = "text-success";

    } else {
        text_type_operation = "text-danger";
    }

    if (statusOp != "OK") {
        text_type = "text-danger";
    } else {
        text_type = "text-success"
    }

    result.innerHTML +=
        `<div class="result filter-preset-1" data-id="${id}" data-type="${type}" data-value="${value}" data-status="${statusOp}" data-date="${date}" data-description="${data}" data-category="${categoryName}" data-account="${account}" >
            <span class="icon-category">
                ${iconCategory}
            </span>
            <span class="description">
                <span class="category">
                    <b>${categoryName}</b>
                </span>
                <div class="data text-secondary">
                    ${data} | ${account}
                </div>
                <span class="font-size-14 text-secondary">${date}</span>
            </span>
            <div class="value ${text_type_operation}">
                <b class="text-value">${value}</b> <span class="status mb-1 ms-1"><i class="fas fa-check-circle ${text_type}"></i></span>
            </div>
        </div>`

    sumRevenueAndExpenseAndFormat();
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

function sumRevenueAndExpenseAndFormat() {
    var sumRevenue = 0;
    var sumExpense = 0;

    $('.result').each(function () {
        var value = parseFloat($(this).attr('data-value'));
        var type = $(this).attr('data-type');

        if (!isNaN(value)) {
            if (type === 'revenue') {
                sumRevenue += value;
            } else if (type === 'expense') {
                sumExpense += value;
            }
        }
    });

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

function formatValueNumber(number) {
    // Verifica se o número é um float
    if (Number.isFinite(number) && Number(number) % 1 !== 0) {
        // Formata o número de ponto flutuante como valor monetário
        return number.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' });
    }

    // Converte o número em uma string
    var stringNumber = number.toString();

    // Verifica se o número já possui formatação com vírgula
    if (stringNumber.indexOf('.') !== -1 || stringNumber.indexOf(',') !== -1) {
        return stringNumber; // Retorna o número original
    }

    // Adiciona o separador de milhares e retorna a string formatada como valor monetário
    return stringNumber.replace(/\B(?=(\d{3})+(?!\d))/g, '.').concat(',00');
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

    console.log(date)

    if (status == "OK") {
        $("#update-input-status-operation").prop("checked", true);
    } else {
        $("#update-input-status-operation").prop("checked", false);
    }

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

function fillSelectCategory(id, name, icon) {

    categoria =
        `
        <option value="${id}">${name}</option>
        `

    $("#create-input-category-operation").append(categoria);
}

function fillSelectAccount(id, bank, balance) {

    conta =
        `
        <option value="${id}">${bank}</option>
        `

    $("#create-input-account-operation").append(conta);
}

// REQUEST

function resquestCreateRevenue() {
    disabledButton($('#button-create'), true);

    var accessToken = sessionStorage.getItem('accessToken');
    var objeto = JSON.parse(accessToken);
    token = objeto.token;

    var value = $("#create-input-value-operation").val();
    var status = $("#create-input-status-operation").val();
    var date = $("#create-input-date-operation").val();
    var description = $("#create-input-description-operation").val();
    var category = $("#create-input-category-operation").val();
    var account = $("#create-input-account-operation").val();

    var connect_success = true;

    var xhr = new XMLHttpRequest();

    xhr.open('POST', 'http://localhost:9999/api/v0/finance/');// ALTERAR

    xhr.setRequestHeader('Token', `Bearer ${token}`);

    xhr.onload = function () {
        if (xhr.status === 200 || xhr.status === 201) {
            disabledButton($('#button-create'), false);

            showModalMessage("bg-success", "NOVA RECEITA", `Receita cadastrada com sucesso!`, 0);

        } else {
            disabledButton($('#button-create'), false);

            connect_success = false;

            var objMessage = JSON.parse(xhr.responseText);

            var code = objMessage.code;
            var msg = objMessage.error;

            showModalMessage("bg-danger", "ERROR", msg, code);

            return connect_success
        }
    };

    var data = { // ALTERAR
        "email": email,
        "name": name,
        "password": password,
        "username": username
    }

    var json = JSON.stringify(data);

    xhr.send(json);

    return connect_success
};

function resquestUpdateRevenue() {
    disabledButton($('#button-update-revenue'), true);

    var accessToken = sessionStorage.getItem('accessToken');
    var objeto = JSON.parse(accessToken);
    token = objeto.token;

    var value = $("#create-input-value-operation").val();
    var status = $("#create-input-status-operation").val();
    var date = $("#create-input-date-operation").val();
    var description = $("#create-input-description-operation").val();
    var category = $("#create-input-category-operation").val();
    var account = $("#create-input-account-operation").val();

    var connect_success = true;

    var xhr = new XMLHttpRequest();

    xhr.open('PATCH', 'http://localhost:9999/api/v0/user/');// ALTERAR

    xhr.setRequestHeader('Token', `Bearer ${token}`);

    xhr.onload = function () {
        if (xhr.status === 200 || xhr.status === 201) {
            disabledButton($('#button-update-revenue'), false);

            showModalMessage("bg-success", "EDITAR RECEITA", `Receita editada com sucesso!`, 0);

        } else {
            disabledButton($('#button-update-revenue'), false);

            connect_success = false;

            var objMessage = JSON.parse(xhr.responseText);

            var code = objMessage.code;
            var msg = objMessage.error;

            showModalMessage("bg-danger", "ERROR", msg, code);

            return connect_success
        }
    };

    var data = { // ALTERAR
        "email": email,
        "name": name,
        "password": password,
        "username": username
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

    console.log(token); // Remover na versão final

    var connect_success = true;

    var xhr = new XMLHttpRequest();

    xhr.open('GET', 'http://localhost:9999/api/v0/finance/'); // ALTERAR

    xhr.setRequestHeader('Token', `Bearer ${token}`);

    xhr.onload = function () {
        if (xhr.status === 200) {
            console.log("Status OK");
            var resposta = JSON.parse(xhr.responseText);
            console.log("RESPOSTA -> " + resposta);

            for (var i = 0; i < resposta.length; i++) {
                console.log("RESPOSTA -> " + resposta);
                // tableUserResults(resposta[i].id, resposta[i].name, resposta[i].username, resposta[i].email, true);// ALTERAR
                generateTableOperation(resposta[i].account_id, resposta[i].category_id, resposta[i].status_id, resposta[i].type_id, resposta[i].id, resposta[i].datetime, resposta[i].description, resposta[i].value)

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

function requestListCategory() {
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
                fillSelectCategory(resposta[i].id, resposta[i].name, resposta[i].icon)
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

function requestListAccount() {
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
                fillSelectAccount(resposta[i].id, resposta[i].name, resposta[i].balance)
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