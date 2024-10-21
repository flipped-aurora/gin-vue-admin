<template>
	<view :class="[`mx-${props.margin[0]}`]">
		<tm-sheet :transprent="props.transprent !== null ? props.transprent : tmFormTransprent" :round="props.round"
			:margin="[0, 0]" :padding="props.padding">
			<view :class="[`py-${props.margin[1]}`]">
				<view v-if="props.showError && props.showTopErrorGap" :style="{ height: `${props.errHeight}rpx` }">
				</view>
				<view :class="[
            'flex',
            tmFormLayout == 'horizontal' ? 'flex-row ' : ' ',
            tmFormLayout == 'horizontal' && !props.align ? 'flex-row-center-start' : '',
			<!-- #ifndef APP-NVUE -->
            tmFormLayout == 'vertical' && !props.align ? 'flex-col' : '',
			<!-- #endif -->
            props.align,
            props.parentClass,
          ]">
					<slot name="label">
						<view v-if="_label"
							:style="[tmFormLayout == 'horizontal' ? { width: (props.labelWidth || tmFormLabelWidth) + 'rpx' } : '']"
							class="flex flex-row" :class="[
								tmFormLabelAlign == 'right' ? 'flex-row-center-end' : 'flex-row-center-start',
								tmFormLayout != 'horizontal' ? 'mb-24 flex-1 ' : 'mr-32 '
							]">
							<view v-if="tmFormLabelAlign != 'right'" style="width: 12px " class="flex flex-row flex-row-center-center">
								<tm-text v-if="_required" color="red" :font-size="30" label="*"></tm-text>
							</view>
							<view class="flex " :class="[
							tmFormLayout == 'horizontal'?'flex-1':'',
							tmFormLabelAlign == 'right'&&tmFormLayout == 'horizontal'?'text-align-right':'',
							tmFormLabelAlign == 'center'&&tmFormLayout == 'horizontal'?'text-align-center':'',
							]" :style="[tmFormLayout == 'horizontal'?{width:'0px'}:'']">
								<tm-text
									:color="tmFormFun == 'validate' && item.isRequiredError == true && props.requiredTitleChangeColor ? 'red' : ''"
									:font-size="30" :label="_label"></tm-text>
							</view>
							<view v-if="tmFormLabelAlign == 'right'" style="width: 12px " class="flex flex-row flex-row-center-center">
								<tm-text v-if="_required" color="red" :font-size="30" label="*"></tm-text>
							</view>
						</view>
					</slot>
					<view class="flex-1" :style="[tmFormLayout == 'horizontal' ? { width: '0px' } : '']">
						<slot></slot>
					</view>
				</view>
				<view :class="props.desc ? 'pt-12' : ''">
					<slot name="desc">
						<tm-text color="grey-darken-2" :font-size="22" :label="props.desc"></tm-text>
					</slot>
				</view>
				<view v-if="props.showError" :style="{ height: `${props.errHeight}rpx` }">
					<view v-if="tmFormFun == 'validate' && item.isRequiredError == true">
						<slot name="error" :data="{ message: item.message }">
							<tm-text color="red" :font-size="22" :label="item.message"></tm-text>
						</slot>
					</view>
				</view>
			</view>
		</tm-sheet>
		<view v-if="tmFormBorder">
			<tm-divider :border="2" :padding="[0, 0]" :margin="[0, 0]"></tm-divider>
		</view>
	</view>
