
document.getElementById("alert-error").style.display = "none"

$('#button-eye').on('click', function () {
    showHidePassword("#button-eye i", "#password");
});

$('#button-eye-repeat').on('click', function () {
    showHidePassword("#button-eye-repeat i", "#password-repeat");
});

$('#button-register').on('click', function () {
    $('#error-msg-authentication').text('');

    if (validationField() == true) {
        resquestRegisterUser();
    }
});

// FUNCTION

function validationField() {
    const MIN_LENGHT_NAME = 3;
    const MIN_LENGHT_USERNAME = 3;
    const MIN_LENGHT_PASSWORD = 8;

    const MAX_LENGHT_NAME = 255;
    const MAX_LENGHT_PASSWORD = 255;
    const MAX_LENGHT_USERNAME = 16;
    const MAX_LENGHT_EMAIL = 255;

    const ERROR_EMPTY_NAME = "Informe o nome!";
    const ERROR_EMPTY_USERNAME = "Informe o username!";
    const ERROR_EMPTY_EMAIL = "Informe o email!";
    const ERROR_EMPTY_PASSWORD = "Informe a senha!";
    const ERROR_EMPTY_PASSWORD_REPEAT = "Informe a senha novamente!";

    const ERROR_MIN_LENGHT_NAME = `O nome deve conter ao menos ${MIN_LENGHT_NAME} caracteres!`;
    const ERROR_MIN_LENGHT_USERNAME = `O username deve conter ao menos ${MIN_LENGHT_USERNAME} caracteres!`;
    const ERROR_MIN_LENGHT_PASSWORD = `A senha deve conter ao menos ${MIN_LENGHT_PASSWORD} caracteres!`;

    const ERROR_MAX_LENGHT_NAME = `O nome deve conter até ${MAX_LENGHT_NAME} caracteres!`;
    const ERROR_MAX_LENGHT_USERNAME = `O username deve conter  até ${MAX_LENGHT_USERNAME} caracteres!`;
    const ERROR_MAX_LENGHT_EMAIL = `O email deve conter até ${MAX_LENGHT_EMAIL} caracteres!`;
    const ERROR_MAX_LENGHT_PASSWORD = `A senha deve conter até ${MIN_LENGHT_PASSWORD} caracteres!`;

    const ERROR_DIFFERENT_PASSWORD = "As senhas devem ser iguais!";
    const ERROR_INVALID_USERNAME = "Username com caracteres inválido!"
    const ERROR_INVALID_EMAIL = "Email inválido!"

    name = $('#name').val();
    username = $('#username').val();
    email = $('#email').val();
    password = $('#password').val();
    password_repeat = $('#password-repeat').val();

    let isValid = true;

    if (name == "") {
        $('.error-msg-name').text(ERROR_EMPTY_NAME);
        isValid = false;

    } else if (name.length < MIN_LENGHT_NAME) {
        $('.error-msg-name').text(ERROR_MIN_LENGHT_NAME);
        isValid = false;

    } else if (name.length > MAX_LENGHT_NAME) {
        $('.error-msg-name').text(ERROR_MAX_LENGHT_NAME);
        isValid = false;

    } else {
        $('.error-msg-name').text("");
    }

    if (username == "") {
        $('.error-msg-username').text(ERROR_EMPTY_USERNAME);
        isValid = false;

    } else if (username.length < MIN_LENGHT_USERNAME) {
        $('.error-msg-username').text(ERROR_MIN_LENGHT_USERNAME);
        isValid = false;

    } else if (username.length > MAX_LENGHT_USERNAME) {
        $('.error-msg-username').text(ERROR_MAX_LENGHT_USERNAME);
        isValid = false;

    } else if (!validationUsername(username)) {
        $('.error-msg-username').text(ERROR_INVALID_USERNAME);
        isValid = false;

    } else {
        $('.error-msg-username').text("");
    }

    if (email == "") {
        $('.error-msg-email').text(ERROR_EMPTY_EMAIL);
        isValid = false;

    } else if (email.length > 255) {
        $('.error-msg-email').text(ERROR_MAX_LENGHT_EMAIL);
        isValid = false;

    } else if (!validarEmail(email)) {
        $('.error-msg-email').text(ERROR_INVALID_EMAIL);
        isValid = false;

    } else {
        $('.error-msg-email').text("");
    }

    if (password == "") {
        $('.error-msg-password').text(ERROR_EMPTY_PASSWORD);
        isValid = false;

    } else if (password.length < MIN_LENGHT_PASSWORD) {
        $('.error-msg-password').text(ERROR_MIN_LENGHT_PASSWORD);
        isValid = false;

    } else if (password.length > MAX_LENGHT_PASSWORD) {
        $('.error-msg-password').text(ERROR_MAX_LENGHT_PASSWORD);
        isValid = false;

    } else {
        $('.error-msg-password').text("");
    }

    if (password != "" && password_repeat == "") {
        $('.error-msg-password-repeat').text(ERROR_EMPTY_PASSWORD_REPEAT);
        isValid = false;

    } else if (password_repeat != password) {
        $('.error-msg-password-repeat').text(ERROR_DIFFERENT_PASSWORD);
        isValid = false;

    } else {
        $('.error-msg-password-repeat').text("");
    }

    return isValid;
}

function showHidePassword(button, field) {
    if ($(field).attr("type") == "password") {
        $(field).attr("type", "text");
        $(button).removeClass("fa-eye");
        $(button).addClass("fa-eye-slash");

    } else {
        $(field).attr("type", "password");
        $(button).removeClass("fa-eye-slash");
        $(button).addClass("fa-eye");
    }
}

function validationUsername(username) {
    var regexUsername = /^[a-zA-Z0-9_]{3,16}$/;

    if (regexUsername.test(username)) {
        return true;
    } else {
        return false;
    }
}

function validarEmail(email) {
    var regexEmail = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

    if (regexEmail.test(email)) {
        return true;
    } else {
        return false;
    }
}

$(document).ready(function () {
    var passwordInput = $('#password');
    var passwordRepeatInput = $('#password-repeat');
    var placeholderText = $('.placeholder-text');

    if (passwordInput.length > 0 && placeholderText.length > 0) {
        passwordInput.on('input', function () {
            if (passwordInput.val().length > 0) {
                placeholderText.eq(0).addClass('hide');
            } else {
                placeholderText.eq(0).removeClass('hide');
            }
        });

        passwordRepeatInput.on('input', function () {
            if (passwordRepeatInput.val().length > 0) {
                placeholderText.eq(1).addClass('hide');
            } else {
                placeholderText.eq(1).removeClass('hide');
            }
        });
    }
});

// REQUEST

function resquestRegisterUser() {
    var email = $("#email").val();
    var name = $("#name").val();
    var password = $("#password").val();
    var username = $("#username").val();
    var role = 0; // AQUI

    var connect_success = true;
    var xhr = new XMLHttpRequest();

    xhr.open('POST', 'http://localhost:9999/api/v0/user/');

    xhr.onload = function () {
        if (xhr.status === 200 || xhr.status === 201) {
            console.log("SUCESSO!");
        } else {
            connect_success = false;

            console.log(xhr.responseText);

            return connect_success
        }
    };

    var data = {
        "email": email,
        "name": name,
        "password": password,
        "role": role, // AQUI
        "username": username
    }

    var json = JSON.stringify(data);

    xhr.send(json);

    return connect_success
}