export function loadStat() {
    let curUrl = new URL(location.href)
    document.location.href = `${curUrl.protocol}//${curUrl.hostname}:80/page.php`;
}