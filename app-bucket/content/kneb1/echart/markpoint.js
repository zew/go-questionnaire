let cntrSymbol = 0;

let markP = {

    // data is required for the symbol
    // data is required for the label formatter
    // https://echarts.apache.org/en/option.html#series-line.markPoint.data
    data: [
        // {
        //     name: 'screen coordinate - pixel',
        //     x: 300, y: 10,
        // },        
        // {
        //     name: 'coordinate - axis values',
        //     coord: [2040, 30000],
        // },
        { nameX: 'fv05' , name: 'schlechteste 5%' , coord: [2043,33465.583685182035 ] },
        { nameX: 'fv'   , name: 'Durchschnitt'    , coord: [2043,36126.58482791517 ] },
        { nameX: 'fv95' , name: 'beste 5%'        , coord: [2043,39044.23734467819 ] },
        // { type: 'median', name: 'median' },
        // { x: 70, y: 140 },
        // { type: 'min', name: 'min days' },
        // { type: 'max', name: 'max days' },
        // { type: 'value', name: 'dta_value' },
    ],

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
        // return '{' + val + '| }\n{value|' + val + '}';
        // return `{${val}|}\n{value|${val}}`;
        console.log(val)
        console.log(param)
        cntrSymbol++;
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



    label: {
        show: false,
        show: true,
        backgroundColor: 'rgba(188,188,188,0.8)',
        distance: 12,
        lineHeight: 12,
        position: 'bottom',
        position: 'right',
        position: 'left',
        formatter: function (val) {
            // return '{' + val + '| }\n{value|' + val + '}';
            return `{${val}|}\n{value|${val}}`;
        },
        formatter: '--{style1|{value}}--',
        formatter: '{style1|{a}}--{style2|{b}}--{style3|{c}}--{style4|{d}}',
        formatter: '{style2|{b}}',

        rich: {
            style2: {
                color: '#333',
            },
        },



    },


}


function addMarkPoint(chartObj) {
    chartObj.setOption(
        {
            series: [
                {
                    // data: dataObject.computeData(),
                },
                {
                    markPoint: markP,
                },
                {
                    // data: dataObject.computeData(),
                },
            ]

        }
    );
    console.log(`markpoint set ${chartObj}`, chartObj)
}


