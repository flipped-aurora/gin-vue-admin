import { GenericCallback, ImageCallback } from '@jimp/core';

export interface FontChar {
  id: number;
  x: number;
  y: number;
  width: number;
  height: number;
  xoffset: number;
  yoffset: number;
  xadvance: number;
  page: number;
  chnl: number;
}

export interface FontInfo {
  face: string;
  size: number;
  bold: number;
  italic: number;
  charset: string;
  unicode: number;
  stretchH: number;
  smooth: number;
  aa: number;
  padding: [number, number, number, number];
  spacing: [number, number];
}

export interface FontCommon {
  lineHeight: number;
  base: number;
  scaleW: number;
  scaleH: number;
  pages: number;
  packed: number;
  alphaChnl: number;
  redChnl: number;
  greenChnl: number;
  blueChnl: number;
}

export interface Font {
  chars: {
    [char: string]: FontChar;
  };
  kernings: {
    [firstString: string]: {
      [secondString: string]: number;
    };
  };
  pages: string[];
  common: FontCommon;
  info: FontInfo;
}

type PrintableText =
  | any
  | {
  text: string;
  alignmentX: number;
  alignmentY: number;
};

interface PrintClass {
  // Text methods
  print(
    font: Font,
    x: number,
    y: number,
    text: PrintableText,
    cb?: ImageCallback<this>
  ): this;
  print(
    font: Font,
    x: number,
    y: number,
    text: PrintableText,
    maxWidth?: number,
    cb?: ImageCallback<this>
  ): this;
  print(
    font: Font,
    x: number,
    y: number,
    text: PrintableText,
    maxWidth?: number,
    maxHeight?: number,
    cb?: ImageCallback<this>
  ): this;
}

interface Print {
  constants: {
    measureText(font: Font, text: PrintableText): number;
    measureTextHeight(font: Font, text: PrintableText, maxWidth: number): number;

    // Font locations
    FONT_SANS_8_BLACK: string;
    FONT_SANS_10_BLACK: string;
    FONT_SANS_12_BLACK: string;
    FONT_SANS_14_BLACK: string;
    FONT_SANS_16_BLACK: string;
    FONT_SANS_32_BLACK: string;
    FONT_SANS_64_BLACK: string;
    FONT_SANS_128_BLACK: string;

    FONT_SANS_8_WHITE: string;
    FONT_SANS_16_WHITE: string;
    FONT_SANS_32_WHITE: string;
    FONT_SANS_64_WHITE: string;
    FONT_SANS_128_WHITE: string;

    loadFont(file: string): Promise<Font>;
    loadFont(file: string, cb: GenericCallback<Font, any, any>): Promise<never>;
  }

  class: PrintClass
}

export default function(): Print;
