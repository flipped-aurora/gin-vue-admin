/////////////////////////////////////////////////////////////////////////////////////////////////
//
//  Tencent is pleased to support the open source community by making libpag available.
//
//  Copyright (C) 2021 THL A29 Limited, a Tencent company. All rights reserved.
//
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  unless required by applicable law or agreed to in writing, software distributed under the
//  license is distributed on an "as is" basis, without warranties or conditions of any kind,
//  either express or implied. see the license for the specific language governing permissions
//  and limitations under the license.
//
/////////////////////////////////////////////////////////////////////////////////////////////////

let getTime;
try {
  getTime = performance.now.bind(performance);
} catch (e) {
  getTime = Date.now;
}
class Clock {
  constructor() {
    this.startTime = getTime();
    this.markers = {};
  }
  reset() {
    this.startTime = getTime();
    this.markers = {};
  }
  mark(key) {
    if (!key) {
      console.log("Clock.mark(): An empty marker name was specified!");
      return;
    }
    if (Object.keys(this.markers).find((markerKey) => markerKey === key)) {
      console.log(`Clock.mark(): The specified marker name '${key}' already exists!`);
      return;
    }
    this.markers[key] = getTime();
  }
  measure(makerFrom, makerTo) {
    let start;
    let end;
    if (!makerFrom) {
      start = this.startTime;
    } else {
      if (!Object.keys(this.markers).find((markerKey) => markerKey === makerFrom)) {
        console.log(`Clock.measure(): The specified makerFrom '${makerFrom}' does not exist!`);
        return 0;
      }
      start = this.markers[makerFrom];
    }
    if (!makerTo) {
      end = getTime();
    } else {
      if (!Object.keys(this.markers).find((markerKey) => markerKey === makerTo)) {
        console.log(`Clock.measure(): The specified makerTo '${makerTo}' does not exist!`);
        return 0;
      }
      end = this.markers[makerTo];
    }
    return end - start;
  }
}

var CompositionType = /* @__PURE__ */ ((CompositionType2) => {
  CompositionType2[CompositionType2["Unknown"] = 0] = "Unknown";
  CompositionType2[CompositionType2["Vector"] = 1] = "Vector";
  CompositionType2[CompositionType2["Bitmap"] = 2] = "Bitmap";
  CompositionType2[CompositionType2["Video"] = 3] = "Video";
  return CompositionType2;
})(CompositionType || {});
var TagCode = /* @__PURE__ */ ((TagCode2) => {
  TagCode2[TagCode2["End"] = 0] = "End";
  TagCode2[TagCode2["FontTables"] = 1] = "FontTables";
  TagCode2[TagCode2["VectorCompositionBlock"] = 2] = "VectorCompositionBlock";
  TagCode2[TagCode2["CompositionAttributes"] = 3] = "CompositionAttributes";
  TagCode2[TagCode2["ImageTables"] = 4] = "ImageTables";
  TagCode2[TagCode2["LayerBlock"] = 5] = "LayerBlock";
  TagCode2[TagCode2["LayerAttributes"] = 6] = "LayerAttributes";
  TagCode2[TagCode2["SolidColor"] = 7] = "SolidColor";
  TagCode2[TagCode2["TextSource"] = 8] = "TextSource";
  TagCode2[TagCode2["TextPathOption"] = 9] = "TextPathOption";
  TagCode2[TagCode2["TextMoreOption"] = 10] = "TextMoreOption";
  TagCode2[TagCode2["ImageReference"] = 11] = "ImageReference";
  TagCode2[TagCode2["CompositionReference"] = 12] = "CompositionReference";
  TagCode2[TagCode2["Transform2D"] = 13] = "Transform2D";
  TagCode2[TagCode2["MaskBlock"] = 14] = "MaskBlock";
  TagCode2[TagCode2["ShapeGroup"] = 15] = "ShapeGroup";
  TagCode2[TagCode2["Rectangle"] = 16] = "Rectangle";
  TagCode2[TagCode2["Ellipse"] = 17] = "Ellipse";
  TagCode2[TagCode2["PolyStar"] = 18] = "PolyStar";
  TagCode2[TagCode2["ShapePath"] = 19] = "ShapePath";
  TagCode2[TagCode2["Fill"] = 20] = "Fill";
  TagCode2[TagCode2["Stroke"] = 21] = "Stroke";
  TagCode2[TagCode2["GradientFill"] = 22] = "GradientFill";
  TagCode2[TagCode2["GradientStroke"] = 23] = "GradientStroke";
  TagCode2[TagCode2["MergePaths"] = 24] = "MergePaths";
  TagCode2[TagCode2["TrimPaths"] = 25] = "TrimPaths";
  TagCode2[TagCode2["Repeater"] = 26] = "Repeater";
  TagCode2[TagCode2["RoundCorners"] = 27] = "RoundCorners";
  TagCode2[TagCode2["Performance"] = 28] = "Performance";
  TagCode2[TagCode2["DropShadowStyle"] = 29] = "DropShadowStyle";
  TagCode2[TagCode2["InnerShadowStyle"] = 30] = "InnerShadowStyle";
  TagCode2[TagCode2["OuterGlowStyle"] = 31] = "OuterGlowStyle";
  TagCode2[TagCode2["InnerGlowStyle"] = 32] = "InnerGlowStyle";
  TagCode2[TagCode2["BevelAndEmbossStyle"] = 33] = "BevelAndEmbossStyle";
  TagCode2[TagCode2["SatinStyle"] = 34] = "SatinStyle";
  TagCode2[TagCode2["ColorOverlayStyle"] = 35] = "ColorOverlayStyle";
  TagCode2[TagCode2["GradientOverlayStyle"] = 36] = "GradientOverlayStyle";
  TagCode2[TagCode2["StrokeStyle"] = 37] = "StrokeStyle";
  TagCode2[TagCode2["TintEffect"] = 38] = "TintEffect";
  TagCode2[TagCode2["FillEffect"] = 39] = "FillEffect";
  TagCode2[TagCode2["StrokeEffect"] = 40] = "StrokeEffect";
  TagCode2[TagCode2["TritoneEffect"] = 41] = "TritoneEffect";
  TagCode2[TagCode2["DropShadowEffect"] = 42] = "DropShadowEffect";
  TagCode2[TagCode2["RadialWipeEffect"] = 43] = "RadialWipeEffect";
  TagCode2[TagCode2["DisplacementMapEffect"] = 44] = "DisplacementMapEffect";
  TagCode2[TagCode2["BitmapCompositionBlock"] = 45] = "BitmapCompositionBlock";
  TagCode2[TagCode2["BitmapSequence"] = 46] = "BitmapSequence";
  TagCode2[TagCode2["ImageBytes"] = 47] = "ImageBytes";
  TagCode2[TagCode2["ImageBytes2"] = 48] = "ImageBytes2";
  TagCode2[TagCode2["ImageBytes3"] = 49] = "ImageBytes3";
  TagCode2[TagCode2["VideoCompositionBlock"] = 50] = "VideoCompositionBlock";
  TagCode2[TagCode2["VideoSequence"] = 51] = "VideoSequence";
  TagCode2[TagCode2["LayerAttributesV2"] = 52] = "LayerAttributesV2";
  TagCode2[TagCode2["Count"] = 53] = "Count";
  return TagCode2;
})(TagCode || {});

const readTagHeader = (byteBuffer) => {
  const codeAndLength = byteBuffer.readUint16();
  let length = (codeAndLength & 63) >>> 0;
  const code = codeAndLength >> 6;
  if (length === 63) {
    length = byteBuffer.readUint32();
  }
  if (byteBuffer.context.tagLevel < code) {
    byteBuffer.context.tagLevel = code;
  }
  return { code, length };
};
function readTags(byteArray, parameter, reader) {
  let header = readTagHeader(byteArray);
  while (header.code !== TagCode.End) {
    const tagBytes = byteArray.readBytes(header.length);
    reader(tagBytes, header.code, parameter);
    if (byteArray.context.tagLevel < tagBytes.context.tagLevel) {
      byteArray.context.tagLevel = tagBytes.context.tagLevel;
    }
    header = readTagHeader(byteArray);
  }
}

const ZERO_ID = 0;
const ZERO_TIME = 0;
const OPAQUE = 255;
const WEBGL_CONTEXT_ATTRIBUTES = {
  alpha: true,
  depth: false,
  stencil: false,
  antialias: false
};
var BlendMode = /* @__PURE__ */ ((BlendMode2) => {
  BlendMode2[BlendMode2["Normal"] = 0] = "Normal";
  BlendMode2[BlendMode2["Multiply"] = 1] = "Multiply";
  BlendMode2[BlendMode2["Screen"] = 2] = "Screen";
  BlendMode2[BlendMode2["Overlay"] = 3] = "Overlay";
  BlendMode2[BlendMode2["Darken"] = 4] = "Darken";
  BlendMode2[BlendMode2["Lighten"] = 5] = "Lighten";
  BlendMode2[BlendMode2["ColorDodge"] = 6] = "ColorDodge";
  BlendMode2[BlendMode2["ColorBurn"] = 7] = "ColorBurn";
  BlendMode2[BlendMode2["HardLight"] = 8] = "HardLight";
  BlendMode2[BlendMode2["SoftLight"] = 9] = "SoftLight";
  BlendMode2[BlendMode2["Difference"] = 10] = "Difference";
  BlendMode2[BlendMode2["Exclusion"] = 11] = "Exclusion";
  BlendMode2[BlendMode2["Hue"] = 12] = "Hue";
  BlendMode2[BlendMode2["Saturation"] = 13] = "Saturation";
  BlendMode2[BlendMode2["Color"] = 14] = "Color";
  BlendMode2[BlendMode2["Luminosity"] = 15] = "Luminosity";
  BlendMode2[BlendMode2["DestinationIn"] = 21] = "DestinationIn";
  BlendMode2[BlendMode2["DestinationOut"] = 22] = "DestinationOut";
  BlendMode2[BlendMode2["DestinationATop"] = 23] = "DestinationATop";
  BlendMode2[BlendMode2["SourceIn"] = 24] = "SourceIn";
  BlendMode2[BlendMode2["SourceOut"] = 25] = "SourceOut";
  BlendMode2[BlendMode2["Xor"] = 26] = "Xor";
  return BlendMode2;
})(BlendMode || {});
var KeyframeInterpolationType = /* @__PURE__ */ ((KeyframeInterpolationType2) => {
  KeyframeInterpolationType2[KeyframeInterpolationType2["None"] = 0] = "None";
  KeyframeInterpolationType2[KeyframeInterpolationType2["Linear"] = 1] = "Linear";
  KeyframeInterpolationType2[KeyframeInterpolationType2["Bezier"] = 2] = "Bezier";
  KeyframeInterpolationType2[KeyframeInterpolationType2["Hold"] = 3] = "Hold";
  return KeyframeInterpolationType2;
})(KeyframeInterpolationType || {});
const IS_IOS = navigator && /(ios|ipad|iphone)/.test(navigator.userAgent.toLowerCase());

const Black = { red: 0, green: 0, blue: 0 };
const White = { red: 255, green: 255, blue: 255 };

const verifyFailed = () => {
  console.error("PAG Verify Failed!");
};
const verifyAndReturn = (expression) => {
  if (expression) {
    return true;
  }
  console.error("PAG Verify Failed!");
  return false;
};

const _Composition = class {
  constructor() {
    this.id = ZERO_ID;
    this.width = 0;
    this.height = 0;
    this.duration = ZERO_TIME;
    this.frameRate = 30;
    this.backgroundColor = White;
    this.cacheID = 0;
    this.cacheID = _Composition.cacheIDCount;
    _Composition.cacheIDCount += 1;
  }
  type() {
    return CompositionType.Unknown;
  }
  getStaticTimeRanges() {
    return [];
  }
  verify() {
    return verifyAndReturn(this.width > 0 && this.height > 0 && this.duration > 0 && this.frameRate > 0);
  }
};
let Composition = _Composition;
Composition.cacheIDCount = 1;

class VideoComposition extends Composition {
  constructor() {
    super(...arguments);
    this.hasAlpha = false;
    this.sequences = [];
    this.staticTimeRanges = [];
    this.staticTimeRangeUpdated = false;
  }
  type() {
    return CompositionType.Video;
  }
  getStaticTimeRanges() {
    if (!this.staticTimeRangeUpdated) {
      this.staticTimeRangeUpdated = true;
      this.updateStaticTimeRanges();
    }
    return this.staticTimeRanges;
  }
  updateStaticTimeRanges() {
    if (this.duration <= 1)
      return;
    if (this.sequences.length > 0) {
      let sequence = this.sequences[0];
      for (let i = 1; i < this.sequences.length; i++) {
        const item = this.sequences[i];
        if (item.frameRate > sequence.frameRate)
          sequence = item;
      }
      const timeScale = this.frameRate / sequence.frameRate;
      for (const timeRange of sequence.staticTimeRanges) {
        timeRange.start = Math.round(timeRange.start * timeScale);
        timeRange.end = Math.round(timeRange.end * timeScale);
        this.staticTimeRanges.push(timeRange);
      }
    } else {
      const range = { start: 0, end: this.duration - 1 };
      this.staticTimeRanges.push(range);
    }
  }
  hasImageContent() {
    return true;
  }
  verify() {
    if (!super.verify() || this.sequences.length <= 0) {
      verifyFailed();
      return false;
    }
    for (const sequence of this.sequences) {
      if (!sequence || !sequence.verify()) {
        verifyFailed();
        return false;
      }
    }
    return true;
  }
}

class Ratio {
  constructor(numerator, denominator) {
    this.numerator = 1;
    this.denominator = 1;
    this.numerator = numerator;
    this.denominator = denominator;
  }
  value() {
    return this.numerator / this.denominator;
  }
}
const DefaultRatio = new Ratio(1, 1);

class Point {
  constructor(x, y) {
    this.x = x;
    this.y = y;
  }
}
const ZERO_POINT = new Point(0, 0);

