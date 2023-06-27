<!DOCTYPE html>
<html lang="pt-BR">

<head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no" />
    <meta name="description" content="" />
    <meta name="author" content="" />
    <title>Login - Openfinance</title>

    <link rel="icon" type="image/x-icon" href="../img/dollar.png" />
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-9ndCyUaIbzAi2FUVXJi0CjmCapSmO7SnpJef0486qhLnuZ2cdeRhO02iuK6FUUVM" crossorigin="anonymous">
    <link rel="stylesheet" href="./style.css">
</head>

<body>
    <main>
        <div class="row">
            <div id="login-box">
                <div class="col-md-12 text-center">
                    <img id="login-logo" src="../img/openfinance_white.png" alt="Logotipo Openfinance">
                </div>
                <form method="post" id="form-login" name="form_login">
                    <div class="col-md-12 mt-4 alert alert-danger" id="alert-error">
                        <i class="fa-solid fa-triangle-exclamation"></i> <span class="error error-execute"></span>
                    </div>
                    <div class="col-md-12 mt-3">
                        <label class="form-label">E-mail:</label>
                        <input id="email" type="text" class="form-control" name="email" placeholder="E-mail ou apelido">
                    </div>
                    <div class="col-md-12">
                        <div class="text-danger" id="error-msg-email"></div>
                    </div>
                    <div class="col-md-12 mt-2">
                        <label class="form-label">Senha:</label>
                        <div class="input-group">
                            <input id="password" class="form-control" id="password" name="password" type="password" placeholder="Senha">
                            <button class="btn button-eye" id="button-eye" type="button"><i id="icon-eye" class="fa-solid fa-eye"></i></button>
                        </div>
                    </div>
                    <div class="col-md-12">
                        <div class="text-danger" id="error-msg-password"></div>
                    </div>
                    <div class="col-md-12 mt-1">
                        <div class="text-danger" id="error-msg-authentication"></div>
                    </div>
                    <div class="col-md-12 mt-3 d-grid">
                        <button class="btn btn-success" id="button-login" type="button"><b>ENTRAR</b></button>
                    </div>
                </form>
                <div class="col-md-12 mt-2">
                    <a href="./register.php">Criar uma conta?</a>
                    <a href="">Esqueceu a senha?</a>
                </div>
                <div class="col-md-3 mt-2">
                </div>
            </div>
        </div>
    </main>

    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.6.4/jquery.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/js/bootstrap.bundle.min.js"></script>
    <script src="https://kit.fontawesome.com/52047f2aa2.js" crossorigin="anonymous"></script>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.min.js" integrity="sha384-fbbOQedDUMZZ5KreZpsbe1LCZPVmfTnH7ois6mU1QK+m14rQ1l2bGBq41eYeM/fS" crossorigin="anonymous"></script>
    <script src="./script.js"></script>

</body>

</html>