const RGB_MAX: any = 255;
const HUE_MAX: any = 360;
const SV_MAX: any = 100;

export const cssGrad = (
  hueDiff: number,
  saturation: number = 75,
  lightness: number = 50,
  initialHue: number = Math.round(Math.random() * 360)
) => {
  const a = initialHue;
  const b = a + hueDiff > 360 ? a + hueDiff - 360 : a + hueDiff;
  return `linear-gradient(135deg, ${hslToHex([
    a,
    saturation,
    lightness,
  ])} 0%, ${hslToHex([b, saturation, lightness])} 100%)`;
};

export const hslToHex = (hsl: number[]) => {
  const rgb = hslToRgb(hsl);
  const hex = rgbToHex(rgb);
  return hex;
};

const hslToRgb = (hsl: number[]) => {
  let [h, s, l] = hsl;

  var r, g, b;

  h = normalizeAngle(h);
  h = h === HUE_MAX ? 1 : (h % HUE_MAX) / parseFloat(HUE_MAX);
  s = s === SV_MAX ? 1 : (s % SV_MAX) / parseFloat(SV_MAX);
  l = l === SV_MAX ? 1 : (l % SV_MAX) / parseFloat(SV_MAX);

  if (s === 0) {
    r = g = b = l; // achromatic
  } else {
    var q = l < 0.5 ? l * (1 + s) : l + s - l * s;
    var p = 2 * l - q;
    r = hue2Rgb(p, q, h + 1 / 3);
    g = hue2Rgb(p, q, h);
    b = hue2Rgb(p, q, h - 1 / 3);
  }

  return [
    Math.round(r * RGB_MAX),
    Math.round(g * RGB_MAX),
    Math.round(b * RGB_MAX),
  ];
};

export const rgbToHex = (rgb: Array<number>) => {
  const [r, g, b] = rgb;
  const toHex = (c: number) => `0${c.toString(16)}`.slice(-2);
  return `#${toHex(r)}${toHex(g)}${toHex(b)}`;
};

const normalizeAngle = (degrees: number) => ((degrees % 360) + 360) % 360;

const hue2Rgb = (p: number, q: number, t: number) => {
  if (t < 0) t += 1;
  if (t > 1) t -= 1;
  if (t < 1 / 6) return p + (q - p) * 6 * t;
  if (t < 1 / 2) return q;
  if (t < 2 / 3) return p + (q - p) * (2 / 3 - t) * 6;
  return p;
};
