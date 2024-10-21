import { onResize } from "@dcloudio/uni-app";
import { reactive, getCurrentInstance, onMounted, ref, ComponentPublicInstance } from "vue"

interface elementSizeType {
	left : number,
	top : number,
	bottom : number,
	right : number,
	width : number,
	height : number,
	id : string,
	dataset : {
		[key : string] : any
	}
}
/**
 * 获取元素的尺寸，只能在生命期内使用
 * select:需要查询的选择器，可以是.class,或者#id,
 * isAll:是否查询该节点所有元素，默认为单个。
 * return {size<Ref<elementSize>>,allSize<Ref<elementSize[]>>},解构出来的值，需要使用.value来访问。
 * */
export const useElementSize = (select : string, isAll : boolean = false) => {
	const proxy = getCurrentInstance()?.proxy ?? null;
	let elementSize = ref<elementSizeType>({
		left: 0,
		top: 0,
		bottom: 0,
		right: 0,
		width: 0,
		height: 0,
		dataset: {},
		id: ""
	});

	let elementSIzeList = ref<elementSizeType[]>([]);

	function init() {
		if (isAll) {
			uni.createSelectorQuery()
				.in(proxy)
				.selectAll(select)
				.boundingClientRect()
				.exec((res) => {
					elementSIzeList.value = [...res[0]]
					if (elementSIzeList.value?.length > 0) {
						elementSize.value = elementSIzeList.value[0]
					}
				})
		} else {
			uni.createSelectorQuery()
				.in(proxy)
				.select(select)
				.boundingClientRect()
				.exec((res) => {
					if (res && Array.isArray(res) && res?.length > 0) {
						elementSize.value = res[0]
					}
				})
		}
	}

	onMounted(() => {
		setTimeout(function () {
			init()
		}, 50);
	})

	onResize(() => {
		setTimeout(function () {
			init()
		}, 50);
	})


	return {
		size: elementSize,
		allSize: elementSIzeList,
	}
}