import {upload} from "./upload";
import {flag} from "./state";
import {backPath} from "./backPath";
import {sort} from "./sort";
import {loadStat} from "./loadStat";

// Получаем текущий путь
let currentPath = (<HTMLDivElement>document.getElementById('current-path')).innerHTML

// Вызываем функцию загрузки файлов и директорий
document.addEventListener("DOMContentLoaded", (event) => {
    upload(currentPath, flag)
})

document.querySelector('#back-button').addEventListener('click', backPath);
document.querySelector('#sort-button').addEventListener('click', sort);
document.querySelector('#btn-stat').addEventListener('click', loadStat);