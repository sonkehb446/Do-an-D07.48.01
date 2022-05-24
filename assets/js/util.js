
// ẩn hiện menu avatar
var styleLogin = document.getElementById("hind_bnt_login").innerText;
var bnt_login_menu = document.querySelector("#bnt_login_menu");
var avatars = document.querySelector("#img-avatar")
// var nameUser  = document.querySelector("#iname_User_Menu")

if (styleLogin == "" || styleLogin == undefined) {
    bnt_login_menu.style.display = 'block';
    avatars.style.display = 'none';
    // nameUser.style.display = 'none';
} else {
    bnt_login_menu.style.display = 'none';
    avatars.style.display = 'block';
    // nameUser.style.display = 'block';
}

//Role
var user = document.querySelector("#hint_User")
var category = document.querySelector("#hint_Category")
var sub = document.querySelector("#hint_Sub")
var role = document.getElementById("role_user").innerText;

if(role != "Admin"){
    user.style.display = 'none';
    category.style.display = 'none';
    sub.style.display = 'none';
}else{
    user.style.display = 'block';
    category.style.display = 'block';
    sub.style.display = 'block';
}

function BackHistory() {
    history.back();
}

