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
                <div class="col-md d-flex justify-content-end">
                    <button class="btn btn-success btn-sm mt-1" id="button-new-transaction" type="button">
                        <i class="fa-solid fa-plus"></i> <b>NOVA</b>
                    </button>
                </div>
            </div>
            <div class="col-md-12">
                <?php include_once("../table/transaction-table.php") ?>
            </div>
            <div class="col-md-12 text-empty-transaction mt-2"></div>
        </div>
    </div>

    <div></div>
    <?php include_once("../modal/modal-new-account.php") ?>
    <?php include_once("../modal/modal-update-account.php") ?>
    <?php include_once("../modal/modal-delete-account.php") ?>
    <?php require_once("../modal/modal-message.php"); ?>
    
    <?php include_once("../includes/footer.php") ?>
    <script src="../js/transaction/script.js"></script>
</body>


</html>