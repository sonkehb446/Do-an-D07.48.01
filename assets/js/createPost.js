
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

function formatSpace(string) {
    str = string.replace(/\s+/g, '');
    return str
}

function trim(s) {
    return s.replace(/^\s+|\s+$/, '');
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

function validateForm() {
    category = getCategory();
    subcategory = getSCategory();
    if (subcategory == "Choose" || subcategory == '') {
        errorCreate("Please select Subcategory")
        return false;
    } else if (category == "Choose" || category == '') {
        errorCreate("Please select sub-category")
        return false;
    } else if (getTitle() == '') {
        errorCreate("Please fill in Title")
        return false;
    } else if (getDescription() == '') {
        errorCreate("Please fill in Description")
        return false;
    } else if (document.getElementById("file-input").files.length == 0) {
        errorCreate("Please Choose Image")
        return false;
    } else if (getDataContext() == '') {
        errorCreate("Please fill in DataContext")
        return false;
    }
    return true;
}

function previewFile() {
    const preview = document.querySelector('#imageUpload');
    const file = document.querySelector('input[type=file]').files[0];
    const reader = new FileReader();

    reader.addEventListener("load", function () {
        preview.src = reader.result;
        preview.style.display = 'block';
    }, false);

    if (file) {
        reader.readAsDataURL(file);
    }
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

function JsonParse(data) {
    return JSON.parse(JSON.stringify(data))
}



$("#createPost").submit(function (e) {
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
                    console.log(message, url);
                    successCreate(message, url);
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
        errorCreate("Title length less than 225 characters, more than 6 characters");
        check = false;
    } else if (!checkLength(description, 225, 6)) {
        errorCreate("Description length less than 225 characters, more than 6 characters");
        check = false;
    } else if (context.length < 100) {
        errorCreate("Context length  more than 100 characters");
        check = false;
    }
    return check;
}



