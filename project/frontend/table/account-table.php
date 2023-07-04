<div id="div-table">
    <table class="table table-stripet table-borderless" id="table-account">
        <thead>
            <tr>
                <th class="order-by text-center" onclick="ordenarTabela(0)">
                    Conta
                    <i class="fas fa-sort ml-1"></i>
                </th>
                <th class="order-by text-center" onclick="ordenarTabela(1)">
                    Saldo
                    <i class="fas fa-sort ml-1"></i>
                </th>
                <th class="order-by text-center">
                    Ações
                </th>
            </tr>
        </thead>
        <tbody>
            <tr class="result-table-account text-center">
                <td>Nubank</td>
                <td>R$ 4.500,00</td>
                <td class="text-center">
                    <button class="btn btn-danger button-delete-account" type="button" data-id="1" data-name-account="Nubank" data-balance-account="4500.00"><i class="fa-solid fa-trash"></i></button>
                    <button class="btn btn-primary button-update-account" type="button" data-id="1" data-name-account="Nubank" data-balance-account="4500.00"><i class="fa-solid fa-pen"></i></button>
                </td>
            </tr>
            <tr class="result-table-account text-center">
                <td>Inter</td>
                <td>R$ 2.300,00</td>
                <td class="text-center">
                    <button class="btn btn-danger button-delete-account" type="button" data-id="2" data-name-account="Inter" data-balance-account="2300.00"><i class="fa-solid fa-trash"></i></button>
                    <button class="btn btn-primary button-update-account" type="button" data-id="2" data-name-account="Inter" data-balance-account="2300.00"><i class="fa-solid fa-pen"></i></button>
                </td>
            </tr>
        </tbody>
    </table>
</div