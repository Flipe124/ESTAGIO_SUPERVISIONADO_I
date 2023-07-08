//------------------ MODAL ------------------
//Botão "nova despesa", abre o modal nova despesa.
$("#btn-open-modal-expense").on("click", function () {
    modalAction("new-expense", "show")
})

// Botão "lixeira", abre o modal de exclusão de despesa.
$(".btn-delete-expense").on("click", function () {
    modalAction("delete-expense", "show")
})

// Botão "Lápis", abre o modal de edição de despesa.
$(".btn-update-expense").on("click", function () {
    modalAction("update-expense", "show")
})

//------------------ BTN FECHAR ------------------
//Botão "fechar", fecha o modal nova despesa.
$(".btn-close-modal-expense").on("click", function () {
    modalAction("new-expense", "hide")
})

//Botão "fechar", fecha o modal excluir despesa.
$(".btn-close-modal-delete-expense").on("click", function () {
    modalAction("delete-expense", "hide")
})

//Botão "fechar", fecha o modal excluir despesa.
$(".btn-close-modal-update-expense").on("click", function () {
    modalAction("update-expense", "hide")
})

//----------------- BTN SALVAR ----------------

$("#btn-save-new-expense").on("click", function () {
    // window.location.reload();
})


//------------------ Funções ------------------

$("#btn-update-expense").on("click", function () {
    let el = document.querySelector(".btn-update-expense");

    let dataId = el.getAttribute("data-id");

    console.log(dataId)
})

// Apresentar/esconder modal
function modalAction(modalName, action) {
    $("#modal-" + modalName).modal(action)
}

// Select2
$(document).ready(function () {
    $('#select-category').select2();
    $('#select-category-update').select2();
});




// request 

function requestListExpense() {
    var accessToken = sessionStorage.getItem('accessToken');
    var objeto = JSON.parse(accessToken);
    token = objeto.token;

    console.log(token); // Remover na versão final

    var connect_success = true;

    var xhr = new XMLHttpRequest();

    xhr.open('GET', 'http://localhost:9999/api/v0/finance/'); // ALTERAR

    xhr.setRequestHeader('Token', `Bearer ${token}`);

    xhr.onload = function () {
        if (xhr.status === 200) {
            console.log("Status OK");
            var resposta = JSON.parse(xhr.responseText);

            for (var i = 0; i < resposta.length; i++) {
                console.log("RESPOSTA -> " + resposta);

                if(resposta[i].type_code == "1") {// 1 saida
                    console.log("IMPRIMI")
                    generateTableOperation(resposta[i].account_id, resposta[i].category_id, resposta[i].status_code, resposta[i].type_code, resposta[i].id, resposta[i].datetime, resposta[i].description, resposta[i].value)
                }

            }

        } else if (xhr.status === 204) {
            console.log("Sem receitas registradas!");

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