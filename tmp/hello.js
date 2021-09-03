class ShoutCast {
    static URLS = `
    http://155.138.139.156:8101/
    http://155.138.139.156:9101/
    http://155.138.139.156:8099/
    http://155.138.139.156:9999/
    http://155.138.139.156:8199/
    http://213.239.204.252:8000/
`.split(/\s+/).filter(Boolean);
    static TITLE_PATTERN = /Stream Title:\ +(.+)\n/;
    static CURRENT_TITLE = /Current Song:\ +(.+?)\n\n/m;
    static TD_MATCH = /<td(?:[^>]*)>(.+?)<\/td>/g;
    static TAG_MATCH = /(<([^>]+)>)/gi;
    static TRAILING_COLON = /\:$/;
    static parse(origin_url, raw) {
        const match = raw.matchAll(ShoutCast.TD_MATCH);
        const info = {
        };
        let last_key = "";
        for (let m of match){
            const s = m[1].replace(ShoutCast.TAG_MATCH, "").trim();
            switch(last_key){
                case "Stream Title":
                case "Stream URL":
                case "Current Song":
                    if (!info[last_key]) {
                        info[last_key] = s;
                    }
                    last_key = "";
                    break;
                default:
                    if (s.indexOf(":") > 3) {
                        last_key = s.replace(ShoutCast.TRAILING_COLON, "");
                    } else {
                        last_key = "";
                    }
            }
        }
        info["Play URL"] = origin_url;
        return info;
    }
}
const options = {
    method: "GET",
    headers: {
        "User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36",
        "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
        "Accept-Language": "en-US,en;q=0.9,en-GB;q=0.8"
    }
};
const resp_o = {
    status: 200,
    headers: {
        "content-type": "application/json; charset=utf-8",
        "Access-Control-Allow-Origin": "*"
    }
};
function get_all(event) {
    const fetches = ShoutCast.URLS.map((url)=>{
        return fetch(url, options).then((resp)=>resp.text()
        ).then((txt)=>ShoutCast.parse(url, txt)
        );
    });
    Promise.all(fetches).then((bodies)=>{
        const response = new Response(JSON.stringify(bodies), resp_o);
        event.respondWith(response);
    });
}
addEventListener("fetch", get_all);
