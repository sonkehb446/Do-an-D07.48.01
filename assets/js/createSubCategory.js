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

function getsubCategory() {
    return document.getElementById('subCategory').value;
}

function getDescription() {
    return document.getElementById('description').value;
}

function getCategory() {
    return document.getElementById('categoryID').value;
}

function validateForm() {
    subcCategory = getsubCategory();
    description = getDescription();
    category = getCategory();
    if (category == "<--select-->" || category == '') {
        errorCreate("Please select category")
        return false;
    } else if (subcCategory == '') {
        errorCreate("Please enter category")
        return false;
    } else if (getDescription() == '') {
        errorCreate("Please enter description")
        return false;
    }
    return true;
}

$("#subCategoryForm").submit(function (e) {
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
    var subCategoryName = getsubCategory();
    var description = getDescription();
    var categoryID = getCategory();
    formData.append("name", subCategoryName);
    formData.append("description", description);
    formData.append("categoryID", categoryID);
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
    var subCategoryName = getsubCategory();
    var description = getDescription();
    var categoryID = getCategory();

    if (!checkLength(subCategoryName, 225, 6)) {
        errorCreate("sub category name length less than 225 characters, more than 6 characters");
        check = false;
    } else if (!checkLength(description, 225, 6)) {
        errorCreate("Description length less than 225 characters, more than 6 characters");
        check = false;
    } else if (categoryID == "<--select-->") {
        errorCreate("select category");
        check = false;
    }
    return check;
}



