let chIDs = ['distanceChart', 'forecastChart', 'consensusChart'];

// === Functionality from share treatment v2 ===
document.addEventListener('DOMContentLoaded', () => {

    const userShareInput      = document.getElementById('userShareInput');
    const userShareSlider     = document.getElementById('userShareSlider');
    const shareComparisonText = document.getElementById('shareComparisonText');


    const distanceCh = echarts.init(document.getElementById('distanceChart'));
    const forecastCh = echarts.init(document.getElementById('forecastChart'));
    const consenssCh = echarts.init(document.getElementById('consensusChart'));

    let parsedData = [];

    function getSelectedQuarter() {
        return document.querySelector('input[name="quarter"]:checked').value;
    }

    function updateCharts() {
        function formatDE(value, digits = 2) {
            if (isNaN(value)) return '';
            return value.toLocaleString("de-DE", {
                minimumFractionDigits: digits,
                maximumFractionDigits: digits
            });
        }
        const respondentIndex = parseInt(10, 10);
        const selectedQuarter = getSelectedQuarter();

        if (isNaN(respondentIndex) || !parsedData[respondentIndex]) {
            distanceCh.clear();
            forecastCh.clear();
            consenssCh.clear();
            if (shareComparisonText) shareComparisonText.textContent = '';
            return;
        }

        const respondentData = parsedData[respondentIndex];
        const userShare = parseFloat(userShareInput.value) || 0;
        const actualShareRaw = respondentData[`grshare${selectedQuarter}`];
        const actualShare = parseFloat(actualShareRaw) * 100;
        const forecast = parseFloat(respondentData[`growth${selectedQuarter}`]);
        const consensus = parseFloat(respondentData[`consensus${selectedQuarter}`]);

        if (!isNaN(actualShare)) {
            shareComparisonText.innerHTML = `Sie haben <strong>${formatDE(userShare)}%</strong> angegeben.<br> Tatsächlich lag der <b>Anteil unter allen Befragten</b>, die im August 2025 ein <u>niedrigeres</u> Wachstum als Sie angegeben haben, bei <strong><span style="color:#EE6666">${formatDE(actualShare)}%</span></strong>.`;
        } else {
            shareComparisonText.innerHTML = `Ihr Anteil: <strong>${formatDE(userShare)}%</strong> | Tatsächlicher Anteil: <strong>N/A</strong>`;
        }

        // Distance Chart
        const distance = Math.abs(userShare - actualShare);
        const distanceChartOption = {
            grid: { top: 20, right: 40, bottom: 20, left: 40 },
            xAxis: { type: 'value', min: 0, max: 100, axisLabel: { formatter: '{value}%' }, splitLine: { show: false } },
            yAxis: { type: 'category', data: [''], show: false },
            series: [{
                type: 'scatter',
                symbolSize: 26,
                data: [
                    {
                        name: 'Ihr Anteil',
                        value: [userShare, 'Ihr Anteil'],
                        itemStyle: { color: '#546a7b' },
                        label: {
                            show: true, position: 'top', distance: 8, fontWeight: 'bold', color: '#000000',
                            formatter: (params) => formatDE(params.value[0]) + '%'
                        }
                    },
                    {
                        name: 'Tatsächlicher Anteil',
                        value: [actualShare, ''],
                        itemStyle: { color: '#EE6666' },
                        label: {
                            show: true, position: 'bottom', distance: 8, fontWeight: 'bold', color: '#EE6666',
                            formatter: (params) => formatDE(params.value[0]) + '%'
                        }
                    }
                ],
                markLine: !isNaN(actualShare) ? {
                    symbol: ['none', 'none'],
                    label: { show: true, formatter: `Abstand: ${formatDE(distance)}%`, position: 'middle', fontWeight: 'bold', color: '#111' },
                    lineStyle: { type: 'solid', width: 3, color: '#6b7280' },
                    data: [[{ coord: [userShare, ''] }, { coord: [actualShare, ''] }]]
                } : {}
            }]
        };
        distanceCh.setOption(distanceChartOption);

        // Forecast Chart
        const forecastOption = {
            grid: { top: 10, right: 60, bottom: 30, left: 100 },
            xAxis: { type: 'value', min: -3, max: 3, interval: 0.5, axisLabel: { formatter: '{value}' } },
            yAxis: { type: 'category', data: [selectedQuarter] },
            series: [{
                type: 'bar',
                data: [forecast],
                itemStyle: { color: '#62929e' },
                label: { show: true, position: 'insideRight', color: '#000', backgroundColor: '#fff', borderRadius: 4, padding: [2, 6], formatter: (val) => (isNaN(val.value) ? '' : formatDE(val.value)) }
            }]
        };
        forecastCh.setOption(forecastOption);

        // Consensus Chart
        const consensusOption = {
            grid: { top: 10, right: 60, bottom: 30, left: 100 },
            xAxis: { type: 'value', min: -3, max: 3, interval: 0.5, axisLabel: { formatter: '{value}' } },
            yAxis: { type: 'category', data: [selectedQuarter] },
            series: [{
                type: 'bar',
                data: [consensus],
                itemStyle: { color: '#546a7b' },
                label: { show: true, position: 'insideRight', color: '#000', backgroundColor: '#fff', borderRadius: 4, padding: [2, 6], formatter: (val) => (isNaN(val.value) ? '' : formatDE(val.value)) }
            }]
        };
        consenssCh.setOption(consensusOption);
    }
    window.updateCharts = updateCharts;



    // Sync slider and input
    userShareInput.addEventListener('input', () => {
        userShareSlider.value = userShareInput.value;
        updateCharts();
    });
    userShareSlider.addEventListener('input', () => {
        userShareInput.value = userShareSlider.value;
        updateCharts();
    });


    chIDs.forEach(function (id) {
        let ch = document.getElementById(id);
        ch.resize();
    });



});
