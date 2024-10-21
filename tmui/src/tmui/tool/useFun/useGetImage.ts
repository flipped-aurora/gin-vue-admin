/**随机获取图片 */
export const useGetImage = function () {
	/**随机取得一张图片 */
	function getRandomImage(width:number=500) {
		return `https://store.tmui.design/api_v2/public/random_picture?random=${(Math.random() * 100).toFixed(0)}&width=${width}`
	}
	/**获取图片列表,row为数量，每次的最大值为14 */
	function getImageList(row: number,width:number=500) {
		if (row >= 14) {
			row = 14;
		}
		return `https://store.tmui.design/api_v2/public/random_picture?random=1&row=${row}&width=${width}`
	}

	return {
		getRandomImage,
		getImageList,
	}
}