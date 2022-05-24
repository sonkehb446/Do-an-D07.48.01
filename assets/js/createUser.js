
$("#danger-create-alert").hide();
$("#success-create-alert").hide();

//Thong Bao Loi
function errorCreate(message) {
    var strength = document.getElementById('danger-create');
    strength.innerHTML = message;
    $("#danger-create-alert").fadeTo(2000, 500).slideUp(500, function () {
        $("#danger-create-alert").slideUp(500);
    });

}

function successCreate(message, url) {
    var strength = document.getElementById('succes-create');
    strength.innerHTML = message;
    $("#success-create-alert").fadeTo(2000, 500).slideUp(500, function () {
        $("#success-create-alert").slideUp(500);
        if (url != "") {
            location.href = url;
        }

    });

}


function formatSpace(fullName) {
    str = fullName.replace(/\s+/g, '');
    return str
}

function getName() {
    nameUser = document.getElementById("UserName").value
    return formatSpace(nameUser)
}

function getEmail() {
    return formatSpace(document.getElementById("Email").value)
}

function  getRole(){
    return document.querySelector('input[name="genderS"]:checked').value;
}


function getPassword() {
    return formatSpace(document.getElementById("Password").value)
}

function validateEmail(email) {
    console.log(email);
    var re = /^([a-zA-Z0-9_\.\-])+\@(([a-zA-Z0-9\-])+\.)+([a-zA-Z0-9]{2,4})+$/;
    return re.test(String(email));
}

function validatePass(password) {
    var passw = /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[a-zA-Z]).{8,}$/;
    return passw.test(String(password));
}

function checkLength(str, maxLength, minlength) {
    if (str.length > maxLength || str.length < minlength) {
        return false;
    }
    return true;
}

function checkEmpty() {
    var check = true;
    var email = getEmail()
    var pw = getPassword()
    var name = getName()
    console.log(checkLength("sssss", 5))
    if (email == "" && name == "" && pw == "") {
        errorCreate("Email or userName or password cannot be empty")
        check = false;
    } else if (email == "") {
        errorCreate("Email cannot be empty")
        check = false;
    } else if (name == "") {
        errorCreate("UserName cannot be empty")
        check = false;
    }
    return check;
}


function checkLengthInput(){
    var check = true;
    var email = getEmail();  // maxlength="255" minlength="6"
    var pw = getPassword(); // maxlength="255" minlength="6"
    var name = getName();// minlength="6" maxlength="150"

    if (!checkLength(email,225,6)){
        errorCreate("Email length less than 225 characters, more than 6 characters");
        check = false;
    }else if (!checkLength(pw,225,8)){
        errorCreate("Password length less than 225 characters, more than 8 characters");
        check = false;
    }else if (!checkLength(name,150,6)){
        errorCreate("UserName length less than 150 characters, more than 6 characters");
        check = false;
    }
    return check;
}


$("#create-user").submit(function (e) {
    e.preventDefault();
    var form = $(this);
    var url = form.attr('action');
    var email = getEmail()
    var pw = getPassword()
    var name = getName()
    var role = getRole()
    if (checkEmpty() && checkLengthInput()) {
        if (validateEmail(email) == true) {
            if (validatePass(pw) == true) {
                $.ajax({
                    method: "POST",
                    url: url,
                    dataType: "text",
                    data: { // Danh sách các thuộc tính sẽ gửi đi
                        email: email.toLowerCase(),
                        password: pw,
                        fullName: name,
                        role: role
                    },
                    success: function (data) {
                        console.log(data);
                        var obj = JSON.parse(data);
                        if (obj.Error != "" && obj.Error != undefined) {
                            errorCreate(obj.Error);
                        }
                        if (obj.Success != "" && obj.Success != undefined) {
                            var message = String(obj.Success);
                            var url = String(obj.Url);
                            console.log(message, url);
                            successCreate(message, url);
                        }
                    },
                });
            } else {
                errorCreate("Password has 8 characters, \n has uppercase and lowercase letters, special characters")
            }
        }
        else {
            errorCreate("Email is wrong format")
        }
    }
});