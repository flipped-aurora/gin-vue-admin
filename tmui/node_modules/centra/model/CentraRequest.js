const path = require('path')
const http = require('http')
const https = require('https')
const followRedirects = require('follow-redirects')
const qs = require('querystring')
const zlib = require('zlib')
const {URL} = require('url')

const CentraResponse = require('./CentraResponse.js')

const supportedCompressions = ['gzip', 'deflate', 'br']

const useRequest = (protocol, maxRedirects) => {
	let httpr
	let httpsr
	if (maxRedirects <= 0) {
		httpr = http.request
		httpsr = https.request
	}
	else {
		httpr = followRedirects.http.request
		httpsr = followRedirects.https.request
	}

	if (protocol === 'http:') {
		return httpr
	}
	else if (protocol === 'https:') {
		return httpsr
	}
	else throw new Error('Bad URL protocol: ' + protocol)
}

module.exports = class CentraRequest {
	constructor (url, method = 'GET') {
		this.url = typeof url === 'string' ? new URL(url) : url
		this.method = method
		this.data = null
		this.sendDataAs = null
		this.reqHeaders = {}
		this.streamEnabled = false
		this.compressionEnabled = false
		this.timeoutTime = null
		this.coreOptions = {}
		this.maxRedirects = 0

		this.resOptions = {
			'maxBuffer': 50 * 1000000 // 50 MB
		}

		return this
	}

	followRedirects(n) {
		this.maxRedirects = n

		return this
	}

	query (a1, a2) {
		if (typeof a1 === 'object') {
			Object.keys(a1).forEach((queryKey) => {
				this.url.searchParams.append(queryKey, a1[queryKey])
			})
		}
		else this.url.searchParams.append(a1, a2)

		return this
	}

	path (relativePath) {
		this.url.pathname = path.join(this.url.pathname, relativePath)

		return this
	}

	body (data, sendAs) {
		this.sendDataAs = typeof data === 'object' && !sendAs && !Buffer.isBuffer(data) ? 'json' : (sendAs ? sendAs.toLowerCase() : 'buffer')
		this.data = this.sendDataAs === 'form' ? qs.stringify(data) : (this.sendDataAs === 'json' ? JSON.stringify(data) : data)

		return this
	}

	header (a1, a2) {
		if (typeof a1 === 'object') {
			Object.keys(a1).forEach((headerName) => {
				this.reqHeaders[headerName.toLowerCase()] = a1[headerName]
			})
		}
		else this.reqHeaders[a1.toLowerCase()] = a2

		return this
	}

	timeout (timeout) {
		this.timeoutTime = timeout

		return this
	}

	option (name, value) {
		this.coreOptions[name] = value

		return this
	}

	stream () {
		this.streamEnabled = true

		return this
	}

	compress () {
		this.compressionEnabled = true

		if (!this.reqHeaders['accept-encoding']) this.reqHeaders['accept-encoding'] = supportedCompressions.join(', ')

		return this
	}

	send () {
		return new Promise((resolve, reject) => {
			if (this.data) {
				if (!this.reqHeaders.hasOwnProperty('content-type')) {
					if (this.sendDataAs === 'json') {
						this.reqHeaders['content-type'] = 'application/json'
					}
					else if (this.sendDataAs === 'form') {
						this.reqHeaders['content-type'] = 'application/x-www-form-urlencoded'
					}
				}

				if (!this.reqHeaders.hasOwnProperty('content-length')) {
					this.reqHeaders['content-length'] = Buffer.byteLength(this.data)
				}
			}

			const options = Object.assign({
				'protocol': this.url.protocol,
				'host': this.url.hostname.replace('[', '').replace(']', ''),
				'port': this.url.port,
				'path': this.url.pathname + (this.url.search === null ? '' : this.url.search),
				'method': this.method,
				'headers': this.reqHeaders,
				'maxRedirects': this.maxRedirects
			}, this.coreOptions)

			let req

			const resHandler = (res) => {
				let stream = res

				if (this.compressionEnabled) {
					if (res.headers['content-encoding'] === 'gzip') {
						stream = res.pipe(zlib.createGunzip())
					}
					else if (res.headers['content-encoding'] === 'deflate') {
						stream = res.pipe(zlib.createInflate())
					}
					else if (res.headers['content-encoding'] === 'br') {
						stream = res.pipe(zlib.createBrotliDecompress())
					}
				}

				let centraRes

				if (this.streamEnabled) {
					resolve(stream)
				}
				else {
					centraRes = new CentraResponse(res, this.resOptions)

					stream.on('error', (err) => {
						reject(err)
					})

					stream.on('aborted', () => {
						reject(new Error('Server aborted request'))
					})

					stream.on('data', (chunk) => {
						centraRes._addChunk(chunk)

						if (this.resOptions.maxBuffer !== null && centraRes.body.length > this.resOptions.maxBuffer) {
							stream.destroy()

							reject('Received a response which was longer than acceptable when buffering. (' + this.body.length + ' bytes)')
						}
					})

					stream.on('end', () => {
						resolve(centraRes)
					})
				}
			}

			const request = useRequest(this.url.protocol, this.maxRedirects)

			req = request(options, resHandler)

			if (this.timeoutTime) {
				req.setTimeout(this.timeoutTime, () => {
					req.abort()

					if (!this.streamEnabled) {
						reject(new Error('Timeout reached'))
					}
				})
			}

			req.on('error', (err) => {
				reject(err)
			})

			if (this.data) req.write(this.data)

			req.end()
		})
	}
}
