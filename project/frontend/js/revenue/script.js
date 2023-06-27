generateTableOperation();

sumRevenueAndFormated();


$("#btn-open-modal-revenue").on("click", function () {
    modalAction("new-revenue", "show")
})

$("#button-new-revenue").on("click", function () {
    modalAction("create", "show")
});

$(".button-update-operation").on("click", function () {
    console.log($(this).data("id"))
    console.log($(this).data("category"))
    modalAction("update", "show")
});

$(".button-delete-operation").on("click", function () {
    console.log($(this).data("id"))
    console.log($(this).data("category"))
    modalAction("delete", "show")
});

//------------------ Funções ------------------

// Apresentar/esconder modal
function modalAction(modalName, action) {
    $("#modal-" + modalName).modal(action)
}

function generateTableOperation() {
    var result = document.querySelector('.revenue-table');

    iconCategory = setIconCategory(2)

    result.innerHTML +=
        `<div class="result">
            <span class="icon-category">
                ${iconCategory}
            </span>
            <span class="description">
                <span class="category"><b>Venda</b></span>
                <h6 class="data">
                    Mercado Livre | Nubank
                </h6>
            </span>
            <div class="value text-success">
                <b class="text-value">R$ 99.500,00</b> <span class="status mb-1 ms-1"><i class="fas fa-check-circle text-success"></i></span>
            </div>
            <span class="buttons">
                <button class="btn btn-outline-danger button-delete-operation" type="button" data-id="2" data-category="dois" data-valueDel="99500.00"><i class="fa-solid fa-trash"></i></button>
                <button class="btn btn-primary button-update-operation" type="button" data-id="2" data-category="dois" data-value="99500.00"><i class="fa-solid fa-pen"></i></button>
            </span>
        </div>`
}

function setIconCategory(category) {
    const iconSalary = `<i class="fab fa-sellcast text-success"></i>`;
    const iconCoins = `<i class="fas fa-coins text-warning"></i>`;
    const iconError = `<i class="fas fa-times text-danger"></i>`;

    if (category == 1) {
        return iconSalary;

    } else if (category == 2) {
        return iconCoins;

    } else {
        return iconError;
    }

}


function sumRevenueAndFormated() {
    var sum = 0;

    $('.revenue-table.results .result').each(function () {
        var value = parseFloat($(this).find('.button-update-operation').attr('data-value'));

        if (!isNaN(value)) {
            sum += value;
        }
    });

    var sumFormated = sum.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' });

    var divSum = document.querySelector('.sum-revenue');

    divSum.innerHTML = sumFormated
}
