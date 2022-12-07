<?php

$ch = curl_init('http://localhost:8008/list?table=finance');

curl_setopt_array($ch, [

    CURLOPT_CUSTOMREQUEST => 'GET',

    // Permite obter o resultado
    CURLOPT_RETURNTRANSFER => 1,
]);

$expensess = json_decode(curl_exec($ch), true);
curl_close($ch);
