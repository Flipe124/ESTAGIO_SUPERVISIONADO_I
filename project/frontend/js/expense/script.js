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

//------------------ Funções ------------------

// Apresentar/esconder modal
function modalAction(modalName, action) {
    $("#modal-" + modalName).modal(action)
}

// Select2
$(document).ready(function() {
    $('#select-category').select2();
    $('#select-category-update').select2();
});