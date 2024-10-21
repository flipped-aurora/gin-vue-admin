export interface qrOpts {
	baseColor?: string //背景色,可以是数组颜色，将产生渐变。，渐变方向见下方：linearDir
	, backgroundImage?: string //使用图片作为二维码背景。
	, backgroundColor?: string | Array<string> // 背景色,可以是数组颜色，将产生渐变。，渐变方向见下方：linearDir
	, size?: number //图片大小
	, border?: number //边 width = size * border,比如二维码是200,那么想要让边为10那么 border = 200*0.05
	, str?: string //内容
	, forgroundColor?: string | Array<string> //前景色,也可以是["#FF0000","#FFFF00"]如果提供数组，将会绘制渐变色。
	, logoImage?: string //logo图片
	, logoWidth?: number
	, logoHeight?: number
	, ecc?: string // 容错等级[ 'L', 'M', 'Q', 'H' ]
	//left,right,bottom,top,tlbr:左顶点至底右下点，trbl,右顶点底右左点，bltr右底左点至顶右点。brtl底右点至顶左点。
	, linearDir?: string //如果forgroundColor是数组渐变色，则此可以更改渐变方向。
}
export const qrOptsDefault = {
	baseColor: '#fff'
	, backgroundImage: null
	, backgroundColor: null
	, size: 300
	, border: 0.05
	, str: 'tmui'
	, forgroundColor: '#000'
	, logoImage: null
	, logoWidth: 20
	, logoHeight: 20
	, ecc: 'M'
	, linearDir: "tlbr"
} 