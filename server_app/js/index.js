import 'htmx.org';
import '../css/styles.css'
import htmx from 'htmx.org';

export async function updateHtmxGet() {
    const hash = window.location.hash.substring(1);
    const pageContent = document.getElementById('page-content');

    let path = '/home';
    if (hash) {
        path = '/' + hash;
    }

    console.log(`browsing to partial view: ${path}`);

    await htmx.ajax('get', path, pageContent);
}

console.log("App is running !");

window.onload = updateHtmxGet;
window.onhashchange = updateHtmxGet;

