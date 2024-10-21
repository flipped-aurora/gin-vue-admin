import { Image, Omit } from '@jimp/core';
import { ThrowStatement } from 'typescript';

export function isNodePattern(cb: Function): true;
export function isNodePattern(cb: Omit<any, Function>): false;

export function throwError(error: string | Error, cb?: (err: Error) => void): ThrowStatement;

export function scan(image: Image, x: number, y: number, w: number, h: number, f: (image: Image, _x: number, _y: number, idx: number) => void): Image;
