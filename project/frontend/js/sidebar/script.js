$(document).ready(function () {
    emailUser = sessionStorage.getItem('emailUser');

    hiddenLogo();
    ShowHiddenDropdown();

    $('.button-menu').on('click', function () {
        buttonMenu();
    });

    $('#return').on('click', function () {
        location.replace("../page/index.php");
    })

    $('#user-email-sidebar').text(emailUser);

    $('#button-notify').on('click', function () {
        location.replace('../page/notify.php')
    });

    $('.item-button-return-login').on('click', function () {
        location.replace('../login/index.php')
    });

    checkOptionSelected();

});

function hiddenLogo() {
    var width_device = window.innerWidth;

    if (width_device >= 768) {
        $(".logo-navbar").remove();
    } else {
        $(".logo-navbar").add();
    }
}

function ShowHiddenDropdown() {
    const dropdownButton = document.querySelector('.dropdown-button');
    const dropdownMenu = document.querySelector('.dropdown-menu');

    dropdownButton.addEventListener('click', function () {
        dropdownMenu.classList.toggle('active');
    });
}

function buttonMenu() {
    var buttonMenu = document.getElementsByClassName("button-menu")[0];

    try {
        if (buttonMenu.classList.contains("active")) {
            buttonMenu.classList.remove("active");
        } else {
            buttonMenu.classList.add("active");
        }
    } catch (error) {
        console.log(error);
    }
}

function checkOptionSelected() {
    var links = document.querySelectorAll('a');
    for (var i = 0; i < links.length; i++) {
        if (links[i].href == window.location.href) {
            links[i].classList.remove('option-not-selected');
            links[i].classList.add('option-selected');
        }
    }
}