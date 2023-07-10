$("#button-generate-report-balance").on("click", function () {
    generateBalanceReportPDF();
});

$("#button-generate-report-transfer").on("click", function () {
    generateTransferReportPDF();
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

function generateTransferReportPDF() {
    var accessToken = sessionStorage.getItem('accessToken');
    var objeto = JSON.parse(accessToken);
    token = objeto.token;

    var connect_success = true;

    var xhr = new XMLHttpRequest();

    xhr.open('GET', 'http://localhost:9999/api/v0/transaction/');

    xhr.setRequestHeader('Token', `Bearer ${token}`);

    xhr.onload = function () {
        if (xhr.status === 200) {
            $(".text-empty-transaction").text("");
            var resposta = JSON.parse(xhr.responseText);

            var content = [];
            var totalTransferencias = 0;
            var transferenciasPorConta = {};

            for (var i = 0; i < resposta.length; i++) {
                var beneficiaryName = resposta[i].beneficiary_name;
                var emitterName = resposta[i].emitter_name;
                var value = resposta[i].value;

                var conta = emitterName + '-' + beneficiaryName;
                if (transferenciasPorConta.hasOwnProperty(conta)) {
                    transferenciasPorConta[conta] += value;
                } else {
                    transferenciasPorConta[conta] = value;
                }

                totalTransferencias += value;
            }

            for (var key in transferenciasPorConta) {
                var contaSplit = key.split('-');
                var emitterName = contaSplit[0];
                var beneficiaryName = contaSplit[1];
                var value = transferenciasPorConta[key];

                content.push(
                    [
                        { text: emitterName, style: 'tableData' },
                        { text: beneficiaryName, style: 'tableData' },
                        { text: formatarMoeda(value), style: 'tableData', alignment: 'right' }
                    ]
                );
            }

            content.sort(function (a, b) {
                return b[2].text.replace(/[^\d.-]/g, '') - a[2].text.replace(/[^\d.-]/g, '');
            });

            var tableHeader = [
                { text: 'Emissor', style: 'tableHeader' },
                { text: 'Beneficiário', style: 'tableHeader' },
                { text: 'Valor', style: 'tableHeader', alignment: 'center' }
            ];

            content.unshift(tableHeader);

            content.push(
                [
                    { text: 'Total:', colSpan: 2, alignment: 'right', style: 'tableFooter' },
                    {},
                    { text: formatarMoeda(totalTransferencias), style: 'tableFooter', alignment: 'right' }
                ]
            );

            var documentDefinition = {
                content: [
                    { text: 'Relatório de Transferências', style: 'header' },
                    { text: 'Data: ' + new Date().toLocaleDateString(), style: 'date' },
                    { text: 'Horário: ' + new Date().toLocaleTimeString(), style: 'time' },
                    {
                        style: 'table',
                        table: {
                            widths: ['*', '*', 'auto'],
                            body: content
                        }
                    },
                    { text: 'Gerado por OPENFINANCE', style: 'footer', alignment: 'center' }
                ],
                styles: {
                    header: { fontSize: 18, bold: true, margin: [0, 0, 0, 10] },
                    date: { fontSize: 12, italics: true, margin: [0, 0, 0, 5] },
                    time: { fontSize: 12, italics: true, margin: [0, 0, 0, 10] },
                    table: { margin: [0, 0, 0, 0] },
                    tableHeader: { bold: true, fillColor: '#343a40', color: '#ffffff', margin: [0, 2], alignment: 'left' },
                    tableData: { fontSize: 12, margin: [0, 2], alignment: 'left' },
                    tableFooter: { fontSize: 12, bold: true, fillColor: '#f2f2f2', margin: [0, 2], alignment: 'right' },
                    footer: { fontSize: 10, margin: [0, 10, 0, 0] }
                }
            };

            pdfMake.createPdf(documentDefinition).download('relatorio_transferencias.pdf');

            console.log("Relatório de transferências gerado com sucesso!");

        } else if (xhr.status === 204) {
            $(".text-empty-transaction").text("Sem transferências realizadas!");

        } else {
            connect_success = false;

            var objMessage = JSON.parse(xhr.responseText);

            var code = objMessage.code;
            var msg = objMessage.error;

            showModalMessage("bg-danger", "ERRO", msg, code);
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