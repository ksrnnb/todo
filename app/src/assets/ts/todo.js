var deleteButtons = document.getElementsByClassName("delete-button");
var onclickbtn = function (e) {
    console.log(e);
    // e.preventDefault();
    // const id = e.target.dataset.itemid;
};
Object.keys(deleteButtons).forEach(function (i) {
    deleteButtons[i].onclick = onclickbtn;
});
