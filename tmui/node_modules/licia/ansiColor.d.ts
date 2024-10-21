declare namespace ansiColor {
    type IFn = (str: string) => string;
}
declare const ansiColor: {
    black: ansiColor.IFn;
    red: ansiColor.IFn;
    green: ansiColor.IFn;
    yellow: ansiColor.IFn;
    blue: ansiColor.IFn;
    magenta: ansiColor.IFn;
    cyan: ansiColor.IFn;
    white: ansiColor.IFn;
    gray: ansiColor.IFn;
    grey: ansiColor.IFn;
    bgBlack: ansiColor.IFn;
    bgRed: ansiColor.IFn;
    bgGreen: ansiColor.IFn;
    bgYellow: ansiColor.IFn;
    bgBlue: ansiColor.IFn;
    bgMagenta: ansiColor.IFn;
    bgCyan: ansiColor.IFn;
    bgWhite: ansiColor.IFn;
    blackBright: ansiColor.IFn;
    redBright: ansiColor.IFn;
    greenBright: ansiColor.IFn;
    yellowBright: ansiColor.IFn;
    blueBright: ansiColor.IFn;
    magentaBright: ansiColor.IFn;
    cyanBright: ansiColor.IFn;
    whiteBright: ansiColor.IFn;
    bgBlackBright: ansiColor.IFn;
    bgRedBright: ansiColor.IFn;
    bgGreenBright: ansiColor.IFn;
    bgYellowBright: ansiColor.IFn;
    bgBlueBright: ansiColor.IFn;
    bgMagentaBright: ansiColor.IFn;
    bgCyanBright: ansiColor.IFn;
    bgWhiteBright: ansiColor.IFn;
};

export = ansiColor;
