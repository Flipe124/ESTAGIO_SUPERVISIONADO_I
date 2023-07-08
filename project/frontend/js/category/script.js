requestListAccount();

// BUTTONS
$('#button-new-category').on('click', function () {
    $('#modal-create-category').modal('show');
    $('.error').text("");
    $("#form-create-category")[0].reset();

    // var meuInput = document.getElementById('create-input-balance-account');

    // formatValueInput(meuInput);
});

$(document).on('click', '.button-update-category', function () {
    $('.error').text("");
    buttonOpenUpdateCategoryModal("update", $(this).data("id"), $(this).data("name"), $(this).data("icon"));
});

$(document).on('click', '.button-delete-category', function () {
    buttonOpenDeleteCategoryModal("delete-category", $(this).data("id"), $(this).data("name"), $(this).data("icon"));
});

// BUTTONS DO MODAL

$('#button-create-category').on('click', function () {
    if (validationFormCategory("create") == true) {
        resquestCreateCategory();
    }
});

$('#button-update-category').on('click', function () {
    if (validationFormCategory("update") == true) {
        resquestUpdateCategory();
    }
});

$('#button-confirm-delete-category').on('click', function () {
    requestDeleteCategory();
});

// OPEN MODAL

function buttonOpenUpdateCategoryModal(modalForm, id, name, icon) {
    $(`#modal-${modalForm}-category`).modal("show");
    $(`#form-${modalForm}-category`)[0].reset();

    $('#update-input-id-category').val(id);

    // console.log($('#update-input-id-category').val())

    $(`#form-${modalForm}-category #${modalForm}-input-name-category`).val(name);
    $(`#form-${modalForm}-category #${modalForm}-input-icon-category`).val(icon);

};

function buttonOpenDeleteCategoryModal(modal, id, name, icon) {
    $(`#modal-${modal}`).modal("show");

    $('#delete-id-category').val(id);
    $(`#text-name-category`).text(name);
    $(`#text-icon-category`).text(icon);

};

// VALIDATION

