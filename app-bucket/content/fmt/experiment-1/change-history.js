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
     *   If the most recent existing change is < 20s old, delete ALL entries whose
     *   timestamps are within (now - 20s, now], then insert the new change at now.
     *   Otherwise, just insert the new change at now.
     */
    function recordChange(src, newValue) {
        const nowSec = historyKey();

        // Find the most recent timestamp already in the stack (if any)
        let lastTs = null;
        const keys = Object.keys(historyStack);

        for (let i = 0; i < keys.length; i++) {
            const ts = Number(keys[i]);
            if (!Number.isNaN(ts)) {
                if (lastTs === null || ts > lastTs) {
                    lastTs = ts;
                }
            }
        }

        // If the most recent change is within the last 20s,
        // remove all entries in the last-20s window before appending this one.
        if (lastTs !== null) {
            const age = nowSec - lastTs;
            if (age < 20) {
                const cutoff = nowSec - 20;
                const keys2 = Object.keys(historyStack);
                for (let i = 0; i < keys2.length; i++) {
                    const ts = Number(keys2[i]);
                    if (!Number.isNaN(ts)) {
                        if (ts > cutoff && ts <= nowSec) {
                            delete historyStack[keys2[i]];
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



