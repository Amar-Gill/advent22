import { open } from "node:fs/promises";

const stacks = new Map([
  [1, ["N", "R", "G", "P"]],
  [2, ["J", "T", "B", "L", "F", "G", "D", "C"]],
  [3, ["M", "S", "V"]],
  [4, ["L", "S", "R", "C", "Z", "P"]],
  [5, ["P", "S", "L", "V", "C", "W", "D", "Q"]],
  [6, ["C", "T", "N", "W", "D", "M", "S"]],
  [7, ["H", "D", "G", "W", "P"]],
  [8, ["Z", "L", "P", "H", "S", "C", "M", "V"]],
  [9, ["R", "P", "F", "L", "W", "G", "Z"]],
]);

const file = await open("./instructions.in");

for await (const line of file.readLines()) {
  const tokens = line.split(" ");
  const numCratesToMove = parseInt(tokens[1]);
  const fromStackNum = parseInt(tokens[3]);
  const toStackNum = parseInt(tokens[5]);

  const fromStack = stacks.get(fromStackNum);
  const toStack = stacks.get(toStackNum);

  for (let index = 0; index < numCratesToMove; index++) {
    const pop = fromStack.pop();
    toStack.push(pop);
  }
}

console.log(stacks);
