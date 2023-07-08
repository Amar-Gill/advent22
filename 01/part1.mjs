import { open } from "node:fs/promises";

const file = await open("./input.in");

for await (const line of file.readLines()) {
  console.log(line);
}
