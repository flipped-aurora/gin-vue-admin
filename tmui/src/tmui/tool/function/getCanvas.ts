import {ComponentInternalInstance} from "vue"
import CanvasRenderingContext2D from '../gcanvas/context-2d/RenderingContext.js';
// #ifdef APP-NVUE
import {enable,WeexBridge,} from '../gcanvas/index.js';
// #endif

export function getCanvas(proxy:ComponentInternalInstance,canvasid:string,w?:number,h?:number,is2d=false):Promise<{ctx:CanvasRenderingContext2D,node:HTMLCanvasElement|null}>{
	return new Promise((resv,rejv)=>{
		// #ifdef APP-NVUE
		setTimeout(async function () {
			
			/*获取元素引用*/
			var ganvas = proxy?.$refs[canvasid];
			/*通过元素引用获取canvas对象*/
			var canvasObj = enable(ganvas, {
				bridge: WeexBridge
			});
			resv({ctx:canvasObj.getContext('2d'),node:canvasObj})
		}, 100)
		// #endif
		
		
		// #ifdef MP
		setTimeout(async function () {
			const query = uni.createSelectorQuery().in(proxy)
				// #ifdef MP-ALIPAY
				
				query.select('#'+canvasid).node().exec((res2) => {
					const canvas = res2[0].node;
					let ctxvb:UniApp.CanvasContext = canvas.getContext('2d');
					resv({ctx:ctxvb,node:canvas})
				})
				// #endif
				
				// #ifndef MP-ALIPAY
				query.select('#'+canvasid)
					.fields({
						node: true,
						size: true,
						context:true
					})
					.exec((res) => {
						
						const canvas = res[0]?.node??null;
						
						if(canvas){
							const ctx = canvas.getContext('2d')
							const dpr = uni.getSystemInfoSync().pixelRatio
							canvas.width = res[0].width * dpr
							canvas.height = res[0].height * dpr
							if(ctx?.scale){
								ctx.scale(dpr, dpr)
							}
							resv({ctx:ctx,node:canvas})
						
						}else{
							const ctx =  res[0]?.context??null
							resv({ctx:ctx,node:null})
						}
						
				})
				// #endif
		}, 100)
		// #endif
		
	
		// #ifdef H5 || APP-VUE
		setTimeout(async function () {
			const query = uni.createSelectorQuery().in(proxy)
			
			resv({ctx:uni.createCanvasContext(canvasid, proxy),node:null})
			// return
			// query.select('#'+canvasid)
			// 	.fields({
			// 		node: true,
			// 		size: true,
			// 		context:true
			// 	})
			// 	.exec((res) => {
			// 		const dpr = uni.getSystemInfoSync().pixelRatio
			// 		let canvas = proxy.$el.querySelector('canvas').getContext('2d')
			// 		let canvasEl = proxy.$el.querySelector('canvas')
			// 		canvas.scale(0.1, 0.1)
			// 		resv({ctx:canvas,node:proxy.$el.querySelector('canvas')})
			// })
		}, 50)
		
		// #endif
		
	})
}



