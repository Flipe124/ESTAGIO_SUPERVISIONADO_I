<link href="../css/dashboard/style.css" rel="stylesheet" />
<link rel="icon" type="image/x-icon" href="../img/dollar.png" />
<?php include_once("../includes/header.php") ?>
</head>

<body>
    <?php include_once("../includes/sidebar.php") ?>
    <div class="container-fluid">
        <div class="row">
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