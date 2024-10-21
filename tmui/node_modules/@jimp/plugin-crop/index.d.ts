import { Jimp, ImageCallback } from '@jimp/core';

interface CropClass {
  crop(x: number, y: number, w: number, h: number, cb?: ImageCallback<this>): this;
  cropQuiet(
    x: number,
    y: number,
    w: number,
    h: number,
    cb?: ImageCallback<this>
  ): this;
  autocrop(tolerance?: number, cb?: ImageCallback<this>): this;
  autocrop(cropOnlyFrames?: boolean, cb?: ImageCallback<this>): this;
  autocrop(
    tolerance?: number,
    cropOnlyFrames?: boolean,
    cb?: ImageCallback<this>
  ): this;
  autocrop(
    options: {
      tolerance?: number;
      cropOnlyFrames?: boolean;
      cropSymmetric?: boolean;
      leaveBorder?: number;
    },
    cb?: ImageCallback<this>
  ): this;
}

interface Crop {
  class: CropClass
}

export default function(): Crop;
