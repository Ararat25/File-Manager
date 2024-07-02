const url = "http://localhost:8080/"

currentPath = document.getElementById('current-path').innerHTML 
upload(currentPath)

async function upload(currentPath) {
    document.getElementById('loading-spinner').style.display = 'block';

    await fetch(url + 'path?root=' + currentPath.slice(1, -1) + '&sort=desc', {
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
                                                    <a href="#" class="name" onclick="getCurrPath(event)">${element["Name"]}</a>
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

    document.getElementById('loading-spinner').style.display = 'none';
}

function getCurrPath(event) {
    var clickedElement = event.target;

    let currentPath = document.getElementById('current-path').textContent + clickedElement.textContent + "/"

    document.getElementById('current-path').innerHTML = currentPath

    upload(currentPath)
}

function backPath() {
    let currentPath = document.getElementById('current-path').textContent

    if (currentPath === "/") {
        alert("Это предел")
        return
    }

    let pathArray = currentPath.split('/');

    console.log(pathArray)

    pathArray.pop();
    pathArray.pop();

    let newPath = pathArray.join('/') + "/";

    document.getElementById('current-path').innerHTML = newPath

    upload(newPath)
}