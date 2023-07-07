$('#box-dashboard-revenue').on('click', function () {
    location.replace('./revenue.php');
});

$('#box-dashboard-expense').on('click', function () {
    location.replace('./expense.php');
});

requestListTransaction()

createTableTransaction(1200)
createTableTransaction(14000)

function createTableTransaction(value) {
    var result = document.querySelector('.transaction-table');

    id = 0;
    type = "";
    // value = "";
    statusOp = "";
    date = "";
    data = "";
    categoryName = "";
    account = "";

    text_type = "";
    text_type_operation = "";

    // iconCategory = setIconCategory(2);

    iconCategory = "";

    if (type == "revenue") {
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
                <b class="text-value">${formatValue(value)}</b> <span class="status mb-1 ms-1"><i class="fas fa-check-circle ${text_type}"></i></span>
            </div>
        </div>`

    sumRevenueAndExpenseAndFormat();
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

// Function 

function requestListTransaction() {
    var accessToken = sessionStorage.getItem('accessToken');
    var objeto = JSON.parse(accessToken);
    token = objeto.token;

    console.log(token); // Remover na versão final

    var connect_success = true;

    var xhr = new XMLHttpRequest();

    xhr.open('GET', 'http://localhost:9999/api/v0/transaction/'); // ALTERAR

    xhr.setRequestHeader('Token', `Bearer ${token}`);

    xhr.onload = function () {
        if (xhr.status === 200) {
            console.log("Status OK");
            var resposta = JSON.parse(xhr.responseText);

            for (var i = 0; i < resposta.length; i++) {
                // tableUserResults(resposta[i].value, resposta[i].name, resposta[i].username, resposta[i].email, true);// ALTERAR
                console.log("TESTE")
            }

        } else if (xhr.status === 204) {
            console.log("Sem transações registradas!");

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