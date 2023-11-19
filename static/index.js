var btnclick = [];
var json;
var currentItems = 0;
var loadCountItems = 5;

function readServerString(url, callback) {
    let request = new XMLHttpRequest();
    request.onreadystatechange = function() {
        if (request.readyState === 4) {
            callback(undefined, request.responseText);
        } else {
            callback(new Error(request.status));
        }
    }
    request.open("get", url, true);
    request.setRequestHeader("Content-Type","application/x-www-form-urlencoded; charset=utf-8");
    request.send();
}

function call() {
    let input = document.getElementById("request").value;

    if (input == ""){
        deleteAllCategories();
        deleteAllItems();
        return;
    }

    let url = '/filter/kpgz/' + input;
    deleteAllItems();
    deleteAllCategories();

    readServerString(url, function(err, response) {
        if (!err) {
            current_items = 0;
            json = response;

            objs = JSON.parse(response);

            let items = document.getElementById("items");

            btnclick = [0]
            let params = document.getElementById('parametrs');
            index = 1;
            p = ['Добросовестность исполнения поставок', 'Масштабность поставок', 'Интенсивность заключения контрактов', 'Общий рейтинг']
            createInn(params);
            for (let i in p){
                btnclick.push(1);
                createButton(params, p[i], index);
                index++;
            }


            for (var i = 0; i < objs.data.length && i < (currentItems + loadCountItems); i++){
                createItem(items, objs, i);
            }
            currentItems = i;
        }
    });
}

function compare(){
    deleteAllItems();
    deleteAllCategories();
    let inn = document.getElementById("request_1").value;

    if (inn.length < 3){
        return;
    }
    let url = "filter/inn/" + inn;

    readServerString(url, function(err, response) {
        if (!err) {
            current_items = 0;
            json = response;

            objs = JSON.parse(response);
            let items = document.getElementById("items");
            let params = document.getElementById('parametrs');
            btnclick = [0]
            index = 1;
            p = ['Добросовестность исполнения поставок', 'Масштабность поставок', 'Интенсивность заключения контрактов', 'Общий рейтинг']

            createInn(params);
            for (let i in p){
                btnclick.push(1);
                createButton(params, p[i], index);
                index++;
            }

            for (var i = 0; i < objs.data.length && i < (currentItems + loadCountItems); i++){
                createItem(items, objs, i);
            }
            currentItems = i;
        }
    });
}

function createInn(params){
    let innItem = document.createElement('div')
    innItem.id = 'inn'
    innItem.className = 'inn'
    innItem.textContent = "Инн";
    params.appendChild(innItem);
    params.appendChild(innItem);
}

function createButton(items, obj, index){
    var item = document.createElement('button');
    item.className = "btn";
    item.addEventListener('click', function(){
        sort(index);
    });

    item.textContent = obj;
    items.appendChild(item);
}

function createItem(items, obj, index){
    let item = document.createElement('div');
    item.className = "item";

    let inn = document.createElement('h4');
    let failed_dedlines = document.createElement('h4');
    let avg_udp_contract = document.createElement('h4');
    let activity = document.createElement('h4');
    let total = document.createElement('h4');

    inn.textContent = obj.data[index].inn;
    failed_dedlines.textContent = obj.data[index].failed_dedlines;
    avg_udp_contract.textContent = obj.data[index].avg_udp_contract;
    activity.textContent = obj.data[index].activity;
    total.textContent = obj.data[index].total;

    item.appendChild(inn);
    item.appendChild(failed_dedlines);
    item.appendChild(avg_udp_contract);
    item.appendChild(activity);
    item.appendChild(total);
    items.appendChild(item);
}

function addNextPages(){
    objs = JSON.parse(json);
    let items = document.getElementById("items");

    if (current_items >= objs.data.length){
        return;
    }

    for (var i = currentItems; i < objs.data.length && i < (currentItems + loadCountItems); i++){
        createItem(items, objs, i);
    }

    currentItems = i
}

function deleteAllItems(){
    quaery = document.querySelectorAll('.item')
    for (let i = 0; i < quaery.length; i++){
        quaery[i].remove();
    }
}

function deleteAllCategories(){
    inn = document.getElementById("inn")
    if(inn){
        inn.remove();
    }

    quaery = document.querySelectorAll('.btn');
    for (var i = 0; i < quaery.length; i++){
        quaery[i].remove();
    }
}

function sort(index) {
    deleteAllItems();
    let countItems = 0

    new_json = []
    data = JSON.parse(json)
    let num_i;
    let inx = 0;
    for (var i in data.data[1]){
        if (inx === index){
            num_i = i;
            break;
        }
        inx++;
    }
    let items = document.getElementById('items');

    data.data.sort( function(right, left){
        let rightNum = parseFloat(right[num_i]);
        let leftNum = parseFloat(left[num_i]);

        if (rightNum < leftNum) return -1 * btnclick[index];

        if (rightNum > leftNum) return 1 * btnclick[index];

        return 0;

    }).forEach(function(node) {
        if (countItems < currentItems){
            sortCreateItem(items, node);
        }
        countItems++;
    });

    btnclick[index] *= -1;
}

function sortCreateItem(items, obj){

    let item = document.createElement('div');
    item.className = "item";

    let inn = document.createElement('h4');
    let failed_dedlines = document.createElement('h4');
    let avg_udp_contract = document.createElement('h4');
    let activity = document.createElement('h4');
    let total = document.createElement('h4');

    inn.textContent = obj.inn;
    failed_dedlines.textContent = obj.failed_dedlines;
    avg_udp_contract.textContent = obj.avg_udp_contract;
    activity.textContent = obj.activity;
    total.textContent = obj.total;

    item.appendChild(inn);
    item.appendChild(failed_dedlines);
    item.appendChild(avg_udp_contract);
    item.appendChild(activity);
    item.appendChild(total);
    items.appendChild(item);
}