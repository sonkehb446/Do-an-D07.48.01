


function checkPw() {
    var check = false;
    var pw1 = trim(document.getElementById("pw1").value);
    var pw2 = trim(document.getElementById("pw2").value);
    if (pw1 == "" && pw2 == "") {
        errorRspass('Password and confirm password is not empty')
        check = false;
    } else if (pw1 == "") {
        errorRspass('New password is not empty')
        check = false;
    } else if (pw2 == "") {
        errorRspass('Confirm password  is not empty')
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
        return check;
    }
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



    //Lay cac du lieu tu url ra bang Name
    // function getParameterByName(name, url = window.location.href) {
    //     name = name.replace(/[\[\]]/g, '\\$&');
    //     var regex = new RegExp('[?&]' + name + '(=([^&#]*)|&|#|$)'),
    //         results = regex.exec(url);
    //     if (!results) return null;
    //     if (!results[2]) return '';
    //     return decodeURIComponent(results[2].replace(/\+/g, ' '));
    // }


    $("#reset-pw-form").submit(function (e) {
        e.preventDefault();
        var form = $(this);
        var url = window.location.href;
        if ((checkPw() == true)) {
            var pw = trim(document.getElementById("pw1").value);
            $.ajax({
                method: "POST",
                url: url,
                dataType: "text",
                data: { // Danh sách các thuộc tính sẽ gửi đi
                    password: pw,
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

                }, error: function (xhr, textStatus, error) {
                    console.log(xhr.statusText);
                    console.log(textStatus);
                    console.log(error);
                }
            });
        }
    });
