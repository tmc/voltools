import React from "react";

import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import IconButton from "@material-ui/core/IconButton";
import MenuIcon from "@material-ui/icons/Menu";
import CssBaseline from "@material-ui/core/CssBaseline";
import Container from '@material-ui/core/Container';

import { ErrorBoundary } from "react-error-boundary";
import { makeStyles } from "@material-ui/core/styles";

import ErrorFallback from "./components/ErrorFallback";
import VolunteerTools from "./components/VolunteerTools";

const errorHandler = (error: Error, info: { componentStack: string }) => {
  console.error("got error:", error, info);
};

export default function App(props) {
  const useStyles = makeStyles((theme) => ({
    offset: theme.mixins.toolbar,
  }));
  const classes = useStyles();

  return (
    <React.Fragment>
      <CssBaseline />
      <ErrorBoundary FallbackComponent={ErrorFallback} onError={errorHandler}>
        <AppBar position="static">
          <Toolbar variant="dense">
            <IconButton edge="start" color="inherit" aria-label="menu">
              <MenuIcon />
            </IconButton>
            <Typography variant="h6" color="inherit">
              Volunteer Tools
            </Typography>
          </Toolbar>
        </AppBar>

        <VolunteerTools />

      </ErrorBoundary>
    </React.Fragment>
  );
}
