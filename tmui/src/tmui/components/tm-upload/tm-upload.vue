<template>
	<view class="flex flex-col flex-col-top-start">
		<view class="flex flex-row flex-row-top-start relative" style="flex-wrap: wrap" :style="{ width: width + 'rpx' }">
			<view class="ma-5" v-for="(item, index) in _flist" :key="index" :style="{ width: itemWidth - 10 + 'rpx' }">
				<tm-sheet :round="2" :color="props.color" text :transprent="true" :padding="[0, 0]" :margin="[0, 0]" class=" ">
					<tm-image
						preview
						:round="2"
						:allowDelete="false"
						@delete="deletedFile(item)"
						extra
						:delete="!_disabled"
						:src="item.url"
						:width="itemWidth - 10"
						extraPosition="in"
						:height="itemHeight - 10"
						:model="props.imageModel"
					>
						<template v-slot:extra class="relative flex-1 flex flex-col flex-col-bottom-start">
							<view
								v-if="!_disabled && props.showStatus"
								:style="{
									background: 'rgba(0, 0, 0, 0.5)',
									width: itemWidth - 10 + 'rpx'
								}"
								:class="[`round-b-${2}`]"
								class="py-4 px-4 flex flex-row flex-row-center-start"
							>
								<tm-icon
									:font-size="23"
									v-if="item.statusCode == 0 || item.statusCode == 1"
									color="grey-3"
									name="tmicon-clock-fill"
								></tm-icon>
								<tm-text
									v-if="item.statusCode == 0 || item.statusCode == 1"
									color="grey-3"
									_class="pl-5"
									:font-size="23"
									:label="item.status"
								></tm-text>
								<tm-icon
									:font-size="23"
									v-if="item.statusCode == 2 || item.statusCode == 4"
									color="red"
									name="tmicon-times-circle-fill"
								></tm-icon>
								<tm-text
									v-if="item.statusCode == 2 || item.statusCode == 4"
									color="red"
									_class="pl-5"
									:font-size="23"
									:label="item.status"
								></tm-text>
								<tm-icon :font-size="23" v-if="item.statusCode == 3" color="green" name="tmicon-check-circle-fill"></tm-icon>
								<tm-text v-if="item.statusCode == 3" color="green" _class="pl-5" :font-size="23" :label="item.status"></tm-text>
							</view>

							<view
								v-if="props.showSort"
								class="absolute l-0 t-0 flex flex-row flex-row-center-between flex-1"
								:style="{
									width: itemWidth - 10 + 'rpx',
									height: 44 + 'rpx',
									top: (itemHeight - 44) / 2 + 'rpx'
								}"
							>
								<view
									@click.stop="prevSort(item, index, 'prev')"
									:class="[index == 0 ? 'opacity-0' : '']"
									class="flex flex-row flex-row-center-center px-12 py-4 round-tr-12 round-br-12"
									:style="{ background: 'rgba(0, 0, 0, 0.5)' }"
								>
									<tm-icon :userInteractionEnabled="false" color="white" :font-size="22" name="tmicon-angle-left"></tm-icon>
								</view>
								<view
									@click.stop="prevSort(item, index, 'next')"
									:class="[index == _flist.length - 1 ? 'opacity-0' : '']"
									class="flex flex-row flex-row-center-center px-12 py-4 round-tl-12 round-bl-12"
									:style="{ background: 'rgba(0, 0, 0, 0.5)' }"
								>
									<tm-icon :userInteractionEnabled="false" color="white" :font-size="22" name="tmicon-angle-right"></tm-icon>
								</view>
							</view>
						</template>
					</tm-image>
				</tm-sheet>
			</view>
			<view @click="chooseFile" v-if="!_disabledAdd" class="ma-5" :style="{ width: itemWidth - 10 + 'rpx' }">
				<tm-sheet
					:eventPenetrationEnabled="true"
					:followTheme="props.followTheme"
					:round="2"
					:color="props.color"
					text
					:padding="[0, 0]"
					:margin="[0, 0]"
					_class="flex-center"
					:height="itemHeight - 10"
				>
					<slot name="icon">
						<tm-icon :font-size="42" :userInteractionEnabled="false" name="tmicon-plus"></tm-icon>
					</slot>
				</tm-sheet>
			</view>
		</view>
	</view>
