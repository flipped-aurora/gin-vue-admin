/**
 * 分享配置
 */
import { wxshareConfig } from "./interface"
import { onShareAppMessage, onShareTimeline } from '@dcloudio/uni-app'
export const share = (args : wxshareConfig = {}) => {
	let defaultWxshareConfig : wxshareConfig = {
		...args
	}
	// 分享朋友默认配置  
	let shareAppOptions: wxshareConfig = {}
	// 分享朋友圈默认配置  
	let shareTimeOptions: wxshareConfig = {}
	// onShareAppMessage  
	const shareApp = (options: wxshareConfig = {}) => {
		onShareAppMessage((res): wxshareConfig => {
			return {
				...defaultWxshareConfig,
				...options,
				...shareAppOptions
			}
		})
	}
	// 添加onShareAppMessage参数  
	const setShareApp = (options: wxshareConfig = {}) => {
		shareAppOptions = options
	}
	// onShareTimeline  
	const shareTime = (options: wxshareConfig = {}) => {
		onShareTimeline((): wxshareConfig => {
			return {
				...defaultWxshareConfig,
				...options,
				...shareTimeOptions
			}
		})
	}
	// 添加onShareTimeline参数  
	const setShareTime = (options = {}) => {
		shareTimeOptions = options
	}

	return {
		onShareAppMessage: shareApp,
		onShareTimeline: shareTime,
		setShareApp,
		setShareTime,
	}
}