import { showDetail, timeDetailType, timeArrayType } from "./interface"
import * as dayjs from "../../tool/dayjs/esm/index"
import isSameOrBefore from '../../tool/dayjs/esm/plugin/isSameOrBefore/index';
import isSameOrAfter from '../../tool/dayjs/esm/plugin/isSameOrAfter/index';
import isBetween from '../../tool/dayjs/esm/plugin/isBetween/index';
type dateType = 'year' | 'month' | 'date' | 'hour' | 'minute' | 'second'
const DayJs = dayjs.default;
DayJs.extend(isBetween)
DayJs.extend(isSameOrBefore)
DayJs.extend(isSameOrAfter)
function rangeNumber(from = 0, to = 0) {
	let range: Array<number> = []
	from = from >= 0 ? from : 1
	for (let i = from; i <= to; i++) {
		range.push(i)
	}
	return range
}

/** 根据显示的时间字段返回相关的时间列的可选选项. */
export function rangeTimeArray(dateStr: string | number | Date | dayjs.Dayjs, start: string | number | Date | dayjs.Dayjs, end: string | number | Date | dayjs.Dayjs, detail: showDetail) {
	let startDayjs = DayJs(start);
	let endDayjs = DayJs(end);
	let dateDayjs = DayJs(dateStr);

	// 计算每一列表数组开始和结束的数字.
	let dateArray = {
		year: [] as Array<number>,
		month: [] as Array<number>,
		date: [] as Array<number>,
		hour: [] as Array<number>,
		minute: [] as Array<number>,
		second: [] as Array<number>,
	};

	// 计算年份数组
	let startYear = startDayjs.year();
	let endYear = endDayjs.year();
	for (let year = startYear; year <= endYear; year++) {
		dateArray.year.push(year);
	}

	// 计算月份数组
	let startMonth = dateDayjs.isSame(startDayjs, 'year') ? startDayjs.month() : 0;
	let endMonth = dateDayjs.isSame(endDayjs, 'year') ? endDayjs.month() : 11;
	for (let month = startMonth; month <= endMonth; month++) {
		dateArray.month.push(month);
	}
	
	// 计算日期数组
	let startDate = dateDayjs.isSame(startDayjs, 'month') ? startDayjs.date() : 1;
	let endDate = dateDayjs.isSame(endDayjs, 'month') ? endDayjs.date() : dateDayjs.daysInMonth();
	for (let date = startDate; date <= endDate; date++) {
		dateArray.date.push(date);
	}

	// 计算小时数组
	let startHour = dateDayjs.isSame(startDayjs, 'date') ? startDayjs.hour() : 0;
	let endHour = dateDayjs.isSame(endDayjs, 'date') ? endDayjs.hour() : 23;
	for (let hour = startHour; hour <= endHour; hour++) {
		dateArray.hour.push(hour);
	}

	// 计算分钟数组
	let startMinute = dateDayjs.isSame(startDayjs, 'hour') ? startDayjs.minute() : 0;
	let endMinute = dateDayjs.isSame(endDayjs, 'hour') ? endDayjs.minute() : 59;
	for (let minute = startMinute; minute <= endMinute; minute++) {
		dateArray.minute.push(minute);
	}

	// 计算秒钟数组
	let startSecond = dateDayjs.isSame(startDayjs, 'minute') ? startDayjs.second() : 0;
	let endSecond = dateDayjs.isSame(endDayjs, 'minute') ? endDayjs.second() : 59;
	for (let second = startSecond; second <= endSecond; second++) {
		dateArray.second.push(second);
	}

	return dateArray;
}

/**
 * 根据现有时间取得当前的索引位置顺序
 * @param tmArray
 * @param nowtime
 * @param detail
 */
