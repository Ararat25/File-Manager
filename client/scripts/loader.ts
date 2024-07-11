// Объект для управления отображением загрузочного спиннера
export const Loader = {
    loadingSpinner: <HTMLDivElement>document.getElementById('loading-spinner'),
    show() {
        this.loadingSpinner.style.display = 'block'
    },
    hide() {
        this.loadingSpinner.style.display = 'none'
    }
}