generateTableOperation();

blockInsertDateManual();

// sumRevenueAndFormated();

$("#btn-open-modal-revenue").on("click", function () {
    modalAction("new-revenue", "show")
})

$("#button-new-revenue").on("click", function () {
    modalAction("create", "show")
    var meuInput = document.getElementById('create-input-value-operation');

    formatValue(meuInput);
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
    if (validationField() == true) {
        resquestCreateRevenue();
    }
});

//------------------ Funções ------------------

// Apresentar/esconder modal
function modalAction(modalName, action) {
    $().modal('hide')
    $(".error").text("")
    $("#modal-" + modalName).modal(action)
}

function validationField() {
    const PAY_STATUS = "OK";

    const ERROR_EMPTY_VALUE = "Informe o valor da receita!";
    const ERROR_EMPTY_STATUS = "Informe o status da receita!";
    const ERROR_EMPTY_DESCRIPTION = "Informe a descrição!";
    const ERROR_EMPTY_DATE = "Informe a data!";
    const ERROR_EMPTY_CATEGORY = "Informe a categoria!";
    const ERROR_EMPTY_ACCOUNT = "Informe a conta!";


    value = $("#create-input-value-operation").val();
    status = $("#create-input-status-operation").val();
    description = $("#create-input-description-operation").val();
    date = $("#create-input-date-operation").val();
    category = $("#create-input-category-operation").val();
    account = $("#create-input-account-operation").val();

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

    iconCategory = setIconCategory(2);

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
                    ${data}
                </div>
                <span class="font-size-14 text-secondary">${date}</span>
            </span>
            <div class="value text-success">
                <b class="text-value">R$ 12.300,30</b> <span class="status mb-1 ms-1"><i class="fas fa-check-circle text-danger"></i></span>
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

function formatValue(input) {
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

    var meuInput = document.getElementById('update-input-value-operation');

    formatValue(meuInput);

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
