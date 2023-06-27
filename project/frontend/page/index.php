<link href="../css/dashboard/style.css" rel="stylesheet" />
<link rel="icon" type="image/x-icon" href="../img/dollar.png" />
<?php include_once("../includes/header.php") ?>
</head>

<body>
    <?php include_once("../includes/sidebar.php") ?>
    <div class="container-fluid">
        <div class="row">
            <div class="d-flex justify-content-between mt-2">
                <div class="col-md">
                    <div class="fs-3"><b>Dashboard</b></div>
                </div>
            </div>
            <!-- <div class="col-md-12 fs-2 mt-2"><b>Dashboard</b></div> -->
            <div class="col-md-12 text-center text-light">
                <div class="box-dashboard mt-1 bg-primary">
                    <div class="pt-2 ps-2 text-start">
                        <i class="fa-solid fa-building-columns"></i> Saldo
                    </div>
                    <div class="pt-2 ps-2 fs-6 text-start">
                        <b>R$ 1.600,55</b>
                    </div>
                </div>
                <div class="box-dashboard mt-2 bg-success">
                    <div class="pt-2 ps-2 text-start">
                        <i class="fas fa-chevron-circle-up"></i> Receitas
                    </div>
                    <div class="pt-2 ps-2 fs-6 text-start">
                        <b>R$ 1.600,55</b>
                    </div>
                </div>
                <div class="box-dashboard mt-2 bg-danger">
                    <div class="pt-2 ps-2 text-start">
                        <i class="fas fa-chevron-circle-down"></i> Despesas
                    </div>
                    <div class="pt-2 ps-2 fs-6 text-start">
                        <b>R$ 101.600,55</b>
                    </div>
                </div>
            </div>
            <div class="col-md-12">
                <div class="dashboard">
                    <div>

                    </div>
                </div>
            </div>
            <div class="row">
                <div class="col-md-6">
                    <div class="revenue">
                        <b class="text-light">RECEITA:</b>
                        <div class="text-success">
                            R$ 1.500,00
                        </div>
                    </div>
                    <div class="expense">
                        <b class="text-light">DESPESA:</b>
                        <div class="text-danger">
                            R$ 1.500,00
                        </div>
                    </div>

                    <div>TESTE2</div>
                </div>
                <div class="col-md-6">
                    <div><canvas id="grafico"></canvas></div>
                </div>
            </div>
            <div class="col-md-6">
                teste
            </div>
        </div>
    </div>

    <div></div>
    <script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
    <script src="../js/graph/script.js"></script>

    <?php include_once("../includes/footer.php") ?>
</body>


</html>