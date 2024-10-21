declare const easing: {
    linear(percent: number): number;
    inQuad(percent: number): number;
    outQuad(percent: number): number;
    inOutQuad(percent: number): number;
    outInQuad(percent: number): number;
    inCubic(percent: number): number;
    outCubic(percent: number): number;
    inQuart(percent: number): number;
    outQuart(percent: number): number;
    inQuint(percent: number): number;
    outQuint(percent: number): number;
    inExpo(percent: number): number;
    outExpo(percent: number): number;
    inSine(percent: number): number;
    outSine(percent: number): number;
    inCirc(percent: number): number;
    outCirc(percent: number): number;
    inElastic(percent: number, elasticity?: number): number;
    outElastic(percent: number, elasticity?: number): number;
    inBack(percent: number): number;
    outBack(percent: number): number;
    inOutBack(percent: number): number;
    outInBack(percent: number): number;
    inBounce(percent: number): number;
    outBounce(percent: number): number;
};

export = easing;
