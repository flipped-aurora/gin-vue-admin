<template>
	<tm-sheet :transprent="props.transprent" :round="3" _class="flex flex-col overflow" :padding="props.padding"
		:margin="props.margin">
		<slot></slot>
	</tm-sheet>
</template>

<script lang="ts" setup>
	/** 
 * 表单
 * @description 注意，内部需要放置tm-form-item,不限层级，可随意布局。
 * 对以下表单类组件进行字段收集。
 * 	"tm-radio-group","tm-checkbox-box",
	"tm-input","tm-rate","tm-slider",
	"tm-segtab","tm-switch","tm-upload"
 */
	import { computed, PropType, provide, ref, watch, Ref, toRaw, shallowReadonly, nextTick, isProxy, unref, readonly, watchEffect } from 'vue'
	import { formItem, validateResultListType } from './interface'
	import tmSheet from '../tm-sheet/tm-sheet.vue'
	import { validateFunCall, getObjectVal } from '../tm-form-item/validateFunCall'
	/**
	 * 事件说明
	 * @method submit 提交表单时触发。
	 * @method reset 重置表单时触发
	 * @method validate 校验表单时触发
	 * @method clearValidate 清除校验状态时触发。
	 */
	const emits = defineEmits<{
		(e : 'submit', event : Tmui.tmFormSubmitResult) : void
		/**只要有数据更改就会触发校验并返回相关数据与submit相同 */
		(e : 'validate', event : Tmui.tmFormRules) : void
		(e : 'reset') : void
		(e : 'clearValidate') : void
		(e : 'update:modelValue', event : any) : void
	}>()

	const props = defineProps({
		modelValue: {
			type: Object as PropType<Object>,
			default: () => {
				return {}
			},
			required: true
		},
		margin: {
			type: Array as PropType<Array<number>>,
			default: () => [32, 24]
		},
		padding: {
			type: Array as PropType<Array<number>>,
			default: () => [16, 0]
		},
		//表单标签是竖还是横排列。
		//vertical,horizontal
		layout: {
			type: String as PropType<'vertical' | 'horizontal'>,
			default: 'horizontal'
		},
		//如果为0表示自动宽度。
		labelWidth: {
			type: Number,
			default: 160
		},
		//标签对齐方式
		labelAlign: {
			type: String,
			default: 'left'
		},
		//显示下划线。
		border: {
			type: Boolean,
			default: true
		},
		transprent: {
			type: Boolean,
			default: false
		}
	})
	const _modelVal = ref(uni.$tm.u.deepClone(props.modelValue))
	//备份，重置时，使用。
	const _backModelVal = uni.$tm.u.deepClone(props.modelValue)
	watchEffect(() => {
		_modelVal.value = uni.$tm.u.deepClone(props.modelValue)
	})
	//收集的字段。状态。它与_modelVal是有区别的，用户提供的字段，不一定就会在页面中存在，需要与已经渲染的字段进行匹配
	const _callBackModelVal : Ref<Array<Tmui.tmFormRules>> = ref([])
	const tmFormComnameId = 'tmFormId'
	//允许被推送的组件清单类型.其它的组件不会被收集进行检验。
	const safeFormCom = ref(['tm-radio-group', 'tm-checkbox-box', 'tm-input', 'tm-rate', 'tm-slider', 'tm-segtab', 'tm-switch', 'tm-upload'])
	//需要对子级，响应的方法。
	// 这里为了更好的性能不再使用vue2版本中children方式，而是采用了provide方式与父子间传递。
	const formFunCallBack = ref('validate')
	const validateResultList = ref<validateResultListType[]>([])
	provide(
		'tmFormFun',
		computed(() => formFunCallBack.value)
	)
	provide(
		'tmFormLabelWidth',
		computed(() => props.labelWidth)
	)
	provide(
		'tmFormLabelAlign',
		computed(() => props.labelAlign)
	)
	provide(
		'tmFormLayout',
		computed(() => props.layout)
	)
	provide(
		'tmFormBorder',
		computed(() => props.border)
	)
	provide(
		'tmFormTransprent',
		computed(() => props.transprent)
	)
	provide(
		'formCallFiled',
		computed(() => _modelVal.value)
	)
	provide(
		'validateResultList',
		computed(() => validateResultList.value)
	)
	let timid : any = NaN
	let ptimeId : any = NaN
	watch(
		() => _modelVal.value,
		(newValue, oldValue) => {
			clearTimeout(timid)
			if (formFunCallBack.value == 'validate') {
				timid = setTimeout(function () {
					let vaildFileds = getChangedField(newValue,oldValue)
					const result = validate(vaildFileds,'')
					validateResultList.value = [...result.result]
				}, 100)
			}
		},
		{ deep: true }
	)

	function getChangedField(oldValue : any, newValue : any) : string[] {
		const changedFields : string[] = [];
		for (const key in oldValue) {
			if (oldValue.hasOwnProperty(key)) {
				if (typeof oldValue[key] === 'object' && typeof newValue[key] === 'object' && !Array.isArray(oldValue[key])) {
					const nestedChangedFields = getChangedField(oldValue[key], newValue[key]);
					if (nestedChangedFields.length > 0) {
						changedFields.push(`${key}.${nestedChangedFields.join('.')}`);
					}
				} else if (valToval(oldValue[key]) !== valToval(newValue[key])) {
					changedFields.push(key);
				}
			}
		}
		return changedFields;
	}
	function valToval(value:any){
		if(Array.isArray(value)){
			return value.join("")
		}
		return value
	}
	function reset() {
		formFunCallBack.value = 'reset'
		let dblack = uni.$tm.u.deepClone(_backModelVal)
		emits('update:modelValue', dblack)
		emits('reset')
		_modelVal.value = dblack
		
		validateResultList.value = []
		formFunCallBack.value = 'validate'
	}
	function clearValidate() {
		formFunCallBack.value = 'clearValidate'
		nextTick(() => {
			emits('clearValidate')
		})
		validateResultList.value = []
	}
	function submit() {
		formFunCallBack.value = 'validate'
		let isPass = true
		uni.$tm.u.throttle(
			() => {
				const result = validate()
				validateResultList.value = [...result.result]
				// @ts-ignore
				emits('submit', { data: toRaw(_modelVal.value), ...result })
			},
			220,
			false
		)
	}
	/**
	 * 执行表单检验
	 * type 为all时会省略fileds字段，校验全部
	 * type非all时，会校验fileds字段
	 */
	function validate(fileds:string[] = [],type = 'all') {
		formFunCallBack.value = 'validate'
		let par = toRaw(_callBackModelVal.value)
		let isPass = true
		let list : validateResultListType[] = []
		for (let i = 0, len = par.length; i < len; i++) {
			let item = par[i]
			if(type=='all'){
				let value = getObjectVal(_modelVal.value, item.field)
				const vallist = validateFunCall(item.rules, value)
				let rulstVal = {
					message: '校验通过',
					validator: true as Function | boolean
				}
				for (let j = 0; j < vallist.length; j++) {
					if (!vallist[j].validator) {
						isPass = false
						rulstVal.message = vallist[j]?.message ?? '校验通过'
						rulstVal.validator = vallist[j]?.validator ?? true
						break
					}
				}
				list.push({ field: item.field, ...rulstVal })
			}else if(fileds.length>0&&fileds.includes(item.field)){
				let value = getObjectVal(_modelVal.value, item.field)
				const vallist = validateFunCall(item.rules, value)
				let rulstVal = {
					message: '校验通过',
					validator: true as Function | boolean
				}
				for (let j = 0; j < vallist.length; j++) {
					if (!vallist[j].validator) {
						isPass = false
						rulstVal.message = vallist[j]?.message ?? '校验通过'
						rulstVal.validator = vallist[j]?.validator ?? true
						break
					}
				}
				list.push({ field: item.field, ...rulstVal })
			}
			
		}

		return { result: list, isPass, data: toRaw(_modelVal.value), validate: isPass }
	}
	

	function pushKey(item : formItem) {
		if (!item.field) return
		let idsIndex = _callBackModelVal.value.findIndex((el) => el.id == item.id)
		if (idsIndex == -1) {
			_callBackModelVal.value.push(item)
		} else {
			_callBackModelVal.value[idsIndex] = { ...item }
		}
	}
	function delKey(item : formItem) {
		let idsIndex = _callBackModelVal.value.findIndex((el) => el.id == item.id)
		if (idsIndex > -1) {
			_callBackModelVal.value.splice(idsIndex, 1)
		}
	}
	function setObjectVal(obj : any, field = '', val : any) {
		if (field == '') return obj
		var arr = field.split('.')
		while (arr.length > 1) {
			let key = String(arr.shift())
			obj = isProxy(obj[key]) ? toRaw(obj[key]) : obj[key]
		}
		return (obj[arr[0]] = isProxy(val) ? toRaw(val) : val)
	}
	/**
	 * ref函数
	 * @method submit 提交表单
	 * @method reset 重置表单
	 * @method validate 手动校验表单
	 * @method clearValidate 清除校验状态
	 */
	defineExpose({
		reset,
		validate,
		clearValidate,
		submit,
		pushKey,
		delKey,
		tmFormComnameId
	})
</script>

<style></style>