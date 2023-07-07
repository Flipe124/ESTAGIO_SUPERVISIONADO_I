// TELA PRINCIPAL

requestListAccount();

$('#button-new-account').on('click', function () {
    $('#modal-create-account').modal('show');
    $('.error').text("");
    $("#form-create-account")[0].reset();

    var meuInput = document.getElementById('create-input-balance-account');

    formatValueInput(meuInput);
});

$(document).on('click', '.button-update-account', function () {
    $('.error').text("");
    buttonOpenUpdateAccountModal("update", $(this).data("id"), $(this).data("name-account"), $(this).data("balance-account"));
});

$(document).on('click', '.button-delete-account', function () {
    buttonOpenDeleteAccountModal("delete-account", $(this).data("id"), $(this).data("name-account"), $(this).data("balance-account"));
});

// DENTRO DO MODAL

$('#button-create-account').on('click', function () {
    if (validationFormAccount("create") == true) {
        resquestCreateAccount();
    }
});

$('#button-update-account').on('click', function () {
    if (validationFormAccount("update") == true) {
        resquestUpdateAccount();
    }
});

$('#button-confirm-delete-account').on('click', function () {
    requestDeleteAccount();
});

// FUNCTIONS

function fillTableAccount(id, bank, balance) {

    console.log(id)
    console.log(bank)
    console.log(balance)

    register =
        `
        <tr class="result-table-account text-center">
            <td>${bank}</td>
            <td>${formatValueNumber(balance)}</td>
            <td class="text-center">
                <button class="btn btn-danger button-delete-account" type="button" data-id="${id}" data-name-account="${bank}" data-balance-account="${formatValueNumber(balance)}"><i class="fa-solid fa-trash"></i></button>
                <button class="btn btn-primary button-update-account" type="button" data-id="${id}" data-name-account="${bank}" data-balance-account="${formatValueNumber(balance)}"><i class="fa-solid fa-pen"></i></button>
            </td>
        </tr>
        `

    $("#table-account tbody").append(register);
}

function buttonOpenUpdateAccountModal(modalForm, id, name, value) {
    $(`#modal-${modalForm}-account`).modal("show");
    $(`#form-${modalForm}-account`)[0].reset();

    $('#update-input-id-account').val(id);

    console.log($('#update-input-id-account').val())

    $(`#form-${modalForm}-account #${modalForm}-input-name-account`).val(name);
    $(`#form-${modalForm}-account #${modalForm}-input-balance-account`).val(formatValueFromData(value));

};

function buttonOpenDeleteAccountModal(modal, id, name, value) {
    $(`#modal-${modal}`).modal("show");

    $('#delete-id-account').val(id);
    $(`#text-name-account`).text(name);
    $(`#text-balance-account`).text(formatValueFromData(value));

};

function formatValue(input) {
    var value = input.value.replace(/\D/g, '');

    value = (value / 100).toFixed(2);

    value = value.replace(/\B(?=(\d{3})+(?!\d))/g, '.');
    value = value.replace('.', ',');

    input.value = 'R$ ' + value;
}

