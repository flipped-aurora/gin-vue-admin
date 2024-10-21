import { DecoderFn, EncoderFn } from '@jimp/core';

interface Bmp {
  constants: {
    MIME_BMP: 'image/bmp';
    MIME_X_MS_BMP: 'image/x-ms-bmp';
  }

  mime: {
    'image/bmp': string[]
  }

  decoders: {
    'image/bmp': DecoderFn
    'image/x-ms-bmp': DecoderFn
  }

  encoders: {
    'image/bmp': EncoderFn
    'image/x-ms-bmp': EncoderFn
  }
}

export default function(): Bmp;
