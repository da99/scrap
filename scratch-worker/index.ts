
import 'https://deno.land/x/worker_types/cloudflare-worker-types.ts';

declare const __KV: KVNamespace;
declare const FILE_SECRET_KEY: string;

addEventListener('fetch', event => {
  event.respondWith(handleRequest(event.request))
});

/**
 * Respond with hello worker text
 * @param {Request} request
 */

let VERSION = 1;
function get_html() {
return `
<html>
  <head>
    <title>Scratchpad</title>
  </head>
  <body>
    <p>Hello</p>
  </body>
  <script type="text/javascript">
    const oReq = new XMLHttpRequest();
    oReq.open("POST", "http://127.0.0.1:8787", true);
    oReq.addEventListener("load", function (oEvent) {
      console.log("${VERSION++} UPLOADED: " + this.responseText);
      console.log(oEvent);
      console.log(this);
    });


    const blob = new Blob(['abc123'], {type: 'text/plain'});

    console.log(\`Uploading: \${Math.floor(Date.now() / 1000)}\`);
    oReq.setRequestHeader("X-helper", "2");
    oReq.send(blob);
  </script>
</html>
`;
} // function


async function handleRequest(request) {
  const url = new URL(request.url);
  const target = `${request.method} ${url.pathname}`;
  switch(target) {
    case "GET /da99": {
      return new Response(get_html(), {
        headers: { 'content-type': 'text/html' },
      })
    } // case

    case "GET /a": {
      const {value, metadata}  = await KV.getWithMetadata("a.jpg", {type: "stream"});
      if (value && metadata === null) {
        console.log("Writing meta data...");
        const v = await KV.put("a.jpg", value, {metadata: {"content-type": "image/jpeg"}});
        console.log(v);
      } else {
        console.log(metadata)
      }
      return new Response("done", {
        status: 200,
        headers: { 'content-type': 'text/plain' },
      });
    }

    case "GET /a.jpg": {
      const {value, metadata}  = await KV.getWithMetadata("a.jpg", {type: "stream"});
      if (value === null) {
        return new Response("file Not found.", {
          status: 404,
          headers: { 'content-type': 'text/plain' },
        });
      } else {
        console.log(metadata);
        return new Response(value, {headers: {...metadata, "X-done": "true"}});
      } // if/else
    } // case

    case "POST /da99": {
      console.log("Post headers: ");
      for (const k of request.headers.keys()) {
        console.log(`${k}: ${request.headers.get(k)}`);
      }
      return new Response("NOT DonE", {
        headers: { 'content-type': 'text/plain' },
      })
    }
    default:
      return new Response("Not Ready.", {
        headers: { 'content-type': 'text/plain' },
      })
  } // switch
} // function
