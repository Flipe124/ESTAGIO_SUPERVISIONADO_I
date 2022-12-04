//------------------ MODAL ------------------
//Botão "nova despesa", abre o modal nova despesa.
$("#btn-open-modal-expense").on("click", function () {
    modalFunction("new-expense", "show")
})

// Botão "lixeira", abre o modal de exclusão de despesa.
$(".btn-delete-expense").on("click", function () {
    modalFunction("delete-expense", "show")
})

// Botão "Lápis", abre o modal de edição de despesa.
$(".btn-update-expense").on("click", function () {
    modalFunction("update-expense", "show")
})

//------------------ BTN FECHAR ------------------
//Botão "fechar", fecha o modal nova despesa.
$(".btn-close-modal-expense").on("click", function () {
    modalFunction("new-expense", "hide")
})

//Botão "fechar", fecha o modal excluir despesa.
$(".btn-close-modal-delete-expense").on("click", function () {
    modalFunction("delete-expense", "hide")
})

//Botão "fechar", fecha o modal excluir despesa.
$(".btn-close-modal-update-expense").on("click", function () {
    modalFunction("update-expense", "hide")
})

//------------------ Funções ------------------

// Apresentar/esconder modal
function modalFunction(modalName, action) {
    $("#modal-" + modalName).modal(action)
}

// Select2
$(document).ready(function() {
    $('#select-category').select2();
    $('#select-category-update').select2();
});