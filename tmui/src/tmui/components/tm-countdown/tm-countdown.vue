<template>
	<view class="flex flex-row">
		<slot name="default" :data="{ data: time_data, finish: isfinish }">
			<tm-text :color="props.color" :label="text"></tm-text>
		</slot>
	</view>
</template>
<script lang="ts" setup>
import tmText from '../tm-text/tm-text.vue'
import { computed, onMounted, PropType, ref, watch } from 'vue'

const emits = defineEmits(['start', 'end', 'change'])
const props = defineProps({
	time: {
		type: Number,
		default: 10 * 1000
	},
	/**
	 * 取值的格式字符串
	 * 注意如果当你的formatType设定某值时，里面只能读取到你设定的值。
	 * */
	format: {
		type: String,
		default: 'DD天HH小时MM分SS秒MS毫秒'
	},
	/**
	 * 到计时格式的类型，设定下面的值时，倒计时功能不会进位，而是以指定的值进行倒计时。
	 * 比如分，你设置为MM,那么你到计时如果是200分钟，就从200开始倒计时。而不会进位到小时。
	 * "DD"|"HH"|"MM"|"SS"|"MS"|""
	 * 天|时|分|秒|毫秒
	 * */
	formatType:{
		type:String as PropType<"DD"|"HH"|"MM"|"SS"|"MS"|"">,
		default:""
	},
	autoStart: {
		type: Boolean,
		default: true
	},
	color: {
		type: String,
		default: ''
	}
})
let timid: any = undefined
let now = ref(0)
interface timeFormar {
	day: number | string
	hour: number | string
	minutes: number | string
	seconds: number | string
	millisecond: number | string,
	[key:string]:any
}
let time_data = computed((): timeFormar => formatTime(props.time - now.value))

const isfinish = ref(true)

watch(
	() => now.value,
	() => {
		if (now.value == props.time || now.value === 0) {
			isfinish.value = true
		}
	}
)
const text = computed<string>(() => {
	let ps = props.format
	
	if(!props.formatType){
		ps = ps.replace(/(DD)/g, String(time_data.value.day))
		ps = ps.replace(/(MM)/g, String(time_data.value.minutes))
		ps = ps.replace(/(HH)/g, String(time_data.value.hour))
		ps = ps.replace(/(SS)/g, String(time_data.value.seconds))
		ps = ps.replace(/(MS)/g, String(time_data.value.millisecond))
	}else{
		if(props.formatType=="DD"){
			ps = ps.replace(/(DD)/g, String(time_data.value.DD))
		}
		if(props.formatType=="HH"){
			ps = ps.replace(/(HH)/g, String(time_data.value.HH))
		}
		if(props.formatType=="MM"){
			ps = ps.replace(/(MM)/g, String(time_data.value.MM))
		}
		if(props.formatType=="SS"){
			ps = ps.replace(/(SS)/g, String(time_data.value.SS))
		}
		if(props.formatType=="MS"){
			ps = ps.replace(/(MS)/g, String(time_data.value.MS))
		}
	}
	return ps
})
onMounted(() => {
	formatTime(props.time)
	if (props.autoStart) {
		start()
	}
})
function formatTime(my_time: number): timeFormar {
	var daysRound = Math.floor(my_time / 1000 / 60 / 60 / 24)
	var hoursRound = Math.floor((my_time / 1000 / 60 / 60) % 24)
	var minutesRound = Math.floor((my_time / 1000 / 60) % 60)
	var secondsRound = Math.floor((my_time / 1000) % 60)
	var millisecondRound = Math.floor(my_time % 1000)
	let time: timeFormar = {
		day: daysRound > 9 ? daysRound : '0' + daysRound, //天
		hour: hoursRound > 9 ? hoursRound : '0' + hoursRound, //小时,
		minutes: minutesRound > 9 ? minutesRound : '0' + minutesRound, //分.
		seconds: secondsRound > 9 ? secondsRound : '0' + secondsRound, //秒。
		millisecond: millisecondRound > 9 ? millisecondRound : '00' + millisecondRound, //毫秒。
		DD:Math.floor(my_time / 1000 / 60 / 60 / 24),
		HH:Math.floor(my_time / 1000 / 60 / 60),
		MM:Math.floor(my_time / 1000 / 60),
		SS:Math.floor(my_time / 1000),
		MS:my_time,
	}
	return time
}
// 开始或者继续。
function start() {
	clearInterval(timid)
	emits('start')
	timid = setInterval(() => {
		let lst = now.value + 50
		if (lst > props.time) {
			clearInterval(timid)
			isfinish.value = true
			emits('end')
			return
		}
		isfinish.value = false
		now.value = lst
		emits('change', time_data)
	}, 50)
}
// 停止，直接结束。
function stop() {
	clearInterval(timid)
	now.value = props.time
	emits('end')
}
// 暂停。
function pause() {
	clearInterval(timid)
}
// 重置。
function resinit() {
	clearInterval(timid)
	now.value = 0
	isfinish.value = true
}
//对外暴露的方法和属性。
defineExpose({
	finish: isfinish,
	start: start,
	stop: stop,
	pause: pause,
	resinit: resinit
})
</script>
