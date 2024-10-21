import type { rulesItem } from "./interface"
import type { formItem } from "./../tm-form/interface";
export function validateFunCall(rules: Array<rulesItem>, value: any) {
	rules = rules.map((el) => {
		if (typeof el.validator === "function" && el.required === true) {
			return el;
		} else if (typeof el.validator === "boolean" && el.required === true) {
			return {
				...el,
				validator: (val: any) => {
					if (val === null || val === "" || typeof val == "undefined") return false;
					if (typeof val === "object") {
						if (Array.isArray(val)) {
							if (val.length == 0) return false;
						} else if (Object.keys(val).length === 0 && val.constructor === Object) {
							return false;
						}
					}
					if (typeof val === "boolean") {
						return val;
					}
					if (typeof val === "number") {
						if (isNaN(val)) return false;
						if (Number(val) < 0) return false;
					}
					if (typeof val === "string") {
						if (val.trim().length == 0) return false;
					}
					return true;
				},
			};
		} else {
			return {
				...el,
				validator: (val: string | number) => {
					return true;
				},
			};
		}
	});
	let rules_filter: Array<rulesItem> = rules.filter((el) => {
		return typeof el.validator === "function" && el.required === true;
	});
	let rules_fun: Array<rulesItem> = rules_filter.map((el) => {
		let validator = true;
		if (typeof el.validator === "function") {
			let vr = el.validator(value);
			if (vr) {
				validator = true;
			} else {
				validator = false;
			}
		} else {
			validator = true;
		}

		return {
			message: String(el.message),
			validator: validator,
		}
	});

	return rules_fun;
};



export function getObjectVal(obj: any, field = "") {
	if (field == "") return obj;
	var arr = field.split(".");
	while (arr.length > 1) {
		let key = String(arr.shift());
		obj = obj[key];
	}
	return obj[arr[0]];
}