function validationFormCategory(form) {
    const MAX_LENGTH_NAME = 30;
    const SPECIAL_CHARACTERS_REGEX = /[!@#$%^&*(),.?":{}|<>]/g;
    const WHITESPACE_REGEX = /\s/g;
    const NUMBER_REGEX = /\d/g;

    const ERROR_EMPTY_NAME = "Informe o nome da categoria!";
    const ERROR_MAX_LENGTH_NAME = `Nome pode conter até ${MAX_LENGTH_NAME} caracteres!`;
    const ERROR_SPECIAL_CHARACTERS_NAME = "Nome não pode conter caracteres especiais!";
    const ERROR_WHITESPACE_NAME = "Nome não pode conter espaços em branco!";
    const ERROR_NUMBER_NAME = "Nome não pode conter números!";

    let name = $(`#form-${form}-category #${form}-input-name-category`).val();

    let isValid = true;

    if (name.trim() === "") {
        $(`#form-${form}-category .error-name-category`).text(ERROR_EMPTY_NAME);
        isValid = false;

    } else if (name.trim().length > MAX_LENGTH_NAME) {
        $(`#form-${form}-category .error-name-category`).text(ERROR_MAX_LENGTH_NAME);
        isValid = false;

    } else if (SPECIAL_CHARACTERS_REGEX.test(name) || name.includes('_')) {
        $(`#form-${form}-category .error-name-category`).text(ERROR_SPECIAL_CHARACTERS_NAME);
        isValid = false;

    } else if (WHITESPACE_REGEX.test(name)) {
        $(`#form-${form}-category .error-name-category`).text(ERROR_WHITESPACE_NAME);
        isValid = false;

    } else if (NUMBER_REGEX.test(name)) {
        $(`#form-${form}-category .error-name-category`).text(ERROR_NUMBER_NAME);
        isValid = false;

    } else {
        $(`#form-${form}-category .error-name-category`).text("");
    }

    return isValid;
}

// SHOW MESSAGE 

function showModalMessage(backgroundTitle, title, message, code) {
    if (code != 409) {
        $(".modal").modal("hide");
        $("#modal-message").modal("show");
    }

    const ERROR_BAD_REQUEST = `Requisição inválida, se o erro persistir contate o suporte!`;
    const ERROR_UNAUTHORIZED = `Seu token de acesso expirou, faça o login novamente!`;
    const ERROR_UNPROCESSABLE_ENTITY = `Erro de entidade improcessável, se o erro persisitir contate o suporte!`;
    const ERROR_INTERNAL_SERVER = `Erro interno do servidor, se o erro persistir contate o suporte!`;

    const ERROR_CONFLIT = `Nome de categoria já utilizado!`;

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

// FUNCTION

function ordenarTabela(coluna) {
    var tabela = $('#table-category').DataTable();
    tabela.order([coluna, tabela.order()[0][1]]).draw();
}

function fillTableCategory(id, name, icon) {

    console.log(id)
    console.log(name)
    console.log(icon)

    category =
        `
        <tr class="result-table-category text-start">
            <td class="ps-3">${name}</td>
            <td class="text-center">
                <button class="btn btn-danger button-delete-category" type="button" data-id="${id}" data-name="${name}"><i class="fa-solid fa-trash"></i></button>
                <button class="btn btn-primary button-update-category" type="button" data-id="${id}" data-name="${name}"><i class="fa-solid fa-pen"></i></button>
            </td>
        </tr>
        `

    $("#table-category tbody").append(category);
}

$(document).ready(function () {
    var tabela = $('#table-category').DataTable();

    $('#table-category_length').hide();
    $('#table-category_filter').hide();
    $('#table-category_info').hide();
    $('#table-category_paginate').hide();
    $('.dataTables_empty').text("Sem registros.");

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

// REQUEST

function resquestCreateCategory() {
    // disabledButton($('#button-create'), true);

    var accessToken = sessionStorage.getItem('accessToken');
    var objeto = JSON.parse(accessToken);

    token = objeto.token;


    var name = $("#create-input-name-category").val();

    // var balance = $("#create-input-icon-category").val();

    var connect_success = true;

    var xhr = new XMLHttpRequest();

    xhr.open('POST', 'http://localhost:9999/api/v0/category/');// ALTERAR

    xhr.setRequestHeader('Token', `Bearer ${token}`);

    xhr.onload = function () {
        if (xhr.status === 200 || xhr.status === 201) {
            // disabledButton($('#button-create'), false);

            showModalMessage("bg-success", "NOVA CATEGORIA", `Categoria cadastrada com sucesso!`, 0);

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
        "name": name,
        "icon": ""
    }

    var json = JSON.stringify(data);

    xhr.send(json);

    return connect_success
};

function resquestUpdateCategory() {
    // disabledButton($('#button-update-revenue'), true);

    var accessToken = sessionStorage.getItem('accessToken');
    var objeto = JSON.parse(accessToken);
    token = objeto.token;

    var id = parseInt($("#update-input-id-category").val());

    var name = $("#update-input-name-category").val();

    var connect_success = true;

    var xhr = new XMLHttpRequest();

    xhr.open('PATCH', `http://localhost:9999/api/v0/category/${id}`);

    xhr.setRequestHeader('Token', `Bearer ${token}`);

    xhr.onload = function () {
        if (xhr.status === 200 || xhr.status === 201 || xhr.status === 204) {
            // disabledButton($('#button-update-revenue'), false);

            showModalMessage("bg-success", "EDITAR CATEGORIA", `Categoria editada com sucesso!`, 0);

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
        "icon": ""
    }

    var json = JSON.stringify(data);

    xhr.send(json);

    return connect_success
};

function requestDeleteCategory() {
    // disabledButton($('#button-confirm-delete'), true);

    var accessToken = sessionStorage.getItem('accessToken');
    var objeto = JSON.parse(accessToken);
    token = objeto.token;

    var id = $('#delete-id-category').val();

    var connect_success = true;

    var xhr = new XMLHttpRequest();

    xhr.open('DELETE', `http://localhost:9999/api/v0/category/${id}`);

    xhr.setRequestHeader('Token', `Bearer ${token}`);

    xhr.onload = function () {
        if (xhr.status === 200 || xhr.status === 204) {
            // disabledButton($('#button-confirm-delete'), false);

            showModalMessage("bg-success", "EXCLUIR CATEGORIA", `Categoria excluida com sucesso!`, 0);

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

    xhr.open('GET', 'http://localhost:9999/api/v0/category/');

    xhr.setRequestHeader('Token', `Bearer ${token}`);

    xhr.onload = function () {
        if (xhr.status === 200) {
            var resposta = JSON.parse(xhr.responseText);

            for (var i = 0; i < resposta.length; i++) {
                fillTableCategory(resposta[i].id, resposta[i].name, resposta[i].icon);
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