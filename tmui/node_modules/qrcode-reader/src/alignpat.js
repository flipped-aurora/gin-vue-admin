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


export default function AlignmentPattern(posX, posY,  estimatedModuleSize) {
  this.x = posX;
  this.y = posY;
  this.count = 1;
  this.estimatedModuleSize = estimatedModuleSize;
}

Object.defineProperty(AlignmentPattern.prototype, "X", {
  get: function() {
    return Math.floor(this.x);
  }
});

Object.defineProperty(AlignmentPattern.prototype, "Y", {
  get: function() {
    return Math.floor(this.y);
  }
});

AlignmentPattern.prototype.incrementCount = function() {
  this.count++;
};

AlignmentPattern.prototype.aboutEquals = function(moduleSize,  i,  j) {
  if (Math.abs(i - this.y) <= moduleSize && Math.abs(j - this.x) <= moduleSize) {
    var moduleSizeDiff = Math.abs(moduleSize - this.estimatedModuleSize);
    return moduleSizeDiff <= 1.0 || moduleSizeDiff / this.estimatedModuleSize <= 1.0;
  }
  return false;
};

export function AlignmentPatternFinder(image,  startX,  startY,  width,  height,  moduleSize,  resultPointCallback) {
  this.image = image;
  this.possibleCenters = [];
  this.startX = startX;
  this.startY = startY;
  this.width = width;
  this.height = height;
  this.moduleSize = moduleSize;
  this.crossCheckStateCount = [0, 0, 0];
  this.resultPointCallback = resultPointCallback;
}

AlignmentPatternFinder.prototype.centerFromEnd = function(stateCount,  end) {
  return  (end - stateCount[2]) - stateCount[1] / 2.0;
};

AlignmentPatternFinder.prototype.foundPatternCross = function(stateCount) {
  var moduleSize = this.moduleSize;
  var maxVariance = moduleSize / 2.0;
  for (var i = 0; i < 3; i++) {
    if (Math.abs(moduleSize - stateCount[i]) >= maxVariance) {
      return false;
    }
  }
  return true;
};

AlignmentPatternFinder.prototype.crossCheckVertical = function(startI,  centerJ,  maxCount,  originalStateCountTotal) {
  var image = this.image;

  var maxI = image.height;
  var stateCount = this.crossCheckStateCount;
  stateCount[0] = 0;
  stateCount[1] = 0;
  stateCount[2] = 0;

  // Start counting up from center
  var i = startI;
  while (i >= 0 && image.data[centerJ + i * image.width] && stateCount[1] <= maxCount) {
    stateCount[1]++;
    i--;
  }
  // If already too many modules in this state or ran off the edge:
  if (i < 0 || stateCount[1] > maxCount) {
    return NaN;
  }
  while (i >= 0 && !image.data[centerJ + i * image.width] && stateCount[0] <= maxCount) {
    stateCount[0]++;
    i--;
  }
  if (stateCount[0] > maxCount) {
    return NaN;
  }

  // Now also count down from center
  i = startI + 1;
  while (i < maxI && image.data[centerJ + i * image.width] && stateCount[1] <= maxCount) {
    stateCount[1]++;
    i++;
  }
  if (i == maxI || stateCount[1] > maxCount) {
    return NaN;
  }
  while (i < maxI && !image.data[centerJ + i * image.width] && stateCount[2] <= maxCount) {
    stateCount[2]++;
    i++;
  }
  if (stateCount[2] > maxCount) {
    return NaN;
  }

  var stateCountTotal = stateCount[0] + stateCount[1] + stateCount[2];
  if (5 * Math.abs(stateCountTotal - originalStateCountTotal) >= 2 * originalStateCountTotal) {
    return NaN;
  }

  return this.foundPatternCross(stateCount) ? this.centerFromEnd(stateCount, i) : NaN;
};

AlignmentPatternFinder.prototype.handlePossibleCenter = function(stateCount,  i,  j) {
  var stateCountTotal = stateCount[0] + stateCount[1] + stateCount[2];
  var centerJ = this.centerFromEnd(stateCount, j);
  var centerI = this.crossCheckVertical(i, Math.floor(centerJ), 2 * stateCount[1], stateCountTotal);
  if (!isNaN(centerI)) {
    var estimatedModuleSize = (stateCount[0] + stateCount[1] + stateCount[2]) / 3.0;
    var max = this.possibleCenters.length;
    for (var index = 0; index < max; index++) {
      var center =  this.possibleCenters[index];
      // Look for about the same center and module size:
      if (center.aboutEquals(estimatedModuleSize, centerI, centerJ)) {
        return new AlignmentPattern(centerJ, centerI, estimatedModuleSize);
      }
    }
    // Hadn't found this before; save it
    var point = new AlignmentPattern(centerJ, centerI, estimatedModuleSize);
    this.possibleCenters.push(point);
    if (this.resultPointCallback != null) {
      this.resultPointCallback.foundPossibleResultPoint(point);
    }
  }
  return null;
};

AlignmentPatternFinder.prototype.find = function() {
  var image = this.image;
  var startX = this.startX;
  var height = this.height;
  var maxJ = startX + this.width;
  var middleI = this.startY + (height >> 1);
  // We are looking for black/white/black modules in 1:1:1 ratio;
  // this tracks the number of black/white/black modules seen so far
  var stateCount = [0, 0, 0];
  for (var iGen = 0; iGen < height; iGen++) {
    // Search from middle outwards
    var i = middleI + ((iGen & 0x01) == 0 ? ((iGen + 1) >> 1) : -((iGen + 1) >> 1));
    stateCount[0] = 0;
    stateCount[1] = 0;
    stateCount[2] = 0;
    var j = startX;
    // Burn off leading white pixels before anything else; if we start in the middle of
    // a white run, it doesn't make sense to count its length, since we don't know if the
    // white run continued to the left of the start point
    while (j < maxJ && !image.data[j + image.width * i]) {
      j++;
    }
    var currentState = 0;
    while (j < maxJ) {
      if (image.data[j + i * image.width]) {
        // Black pixel
        if (currentState == 1) {
          // Counting black pixels
          stateCount[currentState]++;
        } else {
          // Counting white pixels
          if (currentState == 2) {
            // A winner?
            if (this.foundPatternCross(stateCount)) {
              // Yes
              var confirmed = this.handlePossibleCenter(stateCount, i, j);
              if (confirmed != null) {
                return confirmed;
              }
            }
            stateCount[0] = stateCount[2];
            stateCount[1] = 1;
            stateCount[2] = 0;
            currentState = 1;
          } else {
            stateCount[++currentState]++;
          }
        }
      } else {
        // White pixel
        if (currentState == 1) {
          // Counting black pixels
          currentState++;
        }
        stateCount[currentState]++;
      }
      j++;
    }
    if (this.foundPatternCross(stateCount)) {
      var confirmed = this.handlePossibleCenter(stateCount, i, maxJ);
      if (confirmed != null) {
        return confirmed;
      }
    }
  }

  // Hmm, nothing we saw was observed and confirmed twice. If we had
  // any guess at all, return it.
  if (!(this.possibleCenters.length == 0)) {
    return  this.possibleCenters[0];
  }

  throw "Couldn't find enough alignment patterns";
};
