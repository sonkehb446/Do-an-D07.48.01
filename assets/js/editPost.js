
$("#danger-create-alert").hide();
$("#success-create-alert").hide();

//Thong Bao Loi
function errorEdit(message) {
    var strength = document.getElementById('danger-create');
    strength.innerHTML = message;
    $("#danger-create-alert").fadeTo(2000, 500).slideUp(500, function () {
        $("#danger-create-alert").slideUp(500);
    });

}

function successEdit(message, url) {
    var strength = document.getElementById('succes-create');
    strength.innerHTML = message;
    $("#success-create-alert").fadeTo(2000, 500).slideUp(500, function () {
        $("#success-create-alert").slideUp(500);
        if (url != "") {
            location.href = url;
        }

    });

}

function formatSpace(string) {
    str = string.replace(/\s+/g, '');
    return str
}

function LoadImage() {
    const file = document.querySelector('#uploadImage').files[0];
    var reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = function () {
        localStorage.setItem("image", reader.result);
        document.getElementById("image").setAttribute("src", localStorage.getItem("image"))
        document.getElementById("image").style.display = 'block';
    };
}

function getCategory() {
    return document.getElementById('category').value;
}

function getSCategory() {
    return trim(document.getElementById('subCategory').value);
}

function getTitle() {
    return document.getElementById('title').value;
}

function getDescription() {
    return document.getElementById('description').value;
}

function getDataContext() {
    var editorText = CKEDITOR.instances.DSC.getData();
    return editorText;
}

function getFile() {
    input = document.querySelector('input[type="file"]');
    return input.files[0];
}


$('#category').change(function () {
    categoryID = getCategory();
    console.log('chane->', categoryID)
    $.ajax({
        type: 'POST',
        url: '/subCategory/choose',
        json: {},
        data: {
            "id": categoryID,
        },
        success: function (data) {
            console.log(data);
            var $sub = $('#subCategory');
            $sub.empty();
            subCategory = JsonParse(data)
            var size = Object.values(subCategory.sub).length
            console.log(size)
            for (var i = 0; i < size; i++) {
                obj = subCategory.sub[i]
                $sub.append('<option value=' + obj.ID + '>' + obj.CategoryName + '</option>');
                console.log(obj.CategoryName)
            }
        }
    });
})


function getSrcImage() {
    return document.getElementById("imageUpload").src;
}

function previewFile() {
    const preview = document.querySelector('#imageUpload');
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


function trim(s) {
    return s.replace(/^\s+|\s+$/, '');
}

function JsonParse(data) {
    return JSON.parse(JSON.stringify(data))
}


function validateForm() {
    category = getCategory();
    subcategory = getSCategory();
    if (subcategory == "Choose" || subcategory == '') {
        errorEdit("Please select Subcategory")
        return false;
    } else if (category == "Choose" || category == '') {
        errorEdit("Please select sub-category")
        return false;
    } else if (getTitle() == '') {
        errorEdit("Please fill in Title")
        return false;
    } else if (getDescription() == '') {
        errorEdit("Please fill in Description")
        return false;
    } else if (getDataContext() == '') {
        errorEdit("Please fill in DataContext")
        return false;
    }
    return true;
}


// Ajlax Login
$("#Editpost").submit(function (e) {
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
                    errorEdit(obj.Error);
                }
                if (obj.Success != "" && obj.Success != undefined) {
                    var message = String(obj.Success);
                    var url = String(obj.Url);
                    console.log(message, url);
                    successEdit(message, url);
                }
            }
        })
    }
})

function DataContext() {
    let formData = new FormData();
    var categoryID = getSCategory();
    var title = getTitle();
    var description = getDescription();
    var context = getDataContext();
    formData.append("categoryID", categoryID);
    formData.append("title", title);
    formData.append("description", description);
    formData.append("context", context);
    if (document.getElementById("file-input").files.length > 0) {
        formData.append("imageU", $('input[type=file]')[0].files[0]);
    }
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
    var title = getTitle();  // minlength="6" maxlength="255"
    var description = getDescription(); // minlength="6" maxlength="255"
    var context = getDataContext();// minlength="6" maxlength="150"

    if (!checkLength(title, 225, 6)) {
        errorEdit("Title length less than 225 characters, more than 6 characters");
        check = false;
    } else if (!checkLength(description, 225, 6)) {
        errorEdit("Description length less than 225 characters, more than 6 characters");
        check = false;
    } else if (context.length < 100) {
        errorEdit("Context length  more than 100 characters");
        check = false;
    }
    return check;
}




