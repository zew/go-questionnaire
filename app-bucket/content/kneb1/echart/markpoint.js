let cntrSymbol = 0;
let cntrLbl    = 0;


function getMarkpointConfig( keyValsCoords ) {

    if (!keyValsCoords.length ||  keyValsCoords.length !== 3) {
        console.error(`getMarkpointConfig needs array of three numbers`,  typeof keyValsCoords, keyValsCoords)
        return {}
    }

    // data is required for the symbol
    // data is required for the label formatter
    // https://echarts.apache.org/en/option.html#series-line.markPoint.data
    let data1 = [
        // {
        //     name: 'screen coordinate - pixel',
        //     x: 300, y: 10,
        // },
        // {
        //     name: 'coordinate - axis values',
        //     coord: [2040, 30000],
        // },

        // example
        
        { name: 'schlechteste 5%', coord: [yr+azV, 33465.583685182035], nameX: 'fv05',  },
        { name: 'Durchschnitt',    coord: [yr+azV, 36126.58482791517] , nameX: 'fv',    },
        { name: 'beste 5%',        coord: [yr+azV, 39044.23734467819] , nameX: 'fv95',  },
        // { type: 'median', name: 'median' },
        // { x: 70, y: 140 },
        // { type: 'min', name: 'min days' },
        // { type: 'max', name: 'max days' },
        // { type: 'value', name: 'dta_value' },
    ]

    data1 = [
        { name: 'schlechteste 5%', coord: [yr+azV, keyValsCoords[0]], nameX: 'fv05',  'itemStyle': {  color: '#a00'   },  },
        { name: 'Durchschnitt',    coord: [yr+azV, keyValsCoords[1]], nameX: 'fv',    'itemStyle': {   },  },
        { name: 'beste 5%',        coord: [yr+azV, keyValsCoords[2]], nameX: 'fv95',  'itemStyle': {  color: '#0a0'  },  },
    ]

    let config = {

        data: data1,

        animation: true,
        animationDelay: 2000,

        symbolSize: 25,
        symbolOffset: ['25%', '50%'],

        symbolSize: 9,
        symbolOffset: ['0%', '0%'],

        symbol: 'pin',
        symbol: 'roundRect',
        // symbol: 'image://' + weatherIcons.Showers,
        // symbol: vectorImgs.Reindeer,
        // an asterisk
        symbol: 'image://data:image/gif;base64,R0lGODlhEAAQAMQAAORHHOVSKudfOulrSOp3WOyDZu6QdvCchPGolfO0o/XBs/fNwfjZ0frl3/zy7////wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACH5BAkAABAALAAAAAAQABAAAAVVICSOZGlCQAosJ6mu7fiyZeKqNKToQGDsM8hBADgUXoGAiqhSvp5QAnQKGIgUhwFUYLCVDFCrKUE1lBavAViFIDlTImbKC5Gm2hB0SlBCBMQiB0UjIQA7',

        // if symbols need to be different => set with callback function in the following format:
        // (value: Array|number, params: Object) => string
        // first  parameter value  is the value in data,
        // second parameter params is the rest parameters of data item.
        symbol: function (val, param) {
            // val is undefined
            // return '{' + val + '| }\n{value|' + val + '}';
            // return `{${val}|}\n{value|${val}}`;
            cntrSymbol++;
            console.log(`sym ${cntrSymbol}`, param );
            if (cntrSymbol % 3 === 0) {
                // return vectorImgs.Reindeer;
                return 'roundRect';
            }
            if (cntrSymbol % 3 === 1) {
                return 'roundRect';
            }
            return 'roundRect';
            //   return 'image://' + weatherIcons.Showers;
        },



        // https://echarts.apache.org/en/option.html#series-line.markPoint.label
        label: {
            show: false,
            show: true,

            // animation does not help
            animation: true,
            animationDelay: 2000,

            backgroundColor: 'rgba(188,188,188,0.8)',
            lineHeight: 12,
            position: 'inside',
            position: 'bottom',
            position: 'left',
            position: 'top',
            position: 'insideBottomRight',
            position: 'insideBottomLeft',
            position: 'insideLeft',
            position: 'right',

            distance: 14,

            // cntrLbl is not incremented
            // offset: [-10, 30*cntrLbl],
            offset: [-10, 0],


            // formatter *attribute* consists of {style|{key}}
            //    where style must be defined in rich {}
            //    and key is {a|b|c}
            //      https://echarts.apache.org/en/option.html#series-line.markPoint.label.formatter
            //          exlains {a}, {b}...
            //       but only {b} works for markPoint.data
            formatter: '--{style1|{value}}--',
            formatter: '{style1|{a}}--{style2|{b}}--{style3|{c}}--{style4|{d}}',
            formatter: '{style1|{@[2]}}',
            formatter: '{style1|{@[dimNamen]}}',
            formatter: '{style2|{b}}',

            // formatter *func* for markpoins
            // different than for data points
            //  {a}, {b}, etc. dont work, but we get the markpoint as an object as argumet
            formatter: function (markPoint) {
                cntrLbl++;
                // console.log(`lbl ${cntrLbl}`, markPoint.data.coord[0], markPoint.data.coord[1] );
                // // console.log(`markPoint`, markPoint)
                // console.log(`markPoint.name`, markPoint.name)
                // // console.log(`markPoint.data`, markPoint.data)
                // console.log(`markPoint.data.coord[0]`, markPoint.data.coord[0])
                // console.log(`markPoint.data.coord[1]`, markPoint.data.coord[1])
                // let rnd = Math.round(markPoint.data.coord[1])
                let rnd = knebelFormat(markPoint.data.coord[1],false)
                
                // rnd = `${rnd}`  // to string
                // rnd = `${rnd.substring(0,rnd.length-3)}.${rnd.substring(rnd.length-3)}`
                // rnd = `${rnd} €`


                return `{style2|${rnd}}`;
                return `{style2|${rnd}}\n{style1|${markPoint.name}}`;
            },

            align: 'right',
            align: 'center',
            align: 'left',
            // fontWeight
            // fontFamily
            verticalAlign: 'bottom',
            verticalAlign: 'middle',
            lineHeight: 13,

            // width: 120,
            // height: 120,
            overflow: 'none', // truncate, break, breakAll
            // ellipsis

            // border properties possible
            // shadow properties possible
            // text border properties possible

            padding: [0, 4, 0, 4], // represents padding of [top, right, bottom, left].

            rich: {
                style1: {
                    color: '#133',
                    fontSize: 10,
                },
                style2: {
                    color: '#333',
                    fontSize: 12,
                    // override
                    // fontWeight
                    // fontFamily
                },
            },

        },
    }

    return config

}

// unused
// function addMarkPoint(chartObj) {
//     chartObj.setOption(
//         {
//             series: [
//                 {
//                     // data: dataObject.computeData(),
//                 },
//                 {
//                     markPoint: markP,
//                 },
//                 {
//                     // data: dataObject.computeData(),
//                 },
//             ]

//         }
//     );
//     console.log(`markpoint set ${chartObj}`, chartObj)
// }


