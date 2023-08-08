
{


  const ROOT_PATH_1 = "https://echarts.apache.org/examples"
  const weatherIcons = {
    Sunny:   ROOT_PATH_1 + '/data/asset/img/weather/sunny_128.png',
    Cloudy:  ROOT_PATH_1 + '/data/asset/img/weather/cloudy_128.png',
    Showers: ROOT_PATH_1 + '/data/asset/img/weather/showers_128.png'
  };


  // const ROOT_PATH_2 = "http://localhost/ecb-watch/img"
  const ROOT_PATH_2 = "http://localhost/echart-examples/img"
  const directors = {
    Schnabel:   ROOT_PATH_2 + '/out-director-schnabel-16-grey-d2.png',
    Heinemann:  ROOT_PATH_2 + '/out-heinemann-16-grey-d2.png',
  };

  // console.log(directors.Schnabel)
  // console.log(directors.Heinemann)


  const showSeriesLabel = {
    show: true
  };

  var cntrSymbol = 0;

  // https://echarts.apache.org/handbook/en/how-to/chart-types/bar/bar-race

  let yr = new Date().getFullYear();

  if (yr / 2 != 0.0) {
    yr++;
  }

  let visualMap =  {
    orient: 'horizontal',
    left:   'center',
    min: 10,
    max: 100,
    text: ['High Score', 'Low Score'],
    // map the score column to color
    dimension: 0,
    inRange: {
      color: ['#65B581', '#FFCE34', '#FD665F']
    }
  };

  // the visual map makes the first stack visible
  visualMap = null;



  let option = {
    visualMap: visualMap,
    dataset: {
      // https://echarts.apache.org/handbook/en/concepts/dataset/
      // 
      // how to cross-reference...
      // https://echarts.apache.org/en/option.html#xAxis.axisPointer.label
      // dimensions: ['product', '2015', '2016', '2017'],
      dimensions: [  'personName',  'startY',    'stopY1',     'stopY2',   'xxxxx'       ],
      source: [
        [         'Duisenberg',      1998,         5,           '-',        'Latte'         ],
        [         'Noyer',           1998,         4,           '-',        'Matcha'        ],
        [         'Papademos',       2002,         8,           '-',        'Pamda'         ],
        [         'Trichet',         2003,         8,           '-',        'Milk'          ],
        [         'Constâncio',      2010,         8,           '-',        'Taccu'         ],
        [         'Draghi',          2011,         8,           '-',        'Cheese Cocoa'  ],
        [         'Lagarde',         2019,         '-',         8,          'Brownie'       ],
        [         'Guindos',         2018,         '-',         8,          'Togo Cocoa'    ],
      ]
    },
    grid: { containLabel: true },
    xAxis: {
      type: 'category',
      type: 'value',

      axisLabel: {
        // compare  axisLabel.formatter
        formatter: function (vl, index) {
            let vl1 = vl.toFixed(0)
            return vl1;
        },
        margin: 5,
      },

      // only in numerical axis, i.e., type: 'value'.
      //    show zero position only, if justified by data
      //    if min and max are set, this setting has no meaning
      scale: true,

      // supercedes everything
      min: 1998 - 1,
      max: yr   + 1,


      // similar to scale - but less buffer at the start and end
      // min: 'dataMin',
      // max: 'dataMax',

      interval: 2,


    },
    yAxis: {
        type: 'category',
        inverse: true, // instead of option.dataset.source.reverse()
        // data: ['Duisenberg', 'Noyer', 'Papademos'],
        // data: ['Duisenberg', 'Noyer', 'Papademos', 'Trichet', 'Constâncio', 'Draghi', 'Lagarde', 'Guindos'],

        axisLabel: {
          // "value" is the placeholder for what
          //        formatter: '{a|{a}\n}{b|{b} }{c|{c}}',

          formatter: function (val) {
            // return '{' + val + '| }\n{value|' + val + '}';
            return `{${val}|}\n{value|${val}}`;
          },

          margin: 25,
          rich: {
            value: {
              lineHeight: 10,
              align: 'center',
              align: 'right',
            },
            Duisenberg: {
              height: 20,
              align: 'right',
              backgroundColor: {
                image: weatherIcons.Sunny,
                image: directors.Schnabel,
              }
            },
            Noyer: {
              height: 20,
              align: 'right',
              backgroundColor: {
                image: weatherIcons.Cloudy,
                image: directors.Heinemann,
              }
            },
            'Papademos': {
              height: 20,
              align: 'right',
              backgroundColor: {
                image: weatherIcons.Showers
              }
            },
            'Trichet': {
              height: 20,
              // widht:  40,
              // align: 'right',
              backgroundColor: {
                // image: vectorImgs.Plane,
                symbol: vectorImgs.Plane,
                symbol: vectorImgs.Reindeer,
              }
            },
          },




        },




    },

    series: [
      {
        // SILENT first stack
        name: 'startY',

        type: 'bar',
        stack: 'stackX',
        silent: true,

        barWidth:       '20%',
        barGap:         '20%',
        barCategoryGap: '40%',


        itemStyle: {
          color:       'yellow',
          color:       'transparent',
          borderColor: 'transparent',
        },

        // light gray underground rail
        showBackground: true,
        backgroundStyle: {
          color: 'rgba(220, 220, 220, 0.05)'
        },

      },

      {
        name: 'stopY1',

        type: 'bar',
        stack: 'stackX',
        barWidth:       '20%',
        barGap:         '20%',
        barCategoryGap: '40%',

        label: showSeriesLabel,

        // itemStyle: {
        //   color:       'rgb(22,22,222)',
        //   // https://echarts.apache.org/en/option.html#series-bar.itemStyle.decal
        //   // does not work
        //   decal: {
        //     symbol: 'image://data:image/gif;base64,R0lGODlhEAAQAMQAAORHHOVSKudfOulrSOp3WOyDZu6QdvCchPGolfO0o/XBs/fNwfjZ0frl3/zy7////wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACH5BAkAABAALAAAAAAQABAAAAVVICSOZGlCQAosJ6mu7fiyZeKqNKToQGDsM8hBADgUXoGAiqhSvp5QAnQKGIgUhwFUYLCVDFCrKUE1lBavAViFIDlTImbKC5Gm2hB0SlBCBMQiB0UjIQA7',
        //     // symbol: pathSymbols.reindeer,
        //     symbolSize: 0.8,
        //   },
        //   borderColor: 'red',
        // },

        markPoint: {

          // data is required for the symbol
          // data is required for the label formatter
          data: [
            { type: 'min',    name: 'min days' },
            { type: 'max',    name: 'max days' },
            { type: 'median', name: 'median'   },
            { type: 'value',  name: 'dta_value' },
          ],

          symbolSize: 25,
          symbolOffset: ['25%', '50%'],

          symbol: 'pin',
          // an asterisk
          symbol: 'image://data:image/gif;base64,R0lGODlhEAAQAMQAAORHHOVSKudfOulrSOp3WOyDZu6QdvCchPGolfO0o/XBs/fNwfjZ0frl3/zy7////wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACH5BAkAABAALAAAAAAQABAAAAVVICSOZGlCQAosJ6mu7fiyZeKqNKToQGDsM8hBADgUXoGAiqhSvp5QAnQKGIgUhwFUYLCVDFCrKUE1lBavAViFIDlTImbKC5Gm2hB0SlBCBMQiB0UjIQA7',
          symbol: 'roundRect',
          symbol: 'image://' + weatherIcons.Showers,
          symbol: vectorImgs.Reindeer,

          // if symbols need to be different => set with callback function in the following format:
          // (value: Array|number, params: Object) => string
          // first  parameter value  is the value in data, 
          // second parameter params is the rest parameters of data item.
          symbol: function (val,param) {
            // return '{' + val + '| }\n{value|' + val + '}';
            // return `{${val}|}\n{value|${val}}`;
            console.log(val)
            console.log(param)
            cntrSymbol++;
            if (cntrSymbol%3===0) {
              return vectorImgs.Reindeer;
            }
            if (cntrSymbol%3===1) {
              return 'roundRect';
            }
            return 'image://' + weatherIcons.Showers;
          },



          label: {
            show: true,
            backgroundColor: 'rgb(188,188,188)',
            distance: 20,
            lineHeight: 26,
            position: 'right',
            formatter: function (val) {
              // return '{' + val + '| }\n{value|' + val + '}';
              return `{${val}|}\n{value|${val}}`;
            },
            formatter: '--{style1|{value}}--',
            formatter: '{style1|{a}}--{style2|{b}}--{style3|{c}}--{style4|{d}}',

            rich: {
              style1: {
                align: 'center',
                color: '#fff',
                fontSize: 18,
                textBorderColor: '#333',
                textBorderWidth: 2,
              },
              style2: {
                color: '#333',
              },
              style3: {
                color: '#ff8811',
                fontSize: 22
              },
              style4: {
                color: 'red',
                fontSize: 22
              },
            },
    
            

          },


        },



        // the mere existence of 'encode' changes things. But
        // encode: {
        //   x: 'stopY1',
        //   y: 'personName',
        //   // x: 'startY',
        //   // y: 'stopY1',
        // },

      },

      {
        name: 'stopY2',

        type: 'bar',
        stack: 'stackX',
        barWidth:       '20%',
        barGap:         '20%',
        barCategoryGap: '40%',


        itemStyle: {
          color:       'rgb(22,222,22)',
          // borderColor: 'red',
        },

        label: {
          distance: 20,
          position: 'right',
        }


      },
    ]
  };

  // instead: just set yAxis inverse: true,
  // option.dataset.source = option.dataset.source.reverse();


  let domContainer = document.getElementById('chart_container_1');
  // let myChart = echarts.init(domContainer);
  var myChart = echarts.init(domContainer);
  option && myChart.setOption(option);

  let optRes = myChart.getOption();
  if (false) {
    console.log(optRes.series[0]);
    console.log(optRes.series[1]);
    console.log(optRes.series[2].encoding);
    // https://stackoverflow.com/questions/66363002/how-to-add-images-or-shapes-to-a-chart
    console.log(myChart.getModel().getSeries()[0]);
  }

}


