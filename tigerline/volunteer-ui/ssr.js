import ReactDOMServer from "react-dom/server";
import { ServerStyleSheets } from "@material-ui/core/styles";

// SSR.
export function ssr(app) {
  const sheets = new ServerStyleSheets();
  const html = ReactDOMServer.renderToString(sheets.collect(app));
  const css = sheets.toString();
  return [html, css];
}
