import type { Nation } from "../bindings/astroproject/astro/structs/models";

export function randomColor(brightness: number): string {
  function randomChannel(brightness: number): string {
    var r = 255-brightness;
    var n = 0|((Math.random() * r) + brightness);
    var s = n.toString(16);
    return (s.length==1) ? '0'+s : s;
  }
  return '#' + randomChannel(brightness) + randomChannel(brightness) + randomChannel(brightness);
}

export const rightBrighness = 240;

export function getNationFlag(nation: Nation | null): string {

  if (!nation) {
    return "https://flagsapi.com/FR/flat/64.png";
  }

  return `https://flagsapi.com/${nation.nation_code}/flat/64.png`
}

export function getNationFlatAlt(nation: Nation | null): string {
  if (!nation) {
    return "FR";
  }

  return nation.nation_code;
}