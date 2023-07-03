
createGraphPie();
createGraphBar();

function createGraphPie() {
    var ctx = document.getElementById('graph-pie').getContext('2d');

    var data = {
        labels: ['Receita', 'Despesa'],
        datasets: [{
            data: [3780.59, 975],
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

