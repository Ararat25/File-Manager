import {upload} from "./upload";
import {flag} from "./state";
import {backPath} from "./backPath";
import {sort} from "./sort";
import {navigateToDirectory} from "./navigateToDirectory";

// Получаем текущий путь
let currentPath = (<HTMLDivElement>document.getElementById('current-path')).innerHTML

// Вызываем функцию загрузки файлов и директорий
upload(currentPath, flag)

// Добавляем обработчик клика на директории
setInterval(() => {
    let dirItems = document.querySelectorAll('#nameDir');
    dirItems.forEach(item => {
        item.addEventListener('click', navigateToDirectory);
    });
}, 1)

document.querySelector('#back-button').addEventListener('click', backPath);
document.querySelector('#sort-button').addEventListener('click', sort);