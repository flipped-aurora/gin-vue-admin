export function addition(arg1: number[], arg2: number[]) {
	const result: number[] = []
	if (arg1.length === arg2.length) {
		// 遍历数组并相加对应的值
		for (let i = 0; i < arg1.length; i++) {
			result.push(arg1[i] + arg2[i]);
		}
	}
	return result;
}
export function subtract(arg1: number[], arg2: number[], reserver: boolean = false) {
	const result: number[] = []
	if (arg1.length === arg2.length) {
		// 遍历数组并相加对应的值
		for (let i = 0; i < arg1.length; i++) {
			if (!reserver) {
				result.push(arg1[i] - arg2[i]);
			} else {
				result.push(arg2[i] - arg1[i]);
			}
		}
	}
	return result;
}
export function multiply(arg1: number[], arg2: number) {
	const result: number[] = []
	// 遍历数组并相加对应的值
	for (let i = 0; i < arg1.length; i++) {
		result.push(arg1[i] * arg2);
	}
	return result;
}
export function division(arg1: number[], arg2: number) {
	const result: number[] = []
	// 遍历数组并相加对应的值
	for (let i = 0; i < arg1.length; i++) {
		result.push(arg1[i] / arg2);
	}
	return result;
}


export function uuid(len = 24, radix: number) {
	var chars = '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz'.split('');
	var uuid = [],
		i;
	radix = radix || chars.length;
	if (len) {
		// Compact form
		for (i = 0; i < len; i++) uuid[i] = chars[0 | Math.random() * radix];
	} else {
		// rfc4122, version 4 form
		var r;
		// rfc4122 requires these characters
		uuid[8] = uuid[13] = uuid[18] = uuid[23] = '-';
		uuid[14] = '4';
		for (i = 0; i < 36; i++) {
			if (!uuid[i]) {
				r = 0 | Math.random() * 16;
				uuid[i] = chars[(i == 19) ? (r & 0x3) | 0x8 : r];

			}

		}

	}

	return uuid.join('');
}