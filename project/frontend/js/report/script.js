$("#button-generate-report-balance").on("click" , function () {
    generateBalanceReportPDF();
});

// RELATÓRIOS

function generateBalanceReportPDF() {
    var accessToken = sessionStorage.getItem('accessToken');
    var objeto = JSON.parse(accessToken);
    token = objeto.token;

    var xhr = new XMLHttpRequest();

    xhr.open('GET', 'http://localhost:9999/api/v0/account/');

    xhr.setRequestHeader('Token', `Bearer ${token}`);

    xhr.onload = function () {
        if (xhr.status === 200) {
            var resposta = JSON.parse(xhr.responseText);

            resposta.sort(function (a, b) {
                return b.balance - a.balance;
            });

            var content = [];

            var totalBalance = 0;
            for (var i = 0; i < resposta.length; i++) {
                totalBalance += resposta[i].balance;
            }

            var currentDate = new Date();
            var headerText = 'Relatório de Saldo em Contas';
            var dateText = 'Data: ' + currentDate.toLocaleDateString();
            var timeText = 'Horário: ' + currentDate.toLocaleTimeString();
            content.push(
                { text: headerText, style: 'header' },
                { text: dateText, style: 'date' },
                { text: timeText, style: 'time' },
                '',
            );

            var tableData = [];
            for (var i = 0; i < resposta.length; i++) {
                var accountName = resposta[i].name;
                var balance = resposta[i].balance;
                var percentage = (balance / totalBalance * 100).toFixed(2);

                var row = [accountName, formatarMoeda(balance), percentage + '%'];
                tableData.push(row);
            }

            var table = {
                style: 'table table-striped',
                table: {
                    headerRows: 1,
                    widths: ['*', 'auto', 'auto'],
                    body: [
                        [
                            { text: 'Nome da Conta', style: 'tableHeader' },
                            { text: 'Saldo', style: 'tableHeader' },
                            { text: 'Porcentagem', style: 'tableHeader' }
                        ],
                        ...tableData.map(row => {
                            return row.map(column => {
                                return { text: column, style: 'tableCell' };
                            });
                        })
                    ]
                }
            };

            content.push(table);

            content.push(
                { text: 'Saldo Total: ' + formatarMoeda(totalBalance), style: 'totalBalance', alignment: 'right' }
            );

            var footer = {
                text: 'Gerado por OPENFINANCE',
                style: 'footer',
                alignment: 'center'
            };

            var documentDefinition = {
                content: [
                    content,
                    footer
                ],
                styles: {
                    header: { fontSize: 18, bold: true, margin: [0, 0, 0, 10] },
                    date: { fontSize: 12, italics: true, margin: [0, 0, 0, 5] },
                    time: { fontSize: 12, italics: true, margin: [0, 0, 0, 10] },
                    table: { margin: [0, 10, 0, 0] },
                    tableHeader: { bold: true, fillColor: '#343a40', color: '#ffffff', margin: [0, 2] },
                    tableCell: { margin: [0, 2] },
                    totalBalance: { bold: true, margin: [0, 10, 0, 0] },
                    footer: { fontSize: 10, margin: [0, 20] }
                },
                defaultStyle: {
                    columnGap: 10
                }
            };

            pdfMake.createPdf(documentDefinition).download('relatorio_saldo_contas.pdf');

            console.log("Relatório de saldo em contas gerado com sucesso!");
        } else {
            var objMessage = JSON.parse(xhr.responseText);
            var code = objMessage.code;
            var msg = objMessage.error;
            console.error("Erro ao gerar o relatório de saldo em contas:", msg, code);
        }
    };

    xhr.send();
}

function formatarMoeda(valor) {
    var formatter = new Intl.NumberFormat('pt-BR', {
        style: 'currency',
        currency: 'BRL'
    });
    return formatter.format(valor);
}