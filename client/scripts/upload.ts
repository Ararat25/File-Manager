import { Loader } from "./loader";
import { url, sortAsc, sortDesc } from "./consts";

// Загружает файлы и директории из указанного пути с сортировкой
export async function upload(currentPath: string, sortFlag: boolean) {
    Loader.on()

    let sort = sortAsc

    if (!sortFlag) {
        sort = sortDesc
    }

    await fetch(url + '/path?root=' + currentPath.slice(1, -1) + '&sort=' + sort, {
        method: "GET",
    })
        .then(resp => {
            if (resp.ok) {
                resp.json()
                    .then(data => {
                        let file_list = <HTMLDivElement>document.getElementById('file-list')
                        file_list.innerHTML = ""
                        data['Files'].forEach((element: any) => {
                            if (element["FileType"] === "dir") {
                                file_list.innerHTML += `<div class="file-item" id="directory-item">
                                                    <div class="directory-icon"></div>
                                                    <a href="#" class="name" id="nameDir">${element["Name"]}</a>
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
                resp.json().then(text => {
                    alert("Status: " + text["Status"] + "\nError: " + text["Error"])
                })
            }

        })
        .catch(error => {
            console.log(error);
        })

    Loader.off()
}