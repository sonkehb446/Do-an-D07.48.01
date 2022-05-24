
$("#danger-edit-alert").hide();
$("#success-edit-alert").hide();

$(document).ready(function () {
    setValueDepartment();
    setValuePosition();
    setValueBranch();
});

//Thong Bao Loi
function errorEdit(message) {
    var strength = document.getElementById('danger-edit');
    strength.innerHTML = message;
    $("#danger-edit-alert").fadeTo(2000, 500).slideUp(500, function () {
        $("#danger-edit-alert").slideUp(500);
    });

}

function successEdit(message, url) {
    var strength = document.getElementById('succes-edit');

    strength.innerHTML = message;
    $("#success-edit-alert").fadeTo(2000, 500).slideUp(500, function () {
        $("#success-edit-alert").slideUp(500);
        location.href = url;
    });

}


function formatFullName(fullName) {
    str = fullName.replace(/\s+/g, '');
    return str
}

function getName() {
    nameUser = document.getElementById("fullname").value
    return formatFullName(nameUser)
}

function getBirthday() {
    return document.getElementById("birthday").value
}

function getDepartment() {
    return document.getElementById("department-select").value
}

function getBranch() {
    return document.getElementById("branch-select").value
}

function getPosition() {
    return document.getElementById("position-select").value
}

function getPhone() {
    return formatFullName(document.getElementById("phone-number").value)
}

function getDescription() {
    return document.getElementById("description").value
}

function getUserID() {
    return document.getElementById("userID").innerText;
}

function getFb() {
    return document.getElementById("link-fb").value
}

function getEmail() {
    return document.getElementById("Email").value
}

function getimage() {
    return 1;
}

function validateEmail(email) {
    console.log(email);
    var re = /^([a-zA-Z0-9_\.\-])+\@(([a-zA-Z0-9\-])+\.)+([a-zA-Z0-9]{2,4})+$/;
    return re.test(String(email));
}

function trim(s) {
    return s.replace(/^\s+|\s+$/, '');
}

function setValueDepartment() {
    var choose = document.getElementById("choose-department").text
    var x = document.getElementById("department-select");
    var i;
    for (i = 0; i < x.options.length; i++) {
        txt = x.options[i].value;
        if (x.options[i].text == choose) {
            document.getElementById("choose-department").value = x.options[i].value;
        }
    }
}


function setValuePosition() {
    var choose = document.getElementById("choose-position").text
    var x = document.getElementById("position-select");
    var i;
    for (i = 0; i < x.options.length; i++) {
        txt = x.options[i].value;
        if (x.options[i].text == choose) {
            document.getElementById("choose-position").value = x.options[i].value;
        }
    }
}



function setValueBranch() {
    var choose = document.getElementById("branch-position").text
    var x = document.getElementById("branch-select");
    var i;
    for (i = 0; i < x.options.length; i++) {
        txt = x.options[i].value;
        if (x.options[i].text == choose) {
            document.getElementById("branch-position").value = x.options[i].value;
        }
    }
}

function previewFile() {
    const preview = document.querySelector('#imageTemp');
    const file = document.querySelector('input[type=file]').files[0];
    const reader = new FileReader();

    reader.addEventListener("load", function () {
        // convert image file to base64 string
        preview.src = reader.result;
    }, false);

    if (file) {
        reader.readAsDataURL(file);
    }
}


function validatePhone()  {
    var check = true;
    var phone = trim(getPhone())
    if (phone != "") {
        var vnf_regex = /((09|03|07|08|05)+([0-9]{8})\b)/g;
        check = vnf_regex.test(phone)
    }
    return check;
}

function checkEmpty() {
    var check = true;
    var email = getEmail();
    var name = getName();
    var  branch =  getBranch();
    var  department =  getDepartment();
    var  position =  getPosition();
    if (email == "" && name == "") {
        errorEdit("Email or User Name cannot be empty");
        check = false;
    } else if (email == "") {
        errorEdit("Email cannot be empty");
        check = false;
    } else if (name == "") {
        errorEdit("User Name cannot be empty");
        check = false;
    }else if (branch == ""){
        errorEdit("Branch cannot be empty");
        check = false;
    }else if (position == ""){
        errorEdit("Position cannot be empty");
        check = false;
    }else if (department == ""){
        errorEdit("Department cannot be empty");
        check = false;
    }
    return check;
}

function checkLength(str, maxLength, minlength) {
    if (str.length > maxLength || str.length < minlength) {
        return false;
    }
    return true;
}

function checkLengthInput() {
    var check = true;
    var email = getEmail(); //minlength="6" maxlength="255"
    var name = getName(); //minlength="6" maxlength="150"
    var fb = getFb(); // maxlength="200"
    var description = getDescription(); //maxlength="255"
    var phone = getPhone(); // minlength="10" maxlength="12"

    if (!checkLength(email, 225, 6)) {
        errorEdit("Email length less than 225 characters, more than 6 characters");
        check = false;
    } else if (!checkLength(fb, 200, -1)) {
        errorEdit("Facebook length less than 200 characters");
        check = false;
    } else if (!checkLength(name, 150, 6)) {
        errorEdit("UserName length less than 150 characters, more than 6 characters");
        check = false;
    } else if (!checkLength(description, 200, -1)) {
        errorEdit("Description length less than 200 characters");
        check = false;
    }
    else if (!checkLength(phone, 12, -1)) {
        errorEdit("Phone length less than 12 characters");
        check = false;
    }
    return check;
}

// Ajlax Login
$("#editProfile").submit(function (e) {
    e.preventDefault();
    var form = $(this);
    var url = form.attr('action');
    var email = getEmail();
    var name = getName();

    if (checkEmpty() && checkLengthInput()) {
        if (validateEmail(email) == true) {
            var a = DataContext()
            if (validatePhone() == true) {
                var phone = trim(getPhone());
                a.append("phone", phone);
                $.ajax({
                    method: "POST",
                    url: url,
                    dataType: "text",
                    data: a,
                    processData: false,
                    contentType: false,
                    success: function (data) {
                        console.log(data);
                        var obj = JSON.parse(data);
                        if (obj.Error != "" && obj.Error != undefined) {
                            errorEdit(obj.Error);
                        }
                        if (obj.Success != "" && obj.Success != undefined) {
                            var message = String(obj.Success);
                            var url = String(obj.Url);
                            console.log(message, url);
                            successEdit(message, url);
                        }
                    },
                });
            } else {
                errorEdit("Phone is wrong format");
            }
        } else {
            errorEdit("Email is wrong format")
        }
    }
});


function DataContext() {
    let formData = new FormData();
    var useriD = getUserID();
    var email = getEmail().toLowerCase();
    var name = getName();
    var description = getDescription();
    var position = getPosition();
    var branch = getBranch();
    var department = getDepartment();
    var fb = getFb();
    var image = getimage();
    var birthday = getBirthday();

    formData.append("userID", useriD);
    formData.append("name", name);
    formData.append("email", email);

    formData.append("image", image);
    formData.append("facebook", fb);
    formData.append("birthday", birthday);
    formData.append("description", description);
    formData.append("position", position);
    formData.append("branch", branch);
    formData.append("department", department);
    formData.append("imageU", $('input[type=file]')[0].files[0]);
    return formData;
}

