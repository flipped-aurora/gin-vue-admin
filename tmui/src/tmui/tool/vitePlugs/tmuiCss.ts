/**
 * 作者 tmui | https://tmui.design
 * 版权:本插件只服务于tmui组件库,不得移植复制到任何地方使用,侵权必究.
 */
import { normalizePath } from 'vite'
import {path} from "path"
import * as file from "fs"
export default function(opts?:any){
    return {
        name:"tmui-vite-css",
		// 该插件在 plugin-vue 插件之前执行，这样就可以直接解析到原模板文件
		enforce: 'pre',
        handleHotUpdate({file,modules,read}){
			
        },
		// 代码转译，这个函数的功能类似于 `webpack` 的 `loader`
		transform(code:string, id:string, opt) {
			
			const rule = /\.[n]{0,}vue$/;
			const customRule = /\<tm-sheet.*\>/g;
			if (!rule.test(id)) return code;
			let fpath = normalizePath(id).toLocaleLowerCase();
			//只处理页面。
			if(fpath.indexOf('tmui/components')>-1) return;
			//过滤不处理的特殊文件
			if(checkFileisTmui(id)) return code;
			const nodelistStringtag = code.match(customRule)
			
			// console.log(nodelistStringtag)
			if(!nodelistStringtag) return;
			/** 正则color属性。 */
			let attrRule = /\s(color)=([\s]{0,}".*")|\s(color)[\s]{0,}=[\s]{0,}('.*')/g
			nodelistStringtag.forEach(el=>{
				el.match(attrRule)
			})
			
			return code;
		}
    }
}


function checkFileisTmui(filepath:string){
	let autoColorCom = [
		'app.vue',
		]
	filepath = normalizePath(filepath).toLocaleLowerCase()
	let filename = filepath.substring(filepath.lastIndexOf('/')+1);
	return autoColorCom.includes(filename)
}

function getFileName(id:string){
	let filepath = normalizePath(id)
	let filename = filepath.substring(filepath.lastIndexOf('/')+1);
	return filename
}
