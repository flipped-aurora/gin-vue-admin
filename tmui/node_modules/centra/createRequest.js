const CentraRequest = require('./model/CentraRequest.js')

module.exports = (url, method) => {
	return new CentraRequest(url, method)
}
