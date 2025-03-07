let btn = document.querySelector('.Filtre');
let div = document.querySelector('.filtrecache');

btn.addEventListener('click', () => {
    if(div.style.display === 'none') {
        div.style.display = 'block';
    } else {
        div.style.display = 'none';
    }
})