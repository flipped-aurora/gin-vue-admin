import configure from '@jimp/custom';
import gif from '@jimp/gif';
import png from '@jimp/png';
import displace from '@jimp/plugin-displace';
import resize from '@jimp/plugin-resize';
import scale from '@jimp/plugin-scale';
import types from '@jimp/types';
import plugins from '@jimp/plugins';
import * as Jimp from 'jimp';

// configure should return a valid Jimp type with addons
const CustomJimp = configure({
  types: [gif, png],
  plugins: [displace, resize]
});

test('should function the same as the `jimp` types', () => {
  const FullCustomJimp = configure({
    types: [types],
    plugins: [plugins]
  });

  const jimpInst = new FullCustomJimp('test');

  // Main Jimp export should already have all of these already applied
  // $ExpectError
  jimpInst.read('Test');
  jimpInst.displace(jimpInst, 2);
  jimpInst.resize(40, 40);
  jimpInst.displace(jimpInst, 2);
  jimpInst.shadow((err, val, coords) => {});
  jimpInst.fishEye({r: 12});
  jimpInst.circle({radius: 12, x: 12, y: 12});
  // $ExpectError
  jimpInst.PNG_FILTER_NONE;

  // $ExpectError
  jimpInst.test;

  // $ExpectError
  jimpInst.func();

  // Main Jimp export should already have all of these already applied
  FullCustomJimp.read('Test');

  // $ExpectType 0
  FullCustomJimp.PNG_FILTER_NONE;

  // $ExpectError
  FullCustomJimp.test;

  // $ExpectError
  FullCustomJimp.func();

  test('can clone properly', async () => {
    const baseImage = await FullCustomJimp.read('filename');
    const cloneBaseImage = baseImage.clone();

    // $ExpectType number
    cloneBaseImage._deflateLevel;

    test('can handle `this` returns on the core type properly', () => {
      // $ExpectType number
      cloneBaseImage.posterize(3)._quality
    });

    test('can handle `this` returns properly', () => {
      cloneBaseImage
        .resize(1, 1)
        .crop(0, 0, 0, 0)
        .mask(cloneBaseImage, 2, 2)
        .print('a' as any, 2, 2, 'a' as any)
        .resize(1, 1)
        .quality(1)
        .deflateLevel(2)
        ._filterType;
    });

    test('can handle imageCallbacks `this` properly', () => {
      cloneBaseImage.rgba(false, (_, jimpCBIn) => {
        // $ExpectError
        jimpCBIn.read('Test');
        jimpCBIn.displace(jimpInst, 2);
        jimpCBIn.resize(40, 40);
        // $ExpectType number
        jimpCBIn._filterType;

        // $ExpectError
        jimpCBIn.test;

        // $ExpectError
        jimpCBIn.func();
      })
    })
  });

  test('Can handle callback with constructor', () => {
    const myBmpBuffer: Buffer = {} as any;

    Jimp.read(myBmpBuffer, (err, cbJimpInst) => {
      // $ExpectError
      cbJimpInst.read('Test');
      cbJimpInst.displace(jimpInst, 2);
      cbJimpInst.resize(40, 40);
      // $ExpectType number
      cbJimpInst._filterType;

      // $ExpectError
      cbJimpInst.test;

      // $ExpectError
      cbJimpInst.func();
    });
  });

});

test('can handle custom jimp', () => {
  // Constants from types should be applied
  // $ExpectType 0
  CustomJimp.PNG_FILTER_NONE;
  
  // Core functions should still work from Jimp
  CustomJimp.read('Test');

  // Constants should not(?) be applied from ill-formed plugins
  // $ExpectError
  CustomJimp.displace(CustomJimp, 2);

  // Methods should be applied from well-formed plugins only to the instance
  // $ExpectError
  CustomJimp.resize(40, 40)
  
  // Constants should be applied from well-formed plugins
  CustomJimp.RESIZE_NEAREST_NEIGHBOR
  
  // $ExpectError
  CustomJimp.test;
  
  // $ExpectError
  CustomJimp.func();

  const Jiimp = new CustomJimp('test');
  // Methods from types should be applied
  Jiimp.deflateLevel(4);
  // Constants from types should be applied to the static only
  // $ExpectError
  Jiimp.PNG_FILTER_NONE;

  // Core functions should still work from Jimp
  Jiimp.getPixelColor(1, 1);

  // Constants should be applied from ill-formed plugins
  Jiimp.displace(Jiimp, 2);

  // Methods should be applied from well-formed plugins
  Jiimp.resize(40, 40)

  // Constants should not be applied to the object
  // $ExpectError
  Jiimp.RESIZE_NEAREST_NEIGHBOR

  // $ExpectError
  Jiimp.test;

  // $ExpectError
  Jiimp.func();
});

