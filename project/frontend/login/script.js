const PASSWORD_MIN_LENGHT = 8;
const REGEX_EMAIL = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

const ERROR_EMPTY_EMAIL_USERNAME = "Informe seu e-mail ou apelido!";
const ERROR_EMPTY_PASSWORD = "Informe sua senha!";

const ERROR_INVALID_EMAIL = "E-mail inválido!";
const ERROR_MIN_CHARACTER_PASSWORD = "Senha deve conter " + PASSWORD_MIN_LENGHT + " ou mais caracteres!";

var accessToken;
var emailUser;

email = document.getElementById("email");
password = document.getElementById("password");

document.getElementById("alert-error").style.display = "none"

$('#button-eye').on('click', function () {
    showHidePassword("#button-eye i", "#password");
});

$('#button-eye-repeat').on('click', function () {
    showHidePassword("#button-eye-repeat i", "#password-repeat");
});


$('#button-login').on('click', function () {
    $('#error-msg-authentication').text('');

    if (validationEmailOrUsername(email.value) && validationPassword(password.value)) {
        event.preventDefault();
        requestAuthentication();
    }
});

// FUNCTION

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

function validationPassword(password) {
    field_error_password = document.getElementById("error-msg-password");

    if (password.length < PASSWORD_MIN_LENGHT && password.length > 0) {
        field_error_password.innerHTML = ERROR_MIN_CHARACTER_PASSWORD;
        field_error_password.style.display = 'block';

        return false;

    } else if (password.length == 0) {
        field_error_password.innerHTML = ERROR_EMPTY_PASSWORD;
        field_error_password.style.display = 'block';

        return false;

    }

    field_error_password.style.display = 'none';

    return true;
}

function validationEmailOrUsername(email) {
    field_error_email = document.getElementById("error-msg-email");

    if (email.includes('@') || email.includes('.')) {
        if (!REGEX_EMAIL.test(email) && email.length > 0) {
            field_error_email.innerHTML = ERROR_INVALID_EMAIL;
            field_error_email.style.display = 'block';

            return false;

        }
    }

    if (email.length == 0) {
        field_error_email.innerHTML = ERROR_EMPTY_EMAIL_USERNAME;
        field_error_email.style.display = 'block';

        return false;

    }

    field_error_email.style.display = 'none';

    return true;
}

function validationUserName(email) {
    field_error_email = document.getElementById("error-msg-email");

    if (email.length == 0) {
        field_error_email.innerHTML = ERROR_EMPTY_EMAIL_USERNAME;
        field_error_email.style.display = 'block';

        return false;
    }

    field_error_email.style.display = 'none';

    return true;
}

// HIDE PLACEHOLDER

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

function requestAuthentication() {
    user = $("#email").val();
    password = $("#password").val();

    var xhr = new XMLHttpRequest();

    xhr.open('POST', 'http://localhost:9999/api/v0/auth/');

    xhr.setRequestHeader('Content-Type', 'application/json');

    xhr.onload = function () {
        if (xhr.status === 200) {
            sessionStorage.setItem('accessToken', xhr.responseText);
            sessionStorage.setItem('emailUser', email.value)

            accessToken = sessionStorage.getItem('accessToken');
            emailUser = sessionStorage.getItem('emailUser');

            if (accessToken) {
                location.replace('http://localhost:8080/page/index.php')
            }

        } else {
            $('#error-msg-authentication').text('Credenciais incorretas!');
        }
    };

    if (email.value.includes('@')) {
        var data = {
            "email": user,
            "password": password,
        }

    } else {
        var data = {
            "username": user,
            "password": password
        }
    }

    var json = JSON.stringify(data);

    xhr.send(json);

}