var TrackMatteType = /* @__PURE__ */ ((TrackMatteType2) => {
  TrackMatteType2[TrackMatteType2["None"] = 0] = "None";
  TrackMatteType2[TrackMatteType2["Alpha"] = 1] = "Alpha";
  TrackMatteType2[TrackMatteType2["AlphaInverted"] = 2] = "AlphaInverted";
  TrackMatteType2[TrackMatteType2["Luma"] = 3] = "Luma";
  TrackMatteType2[TrackMatteType2["LumaInverted"] = 4] = "LumaInverted";
  return TrackMatteType2;
})(TrackMatteType || {});
var LayerType = /* @__PURE__ */ ((LayerType2) => {
  LayerType2[LayerType2["Unknown"] = 0] = "Unknown";
  LayerType2[LayerType2["undefined"] = 1] = "undefined";
  LayerType2[LayerType2["Solid"] = 2] = "Solid";
  LayerType2[LayerType2["Text"] = 3] = "Text";
  LayerType2[LayerType2["Shape"] = 4] = "Shape";
  LayerType2[LayerType2["Image"] = 5] = "Image";
  LayerType2[LayerType2["PreCompose"] = 6] = "PreCompose";
  return LayerType2;
})(LayerType || {});
class Layer {
  constructor() {
    this.id = 0;
    this.parent = void 0;
    this.containingComposition = void 0;
    this.stretch = DefaultRatio;
    this.startTime = ZERO_ID;
    this.duration = ZERO_TIME;
    this.autoOrientation = false;
    this.transform = void 0;
    this.isActive = true;
    this.blendMode = BlendMode.Normal;
    this.trackMatteType = 0 /* None */;
    this.trackMatteLayer = void 0;
    this.timeRemap = void 0;
    this.masks = void 0;
    this.effects = void 0;
    this.layerStyles = void 0;
    this.layerCache = void 0;
    this.maxScale = void 0;
  }
  type() {
    return 0 /* Unknown */;
  }
  excludeVaryingRanges(timeRanges) {
    var _a;
    (_a = this.transform) == null ? void 0 : _a.excludeVaryingRanges(timeRanges);
    if (this.timeRemap !== void 0) {
      this.timeRemap.excludeVaryingRanges(timeRanges);
    }
    if (this.masks !== void 0) {
      for (const mask of this.masks) {
        mask.excludeVaryingRanges(timeRanges);
      }
    }
    if (this.effects !== void 0 && this.effects.length > 0) {
      for (const effect of this.effects) {
        effect.excludeVaryingRanges(timeRanges);
      }
    }
    if (this.layerStyles !== void 0 && this.layerStyles.length > 0) {
      for (const layerStyle of this.layerStyles) {
        layerStyle.excludeVaryingRanges(timeRanges);
      }
    }
  }
  gotoFrame(frame) {
    var _a;
    (_a = this.transform) == null ? void 0 : _a.gotoFrame(frame);
    if (this.timeRemap !== void 0) {
      this.timeRemap.gotoFrame(frame);
    }
    if (this.masks !== void 0 && this.masks.length > 0) {
      for (const mask of this.masks) {
        mask.gotoFrame(frame);
      }
    }
    if (this.effects !== void 0 && this.effects.length > 0) {
      for (const effect of this.effects) {
        effect.gotoFrame(frame);
      }
    }
    if (this.layerStyles !== void 0 && this.layerStyles.length > 0) {
      for (const layerStyle of this.layerStyles) {
        layerStyle.gotoFrame(frame);
      }
    }
  }
  verify() {
    if (!this.containingComposition || this.duration <= 0 || !this.transform) {
      verifyFailed();
      return false;
    }
    if (!this.transform.verify()) {
      verifyFailed();
      return false;
    }
    if (this.masks && this.masks.length > 0) {
      for (const mask of this.masks) {
        if (!mask || !mask.verify()) {
          verifyFailed();
          return false;
        }
      }
    }
    if (this.layerStyles && this.layerStyles.length > 0) {
      for (const layerStyle of this.layerStyles) {
        if (!layerStyle || !layerStyle.verify()) {
          verifyFailed();
          return false;
        }
      }
    }
    if (this.effects && this.effects.length > 0) {
      for (const effect of this.effects) {
        if (!effect || !effect.verify()) {
          verifyFailed();
          return false;
        }
      }
    }
    return true;
  }
  getMaxScaleFactor() {
    if (this.maxScale !== void 0) {
      return this.maxScale;
    }
    this.maxScale = new Point(1, 1);
    const property = this.transform.scale;
    if (property.animatable()) {
      const { keyframes } = property;
      let scaleX = Math.abs(keyframes[0].startValue.x);
      let scaleY = Math.abs(keyframes[0].startValue.y);
      if (keyframes !== void 0 && keyframes.length > 0) {
        for (const keyframe of keyframes) {
          const x = Math.abs(keyframe.endValue.x);
          const y = Math.abs(keyframe.endValue.y);
          if (scaleX < x) {
            scaleX = x;
          }
          if (scaleY < y) {
            scaleY = y;
          }
        }
      }
      this.maxScale.x = scaleX;
      this.maxScale.y = scaleY;
    } else {
      this.maxScale.x = Math.abs(property.value.x);
      this.maxScale.y = Math.abs(property.value.y);
    }
    if (this.parent !== void 0) {
      const parentScale = this.parent.getMaxScaleFactor();
      this.maxScale.x *= parentScale.x;
      this.maxScale.y *= parentScale.y;
    }
    return this.maxScale;
  }
}

const SPATIAL_PRECISION = 0.05;
const BEZIER_PRECISION = 5e-3;
const readRatio = (byteArray) => {
  const numeratorValue = byteArray.readEncodeInt32();
  const denominatorValue = byteArray.readEncodedUint32();
  const ration = new Ratio(numeratorValue, denominatorValue);
  return ration;
};
const readColor = (byteArray) => {
  const redNum = byteArray.readUint8();
  const greenNum = byteArray.readUint8();
  const blueNum = byteArray.readUint8();
  const color = { red: redNum, green: greenNum, blue: blueNum };
  return color;
};
const readTime = (byteArray) => byteArray.readEncodedUint64();
const readLayerID = (byteArray) => {
  const id = byteArray.readEncodedUint32();
  if (id === 0)
    throw new Error("Layer ID is 0");
  const layer = new Layer();
  layer.id = id;
  return layer;
};
const readPoint = (byteArray) => {
  const x = byteArray.readFloat32();
  const y = byteArray.readFloat32();
  return new Point(x, y);
};

const readCompositionAttributes = (byteArray, composition) => {
  composition.width = byteArray.readEncodeInt32();
  composition.height = byteArray.readEncodeInt32();
  composition.duration = readTime(byteArray);
  composition.frameRate = byteArray.readFloat32();
  composition.backgroundColor = readColor(byteArray);
};

class Context$1 {
  constructor() {
    this.tagLevel = 0;
    this.compositions = [];
    this.errorMessages = [];
  }
  throwException(message) {
    this.errorMessages.push(message);
  }
  releaseCompositions() {
    const compositions = this.compositions.slice();
    this.compositions = [];
    return compositions;
  }
}

const ErrorMessage = {
  PAGDecodeError: "PAG file decode error!"
};

const LENGTH_FOR_STORE_NUM_BITS = 5;
class ByteArray {
  constructor(buffer, littleEndian) {
    this._position = 0;
    this.bitPosition = 0;
    this.dataView = new DataView(buffer);
    this.littleEndian = !!littleEndian;
    this.context = new Context$1();
  }
  get length() {
    return this.dataView.byteLength;
  }
  get bytesAvailable() {
    return this.dataView.byteLength - this._position;
  }
  data() {
    return this.dataView.buffer;
  }
  get position() {
    return this._position;
  }
  alignWithBytes() {
    this.bitPosition = this._position * 8;
  }
  readBoolean() {
    const value = this.dataView.getInt8(this._position);
    this._position += 1;
    this.positonChanged();
    return Boolean(value);
  }
  readChar() {
    if (this._position >= this.length)
      throw new Error(ErrorMessage.PAGDecodeError);
    const value = this.dataView.getInt8(this._position);
    this._position += 1;
    this.positonChanged();
    return String.fromCharCode(value);
  }
  readUint8() {
    if (this._position >= this.length)
      throw new Error(ErrorMessage.PAGDecodeError);
    const value = this.dataView.getUint8(this._position);
    this._position += 1;
    this.positonChanged();
    return value;
  }
  readInt8() {
    if (this._position >= this.length)
      throw new Error(ErrorMessage.PAGDecodeError);
    const value = this.dataView.getInt8(this._position);
    this._position += 1;
    this.positonChanged();
    return value;
  }
  readInt16() {
    if (this._position >= this.length - 1)
      throw new Error(ErrorMessage.PAGDecodeError);
    const value = this.dataView.getInt16(this._position, this.littleEndian);
    this._position += 2;
    this.positonChanged();
    return value;
  }
  readUint16() {
    if (this._position >= this.length - 1)
      throw new Error(ErrorMessage.PAGDecodeError);
    const value = this.dataView.getUint16(this._position, this.littleEndian);
    this._position += 2;
    this.positonChanged();
    return value;
  }
  readInt24() {
    if (this._position >= this.length - 2)
      throw new Error(ErrorMessage.PAGDecodeError);
    const left = this.dataView.getInt16(this._position, this.littleEndian);
    const right = this.dataView.getInt8(this._position + 2);
    this._position += 3;
    this.positonChanged();
    return this.littleEndian ? left + 2 ** 16 * right : 2 ** 16 * left + right;
  }
  readUint24() {
    if (this._position >= this.length - 2)
      throw new Error(ErrorMessage.PAGDecodeError);
    const left = this.dataView.getUint16(this._position, this.littleEndian);
    const right = this.dataView.getUint8(this._position + 2);
    this._position += 3;
    this.positonChanged();
    return this.littleEndian ? left + 2 ** 16 * right : 2 ** 16 * left + right;
  }
  readInt32() {
    if (this._position >= this.length - 3)
      throw new Error(ErrorMessage.PAGDecodeError);
    const value = this.dataView.getInt32(this._position, this.littleEndian);
    this._position += 4;
    this.positonChanged();
    return value;
  }
  readUint32() {
    if (this._position >= this.length - 3)
      throw new Error(ErrorMessage.PAGDecodeError);
    const value = this.dataView.getUint32(this._position, this.littleEndian);
    this._position += 4;
    this.positonChanged();
    return value;
  }
  readInt64() {
    if (this._position >= this.length - 7)
      throw new Error(ErrorMessage.PAGDecodeError);
    const left = this.dataView.getInt32(this._position, this.littleEndian);
    const right = this.dataView.getInt32(this._position + 4, this.littleEndian);
    this._position += 8;
    this.positonChanged();
    return this.littleEndian ? left + 2 ** 32 * right : 2 ** 32 * left + right;
  }
  readUint64() {
    if (this._position >= this.length - 7)
      throw new Error(ErrorMessage.PAGDecodeError);
    const left = this.dataView.getUint32(this._position, this.littleEndian);
    const right = this.dataView.getUint32(this._position + 4, this.littleEndian);
    this._position += 8;
    this.positonChanged();
    return this.littleEndian ? left + 2 ** 32 * right : 2 ** 32 * left + right;
  }
  readFloat32() {
    if (this._position >= this.length - 3)
      throw new Error(ErrorMessage.PAGDecodeError);
    const value = this.dataView.getFloat32(this._position, this.littleEndian);
    this._position += 4;
    this.positonChanged();
    return value;
  }
  readDouble() {
    if (this._position >= this.length - 7)
      throw new Error(ErrorMessage.PAGDecodeError);
    const value = this.dataView.getFloat64(this._position, this.littleEndian);
    this._position += 8;
    this.positonChanged();
    return value;
  }
  readUTF8String() {
    if (this._position >= this.length)
      throw new Error(ErrorMessage.PAGDecodeError);
    let encoded = "";
    let dataLength = 0;
    for (let i = this._position; i < this.length; i++) {
      if (this.dataView.getUint8(i) === 0) {
        break;
      }
      encoded += `%${this.dataView.getUint8(i).toString(16)}`;
      dataLength += 1;
    }
    this._position += dataLength;
    this.positonChanged();
    return decodeURIComponent(encoded);
  }
  readEncodedUint32() {
    const valueMask = 127;
    const hasNext = 128;
    let value = 0;
    let byte = 0;
    for (let i = 0; i < 32; i += 7) {
      if (this._position >= this.length) {
        throw Error("readEncodedUint32 End of file was encountered.");
      }
      byte = this.dataView.getUint8(this._position);
      this._position += 1;
      value |= (byte & valueMask) << i;
      if ((byte & hasNext) === 0) {
        break;
      }
    }
    this.positonChanged();
    return value;
  }
  readEncodeInt32() {
    const data = this.readEncodedUint32();
    const value = data >> 1;
    return (data & 1) > 0 ? -value : value;
  }
  readEncodedUint64() {
    const valueMask = 127;
    const hasNext = 128;
    let value = 0;
    let byte = 0;
    for (let i = 0; i < 64; i += 7) {
      if (this._position >= this.length) {
        throw Error("readEncodedUint64 End of file was encountered.");
      }
      byte = this.dataView.getUint8(this._position);
      this._position += 1;
      value |= (byte & valueMask) << i;
      if ((byte & hasNext) === 0) {
        break;
      }
    }
    this.positonChanged();
    return value;
  }
  readEncodeInt64() {
    const data = this.readEncodedUint64();
    const value = data << 0;
    return (data & 1) > 0 ? -value : value;
  }
  readBytes(length) {
    const len = length || this.length - this._position;
    if (this._position > this.length - len)
      throw new Error(ErrorMessage.PAGDecodeError);
    const newBuffer = this.dataView.buffer.slice(this._position, this._position + len);
    this._position += len;
    this.positonChanged();
    return new ByteArray(newBuffer, this.littleEndian);
  }
  readUBits(numBits) {
    const bitMasks = [0, 1, 3, 7, 15, 31, 63, 127, 255];
    let value = 0;
    if (this.bitPosition > this.length * 8 - numBits)
      throw new Error(ErrorMessage.PAGDecodeError);
    let pos = 0;
    while (pos < numBits) {
      const bytePosition = Math.floor(this.bitPosition * 0.125);
      const bitPosition = this.bitPosition % 8;
      let byte = this.dataView.getUint8(bytePosition) >> bitPosition;
      const bitLength = Math.min(8 - bitPosition, numBits - pos);
      byte &= bitMasks[bitLength];
      value |= byte << pos;
      pos += bitLength;
      this.bitPosition += bitLength;
    }
    this.bitPositionChanged();
    return value;
  }
  readBits(numBits) {
    let value = this.readUBits(numBits);
    value <<= 32 - numBits;
    const data = value << 0;
    return data >> 32 - numBits;
  }
  readNumBits() {
    return this.readUBits(LENGTH_FOR_STORE_NUM_BITS) + 1;
  }
  readInt32List(count) {
    const numBits = this.readNumBits();
    const value = new Array(count);
    for (let i = 0; i < count; i++) {
      value[i] = this.readBits(numBits);
    }
    return value;
  }
  readUint32List(count) {
    const numBits = this.readNumBits();
    const value = new Array(count);
    for (let i = 0; i < count; i++) {
      value[i] = this.readUBits(numBits);
    }
    return value;
  }
  readBitBoolean() {
    return this.readUBits(1) !== 0;
  }
  readFloatList(count, precision) {
    const numBits = this.readNumBits();
    const value = new Array(count);
    for (let i = 0; i < count; i++) {
      value[i] = this.readBits(numBits) * precision;
    }
    return value;
  }
  bitPositionChanged() {
    this._position = Math.ceil(this.bitPosition * 0.125);
  }
  positonChanged() {
    this.bitPosition = this._position * 8;
  }
}

class ByteData {
  constructor(data, length) {
    this.length = 0;
    this.data = data;
    this.length = length;
  }
}

class VideoFrame {
  constructor() {
    this.isKeyframe = false;
    this.frame = 0;
    this.fileBytes = new ByteData(new ByteArray(new ArrayBuffer(0)), 0);
  }
}

class Sequence {
  constructor() {
    this.composition = void 0;
    this.id = 0;
    this.width = 0;
    this.height = 0;
    this.frameRate = 0;
    this.frameCount = 0;
    this.isKeyFrameFlags = [];
  }
  verify() {
    return verifyAndReturn(this.composition !== void 0 && this.width > 0 && this.height > 0 && this.frameRate > 0);
  }
}

class VideoSequence extends Sequence {
  constructor() {
    super(...arguments);
    this.alphaStartX = 0;
    this.alphaStartY = 0;
    this.frames = [];
    this.headers = [];
    this.staticTimeRanges = [];
  }
  verify() {
    if (!super.verify() || this.frames.length <= 0) {
      verifyFailed();
      return false;
    }
    for (const frame of this.frames) {
      if (!frame || !frame.fileBytes) {
        verifyFailed();
        return false;
      }
    }
    for (const header of this.headers) {
      if (!header) {
        verifyFailed();
        return false;
      }
    }
    return true;
  }
  getVideoWidth() {
    let videoWidth = this.alphaStartX + this.width;
    if (videoWidth % 2 === 1) {
      videoWidth += 1;
    }
    return videoWidth;
  }
  getVideoHeight() {
    let videoHeight = this.alphaStartY + this.height;
    if (videoHeight % 2 === 1) {
      videoHeight += 1;
    }
    return videoHeight;
  }
}

