$("#danger-Delete-alert").hide();
$("#success-Delete-alert").hide();
$("#danger-Edit-alert").hide();
$("#success-Edit-alert").hide();

function getID() {
    return document.getElementById('DeleteID').value;
}

function getIDEdit() {
    return document.getElementById('EditID').value;
}


//Thong Bao Loi
function errorDelete(message) {
    var strength = document.getElementById('danger-Delete');
    strength.innerHTML = message;
    $("#danger-Delete-alert").fadeTo(2000, 500).slideUp(500, function () {
        $("#danger-Delete-alert").slideUp(500);
    });

}

function successDelete(message) {
    var strength = document.getElementById('succes-Delete');
    strength.innerHTML = message;
    $("#success-Delete-alert").fadeTo(2000, 500).slideUp(500, function () {
        $("#success-Delete-alert").slideUp(500);
        $("#DeleteUser").modal('hide');

        window.location.reload();
    });

}


function errorEdit(message) {
    var strength = document.getElementById('danger-Edit');
    strength.innerHTML = message;
    $("#danger-Edit-alert").fadeTo(2000, 500).slideUp(500, function () {
        $("#danger-Edit-alert").slideUp(500);
    });

}

function successEdit(message, url) {
    var strength = document.getElementById('succes-Edit');
    strength.innerHTML = message;
    $("#success-Edit-alert").fadeTo(2000, 500).slideUp(500, function () {
        $("#success-Edit-alert").slideUp(500);
        $("#Edit-row").modal('hide');
        window.location.reload();
        if (url != "") {
            location.href = url;
        }
    });

}

function formatSpace(name) {
    str = name.replace(/\s+/g, '');
    return str
}

$("#formDelete").submit(function (e) {
    e.preventDefault();
    var form = $(this);
    var id = getID()
    var link = form.attr('action') + id;
    var url = formatSpace(link)
    $.ajax({
        method: "POST",
        url: url,
        dataType: "text",
        success: function (data) {
            console.log(data);
            var obj = JSON.parse(data);
            if (obj.Error != "" && obj.Error != undefined) {
                errorDelete(obj.Error);
            }
            if (obj.Success != "" && obj.Success != undefined) {
                var message = String(obj.Success);
                successDelete(message);
            }
        },
    });
});

var table = document.getElementById('userTable');
for (var i = 1; i < table.rows.length; i++) {
    table.rows[i].onclick = function () {
        var name = this.cells[1].innerHTML;
        var id = this.cells[0].innerHTML;
        var email = this.cells[2].innerHTML;
        var role = this.cells[3].innerHTML;
        document.getElementById("name").innerHTML = name;
        document.getElementById("DeleteID").value = id;
        document.getElementById("EditID").value = id;
        document.getElementById("NameEdit").value = formatSpace(name);
        document.getElementById("EmailEdit").value = formatSpace(email);
        setRole(role)
    };
}

function setRole(role) {
    var radios = document.getElementsByName('genderS');
    for (var radio of radios) {
        if (role.trim() == radio.value.trim()) {
            console.log(radio)
            radio.checked = true;
        }
    }
}

function  getRole(){
    return document.querySelector('input[name="genderS"]:checked').value;
}


function validateEmail(email) {
    console.log(email);
    var re = /^([a-zA-Z0-9_\.\-])+\@(([a-zA-Z0-9\-])+\.)+([a-zA-Z0-9]{2,4})+$/;
    return re.test(String(email));
}


//EDit
$("#formEdit").submit(function (e) {
    e.preventDefault();
    var form = $(this);
    var id = getIDEdit();
    var link = form.attr('action') + id;
    var url = formatSpace(link)
    var email = document.getElementById("EmailEdit").value;
    console.log("fomr", form.serialize(), email)
    if (checkEmpty()) {
        if (validateEmail(String(email)) == true) {
            $.ajax({
                method: "POST",
                url: url,
                data: form.serialize(),
                dataType: "text",
                success: function (data) {
                    console.log(data);
                    var obj = JSON.parse(data);
                    if (obj.Error != "" && obj.Error != undefined) {
                        errorEdit(obj.err);
                    }
                    if (obj.Success != "" && obj.Success != undefined) {
                        var message = String(obj.Success);
                        var url = String(obj.Url);
                        successEdit(message, url);
                    }
                },
            });
        } else {
            errorEdit("Email is wrong format");
        }
    }
});


function checkEmpty() {
    var check = true;
    var name = document.getElementById("NameEdit").value;
    var email = document.getElementById("EmailEdit").value;
    if (name == "" && email == "") {
        errorEdit("Email or Name user is not empty");
        check = false;
    } else if (email == "") {
        errorEdit("Email user is not empty");
        check = false;
    } else if (name == "") {
        errorEdit("Name user is not empty");
        check = false;
    }
    return check;
}
