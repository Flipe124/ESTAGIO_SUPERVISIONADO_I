generateTableOperation();
generateTableOperation();

blockInsertDateManual();

// sumRevenueAndFormated();

$("#btn-open-modal-revenue").on("click", function () {
    modalAction("new-revenue", "show")
})

$("#button-new-revenue").on("click", function () {
    modalAction("create", "show")
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

$('#modal-delete').on('hidden.bs.modal', function (e) {
    modalAction("update", "show")
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

function generateTableOperation() {
    var result = document.querySelector('.revenue-table');

    id = 1;
    type = "revenue";
    value = "12300.30";
    statusOp = "OK";
    date = "21/07/2023";
    data = "QuantumTech";
    categoryName = "Presente";
    account = "Bradesco";

    text_type = "";

    iconCategory = setIconCategory(2);

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
            <div class="value text-success">
                <b class="text-value">${formatValue(value)}</b> <span class="status mb-1 ms-1"><i class="fas fa-check-circle ${text_type}"></i></span>
            </div>
        </div>`

    sumRevenueAndFormated();
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

function sumRevenueAndFormated() {
    var sum = 0;

    $('.result').each(function () {
        var value = parseFloat($(this).attr('data-value'));

        if (!isNaN(value)) {
            sum += value;
        }
    });

    var sumFormated = sum.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' });

    var divSum = document.querySelector('.sum-revenue');

    divSum.innerHTML = sumFormated;
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

// REQUEST

function resquestCreateRevenue() {
    var value = $("#create-input-value-operation").val();
    var status = $("#create-input-status-operation").val();
    var date = $("#create-input-date-operation").val();
    var description = $("#create-input-description-operation").val();
    var category = $("#create-input-category-operation").val();
    var account = $("#create-input-account-operation").val();

    var connect_success = true;
    var xhr = new XMLHttpRequest();

    xhr.open('POST', 'http://localhost:9999/api/v0/user/');// ALTERAR

    xhr.onload = function () {
        if (xhr.status === 200 || xhr.status === 201) {
            console.log("SUCESSO!");
        } else {
            connect_success = false;

            console.log(xhr.responseText);

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
}

function resquestUpdateRevenue() {
    var value = $("#create-input-value-operation").val();
    var status = $("#create-input-status-operation").val();
    var date = $("#create-input-date-operation").val();
    var description = $("#create-input-description-operation").val();
    var category = $("#create-input-category-operation").val();
    var account = $("#create-input-account-operation").val();

    var connect_success = true;
    var xhr = new XMLHttpRequest();

    xhr.open('PATCH', 'http://localhost:9999/api/v0/user/');// ALTERAR

    xhr.onload = function () {
        if (xhr.status === 200 || xhr.status === 201) {
            console.log("SUCESSO!");
        } else {
            connect_success = false;

            console.log(xhr.responseText);

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
}

function requestDeleteRevenue() {

    var id = $('#delete-id').val();

    var connect_success = true;

    var xhr = new XMLHttpRequest();

    xhr.open('DELETE', `http://localhost:8008/api/v2/user/${id}`); // ALTERAR

    xhr.onload = function () {
        if (xhr.status === 200 || xhr.status === 204) {
            console.log("SUCESSO!")

        } else {

            connect_success = false;

            console.log(xhr.responseText);

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