const memcpy = (dst, dstOffset, src, srcOffset, num) => {
  if (dstOffset >= dst.byteLength || srcOffset >= src.byteLength || src.byteLength - srcOffset > dst.byteLength - dstOffset || num > src.byteLength)
    return;
  const dstUint8Array = new Uint8Array(dst);
  const srcUint8Array = new Uint8Array(src, srcOffset, num);
  dstUint8Array.set(srcUint8Array, dstOffset);
};
const concatUint8Arrays = (arrays) => {
  let totalLength = 0;
  for (const arr of arrays) {
    totalLength += arr.byteLength;
  }
  const result = new Uint8Array(totalLength);
  let offset = 0;
  for (const arr of arrays) {
    result.set(arr, offset);
    offset += arr.byteLength;
  }
  return result;
};

const readByteDataWithStartCode = (byteArray) => {
  const length = byteArray.readEncodedUint32();
  const bytes = byteArray.readBytes(length);
  if (length === 0)
    throw new Error("Read start code with length 0!");
  const data = new ArrayBuffer(length + 4);
  const dataView = new DataView(data);
  dataView.setUint32(0, length);
  memcpy(data, 4, bytes.data(), 0, length);
  return new ByteData(new ByteArray(data), length + 4);
};

const readVideoSequence = (byteArray, hasAlpha) => {
  const videoSequence = new VideoSequence();
  videoSequence.width = byteArray.readEncodeInt32();
  videoSequence.height = byteArray.readEncodeInt32();
  videoSequence.frameRate = byteArray.readFloat32();
  if (hasAlpha) {
    videoSequence.alphaStartX = byteArray.readEncodeInt32();
    videoSequence.alphaStartY = byteArray.readEncodeInt32();
  }
  const sps = readByteDataWithStartCode(byteArray);
  const pps = readByteDataWithStartCode(byteArray);
  videoSequence.headers.push(sps, pps);
  videoSequence.frameCount = byteArray.readEncodedUint32();
  for (let i = 0; i < videoSequence.frameCount; i++) {
    const videoFrame = new VideoFrame();
    videoFrame.isKeyframe = byteArray.readBitBoolean();
    videoSequence.frames.push(videoFrame);
  }
  for (let i = 0; i < videoSequence.frameCount; i++) {
    const videoFrame = videoSequence.frames[i];
    videoFrame.frame = readTime(byteArray);
    videoFrame.fileBytes = readByteDataWithStartCode(byteArray);
  }
  if (byteArray.bytesAvailable > 0) {
    const count = byteArray.readEncodedUint32();
    for (let i = 0; i < count; i++) {
      const staticTimeRange = { start: 0, end: 0 };
      staticTimeRange.start = readTime(byteArray);
      staticTimeRange.end = readTime(byteArray);
      videoSequence.staticTimeRanges.push(staticTimeRange);
    }
  }
  return videoSequence;
};

const readVideoComposition = (byteArray) => {
  const composition = new VideoComposition();
  composition.id = byteArray.readEncodedUint32();
  composition.hasAlpha = byteArray.readBoolean();
  const parameter = { composition, hasAlpha: composition.hasAlpha };
  readTags(byteArray, parameter, ReadTagsOfVideoComposition);
  return composition;
};
const ReadTagsOfVideoComposition = (byteArray, code, parameter) => {
  const { composition } = parameter;
  switch (code) {
    case TagCode.CompositionAttributes:
      readCompositionAttributes(byteArray, composition);
      break;
    case TagCode.VideoSequence: {
      const sequence = readVideoSequence(byteArray, parameter.hasAlpha);
      sequence.composition = composition;
      composition.sequences.push(sequence);
      break;
    }
  }
};

var EffectType = /* @__PURE__ */ ((EffectType2) => {
  EffectType2[EffectType2["Unknown"] = 0] = "Unknown";
  EffectType2[EffectType2["Tint"] = 1] = "Tint";
  EffectType2[EffectType2["Fill"] = 2] = "Fill";
  EffectType2[EffectType2["Stroke"] = 3] = "Stroke";
  EffectType2[EffectType2["Tritone"] = 4] = "Tritone";
  EffectType2[EffectType2["DropShadow"] = 5] = "DropShadow";
  EffectType2[EffectType2["RadialWipe"] = 6] = "RadialWipe";
  EffectType2[EffectType2["DisplacementMap"] = 7] = "DisplacementMap";
  return EffectType2;
})(EffectType || {});

function subtractFromTimeRanges(timeRanges, startTime, endTime) {
  if (endTime < startTime) {
    return;
  }
  const size = timeRanges.length;
  for (let i = size - 1; i >= 0; i--) {
    const timeRange = timeRanges[i];
    if (timeRange.end < startTime || timeRange.start > endTime) {
      continue;
    }
    if (timeRange.start < startTime && timeRange.end > endTime) {
      const range = { start: endTime + 1, end: timeRange.end };
      timeRange.end = startTime - 1;
      if (range.end > range.start) {
        timeRanges.splice(i + 1, 0, range);
      }
      if (timeRange.end <= timeRange.start) {
        timeRanges.splice(i, 1);
      }
      break;
    }
    if (timeRange.start >= startTime && timeRange.end <= endTime) {
      timeRanges.splice(i, 1);
    } else if (timeRange.start < startTime) {
      timeRange.end = startTime - 1;
      if (timeRange.end <= timeRange.start) {
        timeRanges.splice(i, 1);
      }
    } else {
      timeRange.start = endTime + 1;
      if (timeRange.end <= timeRange.start) {
        timeRanges.splice(i, 1);
      }
    }
  }
}
function splitTimeRangesAt(timeRanges, startTime) {
  const size = timeRanges.length;
  for (let i = size - 1; i >= 0; i--) {
    const timeRange = timeRanges[i];
    if (timeRange.start === startTime || timeRange.end <= startTime) {
      break;
    }
    if (timeRange.start < startTime && timeRange.end > startTime) {
      const range = { start: startTime, end: timeRange.end };
      timeRange.end = startTime - 1;
      if (range.end > range.start) {
        timeRanges.splice(i + 1, 0, range);
      }
      if (timeRange.end <= timeRange.start) {
        timeRanges.splice(i, 1);
      }
      break;
    }
  }
}

class VectorComposition extends Composition {
  constructor() {
    super(...arguments);
    this.layers = [];
    this.staticTimeRanges = [];
    this.staticTimeRangeUpdated = false;
  }
  type() {
    return CompositionType.Vector;
  }
  getStaticTimeRanges() {
    if (!this.staticTimeRangeUpdated) {
      this.staticTimeRangeUpdated = true;
      this.updateStaticTimeRanges();
    }
    return this.staticTimeRanges;
  }
  verify() {
    if (!super.verify()) {
      verifyFailed();
      return false;
    }
    for (const layer of this.layers) {
      if (!layer || !layer.verify()) {
        verifyFailed();
        return false;
      }
    }
    return true;
  }
  updateStaticTimeRanges() {
    if (this.duration > 1) {
      const range = { start: 0, end: this.duration - 1 };
      this.staticTimeRanges = [range];
      for (const layer of this.layers) {
        if (this.staticTimeRanges.length <= 0) {
          break;
        }
        layer.excludeVaryingRanges(this.staticTimeRanges);
        splitTimeRangesAt(this.staticTimeRanges, layer.startTime);
        splitTimeRangesAt(this.staticTimeRanges, layer.startTime + layer.duration);
      }
    }
  }
}

class Property {
  constructor(value) {
    this.value = value;
  }
  animatable() {
    return false;
  }
  excludeVaryingRanges(_timeRanges) {
  }
  gotoFrame(_time) {
  }
}

class Transform2D {
  static createDefaultTransform2D() {
    const transform = new Transform2D();
    return transform;
  }
  constructor() {
    this.anchorPoint = new Property(ZERO_POINT);
    this.position = new Property(ZERO_POINT);
    this.xPosition = new Property(0);
    this.yPosition = new Property(0);
    this.scale = new Property(new Point(1, 1));
    this.rotation = new Property(0);
    this.opacity = new Property(OPAQUE);
  }
  excludeVaryingRanges(timeRanges) {
    this.anchorPoint.excludeVaryingRanges(timeRanges);
    if (this.position !== void 0) {
      this.position.excludeVaryingRanges(timeRanges);
    } else {
      this.xPosition.excludeVaryingRanges(timeRanges);
      this.yPosition.excludeVaryingRanges(timeRanges);
    }
    this.scale.excludeVaryingRanges(timeRanges);
    this.rotation.excludeVaryingRanges(timeRanges);
    this.opacity.excludeVaryingRanges(timeRanges);
  }
  gotoFrame(frame) {
    this.anchorPoint.gotoFrame(frame);
    if (this.position !== void 0) {
      this.position.gotoFrame(frame);
    } else {
      this.xPosition.gotoFrame(frame);
      this.yPosition.gotoFrame(frame);
    }
    this.scale.gotoFrame(frame);
    this.rotation.gotoFrame(frame);
    this.opacity.gotoFrame(frame);
  }
  verify() {
    return this.anchorPoint !== void 0 && (this.position !== void 0 || this.xPosition !== void 0 && this.yPosition !== void 0) && this.scale !== void 0 && this.rotation !== void 0 && this.opacity !== void 0;
  }
}

class PreComposeLayer extends Layer {
  constructor() {
    super(...arguments);
    this.composition = void 0;
    this.compositionStartTime = ZERO_TIME;
    this.staticTimeRanges = void 0;
    this.staticTimeRangeUpdated = false;
  }
  static wrap(composition) {
    const layer = new PreComposeLayer();
    layer.duration = composition.duration;
    const transform = new Transform2D();
    layer.transform = transform;
    layer.composition = composition;
    return layer;
  }
  type() {
    return LayerType.PreCompose;
  }
  excludeVaryingRanges(timeRanges) {
    super.excludeVaryingRanges(timeRanges);
    if (!timeRanges || timeRanges.length === 0) {
      return;
    }
    this.updateStaticTimeRanges();
  }
  gotoFrame(frame) {
    super.gotoFrame(frame);
  }
  verify() {
    if (!super.verify()) {
      return false;
    }
    if (this.composition) {
      return true;
    }
    return false;
  }
  updateStaticTimeRanges() {
    var _a;
    if (this.staticTimeRangeUpdated) {
      return;
    }
    this.staticTimeRangeUpdated = true;
    const ranges = (_a = this.composition) == null ? void 0 : _a.getStaticTimeRanges();
    if (!ranges)
      return;
    for (let i = ranges.length - 1; i >= 0; i--) {
      const range = ranges[i];
      range.start += this.compositionStartTime;
      range.end += this.compositionStartTime;
      if (range.end <= this.startTime) {
        ranges.pop();
      } else if (range.start < this.startTime) {
        range.start = 0;
      } else if (range.start >= this.startTime + this.duration - 1) {
        ranges.pop();
      } else if (range.end > this.startTime + this.duration - 1) {
        range.end = this.startTime + this.duration - 1;
      }
    }
    this.staticTimeRanges = ranges;
  }
}

class ShapeLayer extends Layer {
  constructor() {
    super(...arguments);
    this.contents = [];
  }
  type() {
    return LayerType.Shape;
  }
  excludeVaryingRanges(timeRanges) {
    super.excludeVaryingRanges(timeRanges);
    for (const element of this.contents) {
      element.excludeVaryingRanges(timeRanges);
    }
  }
  gotoFrame(frame) {
    super.gotoFrame(frame);
    for (const element of this.contents) {
      element.gotoFrame(frame);
    }
  }
  verify() {
    if (!super.verify()) {
      return false;
    }
    for (const element of this.contents) {
      if (element === void 0 || !element.verify()) {
        return false;
      }
    }
    return true;
  }
}

class SolidLayer extends Layer {
  constructor() {
    super(...arguments);
    this.solidColor = Black;
    this.width = 0;
    this.height = 0;
  }
  type() {
    return LayerType.Solid;
  }
  excludeVaryingRanges(timeRanges) {
    super.excludeVaryingRanges(timeRanges);
  }
  gotoFrame(frame) {
    super.gotoFrame(frame);
  }
  verify() {
    if (!super.verify()) {
      verifyFailed();
      return false;
    }
    return verifyAndReturn(this.width > 0 && this.height > 0);
  }
}

class UnDefinedLayer extends Layer {
  type() {
    return LayerType.undefined;
  }
}

class Keyframe {
  constructor() {
    this.startTime = 0;
    this.endTime = 0;
    this.interpolationType = KeyframeInterpolationType.Hold;
    this.bezierOut = [];
    this.bezierIn = [];
    this.spatialOut = ZERO_POINT;
    this.spatialIn = ZERO_POINT;
  }
  initialize() {
  }
  getValue(_time) {
    return this.startValue;
  }
  containsTime(time) {
    return time >= this.startTime && time < this.endTime;
  }
}

class AnimatableProperty extends Property {
  constructor(keyframes) {
    if (!keyframes || keyframes.length === 0)
      throw new Error("keyframes is required");
    if (keyframes[0].startValue === void 0)
      throw new Error("startValue is required");
    super(keyframes[0].startValue);
    this.keyframes = keyframes;
    this.lastKeyframeIndex = 0;
    for (const keyframe of keyframes) {
      keyframe.initialize();
    }
  }
  animatable() {
    return true;
  }
  excludeVaryingRanges(timeRanges) {
    for (const keyframe of this.keyframes) {
      switch (keyframe.interpolationType) {
        case KeyframeInterpolationType.Bezier:
        case KeyframeInterpolationType.Linear:
          subtractFromTimeRanges(timeRanges, keyframe.startTime, keyframe.endTime - 1);
          break;
        default:
          splitTimeRangesAt(timeRanges, keyframe.startTime);
          splitTimeRangesAt(timeRanges, keyframe.endTime);
          break;
      }
    }
  }
  gotoFrame(frame) {
    let lastKeyframe = this.keyframes[this.lastKeyframeIndex];
    if (lastKeyframe.containsTime(frame)) {
      this.value = lastKeyframe.getValue(frame);
      return;
    }
    if (frame < lastKeyframe.startTime) {
      while (this.lastKeyframeIndex > 0) {
        this.lastKeyframeIndex -= 1;
        if (this.keyframes[this.lastKeyframeIndex].containsTime(frame)) {
          break;
        }
      }
    } else {
      while (this.lastKeyframeIndex < this.keyframes.length - 1) {
        this.lastKeyframeIndex += 1;
        if (this.keyframes[this.lastKeyframeIndex].containsTime(frame)) {
          break;
        }
      }
    }
    lastKeyframe = this.keyframes[this.lastKeyframeIndex];
    if (lastKeyframe.startValue !== void 0 && frame <= lastKeyframe.startTime) {
      this.value = lastKeyframe.startValue;
    } else if (lastKeyframe.endValue !== void 0 && frame >= lastKeyframe.endTime) {
      this.value = lastKeyframe.endValue;
    } else {
      this.value = lastKeyframe.getValue(frame);
    }
  }
}

