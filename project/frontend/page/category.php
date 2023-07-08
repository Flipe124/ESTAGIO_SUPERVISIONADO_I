<link rel="icon" type="image/x-icon" href="../img/dollar.png" />
<?php include_once("../includes/header.php") ?>
</head>

<body>
    <?php include_once("../includes/sidebar.php") ?>
    <div class="container-fluid">
        <div class="row">
            <div class="d-flex justify-content-between mt-2">
                <div class="col-md">
                    <div class="fs-3"><b>Categorias</b></div>
                </div>
                <div class="col-md d-flex justify-content-end">
                    <button class="btn btn-success btn-sm mt-1" id="button-new-category" type="button">
                        <i class="fa-solid fa-plus"></i> <b>NOVA</b>
                    </button>
                </div>
            </div>
            <div class="col-md-12 mt-2">
                <?php include_once("../table/category-table.php") ?>
            </div>
        </div>
    </div>

    <div></div>
    <?php include_once("../modal/modal-new-category.php") ?>
    <?php include_once("../modal/modal-update-category.php") ?>
    <?php include_once("../modal/modal-delete-category.php") ?>
    <?php require_once("../modal/modal-message.php"); ?>
    <?php include_once("../includes/footer.php") ?>
    <script src="../js/category/script.js"></script>
    <script src="https://cdn.datatables.net/1.11.2/js/jquery.dataTables.min.js"></script>
</body>


</html>