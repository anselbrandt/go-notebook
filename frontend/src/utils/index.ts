export const swap = (array: any[], moveIndex: number, toIndex: number) => {
  let itemRemovedArray = [
    ...array.slice(0, moveIndex),
    ...array.slice(moveIndex + 1, array.length),
  ];
  return [
    ...itemRemovedArray.slice(0, toIndex),
    array[moveIndex],
    ...itemRemovedArray.slice(toIndex, itemRemovedArray.length),
  ];
};

export const clamp = (number: number, lower: number, upper: number) => {
  number = number <= upper ? number : upper;
  number = number >= lower ? number : lower;
  return number;
};
