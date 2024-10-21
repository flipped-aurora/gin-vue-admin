import { ImageCallback } from '@jimp/core';

type ColorActionName =
  | 'mix'
  | 'tint'
  | 'shade'
  | 'xor'
  | 'red'
  | 'green'
  | 'blue'
  | 'hue';

type ColorAction = {
  apply: ColorActionName;
  params: any;
};

interface Color {
  brightness(val: number, cb?: ImageCallback<this>): this;
  contrast(val: number, cb?: ImageCallback<this>): this;
  posterize(n: number, cb?: ImageCallback<this>): this;
  greyscale(cb?: ImageCallback<this>): this;
  grayscale(cb?: ImageCallback<this>): this;
  opacity(f: number, cb?: ImageCallback<this>): this;
  sepia(cb?: ImageCallback<this>): this;
  fade(f: number, cb?: ImageCallback<this>): this;
  convolution(kernel: number[][], cb?: ImageCallback<this>): this;
  convolution<T>(
    kernel: number[][],
    edgeHandling: string,
    cb?: ImageCallback<this>
  ): this;
  opaque(cb?: ImageCallback<this>): this;
  pixelate(size: number, cb?: ImageCallback<this>): this;
  pixelate(
    size: number,
    x: number,
    y: number,
    w: number,
    h: number,
    cb?: ImageCallback<this>
  ): this;
  convolute(kernel: number[][], cb?: ImageCallback<this>): this;
  convolute(
    kernel: number[][],
    x: number,
    y: number,
    w: number,
    h: number,
    cb?: ImageCallback<this>
  ): this;
  color(actions: ColorAction[], cb?: ImageCallback<this>): this;
  colour(actions: ColorAction[], cb?: ImageCallback<this>): this;
}

export default function(): Color;
