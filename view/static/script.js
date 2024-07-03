const url = "http://localhost:8080/"

// Константы для сортировки по возрастанию и убыванию
const sortAsc = "asc"
const sortDesc = "desc"

// Флаг для переключения сортировки
let flag = true

// Объект для управления отображением загрузочного спиннера
const Loader = {
    loadingSpinner: document.getElementById('loading-spinner'),
    on() {
        this.loadingSpinner.style.display = 'block'
    },
    off() {
        this.loadingSpinner.style.display = 'none'
    }
}

// Получаем текущий путь
let currentPath = document.getElementById('current-path').innerHTML

// Вызываем функцию загрузки файлов и директорий
upload(currentPath, flag)

// Загружает файлы и директории из указанного пути с сортировкой
async function upload(currentPath, sortFlag) {
    Loader.on()

    let sort = sortAsc

    if (!sortFlag) {
        sort = sortDesc
    }

    await fetch(url + 'path?root=' + currentPath.slice(1, -1) + '&sort=' + sort, {
        method: "GET",
    })
    .then(resp => {
        if (resp.ok) {
            resp.json()
            .then(data => {
                let file_list = document.getElementById('file-list')
                file_list.innerHTML = ""
                data.forEach(element => {
                    if (element["FileType"] === "dir") {
                        file_list.innerHTML += `<div class="file-item" id="directory-item">
                                                    <div class="directory-icon"></div>
                                                    <a href="#" class="name" onclick="navigateToDirectory(event)">${element["Name"]}</a>
                                                    <span class="type">директория</span>
                                                    <span class="size">${element["Size"]}</span>
                                                </div>`
                    }
                    if (element["FileType"] === "file") {
                        file_list.innerHTML += `<div class="file-item">
                                                    <div class="file-icon"></div>
                                                    <div class="name">${element["Name"]}</div>
                                                    <span class="type">файл</span>
                                                    <span class="size">${element["Size"]}</span>
                                                </div>`
                    }
                });
            })
        }
        else {
            resp.text().then(text => {
                alert(text)
            })
        }
        
    })
    .catch(error => {
        console.log(error);
    })

    Loader.off()
}

// Обрабатывает клик на директорию и запускает функцию upload() с новым путем
function navigateToDirectory(event) {
    var clickedElement = event.target;

    let currentPath = document.getElementById('current-path').textContent + clickedElement.textContent + "/"

    document.getElementById('current-path').innerHTML = currentPath

    upload(currentPath, flag)
}

// Обрабатывает клик на кнопку "Назад" и запускает функцию upload() с новым путем.
function backPath() {
    let currentPath = document.getElementById('current-path').textContent

    if (currentPath === "/") {
        alert("Это предел")
        return
    }

    let pathArray = currentPath.split('/');

    pathArray.pop();
    pathArray.pop();

    let newPath = pathArray.join('/') + "/";

    document.getElementById('current-path').innerHTML = newPath

    upload(newPath, flag)
}

// запускает функцию upload() в соответствии с выбранной сортировкой
function sort() {
    let currentPath = document.getElementById('current-path').textContent

    flag = !flag

    if (flag) {
        document.querySelector(".sort-button").style.backgroundImage = "url('static/source/icon/sortAsc.svg')";
    } else {
        document.querySelector(".sort-button").style.backgroundImage = "url('static/source/icon/sortDesc.svg')";
    }

    upload(currentPath, flag)
}