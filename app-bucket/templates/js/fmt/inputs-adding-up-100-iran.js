// function block isolates multiple instances
(function () {

    // must be var
    var inp1Id = "{{.inp1}}";
    var inp2Id = "{{.inp2}}";
    var inp3Id = "{{.inp3}}";
    var inp4Id = "{{.inp4}}";

    console.log( `names for inp1-3 -${inp1Id}- -${inp2Id}- -${inp3Id}- `);

    var inp1 = document.getElementById(inp1Id);
    var inp2 = document.getElementById(inp2Id);
    var inp3 = document.getElementById(inp3Id);
    var inp4 = document.getElementById(inp4Id);


    // function funcOnChange(evt) {
    //     console.log("selected:", evt.target.value);
    //     if (evt.target.value){
    //         inp2.value = 100 - parseInt(evt.target.value) ;
    //     } else {
    //         inp2.value = "";
    //     }
    // }


    // function funcOnChecked(evt) {
    //     console.log("selected:", evt.target.value);
    //     if (evt.target.checked){
    //         inp1.value = "";
    //         inp2.value = "";
    //         inp1.placeholder = "";
    //         inp2.placeholder = "";
    //         inp1.disabled = true;
    //     } else {
    //         inp1.disabled = false;
    //         inp1.placeholder = "0";
    //         // inp2.placeholder = "0";
    //     }
    // }







    function demo(evt) {
        if (confirm("Press a button!")) {
            txt = "You pressed OK!";
            console.log(txt);
            return true;
        } else {
            txt = "You pressed Cancel!";
            console.log(txt);
            evt.preventDefault(); // not only return false - but also preventDefault()
            return false;
        }
    }

    function funcOnChange(evt) {

        // console.log("inp1-3: ", inp1, inp2, inp3);
        let vl1 = inp1.value
        let vl2 = inp2.value
        let vl3 = inp3.value


        let i1 = 0
        if (vl1 != "") {
            i1 = parseInt(vl1, 10);
        }
        let i2 = 0
        if (vl2 != "") {
            i2 = parseInt(vl2, 10);
        }
        let  i3 = 0
        if (vl3 != "") {
            i3 = parseInt(vl3, 10);
        }
        // console.log("vl1-3 integer: ", i1, i2, i3);

        var sum = i1 + i2 + i3;

        console.log(`  changed: ${evt.target.name} to val ${evt.target.value} --  sum ${sum}` );


        if (sum > 0) {        

            inp4.value = sum

            if (sum != 100 ) {

                // console.log("event.type", event.type);
                var doAsk = false;
                if (evt.type=="input") {
                    if (vl1 != "" && vl2 != "" && vl3 != "") {
                        // doAsk = true;
                        // console.log("show error msg");
                    }
                } else {
                    // submit
                    doAsk = true;
                }

                evt.source

                if (doAsk) {
                    // alert("{{.msg}}");
                    var doContinue = window.confirm("{{.msg}}");
                    if (doContinue) {
                        return true;
                    }
                    evt.preventDefault(); // not only return false - but also preventDefault()
                    return false;
                }
            }
        }

        return true;

    }






    // addEventListener is cumulative
    window.addEventListener("load", function (evt) {

        inp1.addEventListener('input',  funcOnChange);
        inp2.addEventListener('input',  funcOnChange);
        inp3.addEventListener('input',  funcOnChange);
        // inp1.addEventListener('blur',   funcOnChange);


        inp4.placeholder = "";


        const evtInit = new Event("input");
        inp1.dispatchEvent(evtInit);         

        var frm = document.forms.frmMain;
        if (frm) {
            frm.addEventListener('submit', funcOnChange);
        }


    });


})();
