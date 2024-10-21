// 这是nvue版本的colorview组件专用
var _tmColorView_w = 0
var _tmColorView_h = 0
var dpr = window.devicePixelRatio;
window.tmColorView_showdiv = function(w,h){
	var ele = document.querySelector("#tmColorview");
	var canvasdom = ele.querySelector("canvas")
	ele.style.display = "block"
	ele.style.width = w+'px'
	ele.style.height = h+'px'
	canvasdom.width = w
	canvasdom.height = h
	canvasdom.style.width = w+'px'
	canvasdom.style.height = h+'px'
	_tmColorView_w = w;
	_tmColorView_h = h;
	document.all.addEventListener('touchmove', function(evt) {
		evt.preventDefault();
	});
}
window.tmColorView_getCanvas = function(){
	var ele = document.querySelector("#tmColorview");
	var canvasdom = ele.querySelector("canvas");
	var ctx = canvasdom.getContext("2d")
	return ctx
}
window.tmColorView_renderColorHu = function(){
	var ctx = tmColorView_getCanvas();
	var dy = _tmColorView_h/3;
	var barcolorWidth = 30
	var x = 0
	var gradient = ctx.createLinearGradient( barcolorWidth/2, 0, barcolorWidth/2,dy);
	gradient.addColorStop(0, 'rgba(255,0,0,1)');
	gradient.addColorStop(0.5, 'rgba(255,0,255,1)');
	gradient.addColorStop(1, 'rgba(0,0,255,1)');
	ctx.fillStyle = gradient;
	ctx.fillRect(x, 0, barcolorWidth,  dy);
	if(ctx?.draw){
		ctx?.draw()
	}
	

	var gradient2 = ctx.createLinearGradient( barcolorWidth/2, dy, barcolorWidth/2, dy*2);
	gradient2.addColorStop(0, 'rgba(0,0,255,1)');
	gradient2.addColorStop(0.5, 'rgba(0,255,255,1)');
	gradient2.addColorStop(1, 'rgba(0,255,0,1)');
	ctx.fillStyle = gradient2;
	ctx.fillRect(x, dy, barcolorWidth,  dy);
	if(ctx?.draw){
		ctx?.draw(true)
	}
	var gradient3 = ctx.createLinearGradient( barcolorWidth/2,  dy*2, barcolorWidth/2,dy*3);
	gradient3.addColorStop(0, 'rgba(0,255,0,1)');
	gradient3.addColorStop(0.5, 'rgba(255,255,1,1)');
	gradient3.addColorStop(1, 'rgba(255,0,0,1)');
	ctx.fillStyle = gradient3;
	ctx.fillRect(x, dy*2,barcolorWidth,  dy);
	// ctx.drawImage(colorimgUrl, 0, 0,_width.value,_height.value)
	if(ctx?.draw){
		ctx?.draw(true)
	}
}
function hslaToRgba(scolor) {
	var { h, s, l, a } = scolor;
	h = h / 360;
	s = s / 100;
	l = l / 100;
	var rgb = [];

	if (s == 0) {
		rgb = [Math.round(l * 255), Math.round(l * 255), Math.round(l * 255)];
	} else {
		var q = l >= 0.5 ? (l + s - l * s) : (l * (1 + s));
		var p = 2 * l - q;
		var tr = rgb[0] = h + 1 / 3;
		var tg = rgb[1] = h;
		var tb = rgb[2] = h - 1 / 3;
		for (var i = 0; i < rgb.length; i++) {
			var tc = rgb[i];
			if (tc < 0) {
				tc = tc + 1;
			} else if (tc > 1) {
				tc = tc - 1;
			}
			switch (true) {
				case (tc < (1 / 6)):
					tc = p + (q - p) * 6 * tc;
					break;
				case ((1 / 6) <= tc && tc < 0.5):
					tc = q;
					break;
				case (0.5 <= tc && tc < (2 / 3)):
					tc = p + (q - p) * (4 - 6 * tc);
					break;
				default:
					tc = p;
					break;
			}
			rgb[i] = Math.round(tc * 255);
		}
	}
	
	return { r: rgb[0], g: rgb[1], b: rgb[2], a: a };
}
function rgbaToCss(sColor){
	return `rgba(${sColor.r},${sColor.g},${sColor.b},${sColor.a})`;
}
window.tmColorView_renderRectFill = function(H){
	if(!H) H=0;
	H = Number(H)
	let x = 40;
	let dy = 2;
	let w = _tmColorView_w;
	
	var ctx = tmColorView_getCanvas();
	for(let i=0;i<100;i++){
		let gradient = ctx.createLinearGradient( x, i, w,i);
		gradient.addColorStop(0, rgbaToCss(hslaToRgba({h:H,s:0,l:100-i,a:1})));
		gradient.addColorStop(1, rgbaToCss(hslaToRgba({h:H,s:100,l:50-i/2,a:1})));
		ctx.fillStyle = gradient;
		ctx.fillRect(x, (i+1)*dy, w,  (i+1)*dy);
		if(ctx?.draw){
			ctx?.draw(true)
		}
	}
}

window.tmColorView_getColor =  function(x,y,active){
	var ctx = tmColorView_getCanvas();
	let arg = ctx.getImageData(Number(x),Number(y),1,1);
	uni.postMessage({
	  data: {
	    action: 'tmColorView_getColor',
		tmColor:JSON.stringify({
			r:arg.data["0"],
			g:arg.data["1"],
			b:arg.data["2"],
			a:1
		}),
		tmColorActive:active
	  }
	})
}






