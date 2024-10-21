var expect = require('chai').expect;
var fs = require('fs');
var QrCode = require('../dist/index.js');
var ImageParser = require("image-parser");
const Jimp = require("jimp");

const expectedResult = {
  "result": 'Test',
  "points": [
    {
      "count": 2,
      "estimatedModuleSize": 8,
      "x": 36,
      "y": 148,
    },
    {
      "count": 2,
      "estimatedModuleSize": 8,
      "x": 36,
      "y": 36,
    },
    {
      "count": 2,
      "estimatedModuleSize": 8,
      "x": 148,
      "y": 36,
    }
  ]
};

function copy(input) {
  return JSON.parse(JSON.stringify(input));
}

it("should work with jimp", function(done) {
  var buffer = fs.readFileSync(__dirname + '/image.png');
  Jimp.read(buffer, function(err, image) {
    if (err) {
      return done(err);
    }
    var qr = new QrCode();
    qr.callback = function(err, result) {
      if (err) {
        return done(err);
      }
      expect(copy(result)).to.deep.equal(expectedResult);
      done();
    };
    qr.decode(image.bitmap);
  });
});

it("should work with a zxing qr code with jimp", function(done) {
  var buffer = fs.readFileSync(__dirname + '/image-zxing.png');
  Jimp.read(buffer, function(err, image) {
    if (err) {
      return done(err);
    }
    var qr = new QrCode();
    qr.callback = function(err, result) {
      if (err) {
        return done(err);
      }
      expect(copy(result)).to.deep.equal({
        "result": 'Test',
        "points": [
          {
            "count": 2,
            "estimatedModuleSize": 9,
            "x": 34.5,
            "y": 160.5,
          },
          {
            "count": 3,
            "estimatedModuleSize": 9,
            "x": 34.5,
            "y": 34.5,
          },
          {
            "count": 2,
            "estimatedModuleSize": 9.428571428571429,
            "x": 160.5,
            "y": 34.5,
          }
        ]
      });
      done();
    };
    qr.decode(image.bitmap);
  });
});

it('should work with basic image', function(done) {
  var buffer = fs.readFileSync(__dirname + '/image.png');
  var img = new ImageParser(buffer);
  img.parse(function(err) {
    if (err) {
      return done(err);
    }
    var qr = new QrCode();
    qr.callback = function(err, result) {
      if (err) {
        return done(err);
      }
      expect(copy(result)).to.deep.equal(expectedResult);
      done();
    };
    qr.decode({width: img.width(), height: img.height()}, img._imgBuffer);
  });
});

it('should work with imageData format', function(done) {
  var buffer = fs.readFileSync(__dirname + '/image.png');
  var img = new ImageParser(buffer);
  img.parse(function(err) {
    if (err) {
      return done(err);
    }
    var qr = new QrCode();
    qr.callback = function(err, result) {
      if (err) {
        return done(err);
      }
      expect(copy(result)).to.deep.equal({
        "result": 'Test',
        "points": [
          {
            "count": 2,
            "estimatedModuleSize": 8,
            "x": 36,
            "y": 148,
          },
          {
            "count": 2,
            "estimatedModuleSize": 8,
            "x": 36,
            "y": 36,
          },
          {
            "count": 2,
            "estimatedModuleSize": 8,
            "x": 148,
            "y": 36,
          }
        ]
      });
      done();
    };
    qr.decode({height: img.height(), width: img.width(), data: img._imgBuffer});
  });
});
