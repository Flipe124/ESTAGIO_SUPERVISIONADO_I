$('#button-new-account').on('click', function () {
    $('#modal-create-account').modal('show');
    $("#form-create-account")[0].reset();
});

$('.button-update-account').on('click', function () {
    buttonOpenUpdateAccountModal("update-account", $(this).data("id"), $(this).data("name-account"), $(this).data("balance-account"));
});

$('.button-delete-account').on('click', function () {
    // $('#modal-delete-account').modal('show');
    buttonOpenDeleteAccountModal("delete-account", $(this).data("id"), $(this).data("name-account"), $(this).data("balance-account"))
});


// FUNCTIONS

function buttonOpenUpdateAccountModal(modalForm, id, name, value) {
    $(`#modal-${modalForm}`).modal("show")
    $(`#form-${modalForm}`)[0].reset();

    // $('#update-id').val(id);
    $(`#form-${modalForm} #input-name-account`).val(name);
    $(`#form-${modalForm} #input-balance-account`).val(formatValueFromData(value));

};

function buttonOpenDeleteAccountModal(modal, id, name, value) {
    $(`#modal-${modal}`).modal("show");

    // $('#update-id').val(id);
    $(`#text-name-account`).text(name);
    $(`#text-balance-account`).text(formatValueFromData(value));

};

function formatValue(input) {
    var value = input.value.replace(/\D/g, '');

    value = (value / 100).toFixed(2);

    value = value.replace(/\B(?=(\d{3})+(?!\d))/g, '.');
    value = value.replace('.', ',');

    input.value = 'R$ ' + value;
}

function formatValueFromData(value) {
    if (typeof value !== 'string') {
        value = value.toString();
    }

    value = value.replace(/\D/g, '');

    value = (value / 100).toFixed(2);

    value = value.replace(/\B(?=(\d{3})+(?!\d))/g, '.');
    value = value.replace('.', ',');

    return 'R$ ' + value;
}