var AttributeType = /* @__PURE__ */ ((AttributeType2) => {
  AttributeType2[AttributeType2["Value"] = 0] = "Value";
  AttributeType2[AttributeType2["FixedValue"] = 1] = "FixedValue";
  AttributeType2[AttributeType2["SimpleProperty"] = 2] = "SimpleProperty";
  AttributeType2[AttributeType2["DiscreteProperty"] = 3] = "DiscreteProperty";
  AttributeType2[AttributeType2["MultiDimensionProperty"] = 4] = "MultiDimensionProperty";
  AttributeType2[AttributeType2["SpatialProperty"] = 5] = "SpatialProperty";
  AttributeType2[AttributeType2["BitFlag"] = 6] = "BitFlag";
  AttributeType2[AttributeType2["Custom"] = 7] = "Custom";
  return AttributeType2;
})(AttributeType || {});
const readTagBlock = (byteArray, parameter, blockConfig) => {
  const tagConfig = blockConfig;
  const flags = [];
  if (!tagConfig.configs || tagConfig.configs.length === 0) {
    return parameter;
  }
  for (const config of tagConfig.configs) {
    const flag = readAttributeFlag(byteArray, config);
    flags.push(flag);
  }
  byteArray.alignWithBytes();
  let index = 0;
  for (const config of tagConfig.configs) {
    const flag = flags[index];
    const target = config.key;
    config.readAttribute(byteArray, flag, parameter, target);
    index += 1;
  }
  return parameter;
};
class BaseAttribute {
  constructor(key, attributeType, defaultValue) {
    this.attributeType = attributeType;
    this.defaultValue = defaultValue;
    this.key = key;
  }
  readAttribute(_byteArray, _flag, _targetClass, _target) {
  }
  readValue(_byteArray) {
    return void 0;
  }
  readValueList(_byteArray, _list, _count) {
  }
  dimensionality() {
    return 1;
  }
  newKeyframe(_flag) {
    return new Keyframe();
  }
}
const readAttribute = (byteArray, flag, targetClass, target, config) => {
  if (config.attributeType === 6 /* BitFlag */) {
    targetClass[target] = flag.exist;
  } else if (config.attributeType === 1 /* FixedValue */) {
    targetClass[target] = config.readValue(byteArray);
  } else if (config.attributeType === 0 /* Value */) {
    targetClass[target] = readValue(byteArray, config, flag);
  } else {
    targetClass[target] = readProperty(byteArray, config, flag);
  }
};
const readProperty = (byteArray, config, flag) => {
  let property;
  if (flag.exist) {
    if (flag.animatable) {
      const keyframes = readKeyframes(byteArray, config, flag);
      if (!keyframes || keyframes.length === 0) {
        throw new Error("Wrong number of keyframes!");
      }
      readTimeAndValue(byteArray, keyframes, config);
      readTimeEase(byteArray, keyframes, config);
      if (flag.hasSpatial);
      property = new AnimatableProperty(keyframes);
    } else {
      property = new Property(readValue(byteArray, config, flag));
    }
  } else {
    property = new Property(config.defaultValue);
  }
  return property;
};
const readValue = (byteArray, config, flag) => {
  if (flag.exist) {
    return config.readValue(byteArray);
  }
  return config.defaultValue;
};
const readAttributeFlag = (byteArray, config) => {
  const flag = { exist: false, animatable: false, hasSpatial: false };
  const { attributeType } = config;
  if (attributeType === 1 /* FixedValue */) {
    flag.exist = true;
    return flag;
  }
  flag.exist = byteArray.readBitBoolean();
  if (!flag.exist || attributeType === 0 /* Value */ || attributeType === 6 /* BitFlag */ || attributeType === 7 /* Custom */) {
    return flag;
  }
  flag.animatable = byteArray.readBitBoolean();
  if (!flag.animatable || attributeType !== 5 /* SpatialProperty */) {
    return flag;
  }
  flag.hasSpatial = byteArray.readBitBoolean();
  return flag;
};
const readKeyframes = (byteArray, config, flag) => {
  const keyframes = [];
  const numFrames = byteArray.readEncodedUint32();
  for (let i = 0; i < numFrames; i++) {
    let keyframe;
    if (config.attributeType === 3 /* DiscreteProperty */) {
      keyframe = new Keyframe();
    } else {
      const interpolationType = byteArray.readUBits(2);
      if (interpolationType === KeyframeInterpolationType.Hold) {
        keyframe = new Keyframe();
      } else {
        keyframe = config.newKeyframe(flag);
        keyframe.interpolationType = interpolationType;
      }
    }
    keyframes.push(keyframe);
  }
  return keyframes;
};
const readTimeAndValue = (byteArray, keyframes, config) => {
  const numFrames = keyframes.length;
  keyframes[0].startTime = readTime(byteArray);
  for (let i = 0; i < numFrames; i++) {
    const time = readTime(byteArray);
    keyframes[i].endTime = time;
    if (i < numFrames - 1) {
      keyframes[i + 1].startTime = time;
    }
  }
  const list = [];
  config.readValueList(byteArray, list, numFrames + 1);
  let index = 0;
  keyframes[0].startValue = list[index];
  index += 1;
  for (let i = 0; i < numFrames; i++) {
    const value = list[index];
    index += 1;
    keyframes[i].endValue = value;
    if (i < numFrames - 1) {
      keyframes[i + 1].startValue = value;
    }
  }
};
const readTimeEase = (byteArray, keyframes, config) => {
  const dimensionality = config.attributeType === 4 /* MultiDimensionProperty */ ? config.dimensionality() : 1;
  const numBits = byteArray.readNumBits();
  for (const keyframe of keyframes) {
    if (keyframe.interpolationType !== KeyframeInterpolationType.Bezier) {
      continue;
    }
    let x;
    let y;
    for (let i = 0; i < dimensionality; i++) {
      x = byteArray.readBits(numBits) * BEZIER_PRECISION;
      y = byteArray.readBits(numBits) * BEZIER_PRECISION;
      keyframe.bezierOut.push({ x, y });
      x = byteArray.readBits(numBits) * BEZIER_PRECISION;
      y = byteArray.readBits(numBits) * BEZIER_PRECISION;
      keyframe.bezierIn.push({ x, y });
    }
  }
};

function readSolidColor(byteArray, layer) {
  layer.solidColor = readColor(byteArray);
  layer.width = byteArray.readEncodeInt32();
  layer.height = byteArray.readEncodeInt32();
}

function interpolateFloat(a, b, t) {
  return a + (b - a) * t;
}

class Interpolator {
  getInterpolation(input) {
    return input;
  }
}

class MultiDimensionPointKeyframe extends Keyframe {
  initialize() {
    super.initialize();
    if (this.interpolationType === KeyframeInterpolationType.Bezier); else {
      this.xInterpolator = new Interpolator();
      this.yInterpolator = new Interpolator();
    }
  }
  getValue(time) {
    var _a, _b, _c, _d;
    const progress = (time - this.startTime) / (this.endTime - this.startTime);
    const xProgress = (_b = (_a = this.xInterpolator) == null ? void 0 : _a.getInterpolation(progress)) != null ? _b : progress;
    const yProgress = (_d = (_c = this.yInterpolator) == null ? void 0 : _c.getInterpolation(progress)) != null ? _d : progress;
    const x = interpolateFloat(this.startValue.x, this.endValue.x, xProgress);
    const y = interpolateFloat(this.startValue.y, this.endValue.y, yProgress);
    return { x, y };
  }
}

class SingleEaseKeyframe extends Keyframe {
  initialize() {
    if (this.interpolationType === KeyframeInterpolationType.Bezier); else {
      this.interpolator = new Interpolator();
    }
  }
  getProgress(time) {
    var _a, _b;
    const progress = (time - this.startTime) / (this.endTime - this.startTime);
    return (_b = (_a = this.interpolator) == null ? void 0 : _a.getInterpolation(progress)) != null ? _b : progress;
  }
  getValue(time) {
    const progress = this.getProgress(time);
    return interpolateFloat(this.startValue, this.endValue, progress);
  }
}

class FloatAttributeConfig extends BaseAttribute {
  constructor(key, attributeType, defaultValue) {
    super(key, attributeType, defaultValue);
  }
  readAttribute(byteArray, flag, targetClass, target) {
    readAttribute(byteArray, flag, targetClass, target, this);
  }
  readValue(byteArray) {
    return byteArray.readFloat32();
  }
  readValueList(byteArray, list, count) {
    for (let i = 0; i < count; i++) {
      list.push(this.readValue(byteArray));
    }
  }
  dimensionality() {
    return 1;
  }
  newKeyframe(_flag) {
    return new SingleEaseKeyframe();
  }
}
class BOOLAttributeConfig extends BaseAttribute {
  constructor(key, attributeType, defaultValue) {
    super(key, attributeType, defaultValue);
  }
  readAttribute(byteArray, flag, targetClass, target) {
    readAttribute(byteArray, flag, targetClass, target, this);
  }
  readValue(byteArray) {
    return byteArray.readBoolean();
  }
  readValueList(byteArray, list, count) {
    for (let i = 0; i < count; i++) {
      list.push(byteArray.readBitBoolean());
    }
  }
  dimensionality() {
    return 1;
  }
  newKeyframe(_flag) {
    return new Keyframe();
  }
}
class Uint8AttributeConfig extends BaseAttribute {
  constructor(key, attributeType, defaultValue) {
    super(key, attributeType, defaultValue);
  }
  readAttribute(byteArray, flag, targetClass, target) {
    readAttribute(byteArray, flag, targetClass, target, this);
  }
  readValue(byteArray) {
    return byteArray.readUint8();
  }
  readValueList(byteArray, list, count) {
    const valueList = byteArray.readUint32List(count);
    for (let i = 0; i < count; i++) {
      list.push(valueList[i]);
    }
  }
  dimensionality() {
    return 1;
  }
  newKeyframe(_flag) {
    return new SingleEaseKeyframe();
  }
}
class TimeAttributeConfig extends BaseAttribute {
  constructor(key, attributeType, defaultValue) {
    super(key, attributeType, defaultValue);
  }
  readAttribute(byteArray, flag, targetClass, target) {
    readAttribute(byteArray, flag, targetClass, target, this);
  }
  readValue(byteArray) {
    return readTime(byteArray);
  }
  readValueList(byteArray, list, count) {
    for (let i = 0; i < count; i++) {
      list[i] = this.readValue(byteArray);
    }
  }
  dimensionality() {
    return 1;
  }
  newKeyframe(_flag) {
    return new SingleEaseKeyframe();
  }
}
class PointAttributeConfig extends BaseAttribute {
  constructor(key, attributeType, defaultValue) {
    super(key, attributeType, defaultValue);
  }
  readAttribute(byteArray, flag, targetClass, target) {
    readAttribute(byteArray, flag, targetClass, target, this);
  }
  readValue(byteArray) {
    return readPoint(byteArray);
  }
  readValueList(byteArray, list, count) {
    if (this.attributeType === AttributeType.SpatialProperty) {
      const values = byteArray.readFloatList(count * 2, SPATIAL_PRECISION);
      for (let i = 0; i < count; i++) {
        list[i] || (list[i] = new Point(0, 0));
        list[i].x = values[i];
      }
    } else {
      for (let i = 0; i < count; i++) {
        list[i] = readPoint(byteArray);
      }
    }
  }
  dimensionality() {
    return 2;
  }
  newKeyframe(_flag) {
    switch (this.attributeType) {
      case AttributeType.MultiDimensionProperty:
        return new MultiDimensionPointKeyframe();
      default:
        return new SingleEaseKeyframe();
    }
  }
}
class RatioAttributeConfig extends BaseAttribute {
  constructor(key, attributeType, defaultValue) {
    super(key, attributeType, defaultValue);
  }
  readAttribute(byteArray, flag, targetClass, target) {
    readAttribute(byteArray, flag, targetClass, target, this);
  }
  readValue(byteArray) {
    return readRatio(byteArray);
  }
  readValueList(byteArray, list, count) {
    for (let i = 0; i < count; i++) {
      list[i] = this.readValue(byteArray);
    }
  }
  dimensionality() {
    return 1;
  }
  newKeyframe(_flag) {
    return new SingleEaseKeyframe();
  }
}
class StringAttributeConfig extends BaseAttribute {
  constructor(key, attributeType, defaultValue) {
    super(key, attributeType, defaultValue);
  }
  readAttribute(byteArray, flag, targetClass, target) {
    readAttribute(byteArray, flag, targetClass, target, this);
  }
  readValue(byteArray) {
    return byteArray.readUTF8String();
  }
  readValueList(byteArray, list, count) {
    for (let i = 0; i < count; i++) {
      list[i] = this.readValue(byteArray);
    }
  }
  dimensionality() {
    return 1;
  }
  newKeyframe(_flag) {
    return new SingleEaseKeyframe();
  }
}
class LayerAttributeConfig extends BaseAttribute {
  constructor(key, attributeType, defaultValue) {
    super(key, attributeType, defaultValue);
  }
  readAttribute(byteArray, flag, targetClass, target) {
    readAttribute(byteArray, flag, targetClass, target, this);
  }
  readValue(byteArray) {
    return readLayerID(byteArray);
  }
  readValueList(byteArray, list, count) {
    for (let i = 0; i < count; i++) {
      list[i] = this.readValue(byteArray);
    }
  }
  dimensionality() {
    return 1;
  }
  newKeyframe(_flag) {
    return new SingleEaseKeyframe();
  }
}

const readBlockConfigOfLayerAttributes = {
  tagCode: TagCode.LayerAttributes,
  configs: [
    new BOOLAttributeConfig("isActive", AttributeType.BitFlag, true),
    new BOOLAttributeConfig("autoOrientation", AttributeType.BitFlag, false),
    new LayerAttributeConfig("parent", AttributeType.Value, void 0),
    new RatioAttributeConfig("stretch", AttributeType.Value, DefaultRatio),
    new TimeAttributeConfig("startTime", AttributeType.Value, ZERO_TIME),
    new Uint8AttributeConfig("blendMode", AttributeType.Value, BlendMode.Normal),
    new Uint8AttributeConfig("trackMatteType", AttributeType.Value, TrackMatteType.None),
    new FloatAttributeConfig("timeRemap", AttributeType.SimpleProperty, 0),
    new TimeAttributeConfig("duration", AttributeType.FixedValue, ZERO_TIME)
  ]
};
const readBlockConfigOfLayerAttributesV2 = {
  tagCode: TagCode.LayerAttributesV2,
  configs: [
    new BOOLAttributeConfig("isActive", AttributeType.BitFlag, true),
    new BOOLAttributeConfig("autoOrientation", AttributeType.BitFlag, false),
    new LayerAttributeConfig("parent", AttributeType.Value, void 0),
    new RatioAttributeConfig("stretch", AttributeType.Value, DefaultRatio),
    new TimeAttributeConfig("startTime", AttributeType.Value, ZERO_TIME),
    new Uint8AttributeConfig("blendMode", AttributeType.Value, BlendMode.Normal),
    new Uint8AttributeConfig("trackMatteType", AttributeType.Value, TrackMatteType.None),
    new FloatAttributeConfig("timeRemap", AttributeType.SimpleProperty, 0),
    new TimeAttributeConfig("duration", AttributeType.FixedValue, ZERO_TIME),
    new StringAttributeConfig("name", AttributeType.Value, "")
  ]
};
const readBlockConfigOfTransform2D = {
  tagCode: TagCode.Transform2D,
  configs: [
    new PointAttributeConfig("anchorPoint", AttributeType.SpatialProperty, ZERO_POINT),
    new PointAttributeConfig("position", AttributeType.SpatialProperty, ZERO_POINT),
    new FloatAttributeConfig("xPosition", AttributeType.SimpleProperty, 0),
    new FloatAttributeConfig("yPosition", AttributeType.SimpleProperty, 0),
    new PointAttributeConfig("scale", AttributeType.MultiDimensionProperty, new Point(1, 1)),
    new FloatAttributeConfig("rotation", AttributeType.SimpleProperty, 0),
    new Uint8AttributeConfig("opacity", AttributeType.SimpleProperty, OPAQUE)
  ]
};
({
  tagCode: TagCode.MaskBlock,
  configs: []
});

function readCompositionReference(byteArray, layer) {
  const id = byteArray.readEncodedUint32();
  if (id > 0) {
    layer.composition = new Composition();
    layer.composition.id = id;
  }
  layer.compositionStartTime = readTime(byteArray);
}

