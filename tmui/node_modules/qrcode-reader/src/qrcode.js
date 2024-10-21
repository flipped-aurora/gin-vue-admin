/*
   Copyright 2011 Lazar Laszlo (lazarsoft@gmail.com, www.lazarsoft.info)

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

import Detector from './detector';
import Decoder from './decoder';

export var qrcode = {};
qrcode.sizeOfDataLengthInfo =  [[10, 9, 8, 8], [12, 11, 16, 10], [14, 13, 16, 12]];

export default function QrCode() {

  this.imagedata = null;
  this.width = 0;
  this.height = 0;
  this.qrCodeSymbol = null;
  this.debug = false;

  this.callback = null;
}


QrCode.prototype.decode = function(src, data) {

  var decode = (function() {

    try {
      this.error = undefined;
      this.result = this.process(this.imagedata);
    } catch (e) {
      this.error = e;
      this.result = undefined;
    }

    if (this.callback != null) {
      this.callback(this.error, this.result);
    }

    return this.result;

  }).bind(this);

  if (src != undefined && src.width != undefined) {
    /* decode from canvas canvas.context.getImageData */
    this.width = src.width;
    this.height = src.height;
    this.imagedata = {"data": data || src.data};
    this.imagedata.width = src.width;
    this.imagedata.height = src.height;

    decode();
  } else {
    if (typeof Image === "undefined") {
      throw new Error("This source format is not supported in your environment, you need to pass an image buffer with width and height (see https://github.com/edi9999/jsqrcode/blob/master/test/qrcode.js)");
    }
    /* decode from URL */

    var image = new Image();
    image.crossOrigin = "Anonymous";

    image.onload = (function() {

      var canvas_qr = document.createElement('canvas');
      var context = canvas_qr.getContext('2d');
      var canvas_out = document.getElementById("out-canvas");

      if (canvas_out != null) {

        var outctx = canvas_out.getContext('2d');
        outctx.clearRect(0, 0, 320, 240);
        outctx.drawImage(image, 0, 0, 320, 240);
      }

      canvas_qr.width = image.width;
      canvas_qr.height = image.height;
      context.drawImage(image, 0, 0);
      this.width = image.width;
      this.height = image.height;

      try {
        this.imagedata = context.getImageData(0, 0, image.width, image.height);
      } catch (e) {
        this.result = "Cross domain image reading not supported in your browser! Save it to your computer then drag and drop the file!";
        if (this.callback != null) return this.callback(null, this.result);
      }

      decode();

    }).bind(this);

    image.src = src;
  }
};

QrCode.prototype.decode_utf8 = function(s) {

  return decodeURIComponent(escape(s));
};

QrCode.prototype.process = function(imageData) {

  var start = new Date().getTime();

  var image = this.grayScaleToBitmap(this.grayscale(imageData));

  var detector = new Detector(image);

  var qRCodeMatrix = detector.detect();

  /*for (var y = 0; y < qRCodeMatrix.bits.height; y++)
   {
   for (var x = 0; x < qRCodeMatrix.bits.width; x++)
   {
   var point = (x * 4*2) + (y*2 * imageData.width * 4);
   imageData.data[point] = qRCodeMatrix.bits.get_Renamed(x,y)?0:0;
   imageData.data[point+1] = qRCodeMatrix.bits.get_Renamed(x,y)?0:0;
   imageData.data[point+2] = qRCodeMatrix.bits.get_Renamed(x,y)?255:0;
   }
   }*/

  var reader = Decoder.decode(qRCodeMatrix.bits);
  var data = reader.DataByte;
  var str = "";
  for (var i = 0; i < data.length; i++) {
    for (var j = 0; j < data[i].length; j++)
      str += String.fromCharCode(data[i][j]);
  }

  var end = new Date().getTime();
  var time = end - start;
  if (this.debug) {
    console.log('QR Code processing time (ms): ' + time);
  }

  return {result: this.decode_utf8(str), points: qRCodeMatrix.points};
};