function formatValueFromData(value) {
    if (typeof value !== 'string') {
        value = value.toString();
    }

    value = value.replace(/\D/g, '');

    value = (value / 100).toFixed(2);

    value = value.replace(/\B(?=(\d{3})+(?!\d))/g, '.');
    value = value.replace('.', ',');

    return 'R$ ' + value;
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

function formatValueOniput(input) {
    var value = input.value.replace(/\D/g, '');

    value = (value / 100).toFixed(2);

    value = value.replace(/\B(?=(\d{3})+(?!\d))/g, '.');
    value = value.replace('.', ',');

    input.value = 'R$ ' + value;
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


function validationFormAccount(form) {
    const MAX_LENGHT_NAME = 30;
    const MAX_LENGHT_BALANCE = 20;

    const ERROR_EMPTY_NAME = "Informe o nome da conta!";
    const ERROR_EMPTY_BALANCE = "Informe o saldo da conta!";

    const ERROR_MAX_LENGHT_NAME = `Nome pode conter até ${MAX_LENGHT_NAME} caracteres!`;
    const ERROR_MAX_LENGHT_BALANCE = `Saldo pode conter  até ${MAX_LENGHT_BALANCE} algarismos!`;

    let name = $(`#form-${form}-account #${form}-input-name-account`).val();
    let balance = $(`#form-${form}-account #${form}-input-balance-account`).val();

    let isValid = true;

    if (name == "") {
        $(`#form-${form}-account .error-name-account`).text(ERROR_EMPTY_NAME);
        isValid = false;

    } else if (name != "" && name.length > MAX_LENGHT_NAME) {
        $(`#form-${form}-account .error-name-account`).text(ERROR_MAX_LENGHT_NAME);
        isValid = false;

    } else {
        $(`#form-${form}-account .error-name-account`).text("");
    }

    if (balance == "R$ 0,00" || balance == "") {
        $(`#form-${form}-account .error-balance-account`).text(ERROR_EMPTY_BALANCE);
        isValid = false;

    } else if (balance != "" && balance.length > MAX_LENGHT_BALANCE) {
        $(`#form-${form}-account .error-balance-account`).text(ERROR_MAX_LENGHT_BALANCE);
        isValid = false;

    } else {
        $(`#form-${form}-account .error-balance-account`).text("");
    }

    return isValid;

}

$(document).ready(function () {
    var tabela = $('#table-account').DataTable();

    $('#table-account_length').hide();
    $('#table-account_filter').hide();
    $('#table-account_info').hide();
    $('#table-account_paginate').hide();

    tabela.on('order.dt', function () {
        atualizarOrdenacaoSelecionada(tabela);
    });

    $('.order-by').click(function () {
        $('.order-by').removeClass('selected');
        $(this).addClass('selected');
    });

    function atualizarOrdenacaoSelecionada(tabela) {
        var colunaOrdenada = tabela.order()[0][0];
        $('.order-by').removeClass('selected');
        $('.order-by').eq(colunaOrdenada).addClass('selected');

        $('.order-by i').removeClass('fa-sort-up fa-sort-down').addClass('fa-sort');
        var colunaSelecionada = $('.order-by').eq(colunaOrdenada);
        var icon = colunaSelecionada.find('i');
        if (tabela.order()[0][1] === 'desc') {
            icon.removeClass('fa-sort').addClass('fa-sort-up');
        } else if (tabela.order()[0][1] === 'asc') {
            icon.removeClass('fa-sort').addClass('fa-sort-down');
        }
    }
});

function ordenarTabela(coluna) {
    var tabela = $('#table-account').DataTable();
    tabela.order([coluna, tabela.order()[0][1]]).draw();
}

function showModalMessage(backgroundTitle, title, message, code) {
    if (code != 409) {
        $(".modal").modal("hide");
        $("#modal-message").modal("show");
    }

    const ERROR_BAD_REQUEST = `Requisição inválida, se o erro persistir contate o suporte!`;
    const ERROR_UNAUTHORIZED = `Seu token de acesso expirou, faça o login novamente!`;
    const ERROR_UNPROCESSABLE_ENTITY = `Erro de entidade improcessável, se o erro persisitir contate o suporte!`;
    const ERROR_INTERNAL_SERVER = `Erro interno do servidor, se o erro persistir contate o suporte!`;

    const ERROR_CONFLIT = `Nome de conta já utilizado!`;

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
        $(".error-name-account").text(ERROR_CONFLIT);

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

function resquestCreateAccount() {
    // disabledButton($('#button-create'), true);

    var accessToken = sessionStorage.getItem('accessToken');
    var objeto = JSON.parse(accessToken);

    token = objeto.token;
    
    var balance = $("#create-input-balance-account").val();

    balance = balance.replace(/[^0-9]/g, "");

    balance = parseFloat((parseFloat(balance) / 100).toFixed(2));

    var name = $("#create-input-name-account").val();

    var connect_success = true;

    var xhr = new XMLHttpRequest();

    xhr.open('POST', 'http://localhost:9999/api/v0/account/');// ALTERAR

    xhr.setRequestHeader('Token', `Bearer ${token}`);

    xhr.onload = function () {
        if (xhr.status === 200 || xhr.status === 201) {
            // disabledButton($('#button-create'), false);

            showModalMessage("bg-success", "NOVA CONTA", `Conta cadastrada com sucesso!`, 0);

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
        "balance": balance,
        "name": name,
    }

    var json = JSON.stringify(data);

    xhr.send(json);

    return connect_success
};

function resquestUpdateAccount() {
    // disabledButton($('#button-update-revenue'), true);

    var accessToken = sessionStorage.getItem('accessToken');
    var objeto = JSON.parse(accessToken);
    token = objeto.token;
    
    var balance = $("#update-input-balance-account").val();

    balance = balance.replace(/[^0-9]/g, "");

    balance = parseFloat((parseFloat(balance) / 100).toFixed(2));

    var id = $("#update-input-id-account").val();
    var name = $("#update-input-name-account").val();

    var connect_success = true;

    var xhr = new XMLHttpRequest();

    xhr.open('PATCH', `http://localhost:9999/api/v0/account/${id}`);

    xhr.setRequestHeader('Token', `Bearer ${token}`);

    xhr.onload = function () {
        if (xhr.status === 200 || xhr.status === 201 || xhr.status === 204) {
            // disabledButton($('#button-update-revenue'), false);

            showModalMessage("bg-success", "EDITAR CONTA", `Conta editada com sucesso!`, 0);

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
        "id": id,
        "name": name,
        "balance": balance
    }

    var json = JSON.stringify(data);

    xhr.send(json);

    return connect_success
};

function requestDeleteAccount() {
    // disabledButton($('#button-confirm-delete'), true);

    var accessToken = sessionStorage.getItem('accessToken');
    var objeto = JSON.parse(accessToken);
    token = objeto.token;

    var id = $('#delete-id-account').val();

    var connect_success = true;

    var xhr = new XMLHttpRequest();

    xhr.open('DELETE', `http://localhost:9999/api/v0/account/${id}`);

    xhr.setRequestHeader('Token', `Bearer ${token}`);

    xhr.onload = function () {
        if (xhr.status === 200 || xhr.status === 204) {
            // disabledButton($('#button-confirm-delete'), false);

            showModalMessage("bg-success", "EXCLUIR RECEITA", `Receita excluida com sucesso!`, 0);

        } else {
            // disabledButton($('#button-confirm-delete'), false);

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
                fillTableAccount(resposta[i].id, resposta[i].name, resposta[i].balance);
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