const readLayer = (byteArray) => {
  const layerType = byteArray.readUint8();
  let layer;
  switch (layerType) {
    case LayerType.undefined:
      layer = new UnDefinedLayer();
      break;
    case LayerType.Solid:
      layer = new SolidLayer();
      break;
    case LayerType.Shape:
      layer = new ShapeLayer();
      break;
    case LayerType.PreCompose:
      layer = new PreComposeLayer();
      break;
    default:
      layer = new Layer();
      break;
  }
  layer.id = byteArray.readEncodedUint32();
  readTags(byteArray, layer, readTagsOfLayer);
  return layer;
};
const readTagsOfLayer = (byteArray, code, layer) => {
  switch (code) {
    case TagCode.LayerAttributes:
      readTagBlock(byteArray, layer, readBlockConfigOfLayerAttributes);
      if (layer.duration <= 0)
        layer.duration = 1;
      break;
    case TagCode.LayerAttributesV2:
      readTagBlock(byteArray, layer, readBlockConfigOfLayerAttributesV2);
      if (layer.duration <= 0)
        layer.duration = 1;
      break;
    case TagCode.Transform2D:
      layer.transform = new Transform2D();
      readTagBlock(byteArray, layer.transform, readBlockConfigOfTransform2D);
      if (layer.transform.position.animatable() || layer.transform.position.value !== ZERO_POINT || !(layer.transform.xPosition.animatable() || layer.transform.xPosition.value !== 0) && !(layer.transform.yPosition.animatable() || layer.transform.yPosition.value !== 0)) {
        layer.transform.xPosition = new Property(0);
        layer.transform.yPosition = new Property(0);
      } else {
        layer.transform.position = new Property(ZERO_POINT);
      }
      break;
    case TagCode.SolidColor:
      if (layer.type() === LayerType.Solid) {
        readSolidColor(byteArray, layer);
      }
      break;
    case TagCode.CompositionReference:
      if (layer.type() === LayerType.PreCompose) {
        readCompositionReference(byteArray, layer);
      }
      break;
  }
};

const readVectorComposition = (byteArray) => {
  const composition = new VectorComposition();
  composition.id = byteArray.readEncodedUint32();
  readTags(byteArray, composition, readTagsOfVectorComposition);
  installArrayLayerReference(composition.layers);
  return composition;
};
const readTagsOfVectorComposition = (byteArray, code, composition) => {
  switch (code) {
    case TagCode.CompositionAttributes:
      readCompositionAttributes(byteArray, composition);
      break;
    case TagCode.LayerBlock:
      composition.layers.push(readLayer(byteArray));
      break;
  }
};
const installArrayLayerReference = (layers) => {
  if (layers && layers.length === 0) {
    return;
  }
  const layerMap = /* @__PURE__ */ new Map();
  for (const layer of layers) {
    if (!layer) {
      continue;
    }
    installLayerReference(layer);
    layerMap.set(layer.id, layer);
  }
  let index = 0;
  for (const layer of layers) {
    if (!layer) {
      continue;
    }
    if (layer.parent !== void 0) {
      const ID = layer.parent.id;
      const result = layerMap.get(ID);
      if (result !== void 0) {
        layer.parent = result;
      }
    }
    if (index > 0 && hasTrackMatte(layer.trackMatteType)) {
      layer.trackMatteLayer = layers[index - 1];
    }
    if (layer.effects !== void 0 && layer.effects.length > 0) {
      for (const effect of layer.effects) {
        if (!effect) {
          continue;
        }
        if (effect.type() === EffectType.DisplacementMap);
      }
    }
    index += 1;
  }
};
const installLayerReference = (layer) => {
  var _a;
  if (!layer || !layer.masks || layer.masks.length === 0)
    return;
  const maskMap = /* @__PURE__ */ new Map();
  for (const mask of layer.masks) {
    if (!mask) {
      continue;
    }
    maskMap.set(mask.id, mask);
  }
  (_a = layer.effects) == null ? void 0 : _a.forEach((effect) => {
    if (!effect)
      return;
    if (effect.maskReferences !== void 0 && effect.maskReferences.length > 0) {
      const maskReferences = new Array();
      effect.maskReferences.forEach((mask) => {
        const ID = mask.id;
        const result = maskMap.get(ID);
        if (result !== void 0) {
          maskReferences.push(result);
        }
      });
      effect.maskReferences = maskReferences;
    }
    switch (effect.type()) {
      case EffectType.Fill:
        if (effect.fillMask !== void 0) {
          const ID = effect.fillMask.id;
          const result = maskMap.get(ID);
          if (result !== void 0) {
            effect.fillMask = result;
          }
        }
        break;
      case EffectType.Stroke: {
        const strokeEffect = effect;
        if (strokeEffect.path !== void 0) {
          const ID = strokeEffect.path.id;
          const result = maskMap.get(ID);
          if (result !== void 0) {
            strokeEffect.path = result;
          }
        }
        break;
      }
    }
  });
  if (layer.type() === LayerType.Text) {
    const { pathOption } = layer;
    if (pathOption == null ? void 0 : pathOption.path) {
      const ID = pathOption.path.id;
      const result = maskMap.get(ID);
      if (result !== void 0) {
        pathOption.path = result;
      }
    }
  }
};
const hasTrackMatte = (type) => {
  switch (type) {
    case TrackMatteType.Alpha:
    case TrackMatteType.AlphaInverted:
      return true;
    default:
      return false;
  }
};

function readTagsOfFile(byteArray, code, context) {
  switch (code) {
    case TagCode.VectorCompositionBlock:
      context.compositions.push(readVectorComposition(byteArray));
      break;
    case TagCode.VideoCompositionBlock:
      context.compositions.push(readVideoComposition(byteArray));
      break;
  }
}

const verifyAndMake = (compositions) => {
  let success = compositions.length > 0;
  for (const composition of compositions) {
    if (!composition || !composition.verify()) {
      success = false;
      break;
    }
  }
  if (!success) {
    throw new Error("Verify composition failed!");
  }
};
function installReference(compositions) {
  if (!compositions || compositions.length === 0)
    return;
  const compositionMap = /* @__PURE__ */ new Map();
  compositions.forEach((composition) => {
    if (composition) {
      compositionMap.set(composition.id, composition);
    }
  });
  compositions.forEach((composition) => {
    if (composition && composition.type() === CompositionType.Vector) {
      const vectorComposition = composition;
      if (vectorComposition.layers && vectorComposition.layers.length > 0) {
        vectorComposition.layers.forEach((layer) => {
          layer.containingComposition = vectorComposition;
          const preComposeLayer = layer;
          if (preComposeLayer.type() === LayerType.PreCompose && preComposeLayer.composition) {
            const res = compositionMap.get(preComposeLayer.composition.id);
            if (res) {
              preComposeLayer.composition = res;
            }
          }
        });
      }
    }
  });
}
const decode = (byteArray) => {
  const bodyByteArray = readBodyBytes(byteArray);
  const { context } = bodyByteArray;
  readTags(bodyByteArray, context, readTagsOfFile);
  installReference(context.compositions);
  const compositions = context.releaseCompositions();
  verifyAndMake(compositions);
  return { compositions, tagLevel: context.tagLevel };
};
const readBodyBytes = (byteArray) => {
  if (byteArray.length < 11)
    throw new Error("PAG file is invalid!");
  const P = byteArray.readInt8();
  const A = byteArray.readInt8();
  const G = byteArray.readInt8();
  if (P !== 80 || A !== 65 || G !== 71)
    throw new Error("invalid PAG header!");
  byteArray.readInt8();
  byteArray.readUint32();
  byteArray.readInt8();
  return byteArray.readBytes();
};

class PAGFile {
  constructor(compositions, tagLevel) {
    this.tagLevel = 1;
    this.compositions = [];
    this.numLayers = 0;
    this.scaledTimeRange = { start: 0, end: 0 };
    this.mainComposition = compositions[compositions.length - 1];
    this.scaledTimeRange.start = 0;
    this.scaledTimeRange.end = this.mainComposition.duration;
    this.compositions = compositions;
    this.duration = this.mainComposition.duration;
    this.implDuration = this.mainComposition.duration * 1e3 / this.mainComposition.frameRate;
    for (const composition of compositions) {
      if (composition.type() !== CompositionType.Vector) {
        this.numLayers += 1;
        continue;
      }
      for (const layer of composition.layers) {
        if (layer.type() === LayerType.PreCompose) {
          continue;
        }
        this.numLayers += 1;
      }
    }
    this.tagLevel = tagLevel;
  }
  static fromArrayBuffer(arrayBuffer) {
    if (!arrayBuffer || arrayBuffer.byteLength === 0)
      throw new Error("Can't read empty array buffer!");
    const byteArray = new ByteArray(arrayBuffer, true);
    const { compositions, tagLevel } = decode(byteArray);
    return new PAGFile(compositions, tagLevel);
  }
  getVideoSequence() {
    const compositionType = this.mainComposition.type();
    if (compositionType === CompositionType.Video) {
      return getVideoSequenceFromVideoComposition(this.mainComposition);
    } else if (compositionType === CompositionType.Vector) {
      return getVideoSequenceFromVectorComposition(this.mainComposition);
    }
  }
}
const getVideoSequenceFromVideoComposition = (videoComposition) => {
  if (!videoComposition.sequences || videoComposition.sequences.length === 0) {
    throw new Error("PAGFile has no BMP video sequence!");
  }
  return videoComposition.sequences[videoComposition.sequences.length - 1];
};
const getVideoSequenceFromVectorComposition = (vectorComposition) => {
  const videoCompositions = getVideoComposition(vectorComposition);
  if (videoCompositions.length > 1)
    throw new Error("PAGFile has more than one BMP video sequence!");
  if (videoCompositions.length < 1)
    throw new Error("PAGFile has no BMP video sequence!");
  const videoComposition = videoCompositions[0];
  return getVideoSequenceFromVideoComposition(videoComposition);
};
const getVideoComposition = (vectorComposition) => {
  const videoCompositions = [];
  vectorComposition.layers.forEach((layer) => {
    if (layer.type() !== LayerType.PreCompose)
      return;
    const { composition } = layer;
    if ((composition == null ? void 0 : composition.type()) === CompositionType.Video) {
      videoCompositions.push(composition);
      return;
    }
    if ((composition == null ? void 0 : composition.type()) === CompositionType.Vector) {
      videoCompositions.push(...getVideoComposition(composition));
    }
  });
  return videoCompositions;
};

var RenderingMode = /* @__PURE__ */ ((RenderingMode2) => {
  RenderingMode2["Canvas2D"] = "2d";
  RenderingMode2["WebGL"] = "WebGL";
  return RenderingMode2;
})(RenderingMode || {});
var EventName = /* @__PURE__ */ ((EventName2) => {
  EventName2["onAnimationStart"] = "onAnimationStart";
  EventName2["onAnimationEnd"] = "onAnimationEnd";
  EventName2["onAnimationCancel"] = "onAnimationCancel";
  EventName2["onAnimationRepeat"] = "onAnimationRepeat";
  EventName2["onAnimationUpdate"] = "onAnimationUpdate";
  EventName2["onAnimationPlay"] = "onAnimationPlay";
  EventName2["onAnimationPause"] = "onAnimationPause";
  return EventName2;
})(EventName || {});
var ScaleMode = /* @__PURE__ */ ((ScaleMode2) => {
  ScaleMode2["None"] = "None";
  ScaleMode2["Stretch"] = "Stretch";
  ScaleMode2["LetterBox"] = "LetterBox";
  ScaleMode2["Zoom"] = "Zoom";
  return ScaleMode2;
})(ScaleMode || {});

var types = /*#__PURE__*/Object.freeze({
  __proto__: null,
  RenderingMode: RenderingMode,
  EventName: EventName,
  ScaleMode: ScaleMode
});

function destroyVerify(constructor) {
  let functions = Object.getOwnPropertyNames(constructor.prototype).filter(
    (name) => name !== "constructor" && typeof constructor.prototype[name] === "function"
  );
  const proxyFn = (target, methodName) => {
    const fn = target[methodName];
    target[methodName] = function (...args) {
      if (this["destroyed"]) {
        console.error(`Don't call ${methodName} of the PAGView that is destroyed.`);
        return;
      }
      return fn.call(this, ...args);
    };
  };
  functions.forEach((name) => proxyFn(constructor.prototype, name));
}

const detectWebGLContext = () => {
  const canvas = document.createElement("canvas");
  const gl = canvas.getContext("webgl") || canvas.getContext("experimental-webgl");
  return !!gl;
};
const createProgram = (gl, vertexShaderSource, fragmentShaderSource) => {
  const program = gl.createProgram();
  if (!program)
    throw new Error("Failed to create program.");
  const vShader = createShader(gl, vertexShaderSource, gl.VERTEX_SHADER);
  if (!vShader)
    throw new Error("Failed to create vertex shader.");
  gl.attachShader(program, vShader);
  const fShader = createShader(gl, fragmentShaderSource, gl.FRAGMENT_SHADER);
  if (!fShader)
    throw new Error("Failed to create fragment shader.");
  gl.attachShader(program, fShader);
  gl.linkProgram(program);
  const programMessage = gl.getProgramInfoLog(program);
  if (programMessage)
    console.log(programMessage);
  const vShaderMessage = gl.getShaderInfoLog(vShader);
  if (vShaderMessage)
    console.log(vShaderMessage);
  const fShaderMessage = gl.getShaderInfoLog(fShader);
  if (fShaderMessage)
    console.log(fShaderMessage);
  return program;
};
const createShader = (gl, source, type) => {
  const shader = gl.createShader(type);
  if (!shader)
    throw new Error("Failed to create shader.");
  gl.shaderSource(shader, source);
  gl.compileShader(shader);
  return shader;
};
const getShaderSourceFromString = (str) => str.replace(/^\s+|\s+$/g, "");
const getVideoParam = (pagFile, videoSequence) => {
  const attribute = {
    width: pagFile.mainComposition.width,
    height: pagFile.mainComposition.height,
    hasAlpha: videoSequence.alphaStartX > 0 || videoSequence.alphaStartY > 0,
    alphaStartX: videoSequence.alphaStartX,
    alphaStartY: videoSequence.alphaStartY,
    sequenceWidth: videoSequence.width,
    sequenceHeight: videoSequence.height,
    MP4Width: (videoSequence.width + videoSequence.alphaStartX) % 2 === 0 ? videoSequence.width + videoSequence.alphaStartX : videoSequence.width + videoSequence.alphaStartX + 1,
    MP4Height: (videoSequence.height + videoSequence.alphaStartY) % 2 === 0 ? videoSequence.height + videoSequence.alphaStartY : videoSequence.height + videoSequence.alphaStartY + 1
  };
  return attribute;
};
const createAndSetupTexture = (gl) => {
  const texture = gl.createTexture();
  if (!texture)
    throw new Error("Failed to create texture.");
  gl.bindTexture(gl.TEXTURE_2D, texture);
  gl.texParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE);
  gl.texParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE);
  gl.texParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST);
  gl.texParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST);
  return texture;
};
const getWechatNetwork = () => {
  return new Promise((resolve) => {
    window.WeixinJSBridge.invoke(
      "getNetworkType",
      {},
      () => {
        resolve();
      },
      () => {
        resolve();
      }
    );
  });
};

class EventManager {
  constructor() {
    this.listenersMap = {};
  }
  on(eventName, listener) {
    if (this.listenersMap[eventName] === void 0) {
      this.listenersMap[eventName] = [];
    }
    this.listenersMap[eventName].push(listener);
    return;
  }
  off(eventName, listener) {
    const listenerList = this.listenersMap[eventName];
    if (listenerList === void 0)
      return;
    if (listener === void 0) {
      delete this.listenersMap[eventName];
      return;
    }
    const index = listenerList.findIndex((fn) => fn === listener);
    listenerList.splice(index, 1);
    return;
  }
  emit(eventName, ...payload) {
    const listenerList = this.listenersMap[eventName];
    if (listenerList === void 0 || listenerList.length < 1)
      return false;
    for (const listener of listenerList) {
      listener(...payload);
    }
    return true;
  }
}

