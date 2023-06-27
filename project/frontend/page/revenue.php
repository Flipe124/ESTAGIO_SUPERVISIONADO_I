<link rel="icon" type="image/x-icon" href="../img/dollar.png" />
<?php include_once("../includes/header.php") ?>
</head>

<body>
    <?php include_once("../includes/sidebar.php") ?>
    <div class="container-fluid">
        <div class="row">
            <div class="d-flex justify-content-between mt-2">
                <div class="col-md">
                    <div class="fs-3"><b>Receitas</b></div>
                </div>
                <div class="col-md d-flex justify-content-end">
                    <button class="btn btn-success btn-sm mt-1" id="button-new-revenue" type="button">
                        <i class="fa-solid fa-plus"></i> <b>NOVA</b>
                    </button>
                </div>
            </div>
            <div class="col-md-12 text-center">
                <div class="box-dashboard-unic mt-1 bg-success text-light">
                    <div class="pt-2 ps-2 text-start">
                        <i class="fas fa-chevron-circle-up"></i> Receitas
                    </div>
                    <div class="pt-2 ps-2 fs-6 text-start">
                        <b class="sum-revenue"></b>
                    </div>
                </div>
            </div>
            <div class="col-md-12">
                <div class="pagination mt-2 text-center">
                    <button class="btn" type="button"><i class="fas fa-arrow-left"></i></button>
                    <b class="date mt-1">Agosto</b>
                    <button class="btn" type="button"><i class="fas fa-arrow-right"></i></button>
                </div>
            </div>
            <div class="col-md-12">
                <?php include_once("../table/revenue-table.php") ?>
            </div>
        </div>
        <?php require_once("../modal/modal-new.php"); ?>
        <?php require_once("../modal/modal-update.php"); ?>
        <?php require_once("../modal/modal-delete.php"); ?>
    </div>

    <div></div>
    <?php include_once("../includes/footer.php") ?>
    <script src="../js/revenue/script.js"></script>
</body>


</html>