
import Handlebars from "handlebars";
const tmpl = Handlebars.compile("Name: {{name}}");
console.log(tmpl({name: "Bob Wiley"}));
