// https://stackoverflow.com/questions/66363002/how-to-add-images-or-shapes-to-a-chart


let symbolGuindos = {
    value: 'Guindos',
    symbol: vectorImgs.Reindeer,
    symbolSize: [60, 60]
};




// https://echarts.apache.org/en/option.html#graphic.elements-rect
let graphicRect = {
  id: 'rect01',
  type: 'rect',
  // translation; default [0, 0]
  position: [100, 60],
  // scale; default  [1, 1]
  scale: [1.2, 0.4],
  // rotation; default  0. Negative -> rotating clockwise.
  rotation: Math.PI / 4,
  // origin point of rotation and scale; default [0, 0]
  origin: [250, 50],

  z: 101,
  shape: {
    // rect props
    x:  100,
    y:  50,
    width: 50,
    height: 150,

    // circle props
    cx: 150,
    cy: 100,
    r:   60,
  },
  style: {
    fill:   'green',
    stroke: 'blue',
    lineWidth: 2,
  }
};



let graphicImg = {
  id: 'img1',
  type: 'image',
  symbol: vectorImgs.Reindeer,
  textContent: {
    style: {
      x: 550,
      y: 150,
      text: [
        'image',
        'caption',
        'multiline '
      ].join('\n'),
      font: '20px "STHeiti", sans-serif'
    },
  },
  style: {
    fill:   'red',
    stroke: 'blue',
    lineWidth: 2,
  }
};





let graphicElements = [{
  elements: [
    graphicRect,
    graphicImg,
    {
      id: 'text1',
      type: 'text',
    },
    {
      id: 'small_circle',
      type: 'circle',
      z: 100,
      shape: {
        cx: 350,
        cy: 200,
        r: 20,
      },
      style: {
        fill: 'rgba(0, 140, 250, 0.5)',
        stroke: 'rgba(0, 50, 150, 0.5)',
        lineWidth: 2,
      }
    },
  ]
}];

// option.graphic = graphicElements;
function setIt(chartObj) {
  chartObj.setOption(
    {
      'graphic': graphicElements,
      'yAxis': {},
    }
  );
}


let lk = document.createElement("A");
lk.setAttribute("href","#")
lk.innerHTML = "graphic elements"
lk.setAttribute("onclick","setIt(myChart)")

var body = document.getElementsByTagName("body")[0];
var lastchild = body.lastChild;
document.body.insertBefore(lk,lastchild);

// document.body.insertBefore(lk,document.body.childNodes[0]);
// document.appendChild(lk);


console.log("graphic-elements.js end")