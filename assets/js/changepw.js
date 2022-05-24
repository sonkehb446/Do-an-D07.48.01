function checkPw() {
    var check = false;
    var pw1 = trim(document.getElementById("pw1").value);
    var pw2 = trim(document.getElementById("pw2").value);
    var currentpw = getCurrentpw();
    if (currentpw == "") {
        errorRspass('Current Password is not empty')
        check = false;
    } else if ((pw1 == "") || (pw2 == "")) {
        errorRspass('Password and confirm password is not empty')
        check = false;
    } else if (pw1 != "" || pw2 != "") {
        if (pw1 != pw2) {
            errorRspass('Password must be duplicate')
            check = false;
        } else {
            if (validatePass(pw1) == true) {
                return true;
            }
            else {
                errorRspass('Password has 8 characters, \n has uppercase and lowercase letters, special characters')
                return false;
            }
        }
    }
    return check;
}


function checkCurrentPw() {
    var pw1 = trim(document.getElementById("pw1").value);
    var currentpw = getCurrentpw();
    if (pw1 == currentpw) {
        errorRspass('The old password cannot be the same as the new password')
        return false;
    }
    return true;
}

function getCurrentpw() {
    return document.getElementById("currentpw").value;
}
$("#success-resetpass-alert").hide();
$("#danger-resetpass-alert").hide();

function validatePass(password) {
    var passw = /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[a-zA-Z]).{8,}$/;
    return passw.test(String(password));
}

function errorRspass(message) {
    var strength = document.getElementById("danger-reset");
    strength.innerHTML = message;
    $("#danger-resetpass-alert").fadeTo(2000, 500).slideUp(500, function () {
        $("#danger-resetpass-alert").slideUp(500);

    });

}

function successRepass(message, url) {
    var strength = document.getElementById('succes-reset');
    strength.innerHTML = message;
    $("#success-resetpass-alert").fadeTo(2000, 500).slideUp(500, function () {
        $("#success-resetpass-alert").slideUp(500);
        location.href = url;
    });

}

function trim(s) {
    return s.replace(/^\s+|\s+$/, '');
}

$("#changepass-form").submit(function (e) {
    e.preventDefault();
    var form = $(this);
    var url = window.location.href;
    if ((checkPw() == true)) {
        if (checkCurrentPw()) {
            var pw = trim(document.getElementById("pw1").value);
            var currentPw = getCurrentpw();
            $.ajax({
                method: "POST",
                url: url,
                dataType: "text",
                data: { // Danh sách các thuộc tính sẽ gửi đi
                    password: pw,
                    currentpw: currentPw,
                },
                success: function (mapValue) {
                    var obj = JSON.parse(mapValue);
                    console.log(obj);
                    if (obj.Error != "" && obj.Error != undefined) {
                        errorRspass(obj.Error);
                    }
                    if (obj.Success != "" && obj.Success != undefined) {
                        var message = String(obj.Success);
                        var url = String(obj.Url);
                        console.log(message, url)
                        successRepass(message, url);
                    }
                }
            });
        }
    }
}); 