createGraphBar();

requestListRevenue(function (somaReceita) {
    requestListExpense(function (somaDespesas) {
        createGraphPie(somaReceita, somaDespesas);
        $(".receita").text(formatarMoeda(somaReceita));
        $(".despesa").text(formatarMoeda(somaDespesas));
    });
});

function createGraphPie(revenue, expense) {
    var ctx = document.getElementById('graph-pie').getContext('2d');

    var data = {
        labels: ['Receita', 'Despesa'],
        datasets: [{
            data: [revenue, expense],
            backgroundColor: ['#198754', '#DC3547'],
        }]
    };

    var options = {
        responsive: true,
        legend: {
            fontSize: 30
        }
    };

    var chart = new Chart(ctx, {
        type: 'pie',
        data: data,
        options: options
    });
}

function createGraphBar() {
    var ctx = document.getElementById('graph-bar').getContext('2d');

    var data = {
        labels: ['Janeiro', 'Fevereiro', 'Março', 'Abril', 'Maio', 'Junho'],
        datasets: [{
            label: 'Evolução do saldo',
            data: [1200.00, 1880.00, 1600.00, 2200.00, 2300.50, 2750.20],
            backgroundColor: '#0D6EFD'
        }]
    };

    var options = {
        responsive: true,
        scales: {
            y: {
                beginAtZero: true
            }
        }
    };

    var chart = new Chart(ctx, {
        type: 'bar',
        data: data,
        options: options
    });
}

function formatarMoeda(valor) {
    var formatter = new Intl.NumberFormat('pt-BR', {
        style: 'currency',
        currency: 'BRL'
    });

    return formatter.format(valor);
}

// REQUEST

function requestListRevenue(callback) {
    var accessToken = sessionStorage.getItem('accessToken');
    var objeto = JSON.parse(accessToken);
    token = objeto.token;

    var connect_success = true;

    var xhr = new XMLHttpRequest();

    xhr.open('GET', 'http://localhost:9999/api/v0/finance/');

    xhr.setRequestHeader('Token', `Bearer ${token}`);

    xhr.onload = function () {
        if (xhr.status === 200) {
            var resposta = JSON.parse(xhr.responseText);
            var somaReceita = 0.00;

            for (var i = 0; i < resposta.length; i++) {
                if (resposta[i].type_code == 0 && resposta[i].status_code == 1) {
                    somaReceita += parseFloat(resposta[i].value.toFixed(2));
                }
            }

            callback(somaReceita);

        } else if (xhr.status === 204) {
            console.log("Sem receitas registradas!");

        } else {
            connect_success = false;

            var objMessage = JSON.parse(xhr.responseText);

            var code = objMessage.code;
            var msg = objMessage.error;

            showModalMessage("bg-danger", "ERRO", msg, code);

            callback(connect_success);
        }
    };

    xhr.send();
};

function requestListExpense(callback) {
    var accessToken = sessionStorage.getItem('accessToken');
    var objeto = JSON.parse(accessToken);
    token = objeto.token;

    var connect_success = true;

    var xhr = new XMLHttpRequest();

    xhr.open('GET', 'http://localhost:9999/api/v0/finance/');

    xhr.setRequestHeader('Token', `Bearer ${token}`);

    xhr.onload = function () {
        if (xhr.status === 200) {
            var resposta = JSON.parse(xhr.responseText);
            var somaDespesas = 0.00;

            for (var i = 0; i < resposta.length; i++) {
                if (resposta[i].type_code == 1 && resposta[i].status_code == 1) {
                    somaDespesas += parseFloat(resposta[i].value.toFixed(2));
                }
            }

            callback(somaDespesas);

        } else if (xhr.status === 204) {
            console.log("Sem despesas registradas!");

        } else {
            connect_success = false;

            var objMessage = JSON.parse(xhr.responseText);

            var code = objMessage.code;
            var msg = objMessage.error;

            showModalMessage("bg-danger", "ERRO", msg, code);

            callback(connect_success);
        }
    };

    xhr.send();
};

