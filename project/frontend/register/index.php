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
    <link rel="stylesheet" href="../login/style.css">
</head>

<body>
    <main>
        <div class="mt-4 ms-4">
            <a class="btn btn-light" href="../login/index.php"><b><i class="fa-solid fa-chevron-left"></i> VOLTAR</b></a>
        </div>
        <div class="row">
            <div id="login-box">
                <div class="col-md-12 text-center">
                    <img id="login-logo" src="../img/openfinance_white.png" alt="Logotipo Openfinance">
                </div>
                <form method="post" id="form-login" name="form_login">
                    <div class="col-md-12 mt-4 alert alert-danger" id="alert-error">
                        <i class="fa-solid fa-triangle-exclamation"></i> <span class="error error-execute"></span>
                    </div>

                    <div class="form-floating mt-1">
                        <input type="text" class="form-control" id="name" placeholder="Nome" maxlength="100">
                        <label class="text-secondary" for="name">Nome: <span class="text-danger">*</span></label>
                    </div>
                    <span class="text-danger error error-msg-name"></span>

                    <div class="form-floating mt-3">
                        <input type="text" class="form-control" id="username" placeholder="Nome de usuário" maxlength="100">
                        <label class="text-secondary" for="username">Nome de usuário: <span class="text-danger">*</span></label>
                    </div>
                    <span class="text-danger error error-msg-username"></span>

                    <div class="form-floating mt-3">
                        <input type="text" class="form-control" id="email" placeholder="Email" maxlength="100">
                        <label class="text-secondary" for="email">Email: <span class="text-danger">*</span></label>
                    </div>
                    <span class="text-danger error error-msg-email"></span>

                    <div class="form-floating mt-3">
                        <div class="input-group">
                            <input class="form-control password-field" id="password" name="password" type="password">
                            <button class="btn button-eye" id="button-eye" type="button"><i id="icon-eye" class="fa-solid fa-eye"></i></button>
                        </div>
                        <label class="placeholder-text text-secondary" for="password">Senha: <span class="placeholder-red">*</span></label>
                    </div>
                    <span class="text-danger error error-msg-password"></span>

                    <div class="form-floating mt-3">
                        <div class="input-group">
                            <input class="form-control password-field" id="password-repeat" name="password_repeat" type="password">
                            <button class="btn button-eye" id="button-eye-repeat" type="button"><i id="icon-eye-repeat" class="fa-solid fa-eye"></i></button>
                        </div>
                        <label class="placeholder-text text-secondary" for="password-repeat">Confirme a senha: <span class="placeholder-red">*</span></label>
                    </div>
                    <span class="text-danger error error-msg-password-repeat"></span>
                    <div class="col-md-12 mt-1">
                        <span class="text-danger">Campos obrigatório *</span>
                    </div>
                    <div class="col-md-12 mt-1">
                        <div class="text-danger" id="error-msg-authentication"></div>
                    </div>
                    <div class="col-md-12 mt-3 d-grid">
                        <button class="btn btn-success" id="button-register" type="button"><b>REGISTRAR</b></button>
                    </div>
                </form>
            </div>
        </div>
        <?php require_once("../modal/modal-message.php"); ?>
    </main>

    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.6.4/jquery.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/js/bootstrap.bundle.min.js"></script>
    <script src="https://kit.fontawesome.com/52047f2aa2.js" crossorigin="anonymous"></script>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.min.js" integrity="sha384-fbbOQedDUMZZ5KreZpsbe1LCZPVmfTnH7ois6mU1QK+m14rQ1l2bGBq41eYeM/fS" crossorigin="anonymous"></script>
    <script src="./script.js"></script>

</body>

</html>