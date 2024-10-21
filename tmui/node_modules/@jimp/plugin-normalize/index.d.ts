import { ImageCallback } from '@jimp/core';

interface Normalize {
  normalize(cb ?: ImageCallback<this>): this;
}

export default function(): Normalize;
