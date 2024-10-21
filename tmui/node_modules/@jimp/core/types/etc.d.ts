import {Jimp} from './jimp';

export interface Image {
  bitmap: Bitmap;
}

export type DecoderFn = (data: Buffer) => Bitmap;
export type EncoderFn<ImageType extends Image = Image> = (
  image: ImageType
) => Buffer;

export type GenericCallback<T, U = any, TThis = any> = (
  this: TThis,
  err: Error | null,
  value: T
) => U;

/**
 * `jimp` must be defined otherwise `this` will not apply properly
 * for custom configurations where plugins and types are needed
 */
export type ImageCallback<jimp = Jimp> = (
  this: jimp,
  err: Error | null,
  value: jimp,
  coords: {
    x: number;
    y: number;
  }
) => any;

type BlendMode = {
  mode: string;
  opacitySource: number;
  opacityDest: number;
};

type ChangeName = 'background' | 'scan' | 'crop';

type ListenableName =
  | 'any'
  | 'initialized'
  | 'before-change'
  | 'changed'
  | 'before-clone'
  | 'cloned'
  | ChangeName;

type ListenerData<T extends ListenableName> = T extends 'any'
  ? any
  : T extends ChangeName
    ? {
        eventName: 'before-change' | 'changed';
        methodName: T;
        [key: string]: any;
      }
    : {
        eventName: T;
        methodName: T extends 'initialized'
          ? 'constructor'
          : T extends 'before-change' | 'changed'
            ? ChangeName
            : T extends 'before-clone' | 'cloned' ? 'clone' : any;
      };

type URLOptions = {
  url: string;
  compression?: boolean;
  headers: {
    [key: string]: any;
  };
};

export interface Bitmap {
  data: Buffer;
  width: number;
  height: number;
}

export interface RGB {
  r: number;
  g: number;
  b: number;
}

export interface RGBA {
  r: number;
  g: number;
  b: number;
  a: number;
}
