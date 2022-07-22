// import * as echarts from 'echarts';

var chartDom = document.getElementById('chart_container');
var myChart = echarts.init(chartDom);
var option;

function run() {
    console.log("run - start");

    const countries = [
        'Finland',
        'France',
        'Germany',
        'Iceland',
        'Norway',
        'Poland',
        'Russia',
        'United Kingdom'
    ];
    
    const datasetWithFilters = [];
    const seriesList = [];
    let counter = -1

    echarts.util.each( countries, function(country){
        counter++;
        var datasetId = 'dataset_' + country;
        console.log(`counter ${counter} - country ${country}`);

        // restriction to certain countries, values greater 1950
        datasetWithFilters.push({
            id: datasetId,
            fromDatasetId: 'dataset_raw',
            transform: {
                type: 'filter',
                config: {
                    and: [
                        // { dimension: 'Year', gte: 1950 },
                        { dimension: 'Year', gte: 1980 },
                        { dimension: 'Country', '=': country }
                    ]
                }
            }
        });

        seriesList.push({
            type: 'line',
            datasetId: datasetId,
            showSymbol: false,
            name: country,
            // dynamic - last data point is labeled
            endLabel: {
                show: true,
                formatter: function (params) {
                    return params.value[3] + ': ' + params.value[0];
                }
            },
            labelLayout: {
                moveOverlap: 'shiftY'
            },
            emphasis: {
                focus: 'series'
            },
            encode: {
                x: 'Year',
                y: 'Income',
                label: ['Country', 'Income'],
                itemName: 'Year',
                tooltip: ['Income']
            }
        });
    }  // countries each - func func body
    ); // countries each - func argument


    // now the full dataset is prepared

    // console.log(datasetWithFilters);
    // console.log(seriesList);

    option = {
        // animationDuration: 10*1000,
        animationDuration: 1.2*1000,
        dataset: [
            {
                id: 'dataset_raw',
                source: lifeExp
            },
            ...datasetWithFilters
            // referencing dataset raw - with different filters
            // ids are 'dataset_' + country
        ],
        title: {
            text: 'Income since 1950'
        },
        tooltip: {
            order: 'valueDesc',
            trigger: 'axis'
        },
        xAxis: {
            type: 'category',
            nameLocation: 'middle'
        },
        yAxis: {
            name: 'Income'
        },
        grid: {
            right: 140
        },
        series: seriesList
    };
    myChart.setOption(option);

    console.log("run - end");

}

option && myChart.setOption(option);

run();


console.log("end of echart");