var chartDom = null;
window.mychart = null;
window.echart_createDom = function (w, h) {
    w = Number(w);
    h = Number(h);
    chartDom = document.createElement("div");
    chartDom.style.width = w + 'px';
    chartDom.style.height = h + 'px';
    chartDom.style.display = 'block';
    document.body.appendChild(chartDom);
    document.all.addEventListener('touchmove', function(evt) {
    	evt.preventDefault();
    });
	
    return chartDom;
}
window.echart_createChart = function (opts) {
    if (!opts) {
        opts =  {}
    }
    window.mychart = echarts.init(chartDom, undefined, opts)
    return window.mychart;
}



window.echart_setOption= function (opts,ops) {
    window.mychart.setOption(opts,ops)
}