    // Progressive disclosure: each .next button reveals the next section
    (function() {

        function showSection(n) {
        var current = document.querySelector('[data-section="' + n + '"]');
        if (current) current.style.display = 'none';
        var next = document.querySelector('[data-section="' + (n + 1) + '"]');
        if (next) {
          next.style.display = 'block';
          var firstInput = next.querySelector('input[type="text"], input[type="number"]');
          if (firstInput) firstInput.focus();

          // On showing section 2, (re)render charts so sizes are correct
          if (next.id === 'section2') {
            setTimeout(function() {
              if (window.updateCharts) window.updateCharts();
              ['distanceChart','forecastChart','consensusChart'].forEach(function(id){
                var el = document.getElementById(id);
                if (!el) return;
                var inst = echarts.getInstanceByDom(el);
                if (inst) inst.resize();
              });
            }, 0);
          }
        }
      }

	  // Attach handlers to all .next buttons
	  document.querySelectorAll('.next').forEach(function(btn) {
		btn.addEventListener('click', function() {
		  var n = parseInt(this.getAttribute('data-next'), 10) - 1; // current section index
		  showSection(n);

		  // hide the Section-2 "Weiter" button after clicking
		  if (this.getAttribute('data-next') === '7') {
			// hide just the button:
			this.style.display = 'none';
			// (optional) hide the whole button row instead:
			// const wrapper = this.closest('.buttons'); if (wrapper) wrapper.style.display = 'none';
		  }
		});
	  });

      // Optional: allow Enter to act like clicking Weiter if a single text input is in focus
      document.addEventListener('keydown', function(e) {
        if (e.key === 'Enter') {
          var active = document.activeElement;
          if (active && (active.tagName === 'INPUT' || active.tagName === 'SELECT')) {
          // if (true) {
            var visibleSection = Array.from(document.querySelectorAll('[data-section]')).find(function(sec){
              return sec.style.display !== 'none';
            });
            if (visibleSection) {
              var nextBtn = visibleSection.querySelector('.next');
              if (nextBtn) {
                e.preventDefault();
                nextBtn.click();
              }
            }
          }
        }
      });
    })();

    // === Functionality from share treatment v2 ===
    document.addEventListener('DOMContentLoaded', () => {
      const respondentSelect = document.getElementById('respondentSelect');
      const quarterSelectRadios = document.querySelectorAll('input[name="quarter"]');
      const userShareInput = document.getElementById('userShareInput');
      const userShareSlider = document.getElementById('userShareSlider');
      const shareComparisonText = document.getElementById('shareComparisonText');
      const csvFileInput = document.getElementById('csvFileInput');
      const downloadBtn = document.getElementById('downloadBtn');

      const distanceChart = echarts.init(document.getElementById('distanceChart'));
      const forecastChart = echarts.init(document.getElementById('forecastChart'));
      const consensusChart = echarts.init(document.getElementById('consensusChart'));

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
        const respondentIndex = parseInt(respondentSelect.value, 10);
        const selectedQuarter = getSelectedQuarter();

        if (isNaN(respondentIndex) || !parsedData[respondentIndex]) {
          distanceChart.clear();
          forecastChart.clear();
          consensusChart.clear();
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
              data: [[ { coord: [userShare, ''] }, { coord: [actualShare, ''] } ]]
            } : {}
          }]
        };
        distanceChart.setOption(distanceChartOption);

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
        forecastChart.setOption(forecastOption);

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
        consensusChart.setOption(consensusOption);
      }
      window.updateCharts = updateCharts;

      function handleFile(file) {
        if (file) {
          Papa.parse(file, {
            header: true,
            skipEmptyLines: true,
            dynamicTyping: true,
            complete: (results) => {
              parsedData = results.data;
              respondentSelect.innerHTML = '';
              parsedData.forEach((_, index) => {
                const option = document.createElement('option');
                option.value = index;
                option.textContent = `Befragte/r ${index + 1}`;
                respondentSelect.appendChild(option);
              });
              updateCharts();
            }
          });
        }
      }

      function downloadData() {
        const respondentIndex = parseInt(respondentSelect.value, 10);
        const selectedQuarter = getSelectedQuarter();
        if (isNaN(respondentIndex) || !parsedData[respondentIndex]) {
          alert('Bitte zuerst eine/n Befragte/n auswählen.');
          return;
        }
        const respondentData = parsedData[respondentIndex];
        const filteredData = {
          [`growth${selectedQuarter}`]: respondentData[`growth${selectedQuarter}`],
          [`consensus${selectedQuarter}`]: respondentData[`consensus${selectedQuarter}`],
          [`grshare${selectedQuarter}`]: respondentData[`grshare${selectedQuarter}`]
        };

        const csvContent = Papa.unparse([filteredData], { header: true });
        const blob = new Blob([csvContent], { type: 'text/csv;charset=utf-8;' });
        const link = document.createElement('a');
        const url = URL.createObjectURL(blob);
        link.setAttribute('href', url);
        link.setAttribute('download', `befragte_${respondentIndex + 1}_${selectedQuarter}.csv`);
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
      }

      // Sync slider and input
      userShareInput.addEventListener('input', () => {
        userShareSlider.value = userShareInput.value;
        updateCharts();
      });
      userShareSlider.addEventListener('input', () => {
        userShareInput.value = userShareSlider.value;
        updateCharts();
      });

      csvFileInput.addEventListener('change', (e) => handleFile(e.target.files[0]));
      respondentSelect.addEventListener('change', updateCharts);
      quarterSelectRadios.forEach(radio => radio.addEventListener('change', updateCharts));
      downloadBtn.addEventListener('click', downloadData);

      window.addEventListener('resize', () => {
        distanceChart.resize();
        forecastChart.resize();
        consensusChart.resize();
      });
    });