</template>
<script lang="ts" setup>
/**
 * 上传
 * @description 图片上传组件。
 */
import { computed, ref, PropType, Ref, watch, toRaw, nextTick, getCurrentInstance, inject, reactive, onMounted, isRef, isProxy } from 'vue'
import { inputPushItem, rulesItem } from './../tm-form-item/interface'
import { file, fileConfig, statusCode, uploadfile } from './upload'
import tmImage from '../tm-image/tm-image.vue'
import tmText from '../tm-text/tm-text.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import tmSheet from '../tm-sheet/tm-sheet.vue'
const proxy = getCurrentInstance()?.proxy ?? null
const props = defineProps({
	//是否跟随全局主题的变换而变换
	followTheme: {
		type: [Boolean, String],
		default: true
	},
	width: {
		type: Number,
		default: 700
	},
	//一行排列几个。
	rows: {
		type: Number,
		default: 5
	},
	//图片的高度
	imageHeight: {
		type: Number,
		default: 140
	},
	imageModel: {
		type: String,
		default: 'scaleToFill'
	},
	defaultValue: {
		type: Array as PropType<Array<file>>,
		default: () => []
	},
	//可以是双向绑定
	modelValue: {
		type: Array as PropType<Array<file>>,
		default: () => []
	},
	color: {
		type: String,
		default: 'primary'
	},
	header: {
		type: Object,
		default: () => {}
	},
	formData: {
		type: Object,
		default: () => {}
	},
	maxFile: {
		type: Number,
		default: 9
	},
	maxSize: {
		type: Number,
		default: 10 * 1024 * 1024
	},
	url: {
		type: String,
		default: '',
		required: true
	},
	formName: {
		type: String,
		default: 'file'
	},
	autoUpload: {
		type: Boolean,
		default: true
	},
	disabled: {
		type: Boolean,
		default: false
	},
	//删除前执行，如果返回false,将阻止删除。
	onRemove: {
		type: [Function, Boolean],
		default: () => {
			return (item: file) => true
		}
	},
	//开始上传前执行，如果返false，将阻止上传，
	onStart: {
		type: [Function, Boolean],
		default: () => {
			return (item: file) => true
		}
	},
	//上传成功后，从服务器响应后立即执行，此时，还未更改当前文件上传的状态，是成功还是失败
	//如果此时返回false,将会让文件状态从成功改为上传失败，尽管 从服务器正确返回，但仍然显示失败。
	onSuccessAfter: {
		type: [Function, Boolean],
		default: () => {
			return (item: file) => true
		}
	},
	//选择文件前执行，如果此时返回false,将阻止选择文件。你可以在这里做一些上传前的配置。
	beforeChooese: {
		type: [Function, Boolean],
		default: () => {
			return (item: file) => true
		}
	},
	chooesefileAfter: {
		type: [Function, Boolean],
		default: () => {
			return (item: file[]) => item
		}
	},
	//是否显示排序功能
	showSort: {
		type: Boolean,
		default: false
	},
	/**文件选择的类型
	 * album,相册，camera：相机
	 */
	fileType: {
		type: Array as PropType<Array<'album' | 'camera'>>,
		default: ['album', 'camera']
	},
	/**服务器返回的状态码，默认200成功 */
	statusCode: {
		type: Number,
		default: 200
	},
	showStatus: {
		type: Boolean,
		default: true
	}
})
/**
 * emits 事件说明
 * @method succcess 一个文件上传成功后立即执行， 返回：file,fileList
 * @method fail 一个文件上传失败后立即执行， 返回：file,fileList
 * @method complete 所有文件上传完成， 返回：file,fileList
 * @method change 文件列表增加或者减少,文件状态的改变， 返回：fileList
 * @method remove 一个文件被移除时触发。file,fileList
 * @method uploadComplete i当前任务所有文件上传结束时触发。fileList
 */
