requestListAccount();


$('#box-dashboard-balance').on('click', function () {
    location.replace('./account.php')
});

$('#box-dashboard-revenue').on('click', function () {
    location.replace('./revenue.php')
});

$('#box-dashboard-expense').on('click', function () {
    location.replace('./expense.php')
});

$(document).ready(function () {
    var tabela = $('#table-account-balance').DataTable();

    $('#table-account-balance_length').hide();
    $('#table-account-balance_filter').hide();
    $('#table-account-balance_info').hide();
    $('#table-account-balance_paginate').hide();

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
    var tabela = $('#table-account-balance').DataTable();
    tabela.order([coluna, tabela.order()[0][1]]).draw();
}

function fillTableAccount(id, bank, balance) {

    console.log(id)
    console.log(bank)
    console.log(balance)

    text_color = "";

    if (balance < 0) {
        text_color = "text-danger";
    }

    register =
        `
        <tr class="result-table-account text-center">
            <td class="text-start ps-3">${bank}</td>
            <td class="text-end ${text_color}">R$ ${formatValueNumber(balance)}</td>
            <td class="text-center">
            </td>
        </tr>
        `

    // $("#table-account-balancet tbody").append(register);
    $("#table-account-balance tbody").append(register);

    console.log("APOS")
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

function sumBalance(balance) {
    saldo = formatValueNumber(balance);

    $(".saldo").text(`R$ ${saldo}`);
}

// REQUEST

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

            var saldo = 0

            for (var i = 0; i < resposta.length; i++) {
                fillTableAccount(resposta[i].id, resposta[i].name, resposta[i].balance);

                saldo += resposta[i].balance;
            }

            console.log(saldo)
            sumBalance(saldo)

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


