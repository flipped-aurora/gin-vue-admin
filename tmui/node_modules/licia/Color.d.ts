declare namespace Color {
    interface IColor {
        val: number[];
        model: string;
    }
}
declare class Color {
    constructor(color: string | Color.IColor);
    toRgb(): string;
    toHex(): string;
    toHsl(): string;
    static parse(colorStr: string): Color.IColor;
}

export = Color;
