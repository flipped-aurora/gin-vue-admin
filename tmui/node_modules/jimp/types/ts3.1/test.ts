import * as Jimp from 'jimp';

const jimpInst: Jimp = new Jimp('test');

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
Jimp.read('Test');

// $ExpectType 0
Jimp.PNG_FILTER_NONE;

// $ExpectError
Jimp.test;

// $ExpectError
Jimp.func();

test('can clone properly', async () => {
  const baseImage = await Jimp.read('filename');
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
