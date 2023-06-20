
// Obtém a referência do elemento canvas no HTML
var ctx = document.getElementById('grafico').getContext('2d');

// Dados para o gráfico de pizza
var data = {
    labels: ['Receita', 'Despesa'], // Rótulos das fatias
    datasets: [{
        data: [333, 333], // Valores das fatias
        backgroundColor: ['green', 'red'], // Cores das fatias
    }]
};

// Configurações do gráfico
var options = {
    responsive: true
};

// Cria o gráfico de pizza
var chart = new Chart(ctx, {
    type: 'pie',
    data: data,
    options: options
});
