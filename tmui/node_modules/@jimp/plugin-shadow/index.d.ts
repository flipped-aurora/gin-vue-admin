import { ImageCallback } from '@jimp/core';

interface Shadow {
  shadow(options?: {
           size?: number,
           opacity?: number,
           blur: number,
           x?: number,
           y?: number
         },
         cb?: ImageCallback<this>): this;
  shadow(cb?: ImageCallback<this>): this;
}

export default function(): Shadow;
