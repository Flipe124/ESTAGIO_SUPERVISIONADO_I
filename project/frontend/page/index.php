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
            <div class="col-md-12 text-center text-light">
                <div class="box-dashboard mt-1 bg-primary filter-preset-1" id="box-dashboard-balance">
                    <div class="pt-2 ps-2 text-start">
                        <i class="fa-solid fa-building-columns fa-lg"></i> Saldo
                    </div>
                    <div class="pt-2 ps-2 fs-6 text-start">
                        <b class="saldo">R$ 0,00</b>
                    </div>
                </div>
                <div class="box-dashboard mt-1 bg-success filter-preset-1" id="box-dashboard-revenue">
                    <div class="pt-2 ps-2 text-start">
                        <i class="fas fa-chevron-circle-up fa-lg"></i> Receitas
                    </div>
                    <div class="pt-2 ps-2 fs-6 text-start">
                        <b class="receita">R$ 0,00</b>
                    </div>
                </div>
                <div class="box-dashboard mt-1 bg-danger filter-preset-1" id="box-dashboard-expense">
                    <div class="pt-2 ps-2 text-start">
                        <i class="fas fa-chevron-circle-down fa-lg"></i> Despesas
                    </div>
                    <div class="pt-2 ps-2 fs-6 text-start">
                        <b class="despesa">R$ 0,00</b>
                    </div>
                </div>
            </div>
            <!-- <div class="col-md-12">
                <div class="mt-2 text-center">
                    <div class="btn-group" role="group" aria-label="Basic example">
                        <button type="button" class="btn btn-light"><i class="fas fa-arrow-left"></i></button>
                        <span class="date btn btn-light">Agosto</span>
                        <button type="button" class="btn btn-light"><i class="fas fa-arrow-right"></i></button>
                    </div>
                </div>
            </div> -->
            <div class="col-md-6 mt-2">
                <div class="dashboard dashboard-preset-1 filter-preset-1">
                    <div class="values">
                        <div class="revenue">
                            <span class="fs-6">Receita:
                                <span class="text-success">
                                    <b class="amount receita">R$ 0,00</b>
                                </span>
                            </span>
                        </div>
                        <div class="expense">
                            <span class="">Despesa:
                                <span class="text-danger">
                                    <b class="amount despesa">R$ 0,00</b>
                                </span>
                            </span>
                        </div>
                    </div>
                    <div class="text-end">
                        <canvas id="graph-pie"></canvas>
                    </div>
                </div>
            </div>
            <div class="col-md-6 mt-2">
                <div class="dashboard dashboard-preset-1 filter-preset-1">
                    <canvas id="graph-bar"></canvas>
                </div>
            </div>
            <div class="col-md-12 mt-2">
                <?php include_once("../table/account-balance-table.php") ?>
            </div>
        </div>
    </div>
    </div>
    </div>
    <?php include_once("../includes/footer.php") ?>
    <script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
    <script src="../js/graph/script.js"></script>
    <script src="../js/dashboard/script.js"></script>
</body>

</html>