function toggleSubmenu(id) {
    let submenu = document.getElementById(id)
    submenu.classList.toggle("open")

}

function showContextMenu(id) {
    document.removeEventListener('click', closeContextMenu)
    let context = document.getElementById(id)
    context.classList.toggle("open")
    setTimeout(listenForClick, 100);
}

function listenForClick() {
    document.addEventListener('click', closeContextMenu);
}

function closeContextMenu(e) {
    let contextMenu = document.querySelector('.context-menu.open');
    if (!(contextMenu === null)) {
        contextMenu.classList.toggle('open');
        document.removeEventListener('click', closeContextMenu)
    }
}

function closeModal() {
    let modal = document.querySelector('.modal');
    setTimeout(() => {
        modal.remove();
    }, "100");
}
function closeModal(id) {
    let modal = document.getElementById(id)
    setTimeout(() => {
        modal.remove();
    }, "100");
}

function openTab(id) {
    let tabs = document.querySelectorAll('.tab');
    tabs.forEach(element => {
        element.classList.remove('open');
    });
    let tabToOpen = document.getElementById(id);
    tabToOpen.classList.toggle('open');
}