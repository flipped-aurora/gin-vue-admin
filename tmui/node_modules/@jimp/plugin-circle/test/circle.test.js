import { Jimp, mkJGD, getTestDir } from '@jimp/test-utils';
import configure from '@jimp/custom';

import circle from '../src';

const jimp = configure({ plugins: [circle] }, Jimp);

describe('Circle', () => {
  it('makes a circle based on image height and width', async () => {
    const expectedImg = await Jimp.read(
      getTestDir(__dirname) + '/images/circled.png'
    );
    const imgSrc = await jimp.read(
      mkJGD(
        '▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦',
        '▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦',
        '▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦',
        '▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦',
        '▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦',
        '▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦',
        '▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦',
        '▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦',
        '▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦',
        '▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦'
      )
    );

    imgSrc.circle().bitmap.data.should.be.deepEqual(expectedImg.bitmap.data);
  });

  it('makes a circle using provided radius', async () => {
    const expectedImg = await Jimp.read(
      getTestDir(__dirname) + '/images/radius-3-circle.png'
    );
    const imgSrc = await jimp.read(
      mkJGD(
        '▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦',
        '▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦',
        '▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦',
        '▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦',
        '▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦',
        '▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦',
        '▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦',
        '▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦',
        '▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦',
        '▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦'
      )
    );

    imgSrc
      .circle({ radius: 3 })
      .bitmap.data.should.be.deepEqual(expectedImg.bitmap.data);
  });

  it('should ', async () => {
    const expectedImg = await Jimp.read(
      getTestDir(__dirname) + '/images/x-y-circle.png'
    );
    const imgSrc = await jimp.read(
      mkJGD(
        '▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦',
        '▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦',
        '▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦',
        '▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦',
        '▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦',
        '▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦',
        '▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦',
        '▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦',
        '▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦',
        '▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦▦'
      )
    );

    imgSrc
      .circle({ radius: 5, x: 5, y: 5 })
      .bitmap.data.should.be.deepEqual(expectedImg.bitmap.data);
  });
});
