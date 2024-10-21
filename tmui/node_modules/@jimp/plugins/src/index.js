import { mergeDeep } from 'timm';

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

const plugins = [
  blit,
  blur,
  circle,
  color,
  contain,
  cover,
  crop,
  displace,
  dither,
  fisheye,
  flip,
  gaussian,
  invert,
  mask,
  normalize,
  print,
  resize,
  rotate,
  scale,
  shadow,
  threshold
];

export default jimpEvChange => {
  const initializedPlugins = plugins.map(pluginModule => {
    let plugin = pluginModule(jimpEvChange) || {};

    if (!plugin.class && !plugin.constants) {
      // Default to class function
      plugin = { class: plugin };
    }

    return plugin;
  });

  return mergeDeep(...initializedPlugins);
};
