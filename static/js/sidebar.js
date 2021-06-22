console.log("sidebar.js loaded")

// function menuBurger() {
//     $(document).ready(function(){
//         $('.burger').click(function(){
//             $('.menu').toggleClass('isOpen');
//             $(this).find('.barre').toggleClass('activate');
//         });
//     });
// }
// menuBurger()

let header1 = document.getElementsByClassName('en-tete')[0]
let header2 = document.getElementsByClassName('en-tete')[1]
let header3 = document.getElementsByClassName('en-tete')[2]
let menu = document.getElementsByClassName('menu')[0]
console.log(header1)

header1.addEventListener('click', () => {
    menu.classList.toggle('isOpen')
})
header2.addEventListener('click', () => {
    menu.classList.toggle('isOpen')
})
header3.addEventListener('click', () => {
    menu.classList.toggle('isOpen')
})
