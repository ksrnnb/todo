function setMethod() {
    var itemButtons = document.getElementsByClassName("item-button");
    var setMethodInput = function (e) {
        var input = document.createElement("input");
        input.setAttribute("name", "_method");
        input.setAttribute("type", "hidden");
        input.setAttribute("value", e.target.dataset.method);
        e.target.parentNode.appendChild(input);
    };
    Object.keys(itemButtons).forEach(function (i) {
        itemButtons[i].onclick = setMethodInput;
    });
}
setMethod();
