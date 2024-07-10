import { Loader } from "./loader";
import { sortAsc, sortDesc } from "./consts";
import {render} from "./render";

// Загружает файлы и директории из указанного пути с сортировкой
export async function upload(currentPath: string, sortFlag: boolean) {
    Loader.on()

    let sort = sortAsc

    if (!sortFlag) {
        sort = sortDesc
    }

    await fetch(`/path?root=${currentPath}&sort=${sort}`, {
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

    Loader.off()
}