var __defProp$6 = Object.defineProperty;
var __getOwnPropDesc$4 = Object.getOwnPropertyDescriptor;
var __getOwnPropSymbols$3 = Object.getOwnPropertySymbols;
var __hasOwnProp$3 = Object.prototype.hasOwnProperty;
var __propIsEnum$3 = Object.prototype.propertyIsEnumerable;
var __defNormalProp$3 = (obj, key, value) => key in obj ? __defProp$6(obj, key, { enumerable: true, configurable: true, writable: true, value }) : obj[key] = value;
var __spreadValues$3 = (a, b) => {
  for (var prop in b || (b = {}))
    if (__hasOwnProp$3.call(b, prop))
      __defNormalProp$3(a, prop, b[prop]);
  if (__getOwnPropSymbols$3)
    for (var prop of __getOwnPropSymbols$3(b)) {
      if (__propIsEnum$3.call(b, prop))
        __defNormalProp$3(a, prop, b[prop]);
    }
  return a;
};
var __decorateClass$4 = (decorators, target, key, kind) => {
  var result = kind > 1 ? void 0 : kind ? __getOwnPropDesc$4(target, key) : target;
  for (var i = decorators.length - 1, decorator; i >= 0; i--)
    if (decorator = decorators[i])
      result = (kind ? decorator(target, key, result) : decorator(result)) || result;
  if (kind && result)
    __defProp$6(target, key, result);
  return result;
};
let Context = class {
  constructor(pagFile, canvas, options) {
    this.canvasSize = { width: 0, height: 0 };
    this.playing = false;
    this.viewportSize = { x: 0, y: 0, width: 0, height: 0, scaleX: 1, scaleY: 1 };
    this.destroyed = false;
    this.renderTimer = null;
    this.repeatCount = 0;
    this.viewScaleMode = ScaleMode.LetterBox;
    this.debugData = {
      FPS: 0,
      decodePAGFile: 0,
      createDir: 0,
      coverMP4: 0,
      writeFile: 0,
      createDecoder: 0,
      getFrame: 0,
      draw: 0
    };
    const videoSequence = pagFile.getVideoSequence();
    if (!videoSequence)
      throw new Error("PAGFile has no BMP video sequence!");
    delete videoSequence.composition;
    this.videoSequence = videoSequence;
    this.canvas = canvas;
    this.videoParam = getVideoParam(pagFile, videoSequence);
    this.eventManager = new EventManager();
    this.renderingMode = options.renderingMode || RenderingMode.WebGL;
    this.updateSize(options.useScale);
    this.setScaleMode();
  }
  isPlaying() {
    return this.playing;
  }
  isDestroyed() {
    return this.destroyed;
  }
  duration() {
    return this.videoSequence.frameCount / this.videoSequence.frameRate;
  }
  frameRate() {
    return this.videoSequence.frameRate;
  }
  setRepeatCount(repeatCount = 1) {
    this.repeatCount = repeatCount < 0 ? -1 : repeatCount - 1;
  }
  addListener(eventName, listener) {
    return this.eventManager.on(eventName, listener);
  }
  removeListener(eventName, listener) {
    return this.eventManager.off(eventName, listener);
  }
  scaleMode() {
    return this.viewScaleMode;
  }
  setScaleMode(scaleMode = ScaleMode.LetterBox) {
    this.viewScaleMode = scaleMode;
    switch (scaleMode) {
      case ScaleMode.None:
        this.viewportSize = {
          x: 0,
          y: this.renderingMode === RenderingMode.WebGL ? this.canvas.height - this.videoParam.height : 0,
          width: this.videoParam.width,
          height: this.videoParam.height,
          scaleX: 1,
          scaleY: 1
        };
        break;
      case ScaleMode.Stretch:
        this.viewportSize = {
          x: 0,
          y: 0,
          width: this.canvas.width,
          height: this.canvas.height,
          scaleX: this.canvas.width / this.videoParam.sequenceWidth,
          scaleY: this.canvas.height / this.videoParam.sequenceHeight
        };
        break;
      case ScaleMode.LetterBox:
        {
          const scaleX = this.canvas.width / this.videoParam.sequenceWidth;
          const scaleY = this.canvas.height / this.videoParam.sequenceHeight;
          const scale = Math.min(scaleX, scaleY);
          this.viewportSize = {
            x: (this.canvas.width - this.videoParam.sequenceWidth * scale) / 2,
            y: (this.canvas.height - this.videoParam.sequenceHeight * scale) / 2,
            width: this.videoParam.sequenceWidth * scale,
            height: this.videoParam.sequenceHeight * scale,
            scaleX: scale,
            scaleY: scale
          };
        }
        break;
      case ScaleMode.Zoom:
        {
          const scaleX = this.canvas.width / this.videoParam.sequenceWidth;
          const scaleY = this.canvas.height / this.videoParam.sequenceHeight;
          const scale = Math.max(scaleX, scaleY);
          this.viewportSize = {
            x: (this.canvas.width - this.videoParam.sequenceWidth * scale) / 2,
            y: (this.canvas.height - this.videoParam.sequenceHeight * scale) / 2,
            width: this.videoParam.sequenceWidth * scale,
            height: this.videoParam.sequenceHeight * scale,
            scaleX: scale,
            scaleY: scale
          };
        }
        break;
    }
  }
  updateSize(useScale = true) {
    if (!this.canvas) {
      throw new Error("Canvas element is not found!");
    }
    let displaySize;
    const styleDeclaration = getComputedStyle(this.canvas);
    const computedSize = {
      width: Number(styleDeclaration.width.replace("px", "")),
      height: Number(styleDeclaration.height.replace("px", ""))
    };
    if (computedSize.width > 0 && computedSize.height > 0) {
      displaySize = computedSize;
    } else {
      const styleSize = {
        width: Number(this.canvas.style.width.replace("px", "")),
        height: Number(this.canvas.style.height.replace("px", ""))
      };
      if (styleSize.width > 0 && styleSize.height > 0) {
        displaySize = styleSize;
      } else {
        displaySize = {
          width: this.canvas.width,
          height: this.canvas.height
        };
      }
    }
    if (!useScale) {
      this.canvas.width = this.canvas.width || displaySize.width;
      this.canvas.height = this.canvas.height || displaySize.height;
      return;
    }
    this.canvas.style.width = `${displaySize.width}px`;
    this.canvas.style.height = `${displaySize.height}px`;
    this.canvas.width = displaySize.width * window.devicePixelRatio;
    this.canvas.height = displaySize.height * window.devicePixelRatio;
  }
  getDebugData() {
    return this.debugData;
  }
  setDebugData(data) {
    this.debugData = __spreadValues$3(__spreadValues$3({}, this.debugData), data);
  }
  loadContext() {
  }
  clearRender() {
  }
};
Context = __decorateClass$4([
  destroyVerify
], Context);

const CORRECTION_UTC = 2082873600;
const DEFAULT_VOLUME = 1;
const getCharCode = (name) => {
  const res = [];
  for (let index = 0; index < name.length; index++) {
    res.push(name.charCodeAt(index));
  }
  return res;
};
const toHexadecimal = (num) => [num >> 24, num >> 16 & 255, num >> 8 & 255, num & 255];
const makeBox = (type, ...payload) => {
  let size = 8;
  let i = payload.length;
  const len = i;
  while (i) {
    i -= 1;
    size += payload[i].byteLength;
  }
  const result = new Uint8Array(size);
  result[0] = size >> 24 & 255;
  result[1] = size >> 16 & 255;
  result[2] = size >> 8 & 255;
  result[3] = size & 255;
  result.set(type, 4);
  for (i = 0, size = 8; i < len; ++i) {
    result.set(payload[i], size);
    size += payload[i].byteLength;
  }
  return result;
};
class MP4Generator {
  constructor(boxParam) {
    this.param = boxParam;
  }
  ftyp() {
    return makeBox(
      getCharCode("ftyp"),
      new Uint8Array(getCharCode("isom")),
      new Uint8Array([0, 0, 0, 1]),
      new Uint8Array(getCharCode("isom")),
      new Uint8Array(getCharCode("iso2")),
      new Uint8Array(getCharCode("avc1")),
      new Uint8Array(getCharCode("mp41"))
    );
  }
  moov() {
    const traks = this.param.tracks.map((track) => this.trak(track)).reverse();
    return makeBox(getCharCode("moov"), this.mvhd(), ...traks, this.mvex());
  }
  moof() {
    return makeBox(getCharCode("moof"), this.mfhd(), this.traf());
  }
  mdat() {
    const buffer = new Uint8Array(this.param.track.len);
    let offset = 0;
    this.param.videoSequence.headers.forEach((header) => {
      buffer.set(new Uint8Array(header.data.data()), offset);
      offset += header.length;
    });
    this.param.videoSequence.frames.forEach((frame, index) => {
      buffer.set(new Uint8Array(frame.fileBytes.data.data()), offset);
      offset += frame.fileBytes.length;
    });
    return makeBox(getCharCode("mdat"), buffer);
  }
  mvhd() {
    return makeBox(
      getCharCode("mvhd"),
      new Uint8Array([
        0,
        0,
        0,
        0,
        ...toHexadecimal(CORRECTION_UTC),
        ...toHexadecimal(CORRECTION_UTC),
        ...toHexadecimal(this.param.timescale),
        ...toHexadecimal(this.param.duration),
        0,
        1,
        0,
        0,
        1,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        1,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        1,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        64,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        2
      ])
    );
  }
  trak(track) {
    return makeBox(getCharCode("trak"), this.tkhd(track), this.edts(track), this.mdia(track));
  }
  tkhd(track) {
    return makeBox(
      getCharCode("tkhd"),
      new Uint8Array([
        0,
        0,
        0,
        1,
        ...toHexadecimal(CORRECTION_UTC),
        ...toHexadecimal(CORRECTION_UTC),
        ...toHexadecimal(track.id),
        0,
        0,
        0,
        0,
        ...toHexadecimal(track.duration),
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        DEFAULT_VOLUME >> 0 & 255,
        DEFAULT_VOLUME % 1 * 10 >> 0 & 255,
        0,
        0,
        0,
        1,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        1,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        64,
        0,
        0,
        0,
        track.width >> 8 & 255,
        track.width & 255,
        0,
        0,
        track.height >> 8 & 255,
        track.height & 255,
        0,
        0
      ])
    );
  }
  edts(track) {
    return makeBox(getCharCode("edts"), this.elst(track));
  }
  elst(track) {
    return makeBox(
      getCharCode("elst"),
      new Uint8Array([
        0,
        0,
        0,
        0,
        ...toHexadecimal(1),
        ...toHexadecimal(track.duration),
        ...toHexadecimal(track.implicitOffset * Math.floor(track.duration / track.samples.length)),
        0,
        1,
        0,
        0
      ])
    );
  }
  mdia(track) {
    return makeBox(getCharCode("mdia"), this.mdhd(), this.hdlr(), this.minf(track));
  }
  mdhd() {
    return makeBox(
      getCharCode("mdhd"),
      new Uint8Array([
        0,
        0,
        0,
        0,
        ...toHexadecimal(CORRECTION_UTC),
        ...toHexadecimal(CORRECTION_UTC),
        ...toHexadecimal(this.param.timescale),
        ...toHexadecimal(0),
        85,
        196,
        0,
        0
      ])
    );
  }
  hdlr() {
    return makeBox(
      getCharCode("hdlr"),
      new Uint8Array([
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        118,
        105,
        100,
        101,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        86,
        105,
        100,
        101,
        111,
        72,
        97,
        110,
        100,
        108,
        101,
        114,
        0
      ])
    );
  }
  minf(track) {
    return makeBox(getCharCode("minf"), this.vmhd(), this.dinf(), this.stbl(track));
  }
  vmhd() {
    return makeBox(
      getCharCode("vmhd"),
      new Uint8Array([
        0,
        0,
        0,
        1,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0
      ])
    );
  }
  dinf() {
    return makeBox(getCharCode("dinf"), this.dref());
  }
  dref() {
    return makeBox(
      getCharCode("dref"),
      new Uint8Array([
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        1,
        0,
        0,
        0,
        12,
        117,
        114,
        108,
        32,
        0,
        0,
        0,
        1
      ])
    );
  }
  stbl(track) {
    return makeBox(
      getCharCode("stbl"),
      this.stsd(track),
      this.stts(track),
      this.ctts(track),
      this.stss(track),
      this.stsc(),
      this.stsz(),
      this.stco()
    );
  }
  stsd(track) {
    const data = [
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      1
    ];
    return makeBox(getCharCode("stsd"), new Uint8Array(data), this.avc1(track));
  }
  avc1(track) {
    const data = [
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      1,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      track.width >> 8 & 255,
      track.width & 255,
      track.height >> 8 & 255,
      track.height & 255,
      0,
      72,
      0,
      0,
      0,
      72,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      1,
      18,
      98,
      105,
      110,
      101,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      24,
      255,
      255
    ];
    return makeBox(getCharCode("avc1"), new Uint8Array(data), this.avcc(track));
  }
  avcc(track) {
    let sps = [];
    let pps = [];
    track.sps.forEach((byteData) => {
      const len = byteData.length - 4;
      sps.push(len >>> 8 & 255);
      sps.push(len & 255);
      sps = sps.concat(Array.prototype.slice.call(new Uint8Array(byteData.data.data(), 4)));
    });
    track.pps.forEach((byteData) => {
      const len = byteData.length - 4;
      pps.push(len >>> 8 & 255);
      pps.push(len & 255);
      pps = pps.concat(Array.prototype.slice.call(new Uint8Array(byteData.data.data(), 4)));
    });
    const data = [
      1,
      sps[3],
      sps[4],
      sps[5],
      252 | 3,
      224 | track.sps.length
    ].concat(sps).concat([track.pps.length]).concat(pps);
    return makeBox(getCharCode("avcC"), new Uint8Array(data));
  }
  stts(track) {
    return makeBox(
      getCharCode("stts"),
      new Uint8Array([
        0,
        0,
        0,
        0,
        ...toHexadecimal(1),
        ...toHexadecimal(track.samples.length),
        ...toHexadecimal(Math.floor(track.duration / track.samples.length))
      ])
    );
  }
  ctts(track) {
    const sampleCount = track.pts.length;
    const sampleDelta = Math.floor(track.duration / sampleCount);
    const data = [
      0,
      0,
      0,
      0,
      ...toHexadecimal(sampleCount)
    ];
    for (let i = 0; i < sampleCount; i++) {
      data.push(...toHexadecimal(1));
      const dts = i * sampleDelta;
      const pts = (track.pts[i] + track.implicitOffset) * sampleDelta;
      data.push(...toHexadecimal(pts - dts));
    }
    return makeBox(getCharCode("ctts"), new Uint8Array(data));
  }
  stss(track) {
    const iFrames = track.samples.filter((sample) => sample.flags.isKeyFrame).map((sample) => sample.index + 1);
    const data = [
      0,
      0,
      0,
      0,
      ...toHexadecimal(iFrames.length)
    ];
    iFrames.forEach((iFrame) => {
      data.push(...toHexadecimal(iFrame));
    });
    return makeBox(getCharCode("stss"), new Uint8Array(data));
  }
  stsc() {
    return makeBox(
      getCharCode("stsc"),
      new Uint8Array([
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0
      ])
    );
  }
  stsz() {
    return makeBox(
      getCharCode("stsz"),
      new Uint8Array([
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0
      ])
    );
  }
  stco() {
    return makeBox(
      getCharCode("stco"),
      new Uint8Array([
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0
      ])
    );
  }
  mvex() {
    const trexs = this.param.tracks.map((track) => this.trex(track)).reverse();
    return makeBox(getCharCode("mvex"), ...trexs);
  }
  trex(track) {
    return makeBox(
      getCharCode("trex"),
      new Uint8Array([
        0,
        0,
        0,
        0,
        track.id >> 24,
        track.id >> 16 & 255,
        track.id >> 8 & 255,
        track.id & 255,
        0,
        0,
        0,
        1,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        1,
        0,
        1
      ])
    );
  }
  mfhd() {
    return makeBox(
      getCharCode("mfhd"),
      new Uint8Array([
        0,
        0,
        0,
        0,
        this.param.sequenceNumber >> 24,
        this.param.sequenceNumber >> 16 & 255,
        this.param.sequenceNumber >> 8 & 255,
        this.param.sequenceNumber & 255
      ])
    );
  }
  traf() {
    const sdtp = this.sdtp();
    this.param.offset = sdtp.length + 72;
    return makeBox(getCharCode("traf"), this.tfhd(), this.tfdt(), this.trun(), sdtp);
  }
  tfhd() {
    return makeBox(
      getCharCode("tfhd"),
      new Uint8Array([
        0,
        0,
        0,
        0,
        ...toHexadecimal(this.param.track.id)
      ])
    );
  }
  tfdt() {
    return makeBox(
      getCharCode("tfdt"),
      new Uint8Array([
        0,
        0,
        0,
        0,
        ...toHexadecimal(this.param.baseMediaDecodeTime)
      ])
    );
  }
  trun() {
    const samples = this.param.track.samples || [];
    const len = samples.length;
    const arraylen = 12 + 16 * len;
    this.param.offset += 8 + arraylen;
    const data = [
      0,
      0,
      15,
      1,
      ...toHexadecimal(len),
      ...toHexadecimal(this.param.offset)
    ];
    this.param.track.samples.forEach((sample) => {
      const paddingValue = 0;
      const { duration, size, flags, cts } = sample;
      data.push(...toHexadecimal(duration));
      data.push(...toHexadecimal(size));
      data.push(flags.isLeading << 2 | flags.dependsOn);
      data.push(flags.isDependedOn << 6 | flags.hasRedundancy << 4 | paddingValue << 1 | flags.isNonSync);
      data.push(flags.degradPrio & 240 << 8);
      data.push(flags.degradPrio & 15);
      data.push(...toHexadecimal(cts));
    });
    return makeBox(getCharCode("trun"), new Uint8Array(data));
  }
  sdtp() {
    const buffer = new Uint8Array(4 + this.param.track.samples.length);
    this.param.track.samples.forEach((sample, index) => {
      buffer[index + 4] = sample.flags.dependsOn << 4 | sample.flags.isDependedOn << 2 | sample.flags.hasRedundancy;
    });
    return makeBox(getCharCode("sdtp"), buffer);
  }
}

