import { open } from "node:fs/promises";

const file = await open("./input.in");

let runningSum = 0;
let max = 0;

for await (const line of file.readLines()) {
  const v = parseInt(line);

  if (!isNaN(v)) {
    runningSum += v;
  }

  if (line.length === 0) {
    max = Math.max(runningSum, max);
    runningSum = 0;
  }
}

console.log(max);
