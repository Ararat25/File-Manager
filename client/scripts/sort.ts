import {flag, toggleFlag} from "./state";
import {upload} from "./upload";
import {SortMethod} from "./consts";

// запускает функцию upload() в соответствии с выбранной сортировкой
export function sort() {
    let currentPath = (<HTMLDivElement>document.getElementById('current-path')).innerHTML

    toggleFlag()

    if (flag == SortMethod.Asc) {
        document.querySelector<HTMLButtonElement>(".sort-button")!.classList.add('asc');
        document.querySelector<HTMLButtonElement>(".sort-button")!.classList.remove('desc');
    } else {
        document.querySelector<HTMLButtonElement>(".sort-button")!.classList.add('desc');
        document.querySelector<HTMLButtonElement>(".sort-button")!.classList.remove('asc');
    }

    upload(currentPath, flag)
}