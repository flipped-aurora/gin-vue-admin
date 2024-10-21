import {
  Bitmap,
  ImageCallback,
  URLOptions,
  ListenableName,
  ListenerData,
  GenericCallback,
  BlendMode,
  RGBA,
  RGB
} from './etc';

interface DiffReturn<This> {
  percent: number;
  image: This;
}

interface ScanIteratorReturn<This> {
  x: number;
  y: number;
  idx: number;
  image: This;
}

export interface JimpConstructors {
  prototype: Jimp;
  // Constants
  AUTO: -1;
  // blend modes
  BLEND_SOURCE_OVER: string;
  BLEND_DESTINATION_OVER: string;
  BLEND_MULTIPLY: string;
  BLEND_SCREEN: string;
  BLEND_OVERLAY: string;
  BLEND_DARKEN: string;
  BLEND_LIGHTEN: string;
  BLEND_HARDLIGHT: string;
  BLEND_DIFFERENCE: string;
  BLEND_EXCLUSION: string;
  // Align modes for cover, contain, bit masks
  HORIZONTAL_ALIGN_LEFT: 1;
  HORIZONTAL_ALIGN_CENTER: 2;
  HORIZONTAL_ALIGN_RIGHT: 4;
  VERTICAL_ALIGN_TOP: 8;
  VERTICAL_ALIGN_MIDDLE: 16;
  VERTICAL_ALIGN_BOTTOM: 32;
  // Edge Handling
  EDGE_EXTEND: 1;
  EDGE_WRAP: 2;
  EDGE_CROP: 3;

  // Constructors
  new(path: string, cb?: ImageCallback<this['prototype']>): this['prototype'];
  new(urlOptions: URLOptions, cb?: ImageCallback<this['prototype']>): this['prototype'];
  new(image: Jimp, cb?: ImageCallback<this['prototype']>): this['prototype'];
  new(data: Buffer, cb?: ImageCallback<this['prototype']>): this['prototype'];
  new(data: Bitmap, cb?: ImageCallback<this['prototype']>): this['prototype'];
  new(w: number, h: number, cb?: ImageCallback<this['prototype']>): this['prototype'];
  new(
    w: number,
    h: number,
    background?: number | string,
    cb?: ImageCallback<this['prototype']>
  ): this['prototype'];
  // For custom constructors when using Jimp.appendConstructorOption
  new(...args: any[]): this['prototype'];

  // Functions
  /**
   * I'd like to make `Args` generic and used in `run` and `test` but alas,
   * it's not possible RN:
   * https://github.com/microsoft/TypeScript/issues/26113
   */
  appendConstructorOption<Args extends any[], J extends Jimp = this['prototype']>(
    name: string,
    test: (...args: any[]) => boolean,
    run: (
      this: J,
      resolve: (jimp?: J) => any,
      reject: (reason: Error) => any,
      ...args: any[]
    ) => any
  ): void;
  read(path: string, cb?: ImageCallback<this['prototype']>): Promise<this['prototype']>;
  read(image: Jimp, cb?: ImageCallback<this['prototype']>): Promise<this['prototype']>;
  read(data: Buffer, cb?: ImageCallback<this['prototype']>): Promise<this['prototype']>;
  read(
    w: number,
    h: number,
    background?: number | string,
    cb?: ImageCallback<this['prototype']>
  ): Promise<this['prototype']>;
  create(path: string): Promise<this['prototype']>;
  create(image: Jimp): Promise<this['prototype']>;
  create(data: Buffer): Promise<this['prototype']>;
  create(w: number, h: number, background?: number | string): Promise<this['prototype']>;
  rgbaToInt(
    r: number,
    g: number,
    b: number,
    a: number,
    cb: GenericCallback<number, any, this['prototype']>
  ): number;
  intToRGBA(i: number, cb?: GenericCallback<RGBA>): RGBA;
  cssColorToHex(cssColor: string): number;
  limit255(n: number): number;
  diff(img1: Jimp, img2: Jimp, threshold?: number): DiffReturn<this['prototype']>;
  distance(img1: Jimp, img2: Jimp): number;
  compareHashes(hash1: string, hash2: string): number;
  colorDiff(rgba1: RGB, rgba2: RGB): number;
  colorDiff(rgba1: RGBA, rgba2: RGBA): number;
}

export interface Jimp {
  // Properties
  bitmap: Bitmap;
  _rgba: boolean;
  _background: number;
  _originalMime: string;
  // Methods
  on<T extends ListenableName>(
    event: T,
    cb: (data: ListenerData<T>) => any
  ): any;
  parseBitmap(
    data: Buffer,
    path: string | null | undefined,
    cb?: ImageCallback<this>
  ): void;
  hasAlpha(): boolean;
  getHeight(): number;
  getWidth(): number;
  inspect(): string;
  toString(): string;
  getMIME(): string;
  getExtension(): string;
  distanceFromHash(hash: string): number;
  write(path: string, cb?: ImageCallback<this>): this;
  writeAsync(path: string): Promise<this>;
  rgba(bool: boolean, cb?: ImageCallback<this>): this;
  getBase64(mime: string, cb: GenericCallback<string, any, this>): this;
  getBase64Async(mime: string): Promise<string>;
  hash(cb?: GenericCallback<string, any, this>): string;
  hash(
    base: number | null | undefined,
    cb?: GenericCallback<string, any, this>
  ): string;
  getBuffer(mime: string, cb: GenericCallback<Buffer>): this;
  getBufferAsync(mime: string): Promise<Buffer>;
  getPixelIndex(
    x: number,
    y: number,
    cb?: GenericCallback<number, any, this>
  ): number;
  getPixelIndex(
    x: number,
    y: number,
    edgeHandling: string,
    cb?: GenericCallback<number, any, this>
  ): number;
  getPixelColor(
    x: number,
    y: number,
    cb?: GenericCallback<number, any, this>
  ): number;
  getPixelColour(
    x: number,
    y: number,
    cb?: GenericCallback<number, any, this>
  ): number;
  setPixelColor(
    hex: number,
    x: number,
    y: number,
    cb?: ImageCallback<this>
  ): this;
  setPixelColour(
    hex: number,
    x: number,
    y: number,
    cb?: ImageCallback<this>
  ): this;
  clone(cb?: ImageCallback<this>): this;
  cloneQuiet(cb?: ImageCallback<this>): this;
  background(hex: number, cb?: ImageCallback<this>): this;
  backgroundQuiet(hex: number, cb?: ImageCallback<this>): this;
  scan(
    x: number,
    y: number,
    w: number,
    h: number,
    f: (this: this, x: number, y: number, idx: number) => any,
    cb?: ImageCallback<this>
  ): this;
  scanQuiet(
    x: number,
    y: number,
    w: number,
    h: number,
    f: (this: this, x: number, y: number, idx: number) => any,
    cb?: ImageCallback<this>
  ): this;
  scanIterator(
    x: number,
    y: number,
    w: number,
    h: number
  ): IterableIterator<ScanIteratorReturn<this>>;

  // Effect methods
  composite(
    src: Jimp,
    x: number,
    y: number,
    options?: BlendMode,
    cb?: ImageCallback<this>
  ): this;
}
