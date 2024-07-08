import {upload} from "./upload";
import {flag} from "./state";

// Обрабатывает клик на директорию и запускает функцию upload() с новым путем
export function navigateToDirectory(event: Event) {
    let clickedElement = event.target;

    let currentPath = (<HTMLDivElement>document.getElementById('current-path')).innerHTML + (<HTMLLinkElement>clickedElement).innerHTML + "/";

    (<HTMLElement>document.getElementById('current-path')).innerHTML = currentPath;

    upload(currentPath, flag);
}