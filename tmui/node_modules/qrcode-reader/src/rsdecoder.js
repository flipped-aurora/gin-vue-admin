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

import GF256 from './gf256';
import GF256Poly from './gf256poly';

export default function ReedSolomonDecoder(field) {
  this.field = field;
}

ReedSolomonDecoder.prototype.decode = function(received, twoS) {
  var poly = new GF256Poly(this.field, received);
  var syndromeCoefficients = new Array(twoS);
  for (var i = 0; i < syndromeCoefficients.length; i++)syndromeCoefficients[i] = 0;
  var dataMatrix = false;//this.field.Equals(GF256.DATA_MATRIX_FIELD);
  var noError = true;
  for (var i = 0; i < twoS; i++) {
    // Thanks to sanfordsquires for this fix:
    var _eval = poly.evaluateAt(this.field.exp(dataMatrix ? i + 1 : i));
    syndromeCoefficients[syndromeCoefficients.length - 1 - i] = _eval;
    if (_eval != 0) {
      noError = false;
    }
  }
  if (noError) {
    return;
  }
  var syndrome = new GF256Poly(this.field, syndromeCoefficients);
  var sigmaOmega = this.runEuclideanAlgorithm(this.field.buildMonomial(twoS, 1), syndrome, twoS);
  var sigma = sigmaOmega[0];
  var omega = sigmaOmega[1];
  var errorLocations = this.findErrorLocations(sigma);
  var errorMagnitudes = this.findErrorMagnitudes(omega, errorLocations, dataMatrix);
  for (var i = 0; i < errorLocations.length; i++) {
    var position = received.length - 1 - this.field.log(errorLocations[i]);
    if (position < 0) {
      throw "ReedSolomonException Bad error location";
    }
    received[position] = GF256.prototype.addOrSubtract(received[position], errorMagnitudes[i]);
  }
};

ReedSolomonDecoder.prototype.runEuclideanAlgorithm = function(a,  b,  R) {
  // Assume a's degree is >= b's
  if (a.Degree < b.Degree) {
    var temp = a;
    a = b;
    b = temp;
  }

  var rLast = a;
  var r = b;
  var sLast = this.field.One;
  var s = this.field.Zero;
  var tLast = this.field.Zero;
  var t = this.field.One;

  // Run Euclidean algorithm until r's degree is less than R/2
  while (r.Degree >= Math.floor(R / 2)) {
    var rLastLast = rLast;
    var sLastLast = sLast;
    var tLastLast = tLast;
    rLast = r;
    sLast = s;
    tLast = t;

    // Divide rLastLast by rLast, with quotient in q and remainder in r
    if (rLast.Zero) {
      // Oops, Euclidean algorithm already terminated?
      throw "r_{i-1} was zero";
    }
    r = rLastLast;
    var q = this.field.Zero;
    var denominatorLeadingTerm = rLast.getCoefficient(rLast.Degree);
    var dltInverse = this.field.inverse(denominatorLeadingTerm);
    while (r.Degree >= rLast.Degree && !r.Zero) {
      var degreeDiff = r.Degree - rLast.Degree;
      var scale = this.field.multiply(r.getCoefficient(r.Degree), dltInverse);
      q = q.addOrSubtract(this.field.buildMonomial(degreeDiff, scale));
      r = r.addOrSubtract(rLast.multiplyByMonomial(degreeDiff, scale));
    }

    s = q.multiply1(sLast).addOrSubtract(sLastLast);
    t = q.multiply1(tLast).addOrSubtract(tLastLast);
  }

  var sigmaTildeAtZero = t.getCoefficient(0);
  if (sigmaTildeAtZero == 0) {
    throw "ReedSolomonException sigmaTilde(0) was zero";
  }

  var inverse = this.field.inverse(sigmaTildeAtZero);
  var sigma = t.multiply2(inverse);
  var omega = r.multiply2(inverse);
  return [sigma, omega];
};

ReedSolomonDecoder.prototype.findErrorLocations = function(errorLocator) {
  // This is a direct application of Chien's search
  var numErrors = errorLocator.Degree;
  if (numErrors == 1) {
    // shortcut
    return new Array(errorLocator.getCoefficient(1));
  }
  var result = new Array(numErrors);
  var e = 0;
  for (var i = 1; i < 256 && e < numErrors; i++) {
    if (errorLocator.evaluateAt(i) == 0) {
      result[e] = this.field.inverse(i);
      e++;
    }
  }
  if (e != numErrors) {
    throw "Error locator degree does not match number of roots";
  }
  return result;
};

ReedSolomonDecoder.prototype.findErrorMagnitudes = function(errorEvaluator, errorLocations, dataMatrix) {
  // This is directly applying Forney's Formula
  var s = errorLocations.length;
  var result = new Array(s);
  for (var i = 0; i < s; i++) {
    var xiInverse = this.field.inverse(errorLocations[i]);
    var denominator = 1;
    for (var j = 0; j < s; j++) {
      if (i != j) {
        denominator = this.field.multiply(denominator, GF256.prototype.addOrSubtract(1, this.field.multiply(errorLocations[j], xiInverse)));
      }
    }
    result[i] = this.field.multiply(errorEvaluator.evaluateAt(xiInverse), this.field.inverse(denominator));
    // Thanks to sanfordsquires for this fix:
    if (dataMatrix) {
      result[i] = this.field.multiply(result[i], xiInverse);
    }
  }
  return result;
};
