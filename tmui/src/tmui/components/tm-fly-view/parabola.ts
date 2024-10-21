export function findXOnQuadraticBezierCurve(startX: number, startY: number, control1X: number, control1Y: number, control2X: number, control2Y: number, endX: number, endY: number, targetY: number): number {
  let tMin = 0;
  let tMax = 1;
  const epsilon = 1e-6;

  while (tMax - tMin > epsilon) {
    const tMid = (tMin + tMax) / 2;
    const yMid = Math.pow(1 - tMid, 3) * startY + 3 * Math.pow(1 - tMid, 2) * tMid * control1Y + 3 * (1 - tMid) * Math.pow(tMid, 2) * control2Y + Math.pow(tMid, 3) * endY;

    if (yMid < targetY) {
      tMin = tMid;
    } else {
      tMax = tMid;
    }
  }

  return Math.pow(1 - tMax, 3) * startX + 3 * Math.pow(1 - tMax, 2) * tMax * control1X + 3 * (1 - tMax) * Math.pow(tMax, 2) * control2X + Math.pow(tMax, 3) * endX;
}
