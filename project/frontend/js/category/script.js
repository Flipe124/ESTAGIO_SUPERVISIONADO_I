requestListCategory();

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