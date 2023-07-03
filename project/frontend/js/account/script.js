// TELA PRINCIPAL

$('#button-new-account').on('click', function () {
    $('#modal-create-account').modal('show');
    $('.error').text("");
    $("#form-create-account")[0].reset();
});

$('.button-update-account').on('click', function () {
    $('.error').text("");
    buttonOpenUpdateAccountModal("update", $(this).data("id"), $(this).data("name-account"), $(this).data("balance-account"));
});

$('.button-delete-account').on('click', function () {
    buttonOpenDeleteAccountModal("delete-account", $(this).data("id"), $(this).data("name-account"), $(this).data("balance-account"))
});

// DENTRO DO MODAL

$('#button-create-account').on('click', function () {
    if (validationFormAccount("create") == true) {
        console.log("Tudo certo!")
    }
});

$('#button-update-account').on('click', function () {
    if (validationFormAccount("update") == true) {
        console.log("Tudo certo!")
    }
});

$('#button-delete-account').on('click', function () {

});

// FUNCTIONS

function buttonOpenUpdateAccountModal(modalForm, id, name, value) {
    $(`#modal-${modalForm}-account`).modal("show")
    $(`#form-${modalForm}-account`)[0].reset();

    // $('#update-id').val(id);
    $(`#form-${modalForm}-account #${modalForm}-input-name-account`).val(name);
    $(`#form-${modalForm}-account #${modalForm}-input-balance-account`).val(formatValueFromData(value));

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

function validationFormAccount(form) {
    const MAX_LENGHT_NAME = 30;
    const MAX_LENGHT_BALANCE = 20;

    const ERROR_EMPTY_NAME = "Informe o nome da conta!";
    const ERROR_EMPTY_BALANCE = "Informe o saldo da conta!";

    const ERROR_MAX_LENGHT_NAME = `Nome pode conter até ${MAX_LENGHT_NAME} caracteres!`;
    const ERROR_MAX_LENGHT_BALANCE = `Saldo pode conter  até ${MAX_LENGHT_BALANCE} algarismos!`;

    let name = $(`#form-${form}-account #${form}-input-name-account`).val();
    let balance = $(`#form-${form}-account #${form}-input-balance-account`).val();

    let isValid = true;

    if (name == "") {
        $(`#form-${form}-account .error-name-account`).text(ERROR_EMPTY_NAME);
        isValid = false;

    } else if (name != "" && name.length > MAX_LENGHT_NAME) {
        $(`#form-${form}-account .error-name-account`).text(ERROR_MAX_LENGHT_NAME);
        isValid = false;

    } else {
        $(`#form-${form}-account .error-name-account`).text("");
    }

    if (balance == "R$ 0,00" || balance == "") {
        $(`#form-${form}-account .error-balance-account`).text(ERROR_EMPTY_BALANCE);
        isValid = false;

    } else if (balance != "" && balance.length > MAX_LENGHT_BALANCE) {
        $(`#form-${form}-account .error-balance-account`).text(ERROR_MAX_LENGHT_BALANCE);
        isValid = false;

    } else {
        $(`#form-${form}-account .error-balance-account`).text("");
    }

    return isValid;

}