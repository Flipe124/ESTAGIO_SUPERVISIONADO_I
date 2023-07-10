<link rel="icon" type="image/x-icon" href="../img/dollar.png" />
<?php include_once("../includes/header.php") ?>
</head>

<body>
    <?php include_once("../includes/sidebar.php") ?>
    <div class="container-fluid">
        <div class="row">
            <div class="d-flex justify-content-between mt-2">
                <div class="col-md">
                    <div class="fs-3"><b>Relatórios</b></div>
                </div>
            </div>
            <div class="col-md-12">
                <button class="btn btn-primary mt-1" id="button-generate-report-balance">Gerar relatório de saldo</button>
                <button class="btn btn-primary mt-1">Relatório de saldo</button>
                <button class="btn btn-primary mt-1">Relatório de saldo</button>
            </div>
        </div>
    </div>

    <div></div>
    <?php require_once("../modal/modal-message.php"); ?>
    <?php include_once("../includes/footer.php") ?>
    <script src="https://cdn.jsdelivr.net/npm/chart.js@2.9.4/dist/Chart.bundle.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/pdfmake/0.1.68/pdfmake.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/pdfmake/0.1.68/vfs_fonts.js"></script>

    <script src="../js/report/script.js"></script>

</body>


</html>