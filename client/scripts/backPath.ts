import {upload} from "./upload";
import {flag} from "./state";

// Обрабатывает клик на кнопку "Назад" и запускает функцию upload() с новым путем.
export function backPath() {
    let currentPath = (<HTMLDivElement>document.getElementById('current-path')).innerHTML

    if (currentPath === null) {
        alert("Это предел")
        return
    }

    if (currentPath === "/") {
        alert("Это предел")
        return
    }

    let pathArray = currentPath.split('/');

    pathArray.pop();
    pathArray.pop();

    let newPath = pathArray.join('/') + "/";

    (<HTMLDivElement>document.getElementById('current-path')).innerHTML = currentPath;

    upload(newPath, flag);
}