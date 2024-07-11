import {Loader} from "./loader";
import {SortMethod} from "./consts";
import {render} from "./render";

// Загружает файлы и директории из указанного пути с сортировкой
export function upload(currentPath: string, sortFlag: string) {
    Loader.show()

    let sort = SortMethod.Asc

    if (sortFlag === SortMethod.Desc) {
        sort = SortMethod.Desc
    }

    fetch(`/path?root=${currentPath}&sort=${sort}`, {
        method: "GET",
    })
        .then(resp => {
            if (resp.ok) {
                resp.json()
                    .then(data => {
                        if (data["Status"] === 200){
                            render(data)
                        }
                        else {
                            alert("Status: " + data["Status"] + "\nError: " + data["Error"])
                        }
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
        .finally(() => {
            Loader.hide()
        })
}