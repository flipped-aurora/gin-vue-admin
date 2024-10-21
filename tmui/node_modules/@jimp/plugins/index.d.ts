import blit from '@jimp/plugin-blit';
import blur from '@jimp/plugin-blur';
import circle from '@jimp/plugin-circle';
import color from '@jimp/plugin-color';
import contain from '@jimp/plugin-contain';
import cover from '@jimp/plugin-cover';
import crop from '@jimp/plugin-crop';
import displace from '@jimp/plugin-displace';
import dither from '@jimp/plugin-dither';
import fisheye from '@jimp/plugin-fisheye';
import flip from '@jimp/plugin-flip';
import gaussian from '@jimp/plugin-gaussian';
import invert from '@jimp/plugin-invert';
import mask from '@jimp/plugin-mask';
import normalize from '@jimp/plugin-normalize';
import print from '@jimp/plugin-print';
import resize from '@jimp/plugin-resize';
import rotate from '@jimp/plugin-rotate';
import scale from '@jimp/plugin-scale';
import shadow from '@jimp/plugin-shadow';
import threshold from '@jimp/plugin-threshold';

type BlitRet = ReturnType<typeof blit>;
type BlurRet = ReturnType<typeof blur>;
type CircleRet = ReturnType<typeof circle>;
type ColorRet = ReturnType<typeof color>;
type ContainRet = ReturnType<typeof contain>;
type CoverRet = ReturnType<typeof cover>;
type CropRet = ReturnType<typeof crop>;
type DisplaceRet = ReturnType<typeof displace>;
type DitherRet = ReturnType<typeof dither>;
type FlipRet = ReturnType<typeof flip>;
type FisheyeRet = ReturnType<typeof fisheye>;
type GaussianRet = ReturnType<typeof gaussian>;
type InvertRet = ReturnType<typeof invert>;
type MaskRet = ReturnType<typeof mask>;
type NormalizeRet = ReturnType<typeof normalize>;
type PrintRet = ReturnType<typeof print>;
type ResizeRet = ReturnType<typeof resize>;
type RotateRet = ReturnType<typeof rotate>;
type ScaleRet = ReturnType<typeof scale>;
type ShadowRet = ReturnType<typeof shadow>;
type ThresholdRet = ReturnType<typeof threshold>;

/**
 * This is made union and not intersection to avoid issues with
 * `IllformedPlugin` and `WellFormedPlugin` when using typings with Jimp
 * generic
 *
 * In reality, this should be an intersection but our type data isn't
 * clever enough to figure out what's a class and what's not/etc
 */
type Plugins =
  | BlitRet
  | BlurRet
  | CircleRet
  | ColorRet
  | ContainRet
  | CoverRet
  | CropRet
  | DisplaceRet
  | DitherRet
  | FlipRet
  | FisheyeRet
  | GaussianRet
  | InvertRet
  | MaskRet
  | NormalizeRet
  | PrintRet
  | ResizeRet
  | RotateRet
  | ScaleRet
  | ShadowRet
  | ThresholdRet;

export default function(): Plugins;
