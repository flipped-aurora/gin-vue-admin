exports = {
    black: genColor([0, 0]),
    red: genColor([31, 39]),
    green: genColor([32, 39]),
    yellow: genColor([33, 39]),
    blue: genColor([34, 39]),
    magenta: genColor([35, 39]),
    cyan: genColor([36, 39]),
    white: genColor([37, 39]),
    gray: genColor([90, 39]),
    grey: genColor([90, 39]),
    bgBlack: genColor([40, 49]),
    bgRed: genColor([41, 49]),
    bgGreen: genColor([42, 49]),
    bgYellow: genColor([43, 49]),
    bgBlue: genColor([44, 49]),
    bgMagenta: genColor([45, 49]),
    bgCyan: genColor([46, 49]),
    bgWhite: genColor([47, 49]),
    blackBright: genColor([90, 39]),
    redBright: genColor([91, 39]),
    greenBright: genColor([92, 39]),
    yellowBright: genColor([93, 39]),
    blueBright: genColor([94, 39]),
    magentaBright: genColor([95, 39]),
    cyanBright: genColor([96, 39]),
    whiteBright: genColor([97, 39]),
    bgBlackBright: genColor([100, 49]),
    bgRedBright: genColor([101, 49]),
    bgGreenBright: genColor([102, 49]),
    bgYellowBright: genColor([103, 49]),
    bgBlueBright: genColor([104, 49]),
    bgMagentaBright: genColor([105, 49]),
    bgCyanBright: genColor([106, 49]),
    bgWhiteBright: genColor([107, 49])
};

function genColor(codes) {
    const open = `\u001b[${codes[0]}m`;
    const close = `\u001b[${codes[1]}m`;

    return input => open + input + close;
}

module.exports = exports;