export function getIndexNowbydate(tmArray: timeArrayType, nowtime: dayjs.Dayjs, detail: showDetail) {
	const d = DayJs(nowtime)
	const intermediate = [
		[timeDetailType.year, detail.year],
		[timeDetailType.month, detail.month],
		[timeDetailType.day, detail.day],
		[timeDetailType.hour, detail.hour],
		[timeDetailType.minute, detail.minute],
		[timeDetailType.second, detail.second]
	];
	
	const order_str = ['year', 'month', 'date', 'hour', 'minute', 'second'];
	let order = []
	for (const key of order_str) {
		let keys = key == 'date'?'day':key
		if (detail[keys]) {
		  order.push(key);
		}
	}
	let idx = []
	for(let i=0;i<order.length;i++){
		let index = tmArray[order[i]].findIndex(n => n === d.get(order[i]))
		idx.push(index === -1 ? 0 : index)
	}
	
	// const idx = intermediate.filter(m => m[1]).map(m => {
	// 	const type = m[0] as timeDetailType;
	// 	let index = tmArray[type].findIndex(n => n === d.get(type))
	// 	return index === -1 ? 0 : index;
	// });

	return [
		...idx,
		...[0, 0, 0, 0, 0, 0]
	].slice(0, 6);
}
/**
 * 根据现有索引值返回当前时间。
 * @param tmArray
 * @param nowtime
 * @param detail
 */
export function getNowbyIndex(tmArray: timeArrayType, nowIndex: Array<number>, detail: showDetail, start:any,end:any) {
	let intermediate = [
		[timeDetailType.year, detail.year],
		[timeDetailType.month, detail.month],
		[timeDetailType.day, detail.day],
		[timeDetailType.hour, detail.hour],
		[timeDetailType.minute, detail.minute],
		[timeDetailType.second, detail.second]
	];
	const order_str = ['year', 'month', 'date', 'hour', 'minute', 'second'];
	let order = []
	
	for (const key of order_str) {
		let keys = key == 'date'?'day':key
		if (detail[keys]) {
		  order.push({
			  type:key,
			  index:nowIndex[order_str.indexOf(key)]
		  });
		}
	}
	
	
	for(let i=0;i<order.length;i++){
		order[i].index = nowIndex[i]
	}
	
	
	function getValue(type: timeDetailType) {
		
		const index = order.findIndex(m => m.type === type);
		if (index > -1) {
			return tmArray[type][order[index].index];
		}
		return tmArray[type][0];
	}
	
	// function getValue(type: timeDetailType) {
	// 	const index = intermediate.filter(m => m[1]).findIndex(m => m[0] === type);
	// 	if (index !== -1) {
	// 		return tmArray[type][nowIndex[index]];
	// 	}
	// 	return tmArray[type][tmArray[type].length - 1];
	// }
	let s = DayJs(start)
	let year = detail.year?getValue(timeDetailType.year):s.year();
	let month =  detail.month?getValue(timeDetailType.month):s.month();
	let date = detail.day?getValue(timeDetailType.day):s.date();
	let hour =  detail.hour?getValue(timeDetailType.hour):s.hour();
	let minute = detail.minute?getValue(timeDetailType.minute):s.minute();
	let second = detail.second?getValue(timeDetailType.second):s.second();

	let str = year
		+ "/" +
		(month + 1)
		+ "/" + date
		+ " " +
		hour
		+ ":" +
		minute
		+ ":" +
		second
	return DayJs(str).format("YYYY/MM/DD HH:mm:ss")
}
/**
 * 检查提供的时候是否是有效的时间段内，并返回正确的可用时间.
 * @param nowdate 
 * @param start 
 * @param end 
 */
export function checkNowDateisBetween(nowdate: string | number | Date | dayjs.Dayjs, start: string | number | Date | dayjs.Dayjs, end: string | number | Date | dayjs.Dayjs) {

	nowdate = DayJs(nowdate).isValid() ? DayJs(nowdate) : DayJs()
	start = DayJs(start)
	end = DayJs(end)
	if (nowdate.isSameOrBefore(start)) {
		return start.format("YYYY/MM/DD HH:mm:ss")
	}
	if (nowdate.isSameOrAfter(end)) {
		return end.format("YYYY/MM/DD HH:mm:ss")
	}
	return nowdate.format("YYYY/MM/DD HH:mm:ss")
}

