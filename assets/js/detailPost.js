
function getContext() {
    return document.getElementById("context").innerText;
}

$(document).ready(function () {
    var tag_id = document.getElementById('Content');
    tag_id.innerHTML = getContext();
});




