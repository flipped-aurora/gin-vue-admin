
import { fetchNet } from '../lib/fetch';
import { isReactive, isRef, ref, Ref, ReactiveEffect, watch } from "vue"

export const useFetch = (url : string, config ?: Ref<Tmui.fetchNetConfigType> | Tmui.fetchNetConfigType | ReactiveEffect<Tmui.fetchNetConfigType>) => {
	//是否加载中
	let loading = ref(false);
	//是否从未开始渲染过
	let pending = ref(true)
	//是否访问错误
	let error = ref(false)
	//数据集,
	let data = ref<any>(undefined)

	//是否空状态。对refresh没有用，会被清空为false,适合列表为空时没有更多数据会禁止getData请求，但通过refresh来重新加载。
	let empty = ref(false)

	let watchRefresh = getRealDataConfig()?.watchRefresh ?? false;
	let showLoading = getRealDataConfig()?.showLoading ?? true;

	let tid : any = null;

	function getRealDataConfig() {
		if (isReactive(config)) {
			return uni.$tm.u.deepClone(config)
		}
		if (isRef(config)) {
			return uni.$tm.u.deepClone(config.value)
		}
		return config
	}
	//重新刷新接口请求
	async function refresh() {
		loading.value = false;
		error.value = false;
		// data.value = undefined;
		empty.value = false;
		await rq();
	}
	function toast(msg : string, err : boolean, cfg : Tmui.fetchNetConfigType) {
		return new Promise((res, rej) => {
			if (cfg.showToast) {
				if (
					cfg.toast == 'success' && !err ||
					cfg.toast == 'fail' && err ||
					cfg.toast == 'all'
				) {
					uni.showToast({
						title: msg,
						icon: "none",
						mask: true
					})
					setTimeout(function () {
						res(true)
					}, 1200);
				} else {
					res(true)
				}

			} else {
				res(true)
			}
		});
	}
	//请求数据。
	async function rq() {
		if (loading.value || empty.value) return Promise.resolve(true);
		error.value = false;
		loading.value = true;

		if (showLoading) {
			uni.showLoading({
				title: "...",
				mask: true,
				icon: 'none'
			})
		}
		let cfg : Tmui.fetchNetConfigType = getRealDataConfig() || {};
		cfg = { showToast: true, toastKey: 'msg', toast: 'all', ...cfg };
		let errmsg = "";
		let newdata = await fetchNet.request({
			url: url,
			...cfg
		}).catch(err => {
			error.value = true;
			errmsg = typeof err == 'object' ? JSON.stringify(err) : err
		}).finally(() => {
			loading.value = false;
			pending.value = false;
			uni.hideLoading();
		})
		if (!newdata) return toast(errmsg || "网络错误", false, cfg);
		if (newdata.statusCode !== (cfg?.statusCode || 200)) {
			error.value = true;
			return toast(newdata.errMsg || "网络错误", false, cfg)
		}
		let d = newdata.data;
		if (Array.isArray(cfg?.pick)) {
			let pd = cfg.pick;

			if (typeof d !== 'object') {
				data.value = d;
				return toast("操作成功", true, cfg)
			}

			let td : { [key : string] : any } = {};
			for (let key in d) {
				if (pd.includes(key)) {
					td[key] = d[key];
				}
			}

			data.value = td;
			return toast(d[cfg.toastKey] || "操作成功", true, cfg)
		}
		data.value = newdata;
		return toast(typeof d === 'object' ? d[cfg.toastKey] || "操作成功" : "操作成功", true, cfg)
	}
	if ((isReactive(config) || isRef(config)) && watchRefresh) {
		watch(config, () => {
			if (pending.value) return;
			if (tid === null) {
				refresh()
				tid = setTimeout(function () { }, 100);
			} else {
				clearTimeout(tid)
				tid = setTimeout(function () {
					refresh()
				}, 100);
			}

		}, { deep: true })
	}


	return {
		loading,
		error,
		data,
		pending,
		refresh,
		getData: async () => await rq(),
		empty
	}
}