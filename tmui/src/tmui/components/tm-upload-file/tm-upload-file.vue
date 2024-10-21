<template>
	<view class="flex">
		<tm-sheet :margin="[0,8]" :padding="[16]" :shadow="1" :round="2" :border="1" v-for="(item,index) in files" :key="index">
			<view class="flex flex-row flex-row-center-between ">
				<view @click="emits('click',item)" class="flex flex-col " style="width:85%">
					<tm-text :font-size="28"
					_class="wrap"
					class="flex-1 wrap mb-5"
					:label="item.name"></tm-text>
					<tm-text 
					:font-size="24"
					:line-height="0" 
					:color="item.statusCode==STATUS_CODE.fail||item.statusCode==STATUS_CODE.max?'red':(item.statusCode==STATUS_CODE.success?'green':'')" :label="item.status"></tm-text>
				</view>
				<tm-icon v-if="!_disabledRemove&&!_disabled" @click="remove(index)" class="pl-32 flex-shrink" color="red" name="tmicon-times"></tm-icon>
			</view>
		</tm-sheet>
		
		<tm-sheet :class="_disabled?'opacity-5':''" :followTheme="props.followTheme" :color="props.color" text :margin="[0,8]"  :padding="[16]" :round="2" :border="1"  v-if="!isLimitCount">
			<view @click="chooseFile" class="flex flex-row flex-row-center-center">
				<tm-icon  name="tmicon-plus"></tm-icon>
				<tm-text class="pl-16" label="添加文件"></tm-text>
			</view>
		</tm-sheet>
	</view>
</template>
<script lang="ts" setup>
/**
 * 文件上传
 * @description 不支持nvue,只支持h5,web,微信，qq,字节，飞书相关平台
 */
import { computed, ref, PropType, Ref, watch, toRaw, nextTick, getCurrentInstance, inject, reactive, onMounted, isRef, isProxy, watchEffect } from 'vue'
import { inputPushItem, rulesItem } from './../tm-form-item/interface'
import { USE_UPLOAD_FILE_CONFIG_TYPE,useUploadFile,FILE_TYPE,STATUS_CODE } from './../../tool/useFun/useUploadFile'
import tmImage from '../tm-image/tm-image.vue';
import tmText from '../tm-text/tm-text.vue';
import tmIcon from '../tm-icon/tm-icon.vue';
import tmSheet from '../tm-sheet/tm-sheet.vue';
import { emit } from 'cluster';
const { config,files,isLimitCount,choose,remove,addFile,addEventListener,uploading } = useUploadFile(null)
const proxy = getCurrentInstance()?.proxy ?? null

const props = defineProps({
	//是否跟随全局主题的变换而变换
	followTheme: {
		type: [Boolean, String],
		default: true
	},
	defaultValue: {
		type: Array as PropType<Array<any>>,
		default: () => []
	},
	//可以是双向绑定
	modelValue: {
		type: Array as PropType<Array<any>>,
		default: () => []
	},
	color: {
		type: String,
		default: 'primary'
	},
	disabled: {
		type: Boolean,
		default: false
	},
	/** 禁用删除功能 */
	disabledRemove:{
		type: Boolean,
		default: false
	},
	/** 当前的上传状态可v-model:uploading,仅单向向外输入状态，不接受双向绑定，通过这个状态来知当前是否所有文件上传完成还是正在运行中。 */
	uploading: {
		type: Boolean,
		default: false
	},
	config:{
		type:Object as PropType<{
			/**media表示只允许图片或者视频选择，file表示允许任意文件，但只支持h5,微信平台  */
			uploadType?:"media"|"file",
			maxCount?: number,//文件最大数量。
			extension?: string[],//文件选择的类型。
			type?: "all" | "image" | "video" | "file" | undefined,
			/**只对h5 */
			sourceType?: Array<'album' | 'camera'>,
			maxSize?: number,//每一个文件上传的最大尺寸，默认为10mb
			hostUrl?: string,//上传文件的服务器地址
			autoUpload?: boolean,
			header?: { [key: string]: any },//头部参数。
			formData?: { [key: string]: any },//额外的表单数据。
			formName?: string,
			code?: number,//服务器响应码，如果不为此码，表示上传失败。
			maxDuration?:number,//如果选择的类型是视频，可以定制此拍摄的最大时长。仅uploadType为media时有效
			sizeType?:Array<'original'|'compressed'>,//仅对 mediaType 为 image 时有效，是否压缩所选文件,仅uploadType为media时有效
			camera?:'back'|'front',//仅在 sourceType 为 camera 时生效，使用前置或后置摄像头,仅uploadType为media时有效
			mediaType?:Array<'image'|'video'>,//注意当uploadType=media时，如果是微信，抖音，飞书这里可以为当前正常类型值；如果是其它平台：只取数组中第一个值，比如要选择图片设置为["image"],视频：["video"],仅uploadType为media时有效
		}>,
		default:()=>{
			return {};
		}
	}

})

const emits = defineEmits(['update:modelValue','update:uploading','click'])
let timeId: any = NaN
const _disabled = computed(()=>props.disabled);
const _disabledRemove = computed(()=>props.disabledRemove);
addEventListener('change',()=>{
	emits('update:uploading',uploading.value)
})

function chooseFile(){
	if(_disabled.value) return;
	choose('file');
}
watchEffect(()=>{
	config.value = uni.$tm.u.deepObjectMerge(uni.$tm.u.deepClone(config.value),props.config);
})
onMounted(()=>{
	initDefault();
})


function initDefault(){
	addFile(props.modelValue);
	emits('update:modelValue',uni.$tm.u.deepClone(files.value))
}


</script>
