<div class="d-flex" id="wrapper">
    <div class="backgroud-black-blue" id="sidebar-wrapper">
        <div class="sidebar-heading text-center">
            <a href="../page/index.php">
                <img id="logotype-sidebar" src="../img/openfinance.png" width="200px" alt="logotipo">
            </a>
        </div>
        <ul id="ul-sidebar-link" class="list-group list-group-flush">
            <li class="option-not-selected">
                <a class="list-group-item p-3 option-not-selected" href="../page/index.php"><i class="fas fa-tools"></i> <b>DASHBOARD</b></a>
            </li>
            <li>
                <a class="list-group-item p-3 option-not-selected" href="../page/client.php"><i class="fa-sharp fa-solid fa-user-tie"></i> <b>RECEITAS</b></a>
            </li>
            <li>
                <a class="list-group-item p-3 option-not-selected" href="../page/user.php"><i class="fa-solid fa-user"></i> <b>DESPESAS</b></a>
            </li>
        </ul>
    </div>
    <div id="page-content-wrapper">
        <nav class="navbar backgroud-black-blue">
            <div class="container-fluid">
                <span>
                    <button class="fs-2 ms-1" id="sidebarToggle">
                        <span class="button-menu">
                            <div class="line line1"></div>
                            <div class="line line2"></div>
                            <div class="line line3"></div>
                        </span>
                    </button>
                </span>
                <span class="navbar-text">
                    <ul class="list-group list-group-horizontal generico">
                        <!-- <button id="button-notify" type="button" class="btn btn-light position-relative me-2">
                            <i id="icon-notify" class="fa-sharp fa-solid fa-bell"></i>
                            <span class="position-absolute top-0 start-100 translate-middle badge rounded-pill bg-danger">
                                0
                                <span class="visually-hidden">unread messages</span>
                            </span>
                        </button> -->
                        <li class="li-group">
                            <button class="dropdown-button"><img class="user-avatar" src="../img/user.png"></button>
                        </li>
                    </ul>
                    <ul class="dropdown-menu">
                        <li class="dropdown-item" id="user-email-SIDEBAR">TESTE</li>
                        <li class="dropdown-item text-danger item-button-return-login"><i class="fa-solid fa-arrow-right-from-bracket"></i> Sair</li>
                    </ul>
                </span>
            </div>
        </nav>