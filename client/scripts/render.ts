import {navigateToDirectory} from "./navigateToDirectory";

export function render(data: any) {
    (<HTMLDivElement>document.getElementById('spent-time')).innerHTML = `Затраченное время: ${data["TimeSpent"]} мс`;
    (<HTMLDivElement>document.getElementById('current-path')).innerHTML = data["Root"];

    let file_list = <HTMLDivElement>document.getElementById('file-list')
    file_list.innerHTML = ""
    data['Files'].forEach((element: any) => {
        if (element["FileType"] === "dir") {
            const fileItem = document.createElement("div");
            fileItem.classList.add("file-item");
            fileItem.id = "directory-item";

            const directoryIcon = document.createElement("div");
            directoryIcon.classList.add("directory-icon");

            const name = document.createElement("div");
            name.classList.add("name");
            name.id = "nameDir";
            name.textContent = element["Name"];
            name.addEventListener("click", navigateToDirectory)

            const typeSpan = document.createElement("span");
            typeSpan.classList.add("type");
            typeSpan.textContent = "директория";

            const sizeSpan = document.createElement("span");
            sizeSpan.classList.add("size");
            sizeSpan.textContent = element["Size"];

            fileItem.appendChild(directoryIcon);
            fileItem.appendChild(name);
            fileItem.appendChild(typeSpan);
            fileItem.appendChild(sizeSpan);

            file_list.appendChild(fileItem);
        }
        if (element["FileType"] === "file") {
            const fileItem = document.createElement("div");
            fileItem.classList.add("file-item");

            const directoryIcon = document.createElement("div");
            directoryIcon.classList.add("file-icon");

            const name = document.createElement("div");
            name.classList.add("name");
            name.textContent = element["Name"];

            const typeSpan = document.createElement("span");
            typeSpan.classList.add("type");
            typeSpan.textContent = "файл";

            const sizeSpan = document.createElement("span");
            sizeSpan.classList.add("size");
            sizeSpan.textContent = element["Size"];

            fileItem.appendChild(directoryIcon);
            fileItem.appendChild(name);
            fileItem.appendChild(typeSpan);
            fileItem.appendChild(sizeSpan);

            file_list.appendChild(fileItem);
        }
    });
}