// --- Configurable colors & order ---
const colorByNode = new Map([
    ['a', '#1f77b4'],
    ['b', '#ff7f0e'],
    ['c', '#2ca02c']
]);
const nodeOrder = new Map([['a', 0], ['b', 1], ['c', 2]]); // <— fixed order

// --- Create a reusable Sankey renderer bound to a container ---
function createSankey(container) {

    // d3.sankeyLinkVertical();

    const svg = d3.select(container).append('svg');
    const gLinks = svg.append('g').attr('class', 'links');
    const gNodes = svg.append('g').attr('class', 'nodes');


    // Base sankey generator
    const sankey = d3.sankey()
    .nodeId(d => d.name)
    .nodeWidth(18)
    .nodePadding(28)
    .iterations(32)
    .nodeSort((a, b) => d3.ascending(nodeOrder.get(a.name) ?? 999, nodeOrder.get(b.name) ?? 999))
    // optional: ensure link z-order is stable (by source order then target)
    .linkSort((a, b) => d3.ascending(a.source.name, b.source.name) || d3.ascending(a.target.name, b.target.name));

    // State
    let bValue = 3;
    let cValue = 4;

    // Data skeleton
    const nodes = [
    { name: 'a', label: '100 €' },
    { name: 'b', label: 'Port 1' },
    { name: 'c', label: 'Port 2' }
    ];

    function getLinks() {
    return [
        { source: 'a', target: 'b', value: bValue },
        { source: 'a', target: 'c', value: cValue }
    ];
    }

    function resize() {
    const { width, height } = container.getBoundingClientRect();
    svg.attr('width', width).attr('height', height);
    sankey.extent([[8, 8], [width - 8, height - 8]]);
    render();
    }

    function render() {
    // Build graph and compute layout
    const graph = sankey({
        nodes: nodes.map(d => Object.assign({}, d)),
        links: getLinks().map(d => Object.assign({}, d))
    });

    // LINKS
    const linkSel = gLinks.selectAll('path').data(graph.links, d => d.source.name + '→' + d.target.name);
    linkSel.enter()
        .append('path')
        .attr('fill', 'none')
        .attr('stroke-opacity', 0.85)
        .attr('stroke-width', d => Math.max(1, d.width))
        .attr('stroke', d => colorByNode.get(d.source.name) || '#999')
        .merge(linkSel)
        .attr('d', d3.sankeyLinkHorizontal())
        .attr('stroke-width', d => Math.max(1, d.width));
    linkSel.exit().remove();

    // NODES
    const nodeSel = gNodes.selectAll('g.node').data(graph.nodes, d => d.name);
    const nodeEnter = nodeSel.enter()
        .append('g')
        .attr('class', 'node');

    nodeEnter.append('rect')
        .attr('rx', 2)
        .attr('ry', 2)
        .attr('stroke', '#000')
        .attr('stroke-opacity', 0.15);

    nodeEnter.append('text')
        .attr('dy', '0.35em')
        .attr('font-size', 12)
        .attr('text-anchor', 'start');

    const nodeMerged = nodeEnter.merge(nodeSel);

    nodeMerged.select('rect')
        .attr('x', d => d.x0)
        .attr('y', d => d.y0)
        .attr('height', d => Math.max(1, d.y1 - d.y0))
        .attr('width', d => Math.max(1, d.x1 - d.x0))
        .attr('fill', d => colorByNode.get(d.name) || '#888');

    nodeMerged.select('text')
        .attr('x', d => d.x1 + 6)
        .attr('y', d => (d.y0 + d.y1) / 2)
        .text(d => d.label ?? d.name);

    nodeSel.exit().remove();
    }

    // Public API
    function setValues(n) {
    bValue = n;
    cValue = 7 - n;
    render();
    }

    // Initial size + render
    resize();
    window.addEventListener('resize', resize);

    return { setValues };
}

// --- Instantiate and wire radios ---
const chart = createSankey(document.getElementById('chart_container'));

const radios = document.querySelectorAll('input[name="flowRatio"]');
for (let i = 0; i < radios.length; i = i + 1) {
    radios[i].addEventListener('change', function (evt) {
    const n = parseInt(evt.target.value, 10);
    chart.setValues(n);
    });
}
// Set initial
chart.setValues(3);
