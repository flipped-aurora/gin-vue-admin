import { tmRoundRect } from './shape/roundRect';
// @ts-expect-error
import * as TWEEN from "@/tmui/tool/lib/tween.min.js";
import { getCanvas } from "@/tmui/tool/function/getCanvas";
// #ifdef APP-NVUE
import { enable, WeexBridge } from "@/tmui/tool/gcanvas/index.js";
const dom = uni.requireNativePlugin("dom");
// #endif
import { ComponentPublicInstance } from "vue"
import { Shape } from "./shape";
export interface tmCvOptsType {
	left?: number,
	top?: number,
	width: number,
	height: number,
	dpr?: number,
	platform?: 'uni' | 'web'
}
var requestAnimationFrameId:IdleRequestCallback|any = 0;
let stopDrawer = false;
let tween:TWEEN|null = null;
/**
 * tmCv
 * 协议：MIT
 * 功能特色：
 * 1/图形为插件式类对象，分模块加载，大小可控，对小程序非常有效
 * 2/保留了高效的内部动画，动画简单有趣，使用超级简单
 * @description tmCv是tmui作者专为uniapp开发轻量级的canvas绘制插件
 * @copyright tmui保留版权，不可随意拿来更改公开发布为自己的产品。不接受二次开发并公布为二次产品。
 * @author tmui3.0附带库
 * @default 2023年3月29日
 */
export class tmCv {
	private requestIdleCallback: Function | null = null
	private requestAnimationFrame: Function | null = null
	/**onFramdi，关闭时需要注销 */
	private timid: any = null;
	ctx: UniApp.CanvasContext | null = null;
	render: HTMLCanvasElement | null = null;
	private canvasId: string = ""
	private proxy: ComponentPublicInstance | null = null;
	graphs: Array<Shape> = [];
	opts: tmCvOptsType = {
		left: 0,
		top: 0,
		width: 0,
		height: 0,
		dpr: 1,
		platform: 'uni'
	}
	constructor(proxy: ComponentPublicInstance, domId: string, opts?: tmCvOptsType) {
		this.canvasId = domId;
		this.proxy = proxy;
		if (typeof opts !== 'undefined') {
			this.opts = { ...this.opts, ...opts }
		}

	}
	init(): Promise<tmCv> {
		return this._getNodes(this.proxy)
	}

	add(graph: any | any[]) {
		if (Array.isArray(graph)) {
			this.graphs.push(...graph)
			return
		}
		this.graphs.push(graph)
	}
	draw() {
		if (!this.ctx) return;
		this.ctx.clearRect(0, 0, this.opts.width, this.opts.height)
		for (const rect of this.graphs) {
			rect.draw()
		}
		if (this.ctx.draw) {
			this.ctx.draw()
		}
	}

	public animation(arg: { duration?: number, repeat?: number, yoyo?: boolean } = { duration: 500, repeat: 0, yoyo: false },
		onUpdate: (progress: number) => void, onStart?: (progress: number) => void): Promise<{ [key: string]: any }> {
		return new Promise((res, rej) => {
			
			tween = new TWEEN.Tween({ progress: 0 })
				.easing(TWEEN.Easing.Linear.None) // 缓动函数
				.to({ progress: 1 }, arg?.duration ?? 0)
				.onUpdate((e: any) => {
					onUpdate(e.progress)
				})
				.onStart((e: any) => {
					if (onStart) {
						onStart(e.progress)
					}
				})
				.onComplete((e: any) => {
					res(e)
				})
				.delay(0)
				.repeat(arg?.repeat ?? 0)
				.yoyo(arg?.yoyo ?? false)
				.start();
		})
	}

	private _getNodes(proxy: ComponentPublicInstance | null): Promise<tmCv> {
		let t = this;
		let sys = uni.getSystemInfoSync()
		
		function initOpts(node: any) {
			try {
				if (typeof window !== 'undefined') {
					t.requestIdleCallback = window.requestIdleCallback;
					t.requestAnimationFrame = window.requestAnimationFrame;
				}
				// @ts-expect-error
				if (node && node !== null && typeof node != 'undefined' && node?.requestAnimationFrame && t.requestAnimationFrame == null) {
					t.requestAnimationFrame = node.requestAnimationFrame
					t.requestIdleCallback = node.cancelAnimationFrame
					t.render = node;
				}
				if (t.requestAnimationFrame == null) {
					t.requestAnimationFrame = uni.$tm.u.requestAnimationFrame;
					t.requestIdleCallback = uni.$tm.u.cancelAnimationFrame
				}
			} catch (error) {
				console.error("tmCv:", error)
			}


		}
		
		return new Promise((res, rej) => {

			if (!this.proxy) {
				res(this)
				return;
			}
			let delay = 10
			// #ifdef APP-NVUE
			let isAndroid = sys.osName == "android";
			if (isAndroid) {
				delay = 250
			} else {
				delay = 100
			}
			// #endif
			// #ifdef MP
			delay = 60
			// #endif
			// #ifdef APP-VUE
			delay = 30
			// #endif
			// #ifdef APP-NVUE
			
			setTimeout(() => {
				/*获取元素引用*/
				let domId = t.canvasId.replace(/[\.|#]/g, "")
				var ganvas: any = proxy?.$refs[domId] ?? null;
				dom?.getComponentRect(ganvas, function (res: any) {
					t.opts.left = res?.size?.left ?? 0;
					t.opts.top = res?.size?.top ?? 0;
					t.opts.dpr = sys.pixelRatio;
					// @ts-ignore
					getCanvas(t.proxy, domId, t.opts.width, t.opts.height).then((e) => {
						t.ctx = e.ctx
						initOpts(e.node)
						res(t)
					}).catch(error => {
						console.error(error)
					})

				});

			}, delay);
			// #endif
			// #ifndef APP-NVUE
			setTimeout(() => {
				
				uni
					.createSelectorQuery()
					.in(proxy)
					.select(".canvas")
					.boundingClientRect()
					.exec(((resNodes) => {
						const result = resNodes[0]
						// @ts-ignore
						this.opts.left = result?.left ?? 0;
						// @ts-ignore
						this.opts.top = result?.top ?? 0;
						this.opts.dpr = sys.pixelRatio;
						// @ts-ignore
						getCanvas(proxy, this.canvasId, this.opts.width, this.opts.height).then((e) => {
							initOpts(e.node)
							this.ctx = e.ctx
							stopDrawer = false;
							animate(this.requestAnimationFrame)
							res(this)
						}).catch(error => {

						})
					}));
			}, delay);
			// #endif
		})

	}

	/**注销tmCv */
	destory() {
		
		if (!this.requestIdleCallback) return;
		stopDrawer = true;
		try{
			this.requestIdleCallback(requestAnimationFrameId);
			requestIdleCallback(requestAnimationFrameId)
		}catch(e){
			//TODO handle the exception
		}
	}
}


function animate(requestAnimationFrames: any) {
	
	if (!requestAnimationFrames) return;
	function animateVC() {
		if(stopDrawer){
			try{
				tween.stop();
			}catch(e){
				//TODO handle the exception
			}
			tween = null;
			clearTimeout(requestAnimationFrameId)
			return;
		}
		requestAnimationFrameId = requestAnimationFrames(animateVC);
		TWEEN.update();
	}
	animateVC()
}

export {
	Shape
}


