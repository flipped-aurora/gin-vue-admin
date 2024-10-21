import { EChartsCoreOption, EChartsOption } from "echarts";

class mytmcharts {
	web = null;
	width = 0;
	height = 0;
	optionNow: EChartsOption | null = null;
	constructor(c: any, w = 0, h = 0) {
		this.web = c;
		this.width = w;
		this.height = h;
	}
	setOption(option: EChartsOption, opts: any = {}) {
		if (!this.web) return;
		
		// 处理echart配置参数，若有函数则将其toString
		const optionString = JSON.stringify(option, (key, value) => {
			if (typeof value === 'function') {
				return JSON.stringify(
					{
						type:'echartCbFn',
						fnString: value.toString()
					}
				)
			}
			return value;
		});
				
				
		this.optionNow = option;
		this.web.evalJs(`echart_setOption(${optionString},${JSON.stringify(opts)})`)
	}
	getWidth() {
		return this.width;
	}
	getHeight() {
		return this.height;
	}
	getDom() {
		return this.web;
	}
	getOption() {
		return this.optionNow;
	}
	resize() {
		if (!this.web) return;
		this.web.evalJs(`mychart.resize()`)
	}
	showLoading(opts: any) {
		if (!this.web) return;
		this.web.evalJs(`mychart.showLoading(${JSON.stringify(opts)})`)
	}
	hideLoading() {
		if (!this.web) return;
		this.web.evalJs(`mychart.hideLoading()`)
	}
	getDataURL() {
		//暂不实现。
	}
}



export default mytmcharts