
let currentStep = 0;
// six sub categories
// const TOTAL_STEPS = 8;
const TOTAL_STEPS = 4;
let companyData = {};
let mainVals = CATS.map(() => 0);
let subVals  = CATS.map( c => c.subs.map(() => 0));
let charts   = {};

function budget(arr) { return arr.length * 10; }
function remaining(arr) { return budget(arr) - arr.reduce((a, b) => a + b, 0); }


function clampSet(arr, idx, val) {
    const max = budget(arr);
    val = Math.max(0, Math.min(max, val));
    const others = arr.reduce((a, b, i) => i === idx ? a : a + b, 0);
    if (val + others > max) val = max - others;
    arr[idx] = val;
}


function submitResults(data) {
    const rows = [];
    const headers = [
        'Zeitstempel',
        ...CATS.map(    c => 'HK_' + c.id),
        ...CATS.flatMap(c => c.subs.map(s => 'UK_' + c.id + '_' + s.id))
    ];
    rows.push(headers.join(';'));
    const vals = [
        new Date().toLocaleString('de-DE'),
        ...data.main,
        ...data.subs.flat()
    ];
    rows.push(
        vals.map(v => '"' + String(v).replace(/"/g, '""') + '"').join(';')
    );

    // console.log(`submitting results`)
    // console.log(rows)

    const blob = new Blob(['\uFEFF' + rows.join('\n')], { type: 'text/csv;charset=utf-8;' });
    const url = URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = 'standortfaktoren_' + Date.now() + '.csv';
    a.click();

    URL.revokeObjectURL(url);
}


function updateProgress() {
    const fill  = document.getElementById('progress-fill');
    const label = document.getElementById('step-label');
    const count = document.getElementById('step-count');
    const pct = Math.round((currentStep / (TOTAL_STEPS - 1)) * 100);
    fill.style.width = pct + '%';
    count.textContent = currentStep + ' / ' + (TOTAL_STEPS - 1);
    const labels = ['Einleitung', 'Hauptkategorien',
        ...CATS.map(c => c.label)];
    label.textContent = labels[currentStep] || '';
}

function buildBudgetBadge(arr) {
    const rem = remaining(arr);
    const tot = budget(arr);
    let cls = rem === 0 ? 'ok' : rem > 0 ? 'warn' : 'over';
    let msg = rem === 0 ? 'Alle ' + tot + ' Punkte vergeben' : rem > 0 ? rem + ' von ' + tot + ' Punkten noch verfügbar' : Math.abs(rem) + ' Punkte zu viel';
    return `<div class="budget-badge ${cls}"><span class="dot"></span>${msg}</div>`;
}

function buildSliders(vals, catIdx, isSub) {
    const items = isSub !== undefined ? CATS[catIdx].subs : CATS;
    const snapPt = 10;
    const dtype = isSub !== undefined ? 'sub' : 'main';
    let html = '<div class="sliders">';
    items.forEach((item, i) => {
        const color = isSub !== undefined ? CATS[catIdx].color : CATS[i].color;
        const val = vals[i];
        const pct = budget(vals) > 0 ? Math.round(val / budget(vals) * 100) : 0;
        const thumbOffset = budget(vals) > 0 ? (val / budget(vals) * 16 - 8) : -8;
        const trackBg = `linear-gradient(to right, ${color} calc(${pct}% - ${thumbOffset.toFixed(2)}px), #E2DED6 calc(${pct}% - ${thumbOffset.toFixed(2)}px))`;
        const snapFrac = 10 / (vals.length * 10);
        const snapPct = snapFrac * 100;
        const snapPos = `calc(${snapPct.toFixed(4)}% - ${(snapFrac * 16).toFixed(4)}px + 8px)`;
        html += `<div class="slider-item">
            <div class="slider-header">
            <div>
                <div class="slider-name">${item.label}${item.tooltip ? `<span class="info-wrap">
                    <button class="info-btn" aria-label="Info">i</button>
                    <span class="info-tooltip">${item.tooltip}</span></span>` : ''}
                </div>
                ${item.hint ? `<div class="slider-hint">${item.hint}</div>` : ''}
            </div>
            </div>
            <div class="slider-controls">
            <input type="number" class="num-box" min="0" max="${budget(vals)}" step="1" value="${val}"
                data-idx="${i}" data-type="${dtype}" data-cat="${catIdx}"
                oninput="handleNumInput(this)">
            <div class="slider-track-wrap">
                <span class="snap-wrap" style="left:${snapPos}">
                <span class="snap-tick"></span>
                <span class="snap-tip">Gleichbewertung (10 Pkt.)</span>
                </span>
                <input type="range" min="0" max="${budget(vals)}" step="1" value="${val}"
                style="background:${trackBg};width:100%"
                data-idx="${i}" data-type="${dtype}" data-cat="${catIdx}" data-snap="${snapPt}"
                oninput="handleSlider(this)">
            </div>
            </div>
        </div>`;
    });
    html += '</div>';
    return html;
}


function handleSlider(el) {
    const idx = +el.dataset.idx;
    const type = el.dataset.type;
    const cat = +el.dataset.cat;
    let val = +el.value;
    const snap = +el.dataset.snap;
    if (Math.abs(val - snap) <= 2) { val = snap; el.value = snap; }
    if (type === 'main') clampSet(mainVals, idx, val);
    else clampSet(subVals[cat], idx, val);
    refreshStep(type === 'main' ? 1 : 2 + cat);
}


function handleNumInput(el) {
    const idx = +el.dataset.idx;
    const type = el.dataset.type;
    const cat = +el.dataset.cat;
    let val = parseInt(el.value) || 0;
    if (type === 'main') clampSet(mainVals, idx, val);
    else clampSet(subVals[cat], idx, val);
    refreshStep(type === 'main' ? 1 : 2 + cat);
}

function refreshStep(stepIdx) {
    const container = document.getElementById('step-' + stepIdx);
    if (!container) return;
    const isSub = stepIdx >= 2;
    const catIdx = isSub ? stepIdx - 2 : null;
    const vals   = isSub ? subVals[catIdx] : mainVals;
    const badge  = container.querySelector('.budget-badge');
    if (badge) {
        const rem = remaining(vals);
        badge.className = 'budget-badge ' + (rem === 0 ? 'ok' : rem > 0 ? 'warn' : 'over');
        badge.querySelector('.dot').className = 'dot';
        const tot2 = budget(vals);
        const msg = rem === 0 ? 'Alle ' + tot2 + ' Punkte vergeben' : rem > 0 ? rem + ' von ' + tot2 + ' Punkten noch verfügbar' : Math.abs(rem) + ' Punkte zu viel';
        badge.lastChild.textContent = msg;
    }
    const items = isSub ? CATS[catIdx].subs : CATS;
    items.forEach((item, i) => {
        const val = vals[i];
        const color  = isSub ? CATS[catIdx].color : CATS[i].color;
        const range  = container.querySelector(`input[type=range][data-idx="${i}"]`);
        const numBox = container.querySelector(`input.num-box[data-idx="${i}"]`);
        if (range) { 
            range.max   = budget(vals); 
            range.value = val; 
            const pct   = budget(vals) > 0 ? (val / budget(vals) * 100) : 0;
            const tOff  = budget(vals) > 0 ? (val / budget(vals) * 16 - 8) : -8; 
            range.style.background = `linear-gradient(to right, ${color} calc(${pct}% - ${tOff.toFixed(2)}px), #E2DED6 calc(${pct}% - ${tOff.toFixed(2)}px))`; 
        }
        if (numBox) numBox.value = val;
    });
    const nextBtn = container.querySelector('.btn.primary');
    if (nextBtn) nextBtn.disabled = remaining(vals) !== 0;
    updateChart(stepIdx, vals, isSub, catIdx);
}

function updateChart(stepIdx, vals, isSub, catIdx) {
    const chartKey = 'chart-' + stepIdx;
    if (!charts[chartKey]) return;
    charts[chartKey].setOption({ series: [{ data: buildChartData(vals, isSub, catIdx) }] });
}

function generateSubColors(hex, n) {
    const colors = [];
    const start = 0.95;
    const step = n > 1 ? (0.5 / (n - 1)) : 0;
    for (let i = 0; i < n; i++) {
        colors.push(blendWithWhite(hex, start - i * step));
    }
    return colors;
}

function blendWithWhite(hex, ratio) {
    const r  = parseInt(hex.slice(1, 3), 16);
    const g  = parseInt(hex.slice(3, 5), 16);
    const b  = parseInt(hex.slice(5, 7), 16);
    const nr = Math.round(r + (255 - r) * (1 - ratio));
    const ng = Math.round(g + (255 - g) * (1 - ratio));
    const nb = Math.round(b + (255 - b) * (1 - ratio));
    return '#' + [nr, ng, nb].map(x => x.toString(16).padStart(2, '0')).join('');
}

function buildChartData(vals, isSub, catIdx) {
    const items = isSub ? CATS[catIdx].subs : CATS;
    const colorList = isSub ? generateSubColors(CATS[catIdx].color, items.length) : CATS.map(c => c.color);
    const total = vals.reduce((a, b) => a + b, 0);
    const rem = budget(vals) - total;
    const data = items
        .map((item, i) => vals[i] > 0 ? { value: vals[i], name: item.label, itemStyle: { color: colorList[i] } } : null)
        .filter(Boolean);
    if (rem > 0) {
        data.push({ value: rem, name: 'Noch zu vergeben', itemStyle: { color: '#E8E5DF' }, emphasis: { disabled: true } });
    }
    return data;
}


// but also refresh
function initChart(stepIdx, vals, isSub, catIdx) {
    const chartKey = 'chart-' + stepIdx;
    const el = document.getElementById('echarts-' + stepIdx);
    if (!el) return;
    const chart = echarts.init(el);
    charts[chartKey] = chart;
    chart.setOption({
        tooltip: { trigger: 'item', formatter: p => p.name === 'Noch zu vergeben' ? `<b>${p.value} Punkte</b> noch zu vergeben` : `${p.name}<br/><b>${p.value} Punkte</b>` },
        series: [{
            type: 'pie', radius: ['38%', '68%'],
            avoidLabelOverlap: false,
            itemStyle: { borderRadius: 8, borderColor: '#FDFCFA', borderWidth: 3 },
            label:     { show: false },
            emphasis:  { scale: true, scaleSize: 6 },
            labelLine: { show: false },
            data: buildChartData(vals, isSub, catIdx)
        }]
    });
}


function buildStep0() {
    return `
        <div class="step active" id="step-0">
            <div class="card">
                <div class="card-title">Über diese Befragung</div>
                <p class="card-desc" style="margin-top:0.5rem">Im Rahmen des jährlichen Länderindex Familienunternehmen werden Standortfaktoren 
                für Unternehmen im internationalen Vergleich systematisch bewertet. 
                Diese Befragung erfasst, welche der bewerteten Faktoren Sie für Ihr Unternehmen bei strategischen Standortüberlegungen 
                als besonders relevant einschätzen.</p>
            </div>
            <div class="card">
                <div class="card-title">So funktioniert die Bewertung</div>
                <div class="intro-feature" style="margin-top:12px">
                <strong>Punktevergabe</strong>Ihnen steht bei jeder Frage ein Budget an Punkten zur Gewichtung der Standortfaktoren zur Verfügung. 
                Vergeben Sie mehr Punkte an Faktoren, die Ihre Standortentscheidung stärker beeinflussen und weniger Punkte an Faktoren, 
                die für Ihre Entscheidung weniger ausschlaggebend sind. 10 Punkte pro Kategorie entsprechen einer Gleichgewichtung der Entscheidungsrelevanz. 
                Die Schaltfläche "Weiter" wird erst aktiv, wenn alle Punkte vergeben sind.
                </div>
                <div class="intro-grid">
                <div class="intro-feature"><strong>Schritt 1 – Hauptkategorien</strong>Verteilen Sie Punkte auf 
                        die Standortfaktoren Steuern, Arbeitskräfte, Finanzierung, Regulierung, Infrastruktur und Energie.</div>
                <div class="intro-feature"><strong>Schritt 2 – Unterkategorien</strong>Verteilen Sie innerhalb jedes Standortfaktors 
                    Punkte auf die 2–3 Unterkategorien.</div>
                </div>
            </div>

            <div class="btn-row">
                <button  accesskey="2" class="btn primary" onclick="goTo(1)">Weiter →</button>
            </div>
        </div>
    `;
}



function buildStep2() {
    return `
        <div class="step" id="step-1">
            <div class="card">
                <div class="card-title">Wie wichtig sind die folgenden Standortfaktoren für Ihr Unternehmen, 
                wenn Sie darüber nachdenken, Aktivitäten in Deutschland fortzuführen und auszubauen oder aber 
                zurückzufahren und gegebenenfalls in andere Länder zu verlagern?
                </div>
                <p class="card-desc">Gewichten Sie die Relevanz folgender Standortfaktoren, indem Sie das Budget 
                    von <strong>60 Punkten</strong> entsprechend aufteilen. 
                        <span class="info-wrap"><button class="info-btn" aria-label="Info">i</button><span class="info-tooltip">Vergeben Sie mehr 
                        Punkte an Faktoren, die Ihre Standortentscheidung stärker beeinflussen und weniger Punkte an Faktoren, 
                        die für Ihre Entscheidung weniger ausschlaggebend sind. 
                        10 Punkte pro Kategorie entsprechen einer Gleichgewichtung der Entscheidungsrelevanz.
                        </span></span>
                </p>
            </div>
            <div class="card">
                ${buildBudgetBadge(mainVals)}
                <div class="pie-layout">
                <div class="pie-container" id="echarts-1"></div>
                <div>${buildSliders(mainVals, null, undefined)}</div>
                </div>
            </div>
            <div class="btn-row">
                <button  accesskey="p" class="btn btn-back" onclick="goTo(0)">← Zurück</button>
                <button  accesskey="2" class="btn primary"  onclick="goTo(2)" ${remaining(mainVals) !== 0 ? 'disabled' : ''}>Weiter →</button>
            </div>
        </div>
    `;
}

function buildSubStep(catIdx) {
    const cat = CATS[catIdx];
    const stepIdx = 2 + catIdx;
    const vals = subVals[catIdx];
    const isLast = catIdx === CATS.length - 1;
    return `
        <div class="step" id="step-${stepIdx}">
            <div class="card">
                <div class="card-title"> Innerhalb des Standortfaktors <b>${cat.label}</b>, 
                    wie wichtig sind die folgenden Faktoren für Ihr Unternehmen, wenn Sie darüber nachdenken, 
                    Aktivitäten in Deutschland fortzuführen und auszubauen oder aber zurückzufahren und gegebenenfalls 
                    in andere Länder zu verlagern?
                </div>
                <p class="card-desc">Gewichten Sie die Relevanz folgender Standortfaktoren, indem Sie das Budget von <strong>${cat.subs.length * 10} Punkten</strong> entsprechend aufteilen. <span class="info-wrap"><button class="info-btn" aria-label="Info">i</button><span class="info-tooltip">Vergeben Sie mehr Punkte an Faktoren, die Ihre Standortentscheidung stärker beeinflussen und weniger Punkte an Faktoren, die für Ihre Entscheidung weniger ausschlaggebend sind. 10 Punkte pro Kategorie entsprechen einer Gleichgewichtung der Entscheidungsrelevanz.</span></span></p>
            </div>
            <div class="card">
                ${buildBudgetBadge(vals)}
                <div class="pie-layout">
                <div class="pie-container" id="echarts-${stepIdx}"></div>
                <div>${buildSliders(vals, catIdx, true)}</div>
                </div>
            </div>
            <div class="btn-row">
                <button accesskey="p" class="btn btn-back" onclick="goTo(${stepIdx - 1})"  >← Zurück</button>
                <button accesskey="2" class="btn primary"  onclick="${isLast ? 'showResults()' : 'goTo(' + (stepIdx + 1) + ')'}" ${remaining(vals) !== 0 ? 'disabled' : ''}>
                ${isLast ? 'Ergebnisse ansehen →' : 'Weiter → '}
                </button>
            </div>
        </div>
    `;
}


function buildResultsStep() {
    return `
        <div class="step" id="step-results">
            <div class="card">
                <div class="thank-you-icon">
                <svg viewBox="0 0 24 24"><polyline points="20 6 9 17 4 12"></polyline></svg>
                </div>
                <div class="card-title">Vielen Dank für Ihre Teilnahme</div>
                <p class="card-desc" style="margin-bottom:1.5rem">Ihre Angaben wurden erfasst. 
                Unten sehen Sie eine Zusammenfassung Ihrer Gewichtungen.</p>
                <div id="results-content"></div>
            </div>
            <div class="btn-row">
                <button class="btn btn-back" onclick="goBackFromResults()">← Zurück</button>
            </div>
        </div>
    `;
}



function goTo(idx) {
    document.querySelectorAll('.step').forEach(s => s.classList.remove('active'));
    const el = document.getElementById('step-' + idx);
    if (el) { 
        el.classList.add('active'); 
        currentStep = idx; 
    }
    updateProgress();
    if (idx === 1) setTimeout(() => { initChart(1, mainVals, false, null); }, 50);
    if (idx >= 2 && idx <= 7) { 
        const catIdx = idx - 2; 
        setTimeout(
            () => { initChart(idx, subVals[catIdx], true, catIdx); }, 50
        ); 
    }
    // xxxx
    console.log(mainVals);
    console.log(subVals);

    window.scrollTo({ top: 0, behavior: 'smooth' });
}



function goBackFromResults() {
    goTo(2 + CATS.length - 1);
}


function showResults() {
    document.querySelectorAll('.step').forEach(s => s.classList.remove('active'));
    const el = document.getElementById('step-results');
    if (el) { el.classList.add('active'); currentStep = TOTAL_STEPS - 1; }
    updateProgress();
    submitResults({ main: mainVals, subs: subVals });

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

    setTimeout(() => {
        const mainChart = echarts.init(document.getElementById('res-main-pie'));
        mainChart.setOption({
            tooltip: { trigger: 'item', formatter: p => `${p.name}<br/><b>${p.value} Punkte</b>` },
            legend: { bottom: 0, left: 'center', textStyle: { fontSize: 12, fontFamily: 'DM Sans' }, itemWidth: 12, itemHeight: 12 },
            series: [{
                type: 'pie', radius: ['32%', '56%'], center: ['50%', '40%'], avoidLabelOverlap: true,
                itemStyle: { borderRadius: 6, borderColor: '#FDFCFA', borderWidth: 3 },
                label:     { show: true, formatter: p => p.value > 0 ? p.value + ' Pkt.' : '', fontSize: 12, fontFamily: 'DM Sans', color: '#2C2A26' },
                labelLine: { show: true, showAbove: false, length: 8, length2: 6 },
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
                tooltip: { trigger: 'item', formatter: p => `${p.name}<br/><b>${p.value} Punkte</b>` },
                series: [{
                    type: 'pie', radius: ['30%', '58%'], avoidLabelOverlap: false,
                    itemStyle: { borderRadius: 6, borderColor: '#FDFCFA', borderWidth: 2 },
                    label: { show: true, formatter: p => p.value > 0 ? p.value + ' Pkt.' : '', fontSize: 11, fontFamily: 'DM Sans', color: '#2C2A26' },
                    labelLine: { show: true, length: 6, length2: 4 },
                    data: cat.subs.map((sub, j) => ({
                        value: subVals[i][j], name: sub.label, itemStyle: { color: subColors[j] },
                        label: { show: subVals[i][j] > 0 }, labelLine: { show: subVals[i][j] > 0 }
                    }))
                }]
            });
        });
    }, 80);

    window.scrollTo({ top: 0, behavior: 'smooth' });
}


function init() {
    const container = document.getElementById('steps-container');
    let html = '';
    html += buildStep0();
    html += buildStep2();
    CATS.forEach((_, i) => { 
        html += buildSubStep(i); 
        console.log(`sub step ${i}`)
    });
    html += buildResultsStep();
    container.innerHTML = html;
    updateProgress();
}

init();