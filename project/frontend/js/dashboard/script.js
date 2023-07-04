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


