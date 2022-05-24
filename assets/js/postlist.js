$("#danger-Delete-alert").hide();
$("#success-Delete-alert").hide();
$("#danger-Edit-alert").hide();
$("#success-Edit-alert").hide();

function formatSpace(name) {
    str = name.replace(/\s+/g, '');
    return str
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



function getID() {
    return document.getElementById('DeleteID').value;
}

function getIDEdit() {
    return document.getElementById('EditID').value;
}


$(document).ready(function () {
    var table = document.getElementById('postTable');
    for (var i = 1; i < table.rows.length; i++) {
        table.rows[i].onclick = function () {
            var title = this.cells[1].innerHTML;
            var id = this.cells[0].innerHTML;
            document.getElementById("title").innerHTML = title;
            document.getElementById("DeleteID").value = formatSpace(id);
        };
    }
});


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


