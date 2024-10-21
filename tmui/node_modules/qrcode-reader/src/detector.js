/*
  Ported to JavaScript by Lazar Laszlo 2011

  lazarsoft@gmail.com, www.lazarsoft.info

*/

/*
*
* Copyright 2007 ZXing authors
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/

import Version from './version';
import {AlignmentPatternFinder} from './alignpat';
import GridSampler from './grid';
import {FinderPatternFinder} from './findpat';

function PerspectiveTransform(a11,  a21,  a31,  a12,  a22,  a32,  a13,  a23,  a33) {
  this.a11 = a11;
  this.a12 = a12;
  this.a13 = a13;
  this.a21 = a21;
  this.a22 = a22;
  this.a23 = a23;
  this.a31 = a31;
  this.a32 = a32;
  this.a33 = a33;
}

PerspectiveTransform.prototype.transformPoints1 = function(points) {
  var max = points.length;
  var a11 = this.a11;
  var a12 = this.a12;
  var a13 = this.a13;
  var a21 = this.a21;
  var a22 = this.a22;
  var a23 = this.a23;
  var a31 = this.a31;
  var a32 = this.a32;
  var a33 = this.a33;
  for (var i = 0; i < max; i += 2) {
    var x = points[i];
    var y = points[i + 1];
    var denominator = a13 * x + a23 * y + a33;
    points[i] = (a11 * x + a21 * y + a31) / denominator;
    points[i + 1] = (a12 * x + a22 * y + a32) / denominator;
  }
};

PerspectiveTransform.prototype.transformPoints2 = function(xValues, yValues) {
  var n = xValues.length;
  for (var i = 0; i < n; i++) {
    var x = xValues[i];
    var y = yValues[i];
    var denominator = this.a13 * x + this.a23 * y + this.a33;
    xValues[i] = (this.a11 * x + this.a21 * y + this.a31) / denominator;
    yValues[i] = (this.a12 * x + this.a22 * y + this.a32) / denominator;
  }
};

PerspectiveTransform.prototype.buildAdjoint = function() {
  // Adjoint is the transpose of the cofactor matrix:
  return new PerspectiveTransform(this.a22 * this.a33 - this.a23 * this.a32, this.a23 * this.a31 - this.a21 * this.a33, this.a21 * this.a32 - this.a22 * this.a31, this.a13 * this.a32 - this.a12 * this.a33, this.a11 * this.a33 - this.a13 * this.a31, this.a12 * this.a31 - this.a11 * this.a32, this.a12 * this.a23 - this.a13 * this.a22, this.a13 * this.a21 - this.a11 * this.a23, this.a11 * this.a22 - this.a12 * this.a21);
};

PerspectiveTransform.prototype.times = function(other) {
  return new PerspectiveTransform(this.a11 * other.a11 + this.a21 * other.a12 + this.a31 * other.a13, this.a11 * other.a21 + this.a21 * other.a22 + this.a31 * other.a23, this.a11 * other.a31 + this.a21 * other.a32 + this.a31 * other.a33, this.a12 * other.a11 + this.a22 * other.a12 + this.a32 * other.a13, this.a12 * other.a21 + this.a22 * other.a22 + this.a32 * other.a23, this.a12 * other.a31 + this.a22 * other.a32 + this.a32 * other.a33, this.a13 * other.a11 + this.a23 * other.a12 + this.a33 * other.a13, this.a13 * other.a21 + this.a23 * other.a22 + this.a33 * other.a23, this.a13 * other.a31 + this.a23 * other.a32 + this.a33 * other.a33);
};

PerspectiveTransform.quadrilateralToQuadrilateral = function(x0,  y0,  x1,  y1,  x2,  y2,  x3,  y3,  x0p,  y0p,  x1p,  y1p,  x2p,  y2p,  x3p,  y3p) {

  var qToS = this.quadrilateralToSquare(x0, y0, x1, y1, x2, y2, x3, y3);
  var sToQ = this.squareToQuadrilateral(x0p, y0p, x1p, y1p, x2p, y2p, x3p, y3p);
  return sToQ.times(qToS);
};

PerspectiveTransform.squareToQuadrilateral = function(x0,  y0,  x1,  y1,  x2,  y2,  x3,  y3) {
  var dy2 = y3 - y2;
  var dy3 = y0 - y1 + y2 - y3;
  if (dy2 == 0.0 && dy3 == 0.0) {
    return new PerspectiveTransform(x1 - x0, x2 - x1, x0, y1 - y0, y2 - y1, y0, 0.0, 0.0, 1.0);
  } else {
    var dx1 = x1 - x2;
    var dx2 = x3 - x2;
    var dx3 = x0 - x1 + x2 - x3;
    var dy1 = y1 - y2;
    var denominator = dx1 * dy2 - dx2 * dy1;
    var a13 = (dx3 * dy2 - dx2 * dy3) / denominator;
    var a23 = (dx1 * dy3 - dx3 * dy1) / denominator;
    return new PerspectiveTransform(x1 - x0 + a13 * x1, x3 - x0 + a23 * x3, x0, y1 - y0 + a13 * y1, y3 - y0 + a23 * y3, y0, a13, a23, 1.0);
  }
};

PerspectiveTransform.quadrilateralToSquare = function(x0,  y0,  x1,  y1,  x2,  y2,  x3,  y3) {
  // Here, the adjoint serves as the inverse:
  return this.squareToQuadrilateral(x0, y0, x1, y1, x2, y2, x3, y3).buildAdjoint();
};

function DetectorResult(bits,  points) {
  this.bits = bits;
  this.points = points;
}

export default function Detector(image) {
  this.image = image;
  this.resultPointCallback = null;
}

Detector.prototype.sizeOfBlackWhiteBlackRun = function(fromX,  fromY,  toX,  toY) {
  // Mild variant of Bresenham's algorithm;
  // see http://en.wikipedia.org/wiki/Bresenham's_line_algorithm
  var steep = Math.abs(toY - fromY) > Math.abs(toX - fromX);
  if (steep) {
    var temp = fromX;
    fromX = fromY;
    fromY = temp;
    temp = toX;
    toX = toY;
    toY = temp;
  }

  var dx = Math.abs(toX - fromX);
  var dy = Math.abs(toY - fromY);
  var error = -dx >> 1;
  var ystep = fromY < toY ? 1 : -1;
  var xstep = fromX < toX ? 1 : -1;
  var state = 0; // In black pixels, looking for white, first or second time
  for (var x = fromX, y = fromY; x != toX; x += xstep) {

    var realX = steep ? y : x;
    var realY = steep ? x : y;
    if (state == 1) {
      // In white pixels, looking for black
      if (this.image.data[realX + realY * this.image.width]) {
        state++;
      }
    } else {
      if (!this.image.data[realX + realY * this.image.width]) {
        state++;
      }
    }

    if (state == 3) {
      // Found black, white, black, and stumbled back onto white; done
      var diffX = x - fromX;
      var diffY = y - fromY;
      return  Math.sqrt((diffX * diffX + diffY * diffY));
    }
    error += dy;
    if (error > 0) {
      if (y == toY) {
        break;
      }
      y += ystep;
      error -= dx;
    }
  }
  var diffX2 = toX - fromX;
  var diffY2 = toY - fromY;
  return  Math.sqrt((diffX2 * diffX2 + diffY2 * diffY2));
};

Detector.prototype.sizeOfBlackWhiteBlackRunBothWays = function(fromX,  fromY,  toX,  toY) {

  var result = this.sizeOfBlackWhiteBlackRun(fromX, fromY, toX, toY);

  // Now count other way -- don't run off image though of course
  var scale = 1.0;
  var otherToX = fromX - (toX - fromX);
  if (otherToX < 0) {
    scale =  fromX /  (fromX - otherToX);
    otherToX = 0;
  } else if (otherToX >= this.image.width) {
    scale =  (this.image.width - 1 - fromX) /  (otherToX - fromX);
    otherToX = this.image.width - 1;
  }
  var otherToY = Math.floor(fromY - (toY - fromY) * scale);

  scale = 1.0;
  if (otherToY < 0) {
    scale =  fromY /  (fromY - otherToY);
    otherToY = 0;
  } else if (otherToY >= this.image.height) {
    scale =  (this.image.height - 1 - fromY) /  (otherToY - fromY);
    otherToY = this.image.height - 1;
  }
  otherToX = Math.floor(fromX + (otherToX - fromX) * scale);

  result += this.sizeOfBlackWhiteBlackRun(fromX, fromY, otherToX, otherToY);
  return result - 1.0; // -1 because we counted the middle pixel twice
};

Detector.prototype.calculateModuleSizeOneWay = function(pattern,  otherPattern) {
  var moduleSizeEst1 = this.sizeOfBlackWhiteBlackRunBothWays(Math.floor(pattern.X), Math.floor(pattern.Y), Math.floor(otherPattern.X), Math.floor(otherPattern.Y));
  var moduleSizeEst2 = this.sizeOfBlackWhiteBlackRunBothWays(Math.floor(otherPattern.X), Math.floor(otherPattern.Y), Math.floor(pattern.X), Math.floor(pattern.Y));
  if (isNaN(moduleSizeEst1)) {
    return moduleSizeEst2 / 7.0;
  }
  if (isNaN(moduleSizeEst2)) {
    return moduleSizeEst1 / 7.0;
  }
  // Average them, and divide by 7 since we've counted the width of 3 black modules,
  // and 1 white and 1 black module on either side. Ergo, divide sum by 14.
  return (moduleSizeEst1 + moduleSizeEst2) / 14.0;
};

Detector.prototype.calculateModuleSize = function(topLeft,  topRight,  bottomLeft) {
  // Take the average
  return (this.calculateModuleSizeOneWay(topLeft, topRight) + this.calculateModuleSizeOneWay(topLeft, bottomLeft)) / 2.0;
};

Detector.prototype.distance = function(pattern1,  pattern2) {
  var xDiff = pattern1.X - pattern2.X;
  var yDiff = pattern1.Y - pattern2.Y;
  return  Math.sqrt((xDiff * xDiff + yDiff * yDiff));
};

Detector.prototype.computeDimension = function(topLeft,  topRight,  bottomLeft,  moduleSize) {
  var tltrCentersDimension = Math.round(this.distance(topLeft, topRight) / moduleSize);
  var tlblCentersDimension = Math.round(this.distance(topLeft, bottomLeft) / moduleSize);
  var dimension = ((tltrCentersDimension + tlblCentersDimension) >> 1) + 7;
  switch (dimension & 0x03) {
  // mod 4
  case 0:
    dimension++;
    break;
  // 1? do nothing

  case 2:
    dimension--;
    break;

  case 3:
    throw "Error";
  }
  return dimension;
};

Detector.prototype.findAlignmentInRegion = function(overallEstModuleSize,  estAlignmentX,  estAlignmentY,  allowanceFactor) {
  // Look for an alignment pattern (3 modules in size) around where it
  // should be
  var allowance = Math.floor(allowanceFactor * overallEstModuleSize);
  var alignmentAreaLeftX = Math.max(0, estAlignmentX - allowance);
  var alignmentAreaRightX = Math.min(this.image.width - 1, estAlignmentX + allowance);
  if (alignmentAreaRightX - alignmentAreaLeftX < overallEstModuleSize * 3) {
    throw "Error";
  }

  var alignmentAreaTopY = Math.max(0, estAlignmentY - allowance);
  var alignmentAreaBottomY = Math.min(this.image.height - 1, estAlignmentY + allowance);

  var alignmentFinder = new AlignmentPatternFinder(this.image, alignmentAreaLeftX, alignmentAreaTopY, alignmentAreaRightX - alignmentAreaLeftX, alignmentAreaBottomY - alignmentAreaTopY, overallEstModuleSize, this.resultPointCallback);
  return alignmentFinder.find();
};

Detector.prototype.createTransform = function(topLeft,  topRight,  bottomLeft, alignmentPattern, dimension) {
  var dimMinusThree =  dimension - 3.5;
  var bottomRightX;
  var bottomRightY;
  var sourceBottomRightX;
  var sourceBottomRightY;
  if (alignmentPattern != null) {
    bottomRightX = alignmentPattern.X;
    bottomRightY = alignmentPattern.Y;
    sourceBottomRightX = sourceBottomRightY = dimMinusThree - 3.0;
  } else {
    // Don't have an alignment pattern, just make up the bottom-right point
    bottomRightX = (topRight.X - topLeft.X) + bottomLeft.X;
    bottomRightY = (topRight.Y - topLeft.Y) + bottomLeft.Y;
    sourceBottomRightX = sourceBottomRightY = dimMinusThree;
  }

  var transform = PerspectiveTransform.quadrilateralToQuadrilateral(3.5, 3.5, dimMinusThree, 3.5, sourceBottomRightX, sourceBottomRightY, 3.5, dimMinusThree, topLeft.X, topLeft.Y, topRight.X, topRight.Y, bottomRightX, bottomRightY, bottomLeft.X, bottomLeft.Y);

  return transform;
};

Detector.prototype.sampleGrid = function(image,  transform,  dimension) {

  var sampler = GridSampler;
  return sampler.sampleGrid3(image, dimension, transform);
};

Detector.prototype.processFinderPatternInfo = function(info) {

  var topLeft = info.topLeft;
  var topRight = info.topRight;
  var bottomLeft = info.bottomLeft;

  var moduleSize = this.calculateModuleSize(topLeft, topRight, bottomLeft);
  if (moduleSize < 1.0) {
    throw "Error";
  }
  var dimension = this.computeDimension(topLeft, topRight, bottomLeft, moduleSize);
  var provisionalVersion = Version.getProvisionalVersionForDimension(dimension);
  var modulesBetweenFPCenters = provisionalVersion.DimensionForVersion - 7;

  var alignmentPattern = null;
  // Anything above version 1 has an alignment pattern
  if (provisionalVersion.alignmentPatternCenters.length > 0) {

    // Guess where a "bottom right" finder pattern would have been
    var bottomRightX = topRight.X - topLeft.X + bottomLeft.X;
    var bottomRightY = topRight.Y - topLeft.Y + bottomLeft.Y;

    // Estimate that alignment pattern is closer by 3 modules
    // from "bottom right" to known top left location
    var correctionToTopLeft = 1.0 - 3.0 /  modulesBetweenFPCenters;
    var estAlignmentX = Math.floor(topLeft.X + correctionToTopLeft * (bottomRightX - topLeft.X));
    var estAlignmentY = Math.floor(topLeft.Y + correctionToTopLeft * (bottomRightY - topLeft.Y));

    // Kind of arbitrary -- expand search radius before giving up
    for (var i = 4; i <= 16; i <<= 1) {
      //try
      //{
      alignmentPattern = this.findAlignmentInRegion(moduleSize, estAlignmentX, estAlignmentY,  i);
      break;
      //}
      //catch (re)
      //{
      // try next round
      //}
    }
    // If we didn't find alignment pattern... well try anyway without it
  }

  var transform = this.createTransform(topLeft, topRight, bottomLeft, alignmentPattern, dimension);

  var bits = this.sampleGrid(this.image, transform, dimension);

  var points;
  if (alignmentPattern == null) {
    points = [bottomLeft, topLeft, topRight];
  } else {
    points = [bottomLeft, topLeft, topRight, alignmentPattern];
  }
  return new DetectorResult(bits, points);
};

Detector.prototype.detect = function() {
  var info =  new FinderPatternFinder().findFinderPattern(this.image);

  return this.processFinderPatternInfo(info);
};