var __defProp$5 = Object.defineProperty;
var __getOwnPropSymbols$2 = Object.getOwnPropertySymbols;
var __hasOwnProp$2 = Object.prototype.hasOwnProperty;
var __propIsEnum$2 = Object.prototype.propertyIsEnumerable;
var __defNormalProp$2 = (obj, key, value) => key in obj ? __defProp$5(obj, key, { enumerable: true, configurable: true, writable: true, value }) : obj[key] = value;
var __spreadValues$2 = (a, b) => {
  for (var prop in b || (b = {}))
    if (__hasOwnProp$2.call(b, prop))
      __defNormalProp$2(a, prop, b[prop]);
  if (__getOwnPropSymbols$2)
    for (var prop of __getOwnPropSymbols$2(b)) {
      if (__propIsEnum$2.call(b, prop))
        __defNormalProp$2(a, prop, b[prop]);
    }
  return a;
};
const SEQUENCE_NUMBER = 1;
const BASE_MEDIA_DECODE_TIME = 0;
const BASE_MEDIA_TIME_SCALE = 6e3;
const coverToMp4 = (videoSequence) => {
  const sequence = IS_IOS ? getVirtualSequence(videoSequence) : videoSequence;
  const mp4Track = makeMp4Track(sequence);
  if (!mp4Track || mp4Track.len === 0)
    throw new Error("mp4Track is empty");
  const boxParam = {
    offset: 0,
    tracks: [mp4Track],
    track: mp4Track,
    duration: mp4Track.duration,
    timescale: mp4Track.timescale,
    sequenceNumber: SEQUENCE_NUMBER,
    baseMediaDecodeTime: BASE_MEDIA_DECODE_TIME,
    nalusBytesLen: mp4Track.len,
    videoSequence: sequence
  };
  const mp4Generator = new MP4Generator(boxParam);
  const ftyp = mp4Generator.ftyp();
  const moov = mp4Generator.moov();
  const moof = mp4Generator.moof();
  const mdat = mp4Generator.mdat();
  return concatUint8Arrays([ftyp, moov, moof, mdat]);
};
const makeMp4Track = (videoSequence) => {
  if (videoSequence.headers.length < 2)
    throw new Error("Bad header data in video sequence!");
  if (videoSequence.frames.length === 0)
    throw new Error("There is no frame data in the video sequence!");
  const mp4Track = {
    id: 1,
    type: "video",
    timescale: BASE_MEDIA_TIME_SCALE,
    duration: Math.floor(videoSequence.frames.length * BASE_MEDIA_TIME_SCALE / videoSequence.frameRate),
    width: videoSequence.getVideoWidth(),
    height: videoSequence.getVideoHeight(),
    sps: [videoSequence.headers[0]],
    pps: [videoSequence.headers[1]],
    implicitOffset: getImplicitOffset(videoSequence.frames),
    len: 0,
    pts: [],
    samples: []
  };
  const headerLen = videoSequence.headers.reduce((pre, cur) => pre + cur.length, 0);
  const sampleDelta = mp4Track.duration / videoSequence.frames.length;
  videoSequence.frames.forEach((frame, index) => {
    var _a;
    let sampleSize = (_a = frame.fileBytes.length) != null ? _a : 0;
    if (index === 0) {
      sampleSize += headerLen;
    }
    mp4Track.len += sampleSize;
    mp4Track.pts.push(frame.frame);
    mp4Track.samples.push({
      index,
      size: sampleSize,
      duration: sampleDelta,
      cts: (frame.frame + mp4Track.implicitOffset - index) * sampleDelta,
      flags: {
        isKeyFrame: frame.isKeyframe,
        isNonSync: frame.isKeyframe ? 0 : 1,
        dependsOn: frame.isKeyframe ? 2 : 1,
        isLeading: 0,
        isDependedOn: 0,
        hasRedundancy: 0,
        degradPrio: 0
      }
    });
  });
  return mp4Track;
};
const getImplicitOffset = (videoFrames) => {
  return Math.max(...videoFrames.map((videoFrame, index) => index - videoFrame.frame));
};
const getVirtualSequence = (videoSequence) => {
  const len = videoSequence.frames.length;
  for (let index = 0; index < videoSequence.frames.length; index++) {
    const frame = __spreadValues$2({}, videoSequence.frames[index]);
    if (frame.isKeyframe && index > 0) {
      break;
    }
    frame.frame += len;
    videoSequence.frames.push(frame);
  }
  return videoSequence;
};

let eventHandlers = {};
const addListener = (node, event, handler, capture = false) => {
  var _a;
  if (!(event in eventHandlers)) {
    eventHandlers[event] = [];
  }
  (_a = eventHandlers[event]) == null ? void 0 : _a.push({ node, handler, capture });
  node.addEventListener(event, handler, capture);
};
const removeListener = (targetNode, event, targetHandler) => {
  var _a, _b, _c, _d;
  if (!(event in eventHandlers))
    return;
  if (targetHandler) {
    (_a = eventHandlers[event]) == null ? void 0 : _a.filter(({ node, handler }) => node === targetNode && handler === targetHandler).forEach(({ node, handler, capture }) => node.removeEventListener(event, handler, capture));
    eventHandlers[event] = (_b = eventHandlers[event]) == null ? void 0 : _b.filter(
      ({ node, handler }) => !(node === targetNode && handler === targetHandler)
    );
  } else {
    (_c = eventHandlers[event]) == null ? void 0 : _c.filter(({ node }) => node === targetNode).forEach(({ node, handler, capture }) => node.removeEventListener(event, handler, capture));
    eventHandlers[event] = (_d = eventHandlers[event]) == null ? void 0 : _d.filter(({ node }) => node !== targetNode);
  }
};
const removeAllListeners = (targetNode) => {
  Object.keys(eventHandlers).forEach((event) => {
    var _a, _b;
    const videoEvent = event;
    (_a = eventHandlers[videoEvent]) == null ? void 0 : _a.filter(({ node }) => node === targetNode).forEach(({ node, handler, capture }) => node.removeEventListener(videoEvent, handler, capture));
    eventHandlers[videoEvent] = (_b = eventHandlers[videoEvent]) == null ? void 0 : _b.filter(({ node }) => node !== targetNode);
  });
};

var __defProp$4 = Object.defineProperty;
var __getOwnPropDesc$3 = Object.getOwnPropertyDescriptor;
var __decorateClass$3 = (decorators, target, key, kind) => {
  var result = kind > 1 ? void 0 : kind ? __getOwnPropDesc$3(target, key) : target;
  for (var i = decorators.length - 1, decorator; i >= 0; i--)
    if (decorator = decorators[i])
      result = (kind ? decorator(target, key, result) : decorator(result)) || result;
  if (kind && result)
    __defProp$4(target, key, result);
  return result;
};
const IS_WECHAT = navigator && /MicroMessenger/i.test(navigator.userAgent);
const playVideoElement = async (videoElement) => {
  if (IS_WECHAT && window.WeixinJSBridge) {
    await getWechatNetwork();
  }
  try {
    await videoElement.play();
  } catch (error) {
    throw new Error(error.message);
  }
};
let VideoReader = class {
  constructor(videoSequence) {
    this.destroyed = false;
    this.frameRate = 0;
    this._duration = videoSequence.frameCount / videoSequence.frameRate;
    this.frameRate = videoSequence.frameRate;
  }
  static create(videoSequence) {
    const videoReader = new VideoReader(videoSequence);
    const debugData = videoReader.load(videoSequence);
    return { videoReader, debugData };
  }
  getVideoElement() {
    return this.videoElement;
  }
  progress() {
    return Math.round(this.videoElement.currentTime / this._duration * 100) / 100;
  }
  duration() {
    return this._duration;
  }
  currentTime() {
    return this.videoElement.currentTime || 0;
  }
  start() {
    return playVideoElement(this.videoElement);
  }
  pause() {
    var _a;
    (_a = this.videoElement) == null ? void 0 : _a.pause();
  }
  seek(time) {
    return new Promise((resolve) => {
      const seekCallback = () => {
        removeListener(this.videoElement, "seeked", seekCallback);
        resolve();
      };
      addListener(this.videoElement, "seeked", seekCallback);
      this.videoElement.currentTime = time;
    });
  }
  addListener(event, handler) {
    addListener(this.videoElement, event, handler);
  }
  removeAllListeners() {
    removeAllListeners(this.videoElement);
  }
  getFrameData(callback) {
  }
  clearCallback() {
  }
  destroy() {
    this.removeAllListeners();
    this.videoElement = void 0;
    this.destroyed = true;
  }
  load(videoSequence) {
    this.videoElement = document.createElement("video");
    this.videoElement.style.display = "none";
    this.videoElement.muted = true;
    this.videoElement.playsInline = true;
    const clock = new Clock();
    const mp4Data = coverToMp4(videoSequence);
    clock.mark("coverMP4");
    this.videoElement.src = URL.createObjectURL(new Blob([mp4Data], { type: "video/mp4" }));
    this.videoElement.load();
    return {
      coverMP4: clock.measure("", "coverMP4")
    };
  }
};
VideoReader = __decorateClass$3([
  destroyVerify
], VideoReader);

var __defProp$3 = Object.defineProperty;
var __getOwnPropDesc$2 = Object.getOwnPropertyDescriptor;
var __decorateClass$2 = (decorators, target, key, kind) => {
  var result = kind > 1 ? void 0 : kind ? __getOwnPropDesc$2(target, key) : target;
  for (var i = decorators.length - 1, decorator; i >= 0; i--)
    if (decorator = decorators[i])
      result = (kind ? decorator(target, key, result) : decorator(result)) || result;
  if (kind && result)
    __defProp$3(target, key, result);
  return result;
};
navigator && /MicroMessenger/i.test(navigator.userAgent);
let View = class extends Context {
  constructor(pagFile, canvas, options) {
    super(pagFile, canvas, options);
    this.fpsBuffer = [];
    this.currentFrame = -1;
    this.needSeek = false;
    this.videoReader = this.createVideoReader(this.videoSequence);
  }
  async play() {
    if (this.playing)
      return;
    this.playing = true;
    await this.videoReader.start();
    await this.flushLoop();
    if (this.getProgress() === 0) {
      this.eventManager.emit(EventName.onAnimationStart);
    }
    this.eventManager.emit(EventName.onAnimationPlay);
  }
  pause() {
    if (!this.playing)
      return;
    this.videoReader.pause();
    this.clearTimer();
    this.playing = false;
    this.eventManager.emit(EventName.onAnimationPause);
  }
  stop() {
    this.videoReader.pause();
    this.videoReader.seek(0);
    this.clearRender();
    this.playing = false;
    this.eventManager.emit(EventName.onAnimationCancel);
  }
  destroy() {
    this.clearTimer();
    this.clearRender();
    this.canvas = null;
    this.videoReader.destroy();
    this.destroyed = true;
  }
  getProgress() {
    return this.currentFrame / this.videoSequence.frameCount;
  }
  setProgress(progress) {
    if (progress < 0 || progress > 1)
      throw new Error("progress must be between 0.0 and 1.0!");
    const currentFrame = Math.round(progress * this.videoSequence.frameCount);
    if (this.currentFrame !== currentFrame) {
      this.needSeek = true;
      this.currentFrame = currentFrame;
    }
  }
  flush() {
    return this.flushInternal(true);
  }
  draw() {
  }
  createVideoReader(videoSequence) {
    const { videoReader, debugData } = VideoReader.create(videoSequence);
    this.setDebugData(debugData);
    if (!IS_IOS) {
      videoReader.addListener("ended", () => {
        this.repeat();
      });
    }
    return videoReader;
  }
  async repeat() {
    if (this.repeatCount === 0) {
      this.setProgress(1);
      await this.flushInternal(true);
      this.videoReader.pause();
      this.clearTimer();
      this.playing = false;
      this.eventManager.emit("onAnimationEnd");
      return false;
    }
    this.repeatCount -= 1;
    if (IS_IOS) {
      await this.videoReader.seek(0);
    } else {
      this.videoReader.start();
    }
    this.eventManager.emit("onAnimationRepeat");
    return true;
  }
  flushLoop() {
    this.renderTimer = window.requestAnimationFrame(() => {
      this.flushLoop();
    });
    if (IS_IOS && this.duration() - this.videoReader.currentTime() <= 1 / this.frameRate()) {
      this.repeat();
    }
    return this.flushInternal(false);
  }
  clearTimer() {
    if (this.renderTimer) {
      window.cancelAnimationFrame(this.renderTimer);
      this.renderTimer = null;
    }
  }
  updateFPS() {
    let now;
    try {
      now = performance.now();
    } catch (e) {
      now = Date.now();
    }
    this.fpsBuffer = this.fpsBuffer.filter((value) => now - value <= 1e3);
    this.fpsBuffer.push(now);
    this.setDebugData({ FPS: this.fpsBuffer.length });
  }
  async flushInternal(sync) {
    const clock = new Clock();
    if (this.needSeek) {
      if (sync) {
        await this.videoReader.seek(this.currentFrame / this.frameRate());
      } else {
        this.videoReader.seek(this.currentFrame / this.frameRate());
      }
      this.needSeek = false;
    } else {
      this.currentFrame = Math.floor(this.videoReader.currentTime() * this.frameRate());
    }
    this.draw();
    clock.mark("draw");
    this.setDebugData({ draw: clock.measure("", "draw") });
    this.updateFPS();
    this.eventManager.emit(EventName.onAnimationUpdate);
  }
};
View = __decorateClass$2([
  destroyVerify
], View);

