

import { constant_time_compare } from "./Utils.ts";
import {encode,decode} from "https://deno.land/std/encoding/base64.ts";

const filePath='./Public/a.jpg';
const origSize=(await Deno.stat(filePath)).size;
const f=await Deno.open(filePath);
const encoded=encode(await Deno.readAll(f));
f.close();
const decoded=decode(encoded);

const json = JSON.stringify([
  {
    "key": "a.jpg",
    "value": encoded,
    "base64": true,
    "metadata": {
      "content-type": "image/jpeg"
    }
  }
]);

console.log(json);
