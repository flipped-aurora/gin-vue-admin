//返回 一个节点从父到子的路径id组。
export const getNodeRouter = function (list = [], id = '', prentId = [], idmap = 'id') {
	let p = [];
	if (typeof prentId == 'undefined') {
		prentId = [];
	}
	if (!Array.isArray(id)) {
		id = [id]
	}
	let arr = Array.from(prentId)
	for (let i = 0, len = list.length; i < len; i++) {
		arr.push(list[i][idmap])
		if (list[i].id === id[0]) {
			return arr
		}
		let children = list[i].children
		if (children && children.length) {
			let result = getNodeRouter(children, id, arr, idmap = 'id')
			if (result) return result
		}
		arr.pop()
	}
	return null
}
//平铺它所有的节点id为一维数组。
export const treeFlat = function (arr = [], idmap = 'id') {
	let res = []
	arr.forEach((item) => {
		res.push(item[idmap])
		if (item.children) {
			res.push(...treeFlat(item.children, idmap = 'id'))
		}
	})
	return res
}

//检查 一个节点是否是否父节点，并返回它下面的所有子节点。
export const queryNodeIsParent = function (arr = [], id = "", idmap = 'id') {

	let res = null;
	for (let i = 0, len = arr.length; i < len; i++) {
		let item = arr[i]

		if (item[idmap] == id && item.children) {
			res = item;
			break;
		} else if (item.children) {
			let rulst = queryNodeIsParent(item.children, id, idmap = 'id');
			if (rulst) {
				res = rulst;
			}
		}
	}
	return res
}
//找出所有父节点id,返回数据
export const queryParentNode = function (arr = [], idmap = 'id'): Array<string | number> {
	let res = [];
	for (let i = 0, len = arr.length; i < len; i++) {
		let item = arr[i]
		if (item.children) {
			res.push(item[idmap])
			res.push(...queryParentNode(item.children))
		}
	}
	return res
}