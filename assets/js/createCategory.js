
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

function getCategory() {
    return document.getElementById('category').value;
}

function getDescription() {
    return document.getElementById('description').value;
}
function getMenu() {
    return document.getElementById('menu').value;
}
function validateForm() {
    category = getCategory();
    description = getDescription();
    menu=getMenu();
    if (menu == "<--select-->" || menu =='') {
        errorCreate("Please select menu")
        return false;
    } else if ( category == '') {
        errorCreate("Please enter category")
        return false;
    }else if (getDescription() == '') {
        errorCreate("Please enter description")
        return false;
    }
    return true;
}

$("#categoryForm").submit(function (e) {
    e.preventDefault();
    var form = $(this);
    var url = form.attr('action');
    var a = DataContext()
    if (validateForm() && checkLengthInput()) {
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
                    errorCreate(obj.Error);
                }
                if (obj.Success != "" && obj.Success != undefined) {
                    var message = String(obj.Success);
                    var url = String(obj.Url);
                    successCreate(message, url);
                }
            }
        })
    }
})

function DataContext() {
    let formData = new FormData();
    var categoryName = getCategory();
    var description = getDescription();
    var idMenu = getMenu();
    formData.append("name", categoryName);
    formData.append("description", description);
    formData.append("menuID", idMenu);
   console.log(idMenu);
    return formData;
}


function checkLength(str, maxLength, minlength) {
    if (str.length > maxLength || str.length < minlength) {
        return false;
    }
    return true;
}

function checkLengthInput() {
    var check = true;
    var categoryName = getCategory();
    var description = getDescription();
    var idMenu = getMenu();

    if (!checkLength(categoryName, 225, 6)) {
        errorCreate("categoryName length less than 225 characters, more than 6 characters");
        check = false;
    } else if (!checkLength(description, 225, 6)) {
        errorCreate("Description length less than 225 characters, more than 6 characters");
        check = false;
    } else if (idMenu == "<--select-->") {
        errorCreate("select menu");
        check = false;
    }
    return check;
}



