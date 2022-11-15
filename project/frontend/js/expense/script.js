//Botão "nova despesa", abre o modal nova despesa.
$("#btn-open-modal-expense").on("click", function () {
    openModal("new-expense")
})

//Botão "fechar", fecha o modal nova despesa.
$(".btn-close-modal-expense").on("click", function () {
    closeModal("new-expense")
})

// Funções
function closeModal(modalName) {
    $("#modal-" + modalName).modal("hide")
}

function openModal(modalName) {
    $("#modal-" + modalName).modal("show")
}