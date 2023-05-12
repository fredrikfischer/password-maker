

$(document).ready(function () {
    $("#buttonCopy").click(function (e) {
        e.preventDefault();
        var copyText = $("#textDisplay").text()
        navigator.clipboard.writeText(copyText)
        $("#buttonCopy").html("Copied!")
        setTimeout(function() {
        $("#buttonCopy").html("Copy Text")
        }, 3000); 
    });
});