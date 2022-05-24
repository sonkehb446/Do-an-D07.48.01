
$("#danger-login-alert").hide();
$("#danger-fwpass-alert").hide();
$("#success-fwpass-alert").hide();

//Thong Bao Loi
function errorLogin(message) {
    var strength = document.getElementById('danger-login');
    strength.innerHTML = message;
    $("#danger-login-alert").fadeTo(2000, 500).slideUp(500, function () {
        $("#danger-login-alert").slideUp(500);
    });

}


function errorFwpass(message) {
    var strength = document.getElementById('danger-email');
    strength.innerHTML = message;
    $("#danger-fwpass-alert").fadeTo(2000, 500).slideUp(500, function () {
        $("#danger-fwpass-alert").slideUp(500);
    });

}

function success(message) {
    var strength = document.getElementById('succes-email');

    strength.innerHTML = message;
    $("#success-fwpass-alert").fadeTo(2000, 500).slideUp(500, function () {
        $("#success-fwpass-alert").slideUp(500);
        $("#myModal").modal('hide');
    });

}

function trim(s) {
    return s.replace(/^\s+|\s+$/, '');
}

function validateEmail(email) {
    console.log(email);
    var re = /^([a-zA-Z0-9_\.\-])+\@(([a-zA-Z0-9\-])+\.)+([a-zA-Z0-9]{2,4})+$/;
    return re.test(String(email));
}

// Ajlax Login
$("#login-form").submit(function (e) {
    e.preventDefault();
    var form = $(this);
    var url = form.attr('action');
    var pw = trim(document.getElementById("password").value);
    var email = trim(document.getElementById("email").value);

    if (checkEmpty()) {
        if (validateEmail(String(email)) == true && pw !== "" && pw !== undefined) {
            $.ajax({
                method: "POST",
                url: url,
                dataType: "text",
                data: { // Danh sách các thuộc tính sẽ gửi đi
                    password: pw,
                    email: email.toLowerCase()
                },
                success: function (mapValue) {
                    console.log(mapValue);
                    var obj = JSON.parse(mapValue);
                    console.log(obj.Url);
                    if ((obj.Error != null && obj.Error !== undefined) || (obj.Error != "" && obj.Error != undefined)) {
                        document.getElementById("password").value = "";
                        errorLogin(obj.Error)
                    } else if ((obj.Url != null && obj.Url != undefined) || (obj.Url != "" && obj.Url != undefined)) {
                        location.href = obj.Url;
                    }

                },
            });
        } else {
            document.getElementById("password").value = "";
            errorLogin("Email is wrong format");
        }
    }
});

function checkEmpty() {
    var check = true;
    var pw = trim(document.getElementById("password").value);
    var email = trim(document.getElementById("email").value);
    if (pw == "" && email == "") {
        errorLogin("Email or password is not empty");
        check = false;
    } else
        if (pw == "") {
            errorLogin("Password is not empty");
            check = false;
        } else if (email == "") {
            errorLogin("Email is not empty");
            check = false;
        }
    return check;
}

// Ajlax SendEmail
$("#sendEmail").submit(function (e) {
    e.preventDefault();
    var form = $(this);
    var url = form.attr('action');
    var email = trim(document.getElementById("tempemail").value);

    if (validateEmail(String(email)) === true) {
        $.ajax({
            method: "POST",
            url: url,
            dataType: "text",
            data: {
                email: email.toLowerCase()
            },
            success: function (mapValue) {
                console.log(mapValue);
                const obj = JSON.parse(mapValue);
                if ((obj.Error != null && obj.Error !== undefined) || (obj.Error != "" && obj.Error != undefined)) {
                    errorFwpass(obj.Error);
                    console.log(obj.Error);
                } else if((obj.Success != null && obj.Success !== undefined) || (obj.Success != "" && obj.Success != undefined)){
                    success(obj.Success);
                    console.log(obj.Success);
                    console.log(obj.Success);
                }
            },
            error: function (e) {
                console.log(e);
            },
        });
    } else if (email == "") {
        errorFwpass("Email is not empty")
    } else {
        errorFwpass("Email is wrong format")
    }
});

