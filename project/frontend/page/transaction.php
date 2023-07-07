<link rel="icon" type="image/x-icon" href="../img/dollar.png" />
<link rel="stylesheet" href="../css/transaction/style.css">
<?php include_once("../includes/header.php") ?>
</head>

<body>
    <?php include_once("../includes/sidebar.php") ?>
    <div class="container-fluid">
        <div class="row">
            <div class="d-flex justify-content-between mt-2">
                <div class="col-md">
                    <div class="fs-3"><b>Transações</b></div>
                </div>
            </div>
            <div class="col-md-12 text-center text-light">
                <div class="box-dashboard mt-1 bg-success filter-preset-1" id="box-dashboard-revenue">
                    <div class="pt-2 ps-2 text-start">
                        <i class="fa-solid fa-building-columns"></i> Receitas
                    </div>
                    <div class="pt-2 ps-2 fs-6 text-start">
                        <b>R$ 2.300,50</b>
                    </div>
                </div>
                <div class="box-dashboard mt-1 bg-danger filter-preset-1" id="box-dashboard-expense">
                    <div class="pt-2 ps-2 text-start">
                        <i class="fa-solid fa-building-columns"></i> Despesas
                    </div>
                    <div class="pt-2 ps-2 fs-6 text-start">
                        <b>R$ 2.300,50</b>
                    </div>
                </div>
            </div>
            <div class="col-md-12 mt-2">
                <div class="mt-2 text-center">
                    <div class="btn-group" role="group" aria-label="Basic example">
                        <button type="button" class="btn btn-success"><i class="fas fa-arrow-left"></i></button>
                        <span class="date btn btn-success">Agosto</span>
                        <button type="button" class="btn btn-success"><i class="fas fa-arrow-right"></i></button>
                    </div>
                </div>
            </div>
            <div class="col-md-12">
                <?php include_once("../table/transaction-table.php") ?>
            </div>
        </div>
    </div>

    <div></div>
    <?php include_once("../modal/modal-new-account.php") ?>
    <?php include_once("../modal/modal-update-account.php") ?>
    <?php include_once("../modal/modal-delete-account.php") ?>
    <?php include_once("../includes/footer.php") ?>
    <script src="../js/transaction/script.js"></script>
</body>


</html>