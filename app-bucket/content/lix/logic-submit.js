function createCsvAutodownload(data) {
    try {
        const rows = [];
        const headers = [
            'Zeitstempel',
            ...CATS.map(    c => 'hk_' + c.id),
            ...CATS.flatMap(c => c.subs.map(s => 'uk_' + c.id + '_' + s.id))
        ];
        rows.push(headers.join(';'));
        const vals = [
            new Date().toLocaleString('de-DE'),
            ...data.main,
            ...data.subs.flat()
        ];

        // escaping quotes to prevent CSV injection or formatting breaks
        rows.push(
            vals.map(v => '"' + String(v).replace(/"/g, '""') + '"').join(';')
        );

        // console.log(`submitting results`)
        // console.log(rows)

        // prepending BOM (\uFEFF) ensuring Excel reads UTF-8 correctly
        const blob = new Blob(['\uFEFF' + rows.join('\n')], { type: 'text/csv;charset=utf-8;' });
        const url = URL.createObjectURL(blob);
        const a = document.createElement('a');
        a.href = url;
        a.download = 'standortfaktoren_' + Date.now() + '.csv';
        a.click();

        // releasing memory allocated for object URL
        URL.revokeObjectURL(url);
    } catch (exc) {
        handleExc(exc, 'submitResults()');
    }
}



function showResults() {

    saveToHiddenInputs();



    document.querySelectorAll('.step').forEach(s => s.classList.remove('active'));
    const el = document.getElementById('step-results');
    if (el) { el.classList.add('active'); currentStep = TOTAL_STEPS - 1; }
    updateProgress();

    if (false) {
        createCsvAutodownload({ main: mainVals, subs: subVals });
    }

    let html = '<div class="results-summary">';
    html += '<div class="res-group-title">Hauptkategorien</div>';
    html += '<div id="res-main-pie" style="width:100%;height:320px;margin-bottom:1.5rem"></div>';
    html += '<div class="res-group-title">Unterkategorien</div>';
    html += '<div class="res-pie-grid">';
    CATS.forEach((cat, i) => {
        html += `
        <div class="res-pie-card">
            <div class="res-pie-card-title">
            <span style="width:10px;height:10px;border-radius:2px;background:${cat.color};display:inline-block;flex-shrink:0"></span>${cat.label}</div>
            <div class="res-pie-container" id="res-sub-pie-${i}"></div>
        </div>`;
    });
    html += '</div></div>';
    document.getElementById('results-content').innerHTML = html;

    // nesting requestAnimationFrame twice ensures the browser has fully completed
    // the layout and paint cycle, guaranteeing container dimensions are available for ECharts
    requestAnimationFrame(() => {
        requestAnimationFrame(() => {
            const mainChart = echarts.init(document.getElementById('res-main-pie'));
            mainChart.setOption({
                tooltip: {
                    trigger: 'item',
                    formatter: p => `${p.name}<br/><b>${p.value} Punkte</b>`,
                },
                legend:  {
                    bottom:    0,
                    left:      'center',
                    textStyle: { fontSize: 12, fontFamily: 'DM Sans' },
                    itemWidth:  12,
                    itemHeight: 12,
                },
                series: [{
                    type: 'pie',
                    // radius: ['32%', '56%'],
                    radius: ['55%', '88%'],
                    center: ['50%', '45%'],
                    avoidLabelOverlap: true,
                    itemStyle: { borderRadius: 6, borderColor: '#FDFCFA', borderWidth: 3 },
                    label:     {
                        show: true, formatter: p => p.value > 0 ? p.value + ' Pkt.' : '',
                        fontSize: 12,
                        fontFamily: 'DM Sans',
                        color: '#2C2A26',
                    },
                    labelLine: {
                        show:      true,
                        showAbove: false,
                        length:  8,
                        length2: 6,
                     },
                    data: CATS.map((cat, i) => ({
                        value: mainVals[i], name: cat.label, itemStyle: { color: cat.color },
                        label: { show: mainVals[i] > 0 }, labelLine: { show: mainVals[i] > 0 }
                    }))
                }]
            });
            CATS.forEach((cat, i) => {
                const subColors = generateSubColors(cat.color, cat.subs.length);
                const el = document.getElementById('res-sub-pie-' + i);
                if (!el) return;
                const chart = echarts.init(el);
                chart.setOption({
                    tooltip: {
                        trigger: 'item',
                        formatter: p => `${p.name}<br/><b>${p.value} Punkte</b>`
                    },
                    series: [{
                        type: 'pie',
                        // radius: ['35%', '58%'],
                        radius: ['45%', '72%'],
                        center: ['50%', '51%'],
                        avoidLabelOverlap: false,
                        itemStyle: {
                            borderRadius: 6,
                            borderColor: '#FDFCFA',
                            borderWidth: 2,
                        },
                        label: {
                            show: true,
                            formatter: p => p.value > 0 ? p.value + ' Pkt.' : '',
                            fontSize: 11,
                            fontFamily: 'DM Sans',
                            color: '#2C2A26',
                        },
                        labelLine: {
                            show: true,
                            length:  8,
                            length2: 6,
                        },
                        data: cat.subs.map((sub, j) => ({
                            value: subVals[i][j], name: sub.label,
                            itemStyle: {
                                color: subColors[j],
                            },
                            label: {
                                show: subVals[i][j] > 0,
                            },
                            labelLine: {
                                show: subVals[i][j] > 0,
                            }
                        }))
                    }]
                });
            });
        });
    });

    window.scrollTo({ top: 0, behavior: 'smooth' });
}

