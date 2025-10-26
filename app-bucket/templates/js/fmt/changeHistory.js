const cutBack = 8;

// const  historyStackInp  = document.getElementById("history_stack_pg2");
let historyStackInp = null;
for (let i = 0; i<20 ; i++) {
    const candidate = document.getElementById(`history_stack_pg${i}`);
    if (candidate) {
        historyStackInp = candidate;
        break;
    }
}


const histJson   =  historyStackInp.value
let historyStack = {}
if (histJson.trim() !== "") {
    historyStack = JSON.parse(histJson);
}



// epoch seconds as the key
function historyKey() {
    const nowMillis = Date.now();
    const nowSeconds = Math.floor(nowMillis / 1000);
    return nowSeconds;
}

/**
 * Record a change into historyStack.
 * - historyStack: plain object mapping timestamp (seconds) -> { src, newValue }
 * Behavior:
 *   If the most recent existing change is < cutBack secs old, delete ALL entries whose
 *   timestamps are within (now - cutBack secs, now], then insert the new change at now.
 *   Otherwise, just insert the new change at now.
 */
function recordChange(src, newValue) {
    const nowSec = historyKey();

    // Find the most recent timestamp already in the stack (if any) — BUT ONLY for this src
    let lastTs = null;
    const keys = Object.keys(historyStack);

    for (let i = 0; i < keys.length; i++) {
        const ts = Number(keys[i]);
        if (!Number.isNaN(ts)) {
            const entry = historyStack[keys[i]];
            if (entry && entry.src === src) {
                if (lastTs === null || ts > lastTs) {
                    lastTs = ts;
                }
            }
        }
    }

    // If the most recent change FOR THIS src is within the last cutBack secs,
    // remove entries in the last-cutBack secs window FOR THIS src before appending this one.
    if (lastTs !== null) {
        const age = nowSec - lastTs;
        if (age < cutBack) {
            const cutoff = nowSec - cutBack;
            const keys2 = Object.keys(historyStack);
            for (let i = 0; i < keys2.length; i++) {
                const ts = Number(keys2[i]);
                if (!Number.isNaN(ts)) {
                    const entry2 = historyStack[keys2[i]];
                    if (entry2 && entry2.src === src) {
                        if (ts > cutoff && ts <= nowSec) {
                            delete historyStack[keys2[i]];
                        }
                    }
                }
            }
        }
    }

    // Append the most recent event at current second
    historyStack[nowSec] = { src: src, value: newValue };

    historyStackInp.value = JSON.stringify(historyStack);
    console.log(`simHistInp.value ${historyStackInp.value}`)

}



let onChangeHandlerRecordHistory = (evt) => {
    let src = evt.srcElement;
    const chVal = src.value;

    // console.log(`   ${evt.srcElement.name} - new val  ${chVal}`)
    const nm = src.name.trim();
    recordChange(nm, chVal);
}


// userShareInp.onchange = paramChange


let elementIds = [
    "pprwbipq1",
    "pprwbipq2",
    "pprwbipq3",
    "pprwbipq4",

    "ssq3a_1",
    "ssq3a_2",
    "ssq3a_3",
    "ssq3a_4",

    // ssq5a_1
    "ssq5a_1",
    "ssq5a_2",
    "ssq5a_3",
    "ssq5a_4",

    "ssq5b_1",
    "ssq5b_2",
    "ssq5b_3",
    "ssq5b_4",

];

for (let i = 0; i < elementIds.length; i++) {
    let el = document.getElementById(elementIds[i]);
    if (!el) {
        // console.log(`Element with ID '${elementIds[i]}' not found.`);
        continue;
    }
    el.addEventListener("change", onChangeHandlerRecordHistory, false);
    console.log(`change history handler attached to  '${elementIds[i]}'.`);
}

// console.log(`change history handlers attached`);