const emits = defineEmits(['success', 'fail', 'complete', 'change', 'remove', 'uploadComplete', 'update:modelValue'])
let timeId: any = NaN
const itemWidth = computed(() => {
	return props.width / props.rows
})
const itemHeight = computed(() => {
	return props.imageHeight
})

const _uploadObj = new uploadfile({
	formName: props.formName,
	hostUrl: props.url,
	autoUpload: props.autoUpload,
	fileList: addSuccess(props.defaultValue),
	header: props.header,
	formData: props.formData,
	maxFile: props.maxFile,
	maxSize: props.maxSize,
	fileType: props.fileType,
	statusCode: props.statusCode
})

const _flist: Ref<Array<file>> = ref([])
const _disabled = computed(() => props.disabled)
const _disabledAdd = computed(() => {
	return props.disabled || _flist.value.length >= props.maxFile
})

// emits('update:modelValue', addSuccess(props.defaultValue))
_flist.value = [..._uploadObj.filelist]

const _blackValue = uni.$tm.u.deepClone(toRaw(_uploadObj.filelist))

watch(
	[() => props.header, () => props.maxFile, () => props.maxSize, () => props.formData],
	() => {
		_uploadObj.setConfig({
			formName: props.formName,
			hostUrl: props.url,
			header: props.header,
			formData: props.formData,
			maxFile: props.maxFile,
			maxSize: props.maxSize
		})
	},
	{ deep: true }
)

function prevSort(item: file, index: number, type: 'prev' | 'next') {
	if ((index == 0 && type == 'prev') || (index == _flist.value.length - 1 && type == 'next')) {
		return
	}
	let nowindex = type == 'prev' ? index - 1 : index + 1
	let nowItem: file = toRaw(_flist.value[index])
	let newnowItem: file = toRaw(_flist.value[nowindex])
	let nowfilelist: Array<file> = uni.$tm.u.deepClone(toRaw(_uploadObj.filelist))
	nowfilelist.splice(index, 1, newnowItem)
	nowfilelist.splice(nowindex, 1, nowItem)
	_uploadObj.filelist = [...nowfilelist]
	_flist.value = [..._uploadObj.filelist]
	emits('update:modelValue', nowfilelist)
}

//添加已上传文件列表。
function addSuccess(fileList: Array<file> = []) {
	let fl = fileList.map((e) => {
		let _itemfile: file = { url: '' }
		if (typeof e == 'string') {
			_itemfile.url = e
		} else {
			_itemfile = { ...e }
		}
		_itemfile = {
			statusCode: e?.statusCode ?? statusCode.success,
			status: e?.status ?? '上传成功',
			progress: e?.progress ?? 100,
			..._itemfile
		}
		return _itemfile
	})
	return fl
}

//选择图片前执行。
_uploadObj.beforeChooesefile = async function () {
	_uploadObj.setConfig({ maxFile: props.maxFile - _uploadObj.filelist.length })
	if (typeof props.beforeChooese === 'function') {
		let p = await props.beforeChooese()
		if (typeof p === 'function') {
			p = await p()
		}
		if (!p) return false
	}
	return true
}

_uploadObj.chooesefileAfter = async function (filelist: file[]) {
	let tep = uni.$tm.u.deepClone(filelist)
	if (typeof props.chooesefileAfter === 'function') {
		let p = await props.chooesefileAfter(tep)
		if (typeof p === 'function') {
			p = await p(tep)
		}
		return p
	}
	return tep
}