test('can compose', () => {
  const OtherCustomJimp = configure({
      plugins: [scale]
    }, CustomJimp);
  // Constants from types should be applied
  // $ExpectType 0
  OtherCustomJimp.PNG_FILTER_NONE;

  // Core functions should still work from Jimp
  OtherCustomJimp.read('Test');

  // Constants should not be applied to the static instance from ill-formed plugins
  // $ExpectError
  OtherCustomJimp.displace(OtherCustomJimp, 2);

  // Methods should not be applied to the static instance from well-formed plugins
  // $ExpectError
  OtherCustomJimp.resize(40, 40);

  // Constants should be applied from well-formed plugins
  OtherCustomJimp.RESIZE_NEAREST_NEIGHBOR;

  // $ExpectError
  OtherCustomJimp.test;

  // $ExpectError
  OtherCustomJimp.func();
  
  const Jiimp = new OtherCustomJimp('test');
  // Methods from types should be applied
  Jiimp.deflateLevel(4);
  // Constants from types should not be applied to objects
  // $ExpectError
  Jiimp.PNG_FILTER_NONE;

  // Methods from new plugins should be applied
  Jiimp.scale(3);

  // Methods from types should be applied
  Jiimp.filterType(4);

  // Core functions should still work from Jimp
  Jiimp.getPixelColor(1, 1);

  // Constants should be applied from ill-formed plugins
  Jiimp.displace(Jiimp, 2);

  // Methods should be applied from well-formed plugins
  Jiimp.resize(40, 40)

  // Constants should not be applied from well-formed plugins to objects
  // $ExpectError
  Jiimp.RESIZE_NEAREST_NEIGHBOR

  // $ExpectError
  Jiimp.test;

  // $ExpectError
  Jiimp.func();
});

test('can handle only plugins', () => {
  const PluginsJimp = configure({
    plugins: [plugins]
  });

  // Core functions should still work from Jimp
  PluginsJimp.read('Test');

  // Constants should not be applied from ill-formed plugins
  // $ExpectError
  PluginsJimp.displace(PluginsJimp, 2);

  // Methods should be not be applied to from well-formed plugins to the top level
  // $ExpectError
  PluginsJimp.resize(40, 40);

  // Constants should be applied from well-formed plugins
  // $ExpectType "nearestNeighbor"
  PluginsJimp.RESIZE_NEAREST_NEIGHBOR;

  // $ExpectError
  PluginsJimp.test;

  // $ExpectError
  PluginsJimp.func();

  const Jiimp = new PluginsJimp('test');
  
  // Core functions should still work from Jimp
  Jiimp.getPixelColor(1, 1);

  // Constants should be applied from ill-formed plugins
  Jiimp.displace(Jiimp, 2);

  // Methods should be applied from well-formed plugins
  Jiimp.resize(40, 40)

  // Constants should be not applied to objects from well-formed plugins
  // $ExpectError
  Jiimp.RESIZE_NEAREST_NEIGHBOR

  // $ExpectError
  Jiimp.test;

  // $ExpectError
  Jiimp.func();
})

test('can handle only all types', () => {
  const TypesJimp = configure({
    types: [types]
  });

  // Methods from types should not be applied
  // $ExpectError
  TypesJimp.filterType(4);
  // Constants from types should be applied
  // $ExpectType 0
  TypesJimp.PNG_FILTER_NONE;

  // $ExpectError
  TypesJimp.test;

  // $ExpectError
  TypesJimp.func();

  const Jiimp = new TypesJimp('test');
  // Methods from types should be applied
  Jiimp.filterType(4);
  // Constants from types should be not applied to objects
  // $ExpectError
  Jiimp.PNG_FILTER_NONE;

  // $ExpectError
  Jiimp.test;

  // $ExpectError
  Jiimp.func();
});

test('can handle only one type', () => {
  const PngJimp = configure({
    types: [png]
  });

  // Constants from other types should be not applied
  // $ExpectError
  PngJimp.MIME_TIFF;

  // Constants from types should be applied
  // $ExpectType 0
  PngJimp.PNG_FILTER_NONE;
  
  // $ExpectError
  PngJimp.test;

  // $ExpectError
  PngJimp.func();


  const Jiimp = new PngJimp('test');
  // Constants from other types should be not applied
  // $ExpectError
  Jiimp.MIME_TIFF;

  // Constants from types should not be applied to objects
  // $ExpectError
  Jiimp.PNG_FILTER_NONE;

  // $ExpectError
  Jiimp.test;

  // $ExpectError
  Jiimp.func();
});


test('can handle only one plugin', () => {
  const ResizeJimp = configure({
    plugins: [resize]
  });

  // Constants from other plugins should be not applied
  // $ExpectError
  ResizeJimp.FONT_SANS_8_BLACK;

  // Constants from plugin should be applied
  // $ExpectType "nearestNeighbor"
  ResizeJimp.RESIZE_NEAREST_NEIGHBOR;

  // $ExpectError
  ResizeJimp.resize(2, 2);

  // $ExpectError
  ResizeJimp.test;

  // $ExpectError
  ResizeJimp.func();


  const Jiimp: InstanceType<typeof ResizeJimp> = new ResizeJimp('test');
  // Constants from other plugins should be not applied
  // $ExpectError
  Jiimp.FONT_SANS_8_BLACK;

  // Constants from plugin should not be applied to the object
  // $ExpectError
  Jiimp.RESIZE_NEAREST_NEIGHBOR;

  Jiimp.resize(2, 2);

  // $ExpectError
  Jiimp.test;

  // $ExpectError
  Jiimp.func();
});


test('Can handle appendConstructorOption', () => {
  const AppendJimp = configure({});

  AppendJimp.appendConstructorOption(
    'Name of Option',
    args => args.hasSomeCustomThing,
    function(resolve, reject, args) {
      // $ExpectError
      this.bitmap = 3;
      // $ExpectError
      AppendJimp.resize(2, 2);
      resolve();
    }
  );
});
