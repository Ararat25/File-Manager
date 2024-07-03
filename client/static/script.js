var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (g && (g = 0, op[0] && (_ = 0)), _) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
var url = "http://localhost:8080/";
var sortAsc = "asc";
var sortDesc = "desc";
var flag = true;
var Loader = {
    loadingSpinner: document.getElementById('loading-spinner'),
    on: function () {
        this.loadingSpinner.style.display = 'block';
    },
    off: function () {
        this.loadingSpinner.style.display = 'none';
    }
};
var currentPath = document.getElementById('current-path').innerHTML;
upload(currentPath, flag);
function upload(currentPath, sortFlag) {
    return __awaiter(this, void 0, void 0, function () {
        var sort;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    Loader.on();
                    sort = sortAsc;
                    if (!sortFlag) {
                        sort = sortDesc;
                    }
                    return [4, fetch(url + 'path?root=' + currentPath.slice(1, -1) + '&sort=' + sort, {
                            method: "GET",
                        })
                            .then(function (resp) {
                            if (resp.ok) {
                                resp.json()
                                    .then(function (data) {
                                    var file_list = document.getElementById('file-list');
                                    file_list.innerHTML = "";
                                    data.forEach(function (element) {
                                        if (element["FileType"] === "dir") {
                                            file_list.innerHTML += "<div class=\"file-item\" id=\"directory-item\">\n                                                    <div class=\"directory-icon\"></div>\n                                                    <a href=\"#\" class=\"name\" onclick=\"navigateToDirectory(event)\">".concat(element["Name"], "</a>\n                                                    <span class=\"type\">\u0434\u0438\u0440\u0435\u043A\u0442\u043E\u0440\u0438\u044F</span>\n                                                    <span class=\"size\">").concat(element["Size"], "</span>\n                                                </div>");
                                        }
                                        if (element["FileType"] === "file") {
                                            file_list.innerHTML += "<div class=\"file-item\">\n                                                    <div class=\"file-icon\"></div>\n                                                    <div class=\"name\">".concat(element["Name"], "</div>\n                                                    <span class=\"type\">\u0444\u0430\u0439\u043B</span>\n                                                    <span class=\"size\">").concat(element["Size"], "</span>\n                                                </div>");
                                        }
                                    });
                                });
                            }
                            else {
                                resp.text().then(function (text) {
                                    alert(text);
                                });
                            }
                        })
                            .catch(function (error) {
                            console.log(error);
                        })];
                case 1:
                    _a.sent();
                    Loader.off();
                    return [2];
            }
        });
    });
}
function navigateToDirectory(event) {
    var clickedElement = event.target;
    var currentPath = document.getElementById('current-path').innerHTML + clickedElement.innerHTML + "/";
    document.getElementById('current-path').innerHTML = currentPath;
    upload(currentPath, flag);
}
function backPath() {
    var currentPath = document.getElementById('current-path').innerHTML;
    if (currentPath === null) {
        alert("Это предел");
        return;
    }
    if (currentPath === "/") {
        alert("Это предел");
        return;
    }
    var pathArray = currentPath.split('/');
    pathArray.pop();
    pathArray.pop();
    var newPath = pathArray.join('/') + "/";
    document.getElementById('current-path').innerHTML = newPath;
    upload(newPath, flag);
}
function sort() {
    var currentPath = document.getElementById('current-path').innerHTML;
    flag = !flag;
    if (flag) {
        document.querySelector(".sort-button").style.backgroundImage = "url('static/source/icon/sortAsc.svg')";
    }
    else {
        document.querySelector(".sort-button").style.backgroundImage = "url('static/source/icon/sortDesc.svg')";
    }
    upload(currentPath, flag);
}
