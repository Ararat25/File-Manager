// Обрабатывает клик на директорию и запускает функцию upload() с новым путем
function navigateToDirectory(event: Event) {
    let clickedElement = event.target;

    let currentPath = (<HTMLDivElement>document.getElementById('current-path')).innerHTML + (<HTMLLinkElement>clickedElement).innerHTML + "/";

    (<HTMLElement>document.getElementById('current-path')).innerHTML = currentPath;

    upload(currentPath, flag);
}

// Обрабатывает клик на кнопку "Назад" и запускает функцию upload() с новым путем.
function backPath() {
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

    (<HTMLDivElement>document.getElementById('current-path')).innerHTML = newPath

    upload(newPath, flag)
}

// Запускает функцию upload() в соответствии с выбранной сортировкой
function sort() {
    let currentPath = (<HTMLDivElement>document.getElementById('current-path')).innerHTML

    flag = !flag

    if (flag) {
        document.querySelector<HTMLButtonElement>(".sort-button")!.style.backgroundImage = "url('static/source/icon/sortAsc.svg')";
    } else {
        document.querySelector<HTMLButtonElement>(".sort-button")!.style.backgroundImage = "url('static/source/icon/sortDesc.svg')";
    }

    upload(currentPath, flag)
}