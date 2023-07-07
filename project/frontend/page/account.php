<link rel="icon" type="image/x-icon" href="../img/dollar.png" />
<?php include_once("../includes/header.php") ?>
</head>

<body>
    <?php include_once("../includes/sidebar.php") ?>
    <div class="container-fluid">
        <div class="row">
            <div class="d-flex justify-content-between mt-2">
                <div class="col-md">
                    <div class="fs-3"><b>Contas</b></div>
                </div>
                <div class="col-md d-flex justify-content-end">
                    <button class="btn btn-success btn-sm mt-1" id="button-new-account" type="button">
                        <i class="fa-solid fa-plus"></i> <b>NOVA</b>
                    </button>
                </div>
            </div>
            <div class="col-md-12 text-center">
                <div class="box-dashboard-unic mt-1 bg-primary text-light filter-preset-1">
                    <div class="pt-2 ps-2 text-start">
                        <i class="fas fa-chevron-circle-up"></i> Saldo
                    </div>
                    <div class="pt-2 ps-2 fs-6 text-start">
                        <!-- <b class="sum-revenue"></b> -->
                        <b>R$ 10.420,78</b>
                    </div>
                </div>
            </div>
            <div class="col-md-12 mt-2">
                <?php include_once("../table/account-table.php") ?>
            </div>
        </div>
    </div>

    <div></div>
    <?php include_once("../modal/modal-new-account.php") ?>
    <?php include_once("../modal/modal-update-account.php") ?>
    <?php include_once("../modal/modal-delete-account.php") ?>
    <?php require_once("../modal/modal-message.php"); ?>
    <?php include_once("../includes/footer.php") ?>
    <script src="../js/account/script.js"></script>
    <script src="https://cdn.datatables.net/1.11.2/js/jquery.dataTables.min.js"></script>
</body>


</html>