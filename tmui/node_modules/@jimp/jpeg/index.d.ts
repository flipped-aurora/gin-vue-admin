import { DecoderFn, EncoderFn, ImageCallback } from '@jimp/core';

interface JpegClass {
  _quality: number;
  quality: (n: number, cb?: ImageCallback<this>) => this;
}

interface Jpeg {
  mime: { 'image/jpeg': string[] },

  constants: {
    MIME_JPEG: 'image/jpeg';
  }

  encoders: {
    'image/jpeg': EncoderFn
  }

  decoders: {
    'image/jpeg': DecoderFn
  }

  class: JpegClass
}

export default function(): Jpeg;