var __defProp$2 = Object.defineProperty;
var __getOwnPropDesc$1 = Object.getOwnPropertyDescriptor;
var __decorateClass$1 = (decorators, target, key, kind) => {
  var result = kind > 1 ? void 0 : kind ? __getOwnPropDesc$1(target, key) : target;
  for (var i = decorators.length - 1, decorator; i >= 0; i--)
    if (decorator = decorators[i])
      result = (kind ? decorator(target, key, result) : decorator(result)) || result;
  if (kind && result)
    __defProp$2(target, key, result);
  return result;
};
let PAG2dView = class extends View {
  constructor(pagFile, canvas, options) {
    var _a;
    super(pagFile, canvas, options);
    const context = (_a = this.canvas) == null ? void 0 : _a.getContext("2d");
    if (!context)
      throw new Error("Can't get 2d context!");
    this.context = context;
    this.renderCanvas2D = document.createElement("canvas");
    this.renderCanvas2D.width = this.videoParam.MP4Width;
    this.renderCanvas2D.height = this.videoParam.MP4Height;
    const renderCanvas2DContext = this.renderCanvas2D.getContext("2d");
    if (!renderCanvas2DContext)
      throw new Error("Can't get 2d context!");
    this.renderCanvas2DContext = renderCanvas2DContext;
  }
  draw() {
    if (this.videoParam.hasAlpha) {
      this.renderCanvas2DContext.clearRect(0, 0, this.renderCanvas2D.width, this.renderCanvas2D.height);
      this.renderCanvas2DContext.drawImage(
        this.videoReader.getVideoElement(),
        0,
        0,
        this.renderCanvas2D.width,
        this.renderCanvas2D.height
      );
      const frameOne = this.renderCanvas2DContext.getImageData(
        0,
        0,
        this.videoParam.sequenceWidth,
        this.videoParam.sequenceHeight
      );
      const frameTwo = this.renderCanvas2DContext.getImageData(
        this.videoParam.alphaStartX,
        this.videoParam.alphaStartY,
        this.videoParam.sequenceWidth,
        this.videoParam.sequenceHeight
      );
      const length = frameOne.data.length / 4;
      for (let i = 0; i < length; i++) {
        frameOne.data[i * 4 + 3] = frameTwo.data[i * 4 + 0];
      }
      this.renderCanvas2DContext.clearRect(0, 0, this.renderCanvas2D.width, this.renderCanvas2D.height);
      this.renderCanvas2DContext.putImageData(
        frameOne,
        0,
        0,
        0,
        0,
        this.videoParam.sequenceWidth,
        this.videoParam.sequenceHeight
      );
      this.context.clearRect(0, 0, this.canvas.width, this.canvas.height);
      this.context.drawImage(
        this.renderCanvas2D,
        0,
        0,
        this.videoParam.sequenceWidth,
        this.videoParam.sequenceHeight,
        this.viewportSize.x,
        this.viewportSize.y,
        this.viewportSize.width,
        this.viewportSize.height
      );
    } else {
      this.context.drawImage(
        this.videoReader.getVideoElement(),
        0,
        0,
        this.videoParam.MP4Width,
        this.videoParam.MP4Height,
        this.viewportSize.x,
        this.viewportSize.y,
        this.viewportSize.width,
        this.viewportSize.height
      );
    }
  }
  clearRender() {
    this.context.clearRect(0, 0, this.canvas.width, this.canvas.height);
  }
};
PAG2dView = __decorateClass$1([
  destroyVerify
], PAG2dView);

const VERTEX_2D_SHADER = `
      attribute vec2 a_position;
      attribute vec2 a_texCoord;
      
      uniform vec2 u_resolution;
      uniform vec2 u_scale;
      
      varying vec2 v_texCoord;
    
      
      void main() {
         // convert the rectangle from pixels to 0.0 to 1.0
         vec2 zeroToOne = a_position / u_resolution;
      
         // convert from 0->1 to 0->2
         vec2 zeroToTwo = zeroToOne * 2.0;
      
         // convert from 0->2 to -1->+1 (clipspace)
         vec2 clipSpace = zeroToTwo - 1.0;
      
         gl_Position = vec4(clipSpace * vec2(1, -1), 0, 1);
      
         // pass the texCoord to the fragment shader
         // The GPU will interpolate this value between points.
         v_texCoord = a_texCoord / u_scale;
      }
        `;
const FRAGMENT_2D_SHADER = `
      precision mediump float;
      // our texture
      uniform sampler2D u_image;
      
      // the texCoords passed in from the vertex shader.
      varying vec2 v_texCoord;
      
      void main() {
         gl_FragColor = texture2D(u_image, v_texCoord);
      }
         `;
const FRAGMENT_2D_SHADER_TRANSPARENT = `
      precision mediump float;
      // our texture
      uniform sampler2D u_image;
      
      // the texCoords passed in from the vertex shader.
      varying vec2 v_texCoord;
      uniform vec2 v_alphaStart;
      
      void main() {
         vec4 color = texture2D(u_image, v_texCoord);
         vec4 alpha = texture2D(u_image, vec2(v_texCoord.x + v_alphaStart.x, v_texCoord.y + v_alphaStart.y));
         gl_FragColor = vec4(color.rgb * alpha.r, alpha.r);
      }     
         `;

var __defProp$1 = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __getOwnPropSymbols$1 = Object.getOwnPropertySymbols;
var __hasOwnProp$1 = Object.prototype.hasOwnProperty;
var __propIsEnum$1 = Object.prototype.propertyIsEnumerable;
var __defNormalProp$1 = (obj, key, value) => key in obj ? __defProp$1(obj, key, { enumerable: true, configurable: true, writable: true, value }) : obj[key] = value;
var __spreadValues$1 = (a, b) => {
  for (var prop in b || (b = {}))
    if (__hasOwnProp$1.call(b, prop))
      __defNormalProp$1(a, prop, b[prop]);
  if (__getOwnPropSymbols$1)
    for (var prop of __getOwnPropSymbols$1(b)) {
      if (__propIsEnum$1.call(b, prop))
        __defNormalProp$1(a, prop, b[prop]);
    }
  return a;
};
var __decorateClass = (decorators, target, key, kind) => {
  var result = kind > 1 ? void 0 : kind ? __getOwnPropDesc(target, key) : target;
  for (var i = decorators.length - 1, decorator; i >= 0; i--)
    if (decorator = decorators[i])
      result = (kind ? decorator(target, key, result) : decorator(result)) || result;
  if (kind && result)
    __defProp$1(target, key, result);
  return result;
};
let PAGWebGLView = class extends View {
  constructor(pagFile, canvas, options) {
    var _a;
    super(pagFile, canvas, options);
    this.scale = { x: 1, y: 1 };
    this.positionLocation = 0;
    this.texcoordLocation = 0;
    this.alphaStartLocation = null;
    this.scaleLocation = null;
    this.resolutionLocation = null;
    this.positionBuffer = null;
    this.texcoordBuffer = null;
    this.originalVideoTexture = null;
    this.renderingTexture = null;
    this.renderingFbo = null;
    const gl = (_a = this.canvas) == null ? void 0 : _a.getContext("webgl", __spreadValues$1({}, WEBGL_CONTEXT_ATTRIBUTES));
    if (!gl)
      throw new Error("Can't get WebGL context!");
    this.gl = gl;
    if (this.videoParam.hasAlpha) {
      this.program = createProgram(
        this.gl,
        getShaderSourceFromString(VERTEX_2D_SHADER),
        getShaderSourceFromString(FRAGMENT_2D_SHADER_TRANSPARENT)
      );
    } else {
      this.program = createProgram(
        this.gl,
        getShaderSourceFromString(VERTEX_2D_SHADER),
        getShaderSourceFromString(FRAGMENT_2D_SHADER)
      );
    }
    this.loadContext();
  }
  loadContext() {
    if (!this.program)
      throw new Error("program is not initialized");
    this.positionLocation = this.gl.getAttribLocation(this.program, "a_position");
    if (this.positionLocation === -1)
      throw new Error("unable to get attribute location for a_position");
    this.scaleLocation = this.gl.getUniformLocation(this.program, "u_scale");
    if (this.scaleLocation === -1)
      throw new Error("unable to get attribute location for u_scale");
    this.texcoordLocation = this.gl.getAttribLocation(this.program, "a_texCoord");
    if (this.texcoordLocation === -1)
      throw new Error("unable to get attribute location for a_texCoord");
    if (this.videoParam.hasAlpha) {
      this.alphaStartLocation = this.gl.getUniformLocation(this.program, "v_alphaStart");
      if (!this.alphaStartLocation)
        throw new Error("unable to get attribute location for v_alphaStart");
    }
    this.resolutionLocation = this.gl.getUniformLocation(this.program, "u_resolution");
    if (this.positionLocation === -1)
      throw new Error("unable to get attribute location for u_resolution");
    this.positionBuffer = this.gl.createBuffer();
    this.gl.bindBuffer(this.gl.ARRAY_BUFFER, this.positionBuffer);
    this.setRectangle(this.gl, 0, 0, this.videoParam.MP4Width, this.videoParam.MP4Height);
    this.texcoordBuffer = this.gl.createBuffer();
    this.gl.bindBuffer(this.gl.ARRAY_BUFFER, this.texcoordBuffer);
    this.gl.bufferData(
      this.gl.ARRAY_BUFFER,
      new Float32Array([0, 0, 1, 0, 0, 1, 0, 1, 1, 0, 1, 1]),
      this.gl.STATIC_DRAW
    );
    this.renderingTexture = createAndSetupTexture(this.gl);
    this.gl.texImage2D(
      this.gl.TEXTURE_2D,
      0,
      this.gl.RGBA,
      this.videoParam.sequenceWidth,
      this.videoParam.sequenceHeight,
      0,
      this.gl.RGBA,
      this.gl.UNSIGNED_BYTE,
      null
    );
    this.renderingFbo = this.gl.createFramebuffer();
    if (!this.renderingFbo)
      throw new Error("unable to create framebuffer");
    this.gl.bindFramebuffer(this.gl.FRAMEBUFFER, this.renderingFbo);
    this.gl.framebufferTexture2D(
      this.gl.FRAMEBUFFER,
      this.gl.COLOR_ATTACHMENT0,
      this.gl.TEXTURE_2D,
      this.renderingTexture,
      0
    );
    this.originalVideoTexture = createAndSetupTexture(this.gl);
  }
  draw() {
    this.gl.bindTexture(this.gl.TEXTURE_2D, this.originalVideoTexture);
    this.texImage2D();
    this.gl.clearColor(0, 0, 0, 0);
    this.gl.clear(this.gl.COLOR_BUFFER_BIT);
    this.gl.useProgram(this.program);
    this.gl.enableVertexAttribArray(this.positionLocation);
    this.gl.bindBuffer(this.gl.ARRAY_BUFFER, this.positionBuffer);
    const size = 2;
    const type = this.gl.FLOAT;
    const normalize = false;
    const stride = 0;
    const offset = 0;
    this.gl.vertexAttribPointer(this.positionLocation, size, type, normalize, stride, offset);
    this.gl.enableVertexAttribArray(this.texcoordLocation);
    this.gl.bindBuffer(this.gl.ARRAY_BUFFER, this.texcoordBuffer);
    this.gl.vertexAttribPointer(this.texcoordLocation, size, type, normalize, stride, offset);
    if (this.videoParam.hasAlpha) {
      this.gl.uniform2f(
        this.alphaStartLocation,
        this.videoParam.alphaStartX / this.videoParam.MP4Width / this.scale.x,
        this.videoParam.alphaStartY / this.videoParam.MP4Height / this.scale.y
      );
    }
    this.gl.bindTexture(this.gl.TEXTURE_2D, this.originalVideoTexture);
    this.gl.bindFramebuffer(this.gl.FRAMEBUFFER, this.renderingFbo);
    this.gl.uniform2f(this.resolutionLocation, this.videoParam.sequenceWidth, this.videoParam.sequenceHeight);
    this.gl.uniform2f(this.scaleLocation, this.scale.x, this.scale.y);
    this.gl.viewport(0, 0, this.videoParam.sequenceWidth, this.videoParam.sequenceHeight);
    const primitiveType = this.gl.TRIANGLES;
    const count = 6;
    this.gl.drawArrays(primitiveType, offset, count);
    this.gl.bindFramebuffer(this.gl.FRAMEBUFFER, null);
    this.gl.uniform2f(this.resolutionLocation, this.videoParam.sequenceWidth, this.videoParam.sequenceHeight);
    this.gl.viewport(this.viewportSize.x, this.viewportSize.y, this.viewportSize.width, this.viewportSize.height);
    this.gl.drawArrays(primitiveType, offset, count);
  }
  clearRender() {
    this.gl.clearColor(0, 0, 0, 0);
    this.gl.clear(this.gl.COLOR_BUFFER_BIT);
  }
  detectWebGLContext() {
    return detectWebGLContext();
  }
  texImage2D() {
    this.gl.texImage2D(
      this.gl.TEXTURE_2D,
      0,
      this.gl.RGBA,
      this.gl.RGBA,
      this.gl.UNSIGNED_BYTE,
      this.videoReader.getVideoElement()
    );
  }
  setRectangle(gl, x, y, width, height) {
    const x1 = x;
    const x2 = x + width;
    const y1 = y;
    const y2 = y + height;
    gl.bufferData(gl.ARRAY_BUFFER, new Float32Array([x1, y1, x2, y1, x1, y2, x1, y2, x2, y1, x2, y2]), gl.STATIC_DRAW);
  }
};
PAGWebGLView = __decorateClass([
  destroyVerify
], PAGWebGLView);

var __defProp = Object.defineProperty;
var __getOwnPropSymbols = Object.getOwnPropertySymbols;
var __hasOwnProp = Object.prototype.hasOwnProperty;
var __propIsEnum = Object.prototype.propertyIsEnumerable;
var __defNormalProp = (obj, key, value) => key in obj ? __defProp(obj, key, { enumerable: true, configurable: true, writable: true, value }) : obj[key] = value;
var __spreadValues = (a, b) => {
  for (var prop in b || (b = {}))
    if (__hasOwnProp.call(b, prop))
      __defNormalProp(a, prop, b[prop]);
  if (__getOwnPropSymbols)
    for (var prop of __getOwnPropSymbols(b)) {
      if (__propIsEnum.call(b, prop))
        __defNormalProp(a, prop, b[prop]);
    }
  return a;
};
class PAGView {
  static init(data, canvas, options = {}) {
    const opts = __spreadValues({
      renderingMode: RenderingMode.WebGL
    }, options);
    const clock = new Clock();
    const pagFile = PAGFile.fromArrayBuffer(data);
    clock.mark("decode");
    let pagView;
    if (opts.renderingMode === RenderingMode.WebGL) {
      pagView = new PAGWebGLView(pagFile, canvas, opts);
    } else {
      pagView = new PAG2dView(pagFile, canvas, opts);
    }
    pagView.setDebugData({ decodePAGFile: clock.measure("", "decode") });
    return pagView;
  }
}

export { PAGView, types };

