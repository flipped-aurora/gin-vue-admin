import { Jimp, mkJGD } from '@jimp/test-utils';
import configure from '@jimp/custom';

import fisheye from '../src';

const jimp = configure({ plugins: [fisheye] }, Jimp);

describe('Fisheye', () => {
  it('should create fisheye lens to image', async () => {
    const imgNormal = await jimp.read(
      mkJGD(
        '0000000000',
        '0001221000',
        '0022222200',
        '0122112210',
        '0221001220',
        '0221001220',
        '0122112210',
        '0022222200',
        '0001221000',
        '0000000000'
      )
    );
    const imgBulged = await jimp.read(
      mkJGD(
        '0001221000',
        '0221112220',
        '0220000121',
        '1100000112',
        '2100000012',
        '2100000012',
        '1200000012',
        '0211000222',
        '0221111220',
        '0012222200'
      )
    );

    imgNormal
      .fisheye()
      .getJGDSync()
      .should.be.sameJGD(imgBulged.getJGDSync());
  });

  it('should create fisheye lens to image with radius', async () => {
    const imgNormal = await jimp.read(
      mkJGD(
        '0000000000',
        '0000000000',
        '0000000000',
        '0000000000',
        '0001111000',
        '0001111000',
        '0000000000',
        '0000000000',
        '0000000000',
        '0000000000'
      )
    );
    const imgBulged = await jimp.read(
      mkJGD(
        '■■■■■■■■■■',
        '■■■■■■■■■■',
        '■■■■■■■■■■',
        '■■■11111■■',
        '■■111111■■',
        '■■111111■■',
        '■■■■111■■■',
        '■■■■■■■■■■',
        '■■■■■■■■■■',
        '■■■■■■■■■■'
      )
    );

    imgNormal
      .fisheye({ r: 1.8 })
      .getJGDSync()
      .should.be.sameJGD(imgBulged.getJGDSync());
  });
});
