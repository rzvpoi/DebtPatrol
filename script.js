const button = document.getElementById("page-button");
const addPage = document.querySelector(".add-wrapper");

const currentDate = document.getElementById("current-date")

const userId = document.getElementById('user-id')
const userVlad = document.querySelector('.vlad')
const userRazvan = document.querySelector('.razvan')


// today date
const date = new Date();
const months = ["January","February","March","April","May","June","July",
"August","September","October","November","December"];
currentDate.innerText = `Today is ${date.getDate()} ${months[date.getMonth()]} ${date.getFullYear()}`

// page switch
let ok = false;
button.addEventListener("click", () => {
    addPage.classList.toggle('add-wrapper-active')
    
    if(!ok) {
        document.body.style.overflow = 'hidden'
        button.innerText = '-'
    }else {
        document.body.style.overflow = 'auto'
        button.innerText = '+'

    }

    ok = !ok;
})

userVlad.addEventListener("click", () => {
    userId.value = 2;

    userVlad.classList.add("user-active")
    userRazvan.classList.remove("user-active")
})

userRazvan.addEventListener("click", () => {
    userId.value = 1;

    userVlad.classList.remove("user-active")
    userRazvan.classList.add("user-active")
})