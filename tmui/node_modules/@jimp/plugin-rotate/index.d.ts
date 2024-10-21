import { ImageCallback } from '@jimp/core';

interface Rotate {
  rotate(deg: number, cb?: ImageCallback<this>): this;
  rotate(deg: number, mode: string | boolean, cb?: ImageCallback<this>): this;
}

export default function(): Rotate;
