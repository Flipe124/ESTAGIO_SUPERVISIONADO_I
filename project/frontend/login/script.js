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
    if ($("#password").attr("type") == "password") {
        $("#password").attr("type", "text");
        $("#button-eye i").removeClass("fa-eye");
        $("#button-eye i").addClass("fa-eye-slash");
    } else {
        $("#password").attr("type", "password");
        $("#button-eye i").removeClass("fa-eye-slash");
        $("#button-eye i").addClass("fa-eye");
    }
});

$('#button-login').on('click', function () {
    $('#error-msg-authentication').text('');

    if (validationEmail(email.value) && validationPassword(password.value)) {
        event.preventDefault();
        requestAuthentication();
    }
})

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

function validationEmail(email) {
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

function requestAuthentication() {
    var xhr = new XMLHttpRequest();

    xhr.open('POST', 'http://localhost:8008/api/v2/auth/');

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
            "email": email.value,
            "password": password.value,
        }

    } else {
        var data = {
            "username": email.value,
            "password": password.value
        }
    }

    var json = JSON.stringify(data);

    xhr.send(json);

}