import { ImageCallback } from '@jimp/core';

interface Contain {
  contain(w: number, h: number, cb?: ImageCallback<this>): this;
  contain(w: number, h: number, mode?: string, cb?: ImageCallback<this>): this;
  contain(w: number, h: number, alignBits?: number, cb?: ImageCallback<this>): this;
  contain(
    w: number,
    h: number,
    alignBits?: number,
    mode?: string,
    cb?: ImageCallback<this>
  ): this;
}

export default function(): Contain;
