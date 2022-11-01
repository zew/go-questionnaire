// non global block
{

    // the wide rectangle, not the `thumb` 

    let baseName = "{{.inputBaseName}}";

    let cols = 3
    let inps = [[], [], [], ];
    let inpsAll = [];

    for (let idx1 = 0; idx1 < cols; idx1++) {
        // const element = inps[idx1];
        let name = baseName + "_prio" + (idx1+1);
        // inps[idx1]  = document.getElementById(name);

        let prioX = document.getElementsByName(name);
        // console.log("prioX", prioX);

        for (let idx2 = 0; idx2 < prioX.length; idx2++) {
            // const element = array[idx2];
            inps[idx1].push(prioX[idx2] );
            inpsAll.push(prioX[idx2] );
        }

    }

    // console.log("inps", inps);

    for (let idx1 = 0; idx1 < inpsAll.length; idx1++) {

        inpsAll[idx1].onclick = function () {
            let n1 = this.name;
            let lastChar = n1.substring(n1.length -1);
            let idxSrc = parseInt(lastChar)-1;
            // console.log(n1, "has value", this.value, lastChar, idxSrc);

            let exists = false;
            for (let idx2 = 0; idx2 < inps.length; idx2++) {
                if (idx2 == idxSrc) {
                    continue;
                }
                const sa = inps[idx2];
                for (let idx3 = 0; idx3 < sa.length; idx3++) {
                    // console.log("   comparing to ", sa[idx3].name);
                    if (this.value == sa[idx3].value && sa[idx3].checked) {
                        exists = true;
                        // console.log("     idx2 - idx3 - equal", idx2, idx3, sa[idx3].value );
                        break;
                    }
                }
            }

            if (exists) {
                this.checked = false;
                alert("{{.msg}}");

            }
        }


    }


}