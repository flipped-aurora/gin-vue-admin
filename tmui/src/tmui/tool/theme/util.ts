export function setDomDarkOrWhite(){
	// #ifdef H5
	let  tmuiNavStyle= localStorage.getItem("tmuiNavStyle")
	let dark = localStorage.getItem('setTmVuetifyBlack');
	try{
		if(tmuiNavStyle&&typeof tmuiNavStyle !='undefined'){
			// dark = JSON.parse(dark||'{"type":"boolean","data":false}')
			// let tmuiNavStyleJson = JSON.parse(tmuiNavStyle);
			// if(document.querySelector("#tmuiBodyId")){
			// 	document.body.removeChild(document.querySelector("#tmuiBodyId"))
			// }
			// let style = document.createElement('style')
			// style.type="text/css"
			// style.id="tmuiBodyId";
			// if(dark?.data==true||dark==true){
			// 	style.append(document.createTextNode(`.uni-page-head{background-color:${tmuiNavStyleJson?.navbarBackground??'none'} !important} body{background:#000}`))
			// }else{
			// 	style.append(document.createTextNode(`.uni-page-head{background-color:${tmuiNavStyleJson?.navbarBackground??'none'} !important} `))
			// }
			// document.body.append(style)
		}
	}catch(e){
		//TODO handle the exception
	}
	// #endif
}