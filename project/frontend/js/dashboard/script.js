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

function fillTableAccount(id, bank, balance) {
    var result = document.querySelector('#table-account-balance');

    text_color = "";

    if (balance < 0) {
        text_color = "text-danger";
    }

    result.innerHTML +=
        `
        <tr class="result-table-account text-center">
            <td class="text-start ps-3">${bank}</td>
            <td class="text-end ${text_color}">${formatarMoeda(balance)}</td>
            <td class="text-center">
            </td>
        </tr>
        `
}

function sumBalance(balance) {
    saldo = formatarMoeda(balance);

    $(".saldo").text(`R$ ${saldo}`);
}

function formatarMoeda(valor) {
    var formatter = new Intl.NumberFormat('pt-BR', {
        style: 'currency',
        currency: 'BRL'
    });

    return formatter.format(valor);
}

// REQUEST

function requestListAccount() {
    var accessToken = sessionStorage.getItem('accessToken');
    var objeto = JSON.parse(accessToken);
    token = objeto.token;

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


