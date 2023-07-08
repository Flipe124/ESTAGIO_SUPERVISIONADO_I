requestListTransaction();


$('#box-dashboard-revenue').on('click', function () {
    location.replace('./revenue.php');
});

$('#box-dashboard-expense').on('click', function () {
    location.replace('./expense.php');
});


// PREENCHER ELEMENTO

function createTableTransaction(beneficiary_id, beneficiary_name, emitter_id, emitter_name, id, value) {
    var result = document.querySelector('.transaction-table');

    result.innerHTML +=
        `<div class="result filter-preset-1" data-baneficiary-id="${beneficiary_id}" data-baneficiary-name="${beneficiary_name}" data-emitter-id="${emitter_id}" data-emitter-name="${emitter_name}" data-id="${id}" data-value="${value}">
            <span class="icon-category text-primary">
                <i class="fa-solid fa-money-bill-transfer"></i>
            </span>
            <span class="description">
                <span class="category">
                    <b>Transferência</b>
                </span>
                <div class="data text-secondary">
                    ${emitter_name} para ${beneficiary_name}
                </div>
                <span class="font-size-14 text-secondary"></span>
            </span>
            <div class="value">
                <b class="text-value">${value}</b> <span class="status mb-1 ms-1"></span>
            </div>
        </div>`

}

// TRATAMENTO DE ERROS

function showModalMessage(backgroundTitle, title, message, code) {
    $(".modal").modal("hide");
    $("#modal-message").modal("show");

    const ERROR_BAD_REQUEST = `Requisição inválida, se o erro persistir contate o suporte!`;
    const ERROR_UNAUTHORIZED = `Seu token de acesso expirou, faça o login novamente!`;
    const ERROR_CONFLIT = `Erro de conflito, registro já existente!`;
    const ERROR_UNPROCESSABLE_ENTITY = `Erro de entidade improcessável, se o erro persisitir contate o suporte!`;
    const ERROR_INTERNAL_SERVER = `Erro interno do servidor, se o erro persistir contate o suporte!`;

    if (code == 400) {
        $(".modal-header").addClass(backgroundTitle);
        $("#modal-message .modal-title").text(title);
        $("#modal-message .message").text(ERROR_BAD_REQUEST);

        $("#modal-message .btn-success").on("click", function () {
            location.reload();
        })
    }
    else if (code == 401 && message == "Unauthorized") {
        $(".modal-header").addClass(backgroundTitle);
        $("#modal-message .modal-title").text(title);
        $("#modal-message .message").text(ERROR_UNAUTHORIZED);

        $("#modal-message .btn-success").on("click", function () {
            location.replace("../login");
        })

    } else if (code == 409) {
        $(".modal-header").addClass(backgroundTitle);
        $("#modal-message .modal-title").text(title);
        $("#modal-message .message").text(ERROR_CONFLIT);

        $("#modal-message .btn-success").on("click", function () {
            location.reload();
        })

    } else if (code == 422) {
        $(".modal-header").addClass(backgroundTitle);
        $("#modal-message .modal-title").text(title);
        $("#modal-message .message").text(ERROR_UNPROCESSABLE_ENTITY);

        $("#modal-message .btn-success").on("click", function () {
            location.reload();
        })
    } else if (code == 500) {
        $(".modal-header").addClass(backgroundTitle);
        $("#modal-message .modal-title").text(title);
        $("#modal-message .message").text(ERROR_INTERNAL_SERVER);

        $("#modal-message .btn-success").on("click", function () {
            location.reload();
        })

    } else {
        $(".modal-header").addClass(backgroundTitle);
        $("#modal-message .modal-title").text(title);
        $("#modal-message .message").text(message);

        $("#modal-message .btn-success").on("click", function () {
            location.reload();
        })
    }
};

// REQUEST 

function requestListTransaction() {
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

            for (var i = 0; i < resposta.length; i++) {
                createTableTransaction(resposta[i].beneficiary_id, resposta[i].beneficiary_name, resposta[i].emitter_id, resposta[i].emitter_name, resposta[i].id, resposta[i].value);

                console.log("TESTE")
            }

        } else if (xhr.status === 204) {
            $(".text-empty-transaction").text("Sem transferências realizadas!");

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