import React from "react";
import ReactDOM from "react-dom";

import { ssr } from "./ssr.js";

import { ThemeProvider } from "@material-ui/core/styles";
import theme from "./theme.js";

import App from "./App";

var app = (
  <ThemeProvider theme={theme}>
    <App version="1.1" />
  </ThemeProvider>
);

// Enable HMR.
if (module.hot) {
  module.hot.accept();
}

// Expose SSR entrypoint.
global.renderServerSide = () => ssr(app);

// If we're running in a browser environment.
if (typeof window !== "undefined") {
  if (
    typeof __prerendered !== "undefined" &&
    __prerendered &&
    __prerendered !== "{{ .IsPrerendered }}"
  ) {
    console.log("pr check", __prerendered);
    ReactDOM.hydrate(app, document.getElementById("root"));
  } else {
    ReactDOM.render(app, document.getElementById("root"));
  }
}
