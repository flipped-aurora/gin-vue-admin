import { ImageCallback } from '@jimp/core';

interface Scale {
  scale(f: number, cb?: ImageCallback<this>): this;
  scale(f: number, mode?: string, cb?: ImageCallback<this>): this;
  scaleToFit(w: number, h: number, cb?: ImageCallback<this>): this;
  scaleToFit(w: number, h: number, mode?: string, cb?: ImageCallback<this>): this;
}

export default function(): Scale;
