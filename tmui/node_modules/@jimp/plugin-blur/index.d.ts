import { ImageCallback } from '@jimp/core';

interface Blur {
  blur(r: number, cb?: ImageCallback<this>): this;
}

export default function(): Blur;
