import * as Jimp from 'jimp';

const jimpInst: Jimp = new Jimp('test');

// Main Jimp export should already have all of these already applied
jimpInst.read('Test');
jimpInst.displace(jimpInst, 2);
jimpInst.resize(40, 40);
// $ExpectType 0
jimpInst.PNG_FILTER_NONE;

// $ExpectError
jimpInst.test;

// $ExpectError
jimpInst.func();

// Main Jimp export should already have all of these already applied
Jimp.read('Test');
Jimp.displace(Jimp, 2);
Jimp.shadow((err, val, coords) => {});
Jimp.resize(40, 40);
// $ExpectType 0
Jimp.PNG_FILTER_NONE;

// $ExpectError
Jimp.test;

// $ExpectError
Jimp.func();

test('can clone properly', async () => {
  const baseImage = await Jimp.read('filename');
  const cloneBaseImage = baseImage.clone();

  // $ExpectType -1
  cloneBaseImage.PNG_FILTER_AUTO;

  test('can handle `this` returns on the core type properly', () => {
    // $ExpectType -1
    cloneBaseImage.diff(jimpInst, jimpInst).image.PNG_FILTER_AUTO
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
      .PNG_FILTER_AUTO;
  });

  test('can handle imageCallbacks `this` properly', () => {
    cloneBaseImage.rgba(false, (_, jimpCBIn) => {
      jimpCBIn.read('Test');
      jimpCBIn.displace(jimpInst, 2);
      jimpCBIn.resize(40, 40);
      // $ExpectType 0
      jimpCBIn.PNG_FILTER_NONE;

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
    cbJimpInst.read('Test');
    cbJimpInst.displace(jimpInst, 2);
    cbJimpInst.resize(40, 40);
    // $ExpectType 0
    cbJimpInst.PNG_FILTER_NONE;

    // $ExpectError
    cbJimpInst.test;

    // $ExpectError
    cbJimpInst.func();
  });
})

test('Can handle appendConstructorOption', () => {
  Jimp.appendConstructorOption(
    'Name of Option',
    args => args.hasSomeCustomThing,
    function(resolve, reject, args) {
      // $ExpectError
      this.bitmap = 3;
      Jimp.resize(2, 2);
      resolve();
    }
  );
});
