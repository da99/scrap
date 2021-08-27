
interface Shoutcast_Station {
  "Current Song"?: string,
  "Stream Title"?: string,
  "Play URL"?: string
} // interface


class ShoutCast {
  static TITLE_PATTERN  = /Stream Title:\ +(.+)\n/;
  static CURRENT_TITLE  = /Current Song:\ +(.+?)\n\n/m;
  static TD_MATCH       = /<td(?:[^>]*)>(.+?)<\/td>/g;
  static TAG_MATCH      = /(<([^>]+)>)/gi;
  static TRAILING_COLON = /\:$/;

  static parse(raw : string) {
    const match = raw.matchAll(ShoutCast.TD_MATCH);
    const info : Shoutcast_Station = {};
    let last_key : string = "";
    for(let m of match) {
      const s = m[1].replace(ShoutCast.TAG_MATCH, "").trim();
      switch (last_key) {
        case "Current Song":
          case "Stream Title":
          if (!info[last_key]) {
          info[last_key] = s;
        }
        last_key = "";
        break;
        default:
          last_key = s.replace(ShoutCast.TRAILING_COLON, "");
      }
    } // for
    return info;
  } // static

} // class

const options = {
  method: "GET",
  headers: {
    "User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36",
    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
    "Accept-Language": "en-US,en;q=0.9,en-GB;q=0.8"
  }
};

class Main {
  static async get() {
    try {
      const textPromise = fetch("http://155.138.139.156:8099/", options);
      const resp = await textPromise;
      const body = await resp.text();
      const result = ShoutCast.parse(body);
      console.log(result);
    } catch (error) {
      console.log(error);
    }
  } // func
} // class


addEventListener("fetch", (event) => {

  const options = {
    status: 200,
    headers: { "content-type": "text/plain" }
  };
  const response = new Response("Hello, Dr. Deno!", options);
  event.respondWith(response);
});