QrCode.prototype.getPixel = function(imageData, x, y) {
  if (imageData.width < x) {
    throw "point error";
  }
  if (imageData.height < y) {
    throw "point error";
  }
  var point = (x * 4) + (y * imageData.width * 4);
  return (imageData.data[point] * 33 + imageData.data[point + 1] * 34 + imageData.data[point + 2] * 33) / 100;
};

QrCode.prototype.binarize = function(th) {
  var ret = new Array(this.width * this.height);
  for (var y = 0; y < this.height; y++) {
    for (var x = 0; x < this.width; x++) {
      var gray = this.getPixel(x, y);

      ret[x + y * this.width] = gray <= th;
    }
  }
  return ret;
};

QrCode.prototype.getMiddleBrightnessPerArea = function(imageData) {
  var numSqrtArea = 4;
  //obtain middle brightness((min + max) / 2) per area
  var areaWidth = Math.floor(imageData.width / numSqrtArea);
  var areaHeight = Math.floor(imageData.height / numSqrtArea);
  var minmax = new Array(numSqrtArea);
  for (var i = 0; i < numSqrtArea; i++) {
    minmax[i] = new Array(numSqrtArea);
    for (var i2 = 0; i2 < numSqrtArea; i2++) {
      minmax[i][i2] = [0, 0];
    }
  }
  for (var ay = 0; ay < numSqrtArea; ay++) {
    for (var ax = 0; ax < numSqrtArea; ax++) {
      minmax[ax][ay][0] = 0xFF;
      for (var dy = 0; dy < areaHeight; dy++) {
        for (var dx = 0; dx < areaWidth; dx++) {
          var target = imageData.data[areaWidth * ax + dx + (areaHeight * ay + dy) * imageData.width];
          if (target < minmax[ax][ay][0])
            minmax[ax][ay][0] = target;
          if (target > minmax[ax][ay][1])
            minmax[ax][ay][1] = target;
        }
      }
    }
  }
  var middle = new Array(numSqrtArea);
  for (var i3 = 0; i3 < numSqrtArea; i3++) {
    middle[i3] = new Array(numSqrtArea);
  }
  for (var ay = 0; ay < numSqrtArea; ay++) {
    for (var ax = 0; ax < numSqrtArea; ax++) {
      middle[ax][ay] = Math.floor((minmax[ax][ay][0] + minmax[ax][ay][1]) / 2);
    }
  }

  return middle;
};

QrCode.prototype.grayScaleToBitmap = function(grayScaleImageData) {
  var middle = this.getMiddleBrightnessPerArea(grayScaleImageData);
  var sqrtNumArea = middle.length;
  var areaWidth = Math.floor(grayScaleImageData.width / sqrtNumArea);
  var areaHeight = Math.floor(grayScaleImageData.height / sqrtNumArea);

  for (var ay = 0; ay < sqrtNumArea; ay++) {
    for (var ax = 0; ax < sqrtNumArea; ax++) {
      for (var dy = 0; dy < areaHeight; dy++) {
        for (var dx = 0; dx < areaWidth; dx++) {
          grayScaleImageData.data[areaWidth * ax + dx + (areaHeight * ay + dy) * grayScaleImageData.width] = (grayScaleImageData.data[areaWidth * ax + dx + (areaHeight * ay + dy) * grayScaleImageData.width] < middle[ax][ay]);
        }
      }
    }
  }
  return grayScaleImageData;
};

QrCode.prototype.grayscale = function(imageData) {
  var ret = new Array(imageData.width * imageData.height);

  for (var y = 0; y < imageData.height; y++) {
    for (var x = 0; x < imageData.width; x++) {
      var gray = this.getPixel(imageData, x, y);

      ret[x + y * imageData.width] = gray;
    }
  }

  return {
    height: imageData.height,
    width: imageData.width,
    data: ret
  };
};

export function URShift(number,  bits) {
  if (number >= 0)
    return number >> bits;
  else
    return (number >> bits) + (2 << ~bits);
}
