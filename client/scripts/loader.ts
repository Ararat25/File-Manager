// Объект для управления отображением загрузочного спиннера
export const Loader = {
    loadingSpinner: <HTMLDivElement>document.getElementById('loading-spinner'),
    on() {
        this.loadingSpinner.style.display = 'block'
    },
    off() {
        this.loadingSpinner.style.display = 'none'
    }
}