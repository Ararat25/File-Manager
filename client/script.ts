const url = "http://localhost:8080/"

// Константы для сортировки по возрастанию и убыванию
const sortAsc = "asc"
const sortDesc = "desc"

// Флаг для переключения сортировки
let flag = true

// Получаем текущий путь
let currentPath = (<HTMLDivElement>document.getElementById('current-path')).innerHTML

// Вызываем функцию загрузки файлов и директорий
upload(currentPath, flag)