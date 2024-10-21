var test = require('tape')
var load = require('./')
var expectedArial = require('./fnt/Arial.json')
var fs = require('fs')
var http = require('http')

var arialBin = fs.readFileSync('fnt/Arial.bin')

test('should load from server URL', function (t) {
  t.plan(1)

  const server = http.createServer((req,res) => {
    res.end(arialBin)
  })
  
  server.listen(8003, () => {
    load({
      url: 'http://localhost:8003',
      binary: true
    }, (err, res) => {
      if (err) t.fail(err)
      else t.deepEqual(res, expectedArial)
      server.close()
    })
  })
})