</template>
<script lang="ts" setup>
	/**
 * 表单项目
 * @description 只能放置在tm-from里面，不限层级。但不能嵌套使用。
 * 从3.0.86开始,验证不与子组件挂勾,只验证绑定到form中的vmodel数据中的formitem中的filed字段.
 * 这样的改变意味着,传统的cell或者只用展示的view组件,只要提供了filed就能验证,无需特定Input,upload这样的
 * 组件才能进行验证.大大的提高了移动端的表单验证效率,特别适合大量的不同类型数据组件的验证规则.
 */
	import {
		computed,
		watch,
		PropType,
		provide,
		ref,
		getCurrentInstance,
		onUnmounted,
		Ref,
		inject,
		isProxy,
		toRaw,
		ComputedRef,
		onMounted,
		nextTick
	} from 'vue'
	import tmSheet from '../tm-sheet/tm-sheet.vue'
	import tmText from '../tm-text/tm-text.vue'
	import tmDivider from '../tm-divider/tm-divider.vue'
	import { rulesItem, inputPushItem } from './interface'
	import { formItem, validateResultListType } from './../tm-form/interface'
	const proxy = getCurrentInstance()?.proxy ?? null
	const tmFormComnameFormItem = 'tmFormComnameFormItem'
	const props = defineProps({
		parentClass: {
			type: String,
			default: ''
		},
		align: {
			type: String,
			default: ''
		},
		label: {
			type: String,
			default: ''
		},
		//表单描述
		desc: {
			type: String,
			default: ''
		},
		margin: {
			type: Array as PropType<Array<number>>,
			default: () => [12, 12]
		},
		padding: {
			type: Array as PropType<Array<number>>,
			default: () => [0, 0]
		},
		//如果在forom绑定的model为深层对象，这里的名称需要如下:
		//比如model = {a:2,b:{c:333}}
		//如果想绑定c,则field = "b.c"
		field: {
			type: String,
			default: ''
		},
		//表彰底部的单项注意说明。
		help: {
			type: String,
			default: ''
		},
		//是否必填
		required: {
			type: Boolean,
			default: false
		},
		//检验规则
		rules: {
			type: [Object, Array] as PropType<Array<rulesItem> | rulesItem>,
			default: () => {
				return [{ validator: false, required: false }]
			}
		},
		//显示下划线。
		border: {
			type: Boolean,
			default: null
		},
		showError: {
			type: Boolean,
			default: true
		},
		/**当显示错误信息标题时，是否隐藏顶部的间隙，当连续的布局时有用，可以减少之间的间隙大小。 */
		showTopErrorGap: {
			type: Boolean,
			default: true
		},
		//校验不通过时，是否让标题跟着变化文字颜色，默认是。
		requiredTitleChangeColor: {
			type: Boolean,
			default: false
		},
		transprent: {
			type: [Boolean, String],
			default: null
		},
		round: {
			type: Number,
			default: 0
		},
		errHeight: {
			type: Number,
			default: 30
		},
		labelWidth: {
			type: Number,
			default: 0
		}
	})
	const item = ref<formItem>({
		label: '', //标签名称。
		field: props.field, //字段名称key.
		value: null,
		isRequiredError: false, //true,错误，false正常 检验状态
		message: '', //检验信息提示语。
		id: uni.$tm.u.getUid(1), //表单唯一标识id
		componentsName: '', //表单组件类型。
		rules: []
	})
	const _required = ref(props.required)
	const tmFormLabelWidth = inject(
		'tmFormLabelWidth',
		computed(() => 100)
	)
	const tmFormLabelAlign = inject(
		'tmFormLabelAlign',
		computed(() => 'left')
	)
	const tmFormLayout = inject(
		'tmFormLayout',
		computed(() => 'horizontal')
	)
	const tmFormBorder_inject = inject(
		'tmFormBorder',
		computed(() => true)
	)
	const tmFormTransprent = inject(
		'tmFormTransprent',
		computed(() => false)
	)
	const tmFormFun = inject(
		'tmFormFun',
		computed(() => '')
	)
	const tmFormValidateResultList = inject(
		'validateResultList',
		computed<validateResultListType[]>(() => [])
	)

	const tmFormBorder = computed(() => {
		if (props.border !== null && typeof props.border === 'boolean') return props.border
		return tmFormBorder_inject.value
	})
	const _label = computed(() => props.label)
	//父级方法。
	let parent : any = proxy?.$parent
	while (parent) {
		if (parent?.tmFormComnameId == 'tmFormId' || !parent) {
			break
		} else {
			parent = parent?.$parent ?? undefined
		}
	}

	//卸载后需要清空。
	onUnmounted(() => {
		delCom()
	})

	// provide("tmFormItemVaildata",item.value.isRequiredError)
	const Rules = computed(() => {
		let defaultrs : Array<rulesItem> = []
		if (Array.isArray(props?.rules)) {
			props?.rules.forEach((el) => {
				let isreq = el?.required || props.required
				defaultrs.push({
					message: el?.message ?? '请填写必要的内容',
					required: isreq,
					validator: el?.validator ?? false
				})
			})
		} else {
			defaultrs = [
				{
					message: props?.rules?.message ?? '请填写必要的内容',
					required: props.rules?.required || props.required,
					validator: props.rules?.validator ?? false
				}
			]
		}
		return defaultrs
	})
	//向父级推表单类组件。
	function pushCom(itemComval ?: inputPushItem) {
		if (parent) {
			item.value = { ...item.value, ...(itemComval ?? {}), rules: Rules.value }
			parent.pushKey({ ...item.value })
		}
	}
	function delCom() {
		if (parent) {
			parent.delKey(item.value)
		}
	}

	defineExpose({ pushCom, delCom, tmFormComnameFormItem })

	const validate = (rules : Array<rulesItem>, value : any) => {
		rules = rules.map((el) => {
			if (typeof el.validator === 'function' && el.required === true) {
				return el
			} else if (typeof el.validator === 'boolean' && el.required === true) {
				return {
					...el,
					validator: (val : any) => {
						if (val === null || val === '' || typeof val == 'undefined') return false
						if (typeof val === 'object') {
							if (Array.isArray(val)) {
								if (val.length == 0) return false
							} else if (Object.keys(val).length === 0 && val.constructor === Object) {
								return false
							}
						}
						if (typeof val === 'boolean') {
							return val
						}
						if (typeof val === 'number') {
							if (isNaN(val)) return false
							if (Number(val) < 0) return false
						}
						if (typeof val === 'string') {
							if (val.trim().length == 0) return false
						}
						return true
					}
				}
			} else {
				return {
					...el,
					validator: (val : string | number) => {
						return true
					}
				}
			}
		})
		let rules_filter : Array<rulesItem> = rules.filter((el) => {
			return typeof el.validator === 'function' && el.required === true
		})
		let rules_fun : Array<Promise<rulesItem>> = rules_filter.map((el) => {
			return new Promise(async (res, rej) => {
				if (typeof el.validator === 'function') {
					let vr = await el.validator(value)
					if (vr) {
						res({
							message: String(el.message),
							validator: true
						})
					} else {
						rej({
							message: el.message,
							validator: false
						})
					}
				} else {
					res({
						message: el.message,
						validator: true
					})
				}
			})
		})

		return Promise.all(rules_fun)
	}

	onMounted(() => {
		//预先推送。
		pushCom()
	})

	watch(
		() => tmFormValidateResultList.value,
		() => {
			const result = tmFormValidateResultList.value.filter((el) => el.field == props.field)
			if (result.length > 0) {
				item.value.message = result[0].message
				item.value.isRequiredError = !result[0].validator
			}else if(tmFormValidateResultList.value.length==0){
				item.value.message=""
				item.value.isRequiredError = true
			}
		},
		{ deep: true, immediate: true }
	)

	function _recursive_parseJson2Form_(resObj : any, parentKey : string, JsonObject : any) {
		var gettype = Object.prototype.toString
		for (var Key in JsonObject) {
			var tmpVal = JsonObject[Key]

			var typeStr = gettype.call(tmpVal)
			/*
	
								"[object String]";
								"[object Number]";
								"[object Boolean]"
								"[object Undefined]"
								"[object Null]"
								"[object Object]"
								"[object Array]"
								"[object Function]"
								*/
			if (tmpVal == null || tmpVal == undefined) {
				resObj[parentKey + '.' + Key] = ''
			} else {
				if (
					typeStr == '[object String]' ||
					typeStr == '[object Number]' ||
					typeStr == '[object Boolean]' ||
					typeStr == '[object Null]' ||
					typeStr == '[object Undefined]'
				) {
					resObj[parentKey + '.' + Key] = tmpVal + ''
				} else if (typeStr == '[object Object]') {
					_recursive_parseJson2Form_(resObj, parentKey + '.' + Key, tmpVal)
				} else if (typeStr == '[object Array]') {
					if (tmpVal.length <= 0) {
						resObj[parentKey + '.' + Key] = ''
					} else {
						var isSimpleArray = false
						for (var arrIndex in tmpVal) {
							var childItem = tmpVal[arrIndex]
							var childTypeStr = gettype.call(childItem)

							if (
								childTypeStr == '[object String]' ||
								childTypeStr == '[object Number]' ||
								childTypeStr == '[object Boolean]' ||
								childTypeStr == '[object Null]' ||
								childTypeStr == '[object Undefined]'
							) {
								isSimpleArray = true
							}
						}
						if (isSimpleArray) {
							resObj[parentKey + '.' + Key] = tmpVal.join(',')
						} else {
							var nowIndex = 0
							for (var arrIndex in tmpVal) {
								var childItem = tmpVal[arrIndex]
								_recursive_parseJson2Form_(resObj, parentKey + '.' + Key + '[' + nowIndex + ']', childItem)
								nowIndex++
							}
						}
					}
				}
			}
		}
	}
</script>

<style></style>