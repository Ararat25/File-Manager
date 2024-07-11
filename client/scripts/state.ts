// Флаг для переключения сортировки
import {SortMethod} from "./consts";

export let flag = SortMethod.Asc;

export function toggleFlag() {
    if (flag === SortMethod.Asc) {
        flag = SortMethod.Desc
    } else {
        flag = SortMethod.Asc
    }
}