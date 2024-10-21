/**
 * These files pertain to the typings of plugins or types
 * 
 * They're not meant as utils to decode types, but rather
 * the type definitons themsleves for plugins and types of various kinds
 */
import {
  Image,
  EncoderFn,
  DecoderFn
} from './etc';

import {
  Omit
} from './utils';

export type IllformedPlugin = Omit<any, 'class' | 'constants'> & {
  class?: never;
  constants?: never;
}

export interface WellFormedPlugin<ImageType extends Image = Image> {
  mime?: {
    [MIME_TYPE: string]: string[];
  };
  hasAlpha?: {
    [MIME_SPECIAL: string]: boolean;
  };
  constants?: {
    // Contants to assign to the Jimp instance
    [MIME_SPECIAL: string]: any;
  };
  decoders?: {
    [MIME_TYPE: string]: DecoderFn;
  };
  encoders?: {
    // Jimp Image
    [MIME_TYPE: string]: EncoderFn<ImageType>;
  };
  // Extend the Jimp class with the following constants, etc
  class?: any;
}

type ClassOrConstantPlugin<T extends Image> = WellFormedPlugin<T> &
  (
    | Required<Pick<WellFormedPlugin<T>, 'class'>>
    | Required<Pick<WellFormedPlugin<T>, 'constants'>>
  );

// A Jimp type requires mime, but not class
export type JimpType<T extends Image = Image> = WellFormedPlugin<T> &
  Required<Pick<WellFormedPlugin<T>, 'mime'>>;

// Jimp plugin either MUST have class OR constant or be illformed
export type JimpPlugin<T extends Image = Image> =
  | ClassOrConstantPlugin<T>
  | IllformedPlugin;
