
function allowDrop(ev) {
    ev.preventDefault();
}

function drag(ev) {
    ev.dataTransfer.setData(" text", ev.target.id);

}

function drop(ev) {
    ev.preventDefault();
    const data = ev.dataTransfer.getData("text");
    const draggedItem = document.getElementById(data);
    const targetItem = ev.target;
    if (!targetItem.classList.contains('item')) {
        ev.target.appendChild(draggedItem);
    } else {
        targetItem.parentNode.insertBefore(draggedItem, targetItem);
    }
}

const init = () => {

    const boxes = document.querySelectorAll("[data-droppable]");
    const inputs = {};

    for (let i = 0; i < boxes.length; i++) {
        const box = boxes[i]
        const input = document.getElementById(box.dataset.target)
        if (!input) {
            throw new Error("Missing input element.")
        }

        inputs[i] = {
            input: input,
            droppable: box
        };
    }

    const handleUpdateValues = () => {
        Object.values(inputs).forEach(({ input, droppable }) => {
            const selected = droppable.querySelectorAll(".item");
            const values = [...selected].map((item) => item.dataset.value)
            input.value = values.join(",")
        })
    }

    const items = document.querySelectorAll("#draggable-list .item");

    if (!boxes.length || !items.length) {
        return
    }
    const handleDrop = (ev) => {
        ev.preventDefault();
        const data = ev.dataTransfer.getData("text");
        const draggedItem = document.getElementById(data);
        const targetItem = ev.target;

        if (!targetItem.classList.contains('item')) {
            ev.target.appendChild(draggedItem);
        } else {
            targetItem.parentNode.insertBefore(draggedItem, targetItem);
        }

        handleUpdateValues();
    }

    const handleDragStart = (ev) => {
        ev.dataTransfer.setData(" text", ev.target.id);
    }

    function handleDragOver(ev) {
        ev.preventDefault();
    }

    boxes.forEach(box => {
        box.addEventListener("drop", handleDrop)
        box.addEventListener("dragover", handleDragOver)
    })

    items.forEach(item => {
        item.addEventListener("dragstart", handleDragStart)
    })
}

init();