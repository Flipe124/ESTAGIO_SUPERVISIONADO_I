generateTableOperation();

blockInsertDateManual();

// sumRevenueAndFormated();


$("#btn-open-modal-revenue").on("click", function () {
    modalAction("new-revenue", "show")
})

$("#button-new-revenue").on("click", function () {
    modalAction("create", "show")
});

$(".result").on("click", function () {
    $(`#form-update`)[0].reset();
    modalAction("update", "show")
});

$("#button-delete-revenue").on("click", function () {
    console.log($(this).data("id"))
    modalAction("update", "hide")
    modalAction("delete", "show")
});

$('#modal-delete').on('hidden.bs.modal', function (e) {
    modalAction("update", "show")
});

//------------------ Funções ------------------

// Apresentar/esconder modal
function modalAction(modalName, action) {
    $().modal('hide')
    $("#modal-" + modalName).modal(action)
}

function generateTableOperation() {
    var result = document.querySelector('.revenue-table');

    id = 1;
    iconCategory = setIconCategory(2);
    type = "revenue";
    value = 12300.30;
    categoryName = "Salário";
    data = "QuantumTech | Nubank";
    date = "21/07/2023";

    result.innerHTML +=
        `<div class="result filter-preset-1" data-id="${id}" data-category="${categoryName}" data-type="${type}" data-value="${value}">
            <span class="icon-category">
                ${iconCategory}
            </span>
            <span class="description">
                <span class="category">
                    <b>${categoryName}</b>
                </span>
                <div class="data text-secondary">
                    ${data}
                </div>
                <span class="font-size-14 text-secondary">${date}</span>
            </span>
            <div class="value text-success">
                <b class="text-value">R$ 12.300,30</b> <span class="status mb-1 ms-1"><i class="fas fa-check-circle text-danger"></i></span>
            </div>
        </div>`

    sumRevenueAndFormated();
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

    $('.result').each(function () {
        var value = parseFloat($(this).attr('data-value'));

        if (!isNaN(value)) {
            sum += value;
        }
    });

    var sumFormated = sum.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' });

    var divSum = document.querySelector('.sum-revenue');

    divSum.innerHTML = sumFormated;
}


function formatValue(input) {
    var value = input.value.replace(/\D/g, '');

    value = (value / 100).toFixed(2);

    value = value.replace(/\B(?=(\d{3})+(?!\d))/g, '.');
    value = value.replace('.', ',');

    input.value = 'R$ ' + value;
}

function blockInsertDateManual() {
    document.getElementById("update-input-date-operation").addEventListener("keydown", function(event) {
        event.preventDefault();
    });
}