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
            <div class="col-md-12 mt-3">
                <h5><i class="fas fa-chevron-circle-up"></i> Receita</h5>
                <button class="btn btn-success mt-1" id="button-generate-report-revenue" style="width: 200px">Receitas</button>
                <button class="btn btn-success mt-1" id="button-generate-report-revenue-for-category" style="width: 200px">Receitas por categoria</button>
            </div>
            <div class="col-md-12 mt-3">
                <h5><i class="fas fa-chevron-circle-down"></i> Despesa</h5>
                <button class="btn btn-danger mt-1" id="button-generate-report-expense" style="width: 200px">Despesas</button>
                <button class="btn btn-danger mt-1" id="button-generate-report-expense-for-category" style="width: 200px">Despesa por categoria</button>
            </div>
            <div class="col-md-12 mt-3">
                <h5><i class="fas fa-exchange-alt"></i> Tranferências</h5>
                <button class="btn btn-secondary mt-1" id="button-generate-report-transfer" style="width: 200px">Tranferências</button>
            </div>
            <div class="col-md-12 mt-3">
                <h5><i class="fa-solid fa-building-columns"></i> Saldo</h5>
                <button class="btn btn-primary mt-1" id="button-generate-report-balance" style="width: 200px">Saldo por contas</button>
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