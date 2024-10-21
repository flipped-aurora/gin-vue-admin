import {Jimp} from './jimp';

export function addConstants(
  constants: [string, string | number],
  jimpInstance?: Jimp
): void;
export function addJimpMethods(
  methods: [string, Function],
  jimpInstance?: Jimp
): void;
export function jimpEvMethod(
  methodName: string,
  evName: string,
  method: Function
): void;
export function jimpEvChange(methodName: string, method: Function): void;
export function addType(mime: string, extensions: string[]): void;
