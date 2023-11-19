var btnclick = [];
var json;
var currentItems = 0;
var loadCountItems = 2;

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
            btnclick = []
            objs = JSON.parse(response);
            let params = document.getElementById('parametrs');
            let items = document.getElementById("items");
            index = 0;
            p = ['inn', 'fail', 'avg', 'active', 'total']
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
    let inn_1 = document.getElementById("request_1").value;
    let inn_2 = document.getElementById("request_2").value;
    
    if (inn_1.length < 3 || inn_2.length < 3){
        return;
    }

    let url = "?inn=inn_1";
    
    readServerString(url, function(err, response) {
        if (!err) {
            current_items = 0;
            json = response;
            btnclick = []
            objs = JSON.parse(response);
            let params = document.getElementById('parametrs');
            let items = document.getElementById("items");
            index = 0;
            p = ['inn', 'fail', 'avg', 'active', 'total']
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

    url = "?inn=inn_2";

    readServerString(url, function(err, response) {
        if (!err) {
            json = response;
            let items = document.getElementById("items");

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
    quaery = document.querySelectorAll('.btn');
    for (var i = 0; i < quaery.length; i++){
        quaery[i].remove();
    }
}

function sort(index) {
    deleteAllItems();
    data = JSON.parse(json)
    let num_i;
    let inx = 0;
    for (var i in data.data){
        if (inx === index){
            num_i = i;
            break;
        }
        inx++;
    }
    let items = document.querySelectorAll('.item')
    /*let itemsList = document.querySelectorAll('.item');
    let itemsArray = [];
    let parent = itemsList[0].parentNode;
    
    for (let i = 0; i < itemsList.length; i++) {
        itemsArray.push(itemsList[i]);
        parent.removeChild(itemsList[i]);
    }*/

    json.s

    /*itemsArray.sort(function(right, left){
        leftChild = Array.from(left.children);
        rightChild = Array.from(right.children);

        let rightNum = parseFloat(rightChild[index].textContent);
        let leftNum = parseFloat(leftChild[index].textContent);

        if (rightNum < leftNum) return -1 * btnclick[index];
        if (rightNum > leftNum) return 1 * btnclick[index];
        return 0;

    }).forEach(function(node) {
        parent.appendChild(node);
    });*/

    btnclick[index] *= -1;*/
}

/*
window.onload = funonload;

function funonload() {
    mainBody = document.body.innerHTML;
}

 function findOnpage(){
    var input = document.getElementById("myfilter").value;
    var table = document.getElementById("items");
    document.body.innerHTML = main_str;
    if (input.length < 3 || input == ' '){
        return
    }
    quaery = document.querySelectorAll(".item")
    let reg = new RegExp(input);
    for (var i = 0; i < quaery.length; i++) {
        var flag = false;
        for (var child of quaery[i].children){                    
            if (String(child.textContent).search(reg) >= 0){
                flag = true;
                break;
            }
        }
        if (!flag){
                quaery[i].style.cssText = 'display: none';
        }
    }
}


function callByCategoria(categoria){
    if (categoria == ""){
        return;
    }
    deleteAllItems();

    var url = "/?categoria=" + categoria;
    readServerString(url, function(err, response){
        var items = document.getElementById("items");
        objs = JSON.parse(response);
        for (let i in objs) {
            addItem(items, objs[i], objs.length - 1 != i);
        }
        funonload();
    });
}
 */