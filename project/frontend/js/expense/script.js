//------------------ MODAL ------------------
//Botão "nova despesa", abre o modal nova despesa.
$("#btn-open-modal-expense").on("click", function () {
    openModal("new-expense")
})

// Botão "lixeira", abre o modal de exclusão de despesa.
$(".btn-delete-expense").on("click", function () {
    openModal("delete-expense")
})

// Botão "Lápis", abre o modal de edição de despesa.
$(".btn-update-expense").on("click", function () {
    openModal("update-expense")
})

//------------------ BTN FECHAR ------------------
//Botão "fechar", fecha o modal nova despesa.
$(".btn-close-modal-expense").on("click", function () {
    closeModal("new-expense")
})

//Botão "fechar", fecha o modal excluir despesa.
$(".btn-close-modal-delete-expense").on("click", function () {
    closeModal("delete-expense")
})

//Botão "fechar", fecha o modal excluir despesa.
$(".btn-close-modal-update-expense").on("click", function () {
    closeModal("update-expense")
})

// Funções
function closeModal(modalName) {
    $("#modal-" + modalName).modal("hide")
}

function openModal(modalName) {
    $("#modal-" + modalName).modal("show")
}

// Select2
$(document).ready(function() {
    $('#select-category').select2();
    $('#select-category-update').select2();
});