//上传成功后，即将更新文件状态前执行，如果返回false，文件将标记为失败，尽管已经上传成功。
_uploadObj.beforeSuccess = async function (item: file) {
	if (typeof props.onSuccessAfter === 'function') {
		let p = await props.onSuccessAfter(item)
		if (typeof p === 'function') {
			p = await p(item)
		}
		if (!p) return false
	}

	return true
}
//开始上传前执行。
_uploadObj.beforeStart = async function (item: file) {
	if (typeof props.onStart === 'function') {
		let p = await props.onStart(item)
		if (typeof p === 'function') {
			p = await p(item)
		}
		if (!p) return false
	}
	return true
}
//任何一个文件上传结束时都会触发。
_uploadObj.complete = function (item: file) {
	_flist.value = [..._uploadObj.filelist]
}
//自动监听加入已上传文件到列表中。
watch(
	() => props.modelValue,
	() => {
		clearTimeout(timeId)

		let pl = isProxy(props.modelValue) ? toRaw(props.modelValue) : props.modelValue

		timeId = setTimeout(function () {
			let fl = Array.isArray(pl) ? pl : []
			_uploadObj.clear()
			if (fl.length == 0) {
				_uploadObj.filelist = []
				_flist.value = []
			} else {
				let nf = addSuccess(fl)
				_uploadObj.addFile(nf)
				let nsf: Array<file> = uni.$tm.u.deepClone(_uploadObj.filelist)
				_flist.value = [...nsf]
			}
		}, 100)
	},
	{ deep: true }
)
_uploadObj.uploadComplete = function (filelist) {
	emits('uploadComplete', filelist)
	emits('update:modelValue', _uploadObj.filelist)
}
_uploadObj.success = function (item, fileList) {
	let files = uni.$tm.u.deepClone(_uploadObj.filelist)
	emits('success', toRaw(item), files)
	emits('change', files)
}
_uploadObj.fail = function (item) {
	let files = uni.$tm.u.deepClone(_uploadObj.filelist)
	emits('fail', toRaw(item), files)
	emits('change', files)
}
function chooseFile() {
	if (_disabled.value) return
	_uploadObj.chooesefile().then((fileList) => {
		_flist.value = [..._uploadObj.filelist]
		emits('update:modelValue', _uploadObj.filelist)
	})
}
_uploadObj.progress = function (item,index) {
	_flist.value[index].status = item.progress+'% '+item.status;
	_flist.value[index].statusCode = item.statusCode
}
async function deletedFile(item: file) {
	if(item.statusCode==1){
		uni.showToast({
			title:'上传中不允许删除',
			icon:'none',
			mask:true
		})
		return;
	}
	if (typeof props.onRemove === 'function') {
		let p = await props.onRemove(item)
		if (typeof p === 'function') {
			p = await p(item)
		}
		if (!p) return false
	}
	_uploadObj.delete(item)
	_flist.value = [..._uploadObj.filelist]
	emits('remove', toRaw(item))
	emits('update:modelValue', _uploadObj.filelist)
	emits('change', toRaw(_uploadObj.filelist))
}

function clear() {
	_uploadObj.clear()
	_uploadObj.filelist = []
	_flist.value = []
	emits('update:modelValue', [])
}
//fileId标识id.
function del(fileId: number | string) {
	let index = _uploadObj.filelist.findIndex((el) => el.uid == fileId)
	if (index > -1) {
		const item = _uploadObj.filelist[index]
		if(item.statusCode==1){
			uni.showToast({
				title:'上传中不允许删除',
				icon:'none',
				mask:true
			})
			return;
		}
		_uploadObj.delete(item)
		_flist.value = _flist.value.splice(index, 1)
		emits('remove', toRaw(item))
		emits('update:modelValue', _uploadObj.filelist)
		emits('change', toRaw(_uploadObj.filelist))
	}
}
function getFailList() {
	return _uploadObj.filelist.filter((el) => el.statusCode != statusCode.fail && el.statusCode != statusCode.max)
}
function getSuccessList() {
	return _uploadObj.filelist.filter((el) => el.statusCode != statusCode.success)
}
function clearFail() {
	const list = _uploadObj.filelist.filter((el) => el.statusCode != statusCode.fail && el.statusCode != statusCode.max)
	_flist.value = [...list]
	emits('update:modelValue', _uploadObj.filelist)
}
/**
 * ref 手动调用的方法
 * start 开始上传
 * stop 停止上传
 * clear 清除所有文件
 * del 删除某一项文件。
 * getFailList 获取上传失败的文件列表。
 * getSuccessList 获取上传成功的文件列表
 * clearFail 清除上传失败的文件（只要标识不是成功的都会删除）
 * 手动
 */
defineExpose({
	start: () => {
		return _uploadObj.start()
	},
	stop: () => {
		return _uploadObj.stop()
	},
	clear,
	del,
	getFailList,
	clearFail
})
</script>
