//倒计时
import { ref, computed, ComputedRef, watch } from 'vue';

interface TimerOptions {
	totalTime : number; // 总时间，单位毫秒
	unit ?: 'ms' | 'ss' | 'mm' | 'hh' | 'dd'; // 时间单位，默认为毫秒
	format ?: string; // 时间格式，例如 'yy年dd天 hh时mm分ss秒ms毫秒'
}

interface FormattedTime {
	year : number;
	day : number;
	hour : number;
	minute : number;
	second : number;
	millisecond : number;
}

interface TimerState {
	times : ComputedRef<number>;
	formattedTime : ComputedRef<(format ?: string) => string>;
	timeObj : ComputedRef<FormattedTime>; // 新增 timeObj 属性
	/**0暂停中，1，运行中，-1从未开始运行过 */
	status : ComputedRef<0 | 1 | -1>;
	start : () => void;
	stop : () => void;
	restart : () => void;
	beforeStart : (call : () => void) => void;
	timeEnd : (call : () => void) => void;
	change : (call : (n : number) => void) => void;
}
/**
 * 倒计时
 * @description 方便计时操作。
 */
export function useTimer(options : TimerOptions) : TimerState {
	const { totalTime, unit = 'ms', format = '' } = options;

	const times = ref(calculateRemainingTime());

	const status = ref<-1 | 0 | 1>(-1); // 'not_started', 'running', 'paused'

	let _beforeStart = () => { };
	let _timeEnd = () => { };
	let _change = (n : number) => { };

	let timerInterval = 50; // 定时器间隔，单位毫秒
	if (unit != 'ms') {
		timerInterval = 1000
	}
	let timer : NodeJS.Timeout | null = null;
	let animationFrameId : number | null = null; // 用于存储 requestAnimationFrame 的 ID

	const timeFomartValue = computed(() => {
		switch (unit) {
			case 'ms':
				return times.value;
			case 'ss':
				return Math.floor(times.value / 1000);
			case 'mm':
				return Math.floor(times.value / 60000);
			case 'hh':
				return Math.floor(times.value / 3600000);
			case 'dd':
				return Math.floor(times.value / 86400000);
			default:
				return times.value;
		}
	})
	// #ifndef H5
	const start = () => {
		if (status.value === -1 || status.value === 0) {
			status.value = 1;
			_beforeStart();
			timer = setInterval(() => {
				times.value -= timerInterval;
				if (times.value >= 0) {
					_change(timeFomartValue.value)
				}
				if (times.value <= 0) {
					times.value = 0;
					stop();
				}
			}, timerInterval);
		}
	};
	// #endif
	// #ifdef H5
	const start = () => {
		if (status.value === -1 || status.value === 0) {
			status.value = 1;
			_beforeStart();
			console.log(timeFomartValue.value)
			let lastTime = performance.now() - timerInterval;
			const update = (currentTime : number) => {
				const deltaTime = currentTime - lastTime;
				if (deltaTime >= timerInterval) {
					lastTime = currentTime;
					const p = times.value - timerInterval
					times.value = Number(p.toFixed(0))

					if (times.value >= 0) {
						_change(timeFomartValue.value)
					}
					if (times.value <= 0) {
						times.value = 0;
						stop();
					}
				}
				if (status.value === 1) {
					animationFrameId = requestAnimationFrame(update);
				}
			};

			animationFrameId = requestAnimationFrame(update);
		}
	};
	// #endif
	const stop = () => {
		if (timer !== null) {
			clearInterval(timer);
			timer = null;
		}
		if (animationFrameId !== null) {
			cancelAnimationFrame(animationFrameId); // 取消动画帧请求
			animationFrameId = null;
		}
		status.value = 0;
		_timeEnd();
	};

	const restart = () => {
		stop();
		times.value = calculateRemainingTime();
		start();
	};

	// 监听 totalTime 和 unit 变化，重新计算 times
	watch([() => options.totalTime, () => options.unit], () => {
		times.value = calculateRemainingTime();
	});

	function beforeStart(call : () => void) {
		_beforeStart = call;
	}
	function timeEnd(call : () => void) {
		_timeEnd = call;
	}
	function change(call : (n : number) => void) {
		_change = call;
	}

	function calculateRemainingTime() : number {
		switch (unit) {
			case 'ms':
				return totalTime;
			case 'ss':
				return totalTime * 1000;
			case 'mm':
				return totalTime * 60000;
			case 'hh':
				return totalTime * 3600000;
			case 'dd':
				return totalTime * 86400000;
			default:
				return totalTime;
		}
	};

	const formatTime = (time : number) : FormattedTime => {
		const years = Math.floor(time / (1000 * 60 * 60 * 24 * 365)); // 计算年份
		const days = Math.floor((time % (1000 * 60 * 60 * 24 * 365)) / (1000 * 60 * 60 * 24));
		const hours = Math.floor((time % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
		const minutes = Math.floor((time % (1000 * 60 * 60)) / (1000 * 60));
		const seconds = Math.floor((time % (1000 * 60)) / 1000);
		const milliseconds = time % 1000;
		return { year: years, day: days, hour: hours, minute: minutes, second: seconds, millisecond: milliseconds };
	};

	const formattedTime = computed(() => {
		let p = formatTime(times.value);
		return (temForamt ?: string) => {
			let fat = temForamt || format
			return fat
				.replace('yy', String(p.year))
				.replace('dd', String(p.day))
				.replace('hh', String(p.hour))
				.replace('mm', String(p.minute))
				.replace('ss', String(p.second))
				.replace('ms', String(p.millisecond))
		}
	});
	const timeObj = computed(() => formatTime(times.value)); // 计算 timeObj

	return {
		times: timeFomartValue,
		formattedTime,
		timeObj,
		status: computed(() => status.value),
		start,
		stop,
		restart,
		beforeStart,
		timeEnd,
		change
	};
}