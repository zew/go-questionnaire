
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
    const items = document.querySelectorAll("#draggable-list .item");

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

    Object.values(inputs).forEach(({ input, droppable }) => {
        const values = input.value ? input.value.split(",") : []
        values.forEach(value => {
            const item = [...items].find(el => el.dataset.value === value)
            if (item) {
                droppable.appendChild(item);
            }
        })
    })

    const handleUpdateValues = () => {
        Object.values(inputs).forEach(({ input, droppable }) => {
            const selected = droppable.querySelectorAll(".item");
            const values = [...selected].map((item) => item.dataset.value)
            input.value = values.join(",")
        })
    }

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

    const handleDragOver = (ev) => {
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
