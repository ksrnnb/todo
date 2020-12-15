
function setMethod() {
    const itemButtons = document.getElementsByClassName("item-button");

    const setMethodInput = e => {
        const input = document.createElement("input");
        input.setAttribute("name", "_method");
        input.setAttribute("type", "hidden");
        input.setAttribute("value", e.target.dataset.method);

        e.target.parentNode.appendChild(input);
    }

    Object.keys(itemButtons).forEach(i => {
        itemButtons[i].onclick = setMethodInput;
    });
}

setMethod();