const menu = document.getElementById('menu-label');
const sidebar = document.getElementsByClassName('sidebar')[0];

menu.addEventListener('click', function() {
    sidebar.classList.toggle('hide');
    console.log('ok');
})

document.addEventListener('DOMContentLoaded', function() {
    // Set tinggi main content sesuai viewport
    const sidebar = document.querySelector('.sidebar');
    const mainContent = document.querySelector('.main-content');
    
    function setHeights() {
        const windowHeight = window.innerHeight + 'px';
        if (sidebar) sidebar.style.height = windowHeight;
        if (mainContent) mainContent.style.height = windowHeight;
    }
    
    setHeights();
    window.addEventListener('resize', setHeights);
    
    // Atur margin untuk mencegah overflow
    const container = document.querySelector('.container');
    if (container) {
        container.style.margin = '0';
        container.style.padding = '0';
        container.style.maxWidth = '100%';
        container.style.overflowX = 'hidden';
    }

});