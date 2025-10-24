
import van from "./van-1.6.0.min.js"
const {a, div, li, p, ul, span, strong} = van.tags
const total = van.state(1);
const the_link = a("Hello");
the_link.setAttribute("href", "/a\">ddf</a>")

const Hello = () => div(
  p("👋Hello"),
  ul(
    li("🗺️World"),
    li("🗺️  ...again"),
  ),
  span(strong("This is a span."), "and this is normal.")
)

van.add(document.body, Hello())
