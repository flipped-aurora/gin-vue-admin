module.exports = class CentraResponse {
	constructor (res, resOptions) {
		this.coreRes = res
		this.resOptions = resOptions

		this.body = Buffer.alloc(0)

		this.headers = res.headers
		this.statusCode = res.statusCode
	}

	_addChunk (chunk) {
		this.body = Buffer.concat([this.body, chunk])
	}

	async json () {
		return this.statusCode === 204 ? null : JSON.parse(this.body)
	}

	async text () {
		return this.body.toString()
	}
}
