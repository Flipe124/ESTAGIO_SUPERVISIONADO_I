
$("#btn-open-modal-revenue").on("click", function () {
    modalAction("new-revenue", "show")
})


//------------------ Funções ------------------

// Apresentar/esconder modal
function modalAction(modalName, action) {
    $("#modal-" + modalName).modal(action)
}
