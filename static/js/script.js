// accordion
var acc = document.getElementsByClassName("accordion");
var i;
for (i = 0; i < acc.length; i++) {
    acc[i].addEventListener("click", function() {
        this.classList.toggle("active");
        var panel = this.nextElementSibling;
        if (panel.style.maxHeight) {
            panel.style.maxHeight = null;
        } else {
            panel.style.maxHeight = panel.scrollHeight + "px";
        } 
    });
}

var speed = 1;
function incEltNbr(id, count) {
    elt = document.getElementById(id);
    endNbr = Number(document.getElementById(id).innerHTML);
    incNbrRec(0, endNbr, elt, count);
}
function incNbrRec(i, endNbr, elt, count) {
    if (i <= endNbr) {
        elt.innerHTML = i + '+';
        setTimeout(function() {
            incNbrRec(i + count, endNbr, elt, count);
        }, speed);
    }
}