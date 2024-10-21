import random = require('./random');
import Color = require('./Color');

declare function randomColor(): string;
declare function randomColor(options: {
    count?: number;
    hue?: number;
    lightness?: number;
    format?: string;
    seed?: number;
}): string | string[];

export = randomColor;
