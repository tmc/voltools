import React from "react";
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import Typography from "@material-ui/core/Typography";

import { makeStyles } from "@material-ui/core/styles";

import ZipInfo from "./ZipInfo";

export default function VolunteerTools(props) {
  const useStyles = makeStyles((theme) => ({
    offset: theme.mixins.toolbar,
    root: {
      padding: theme.spacing(2),
      flexGrow: 1,
    },
    paper: {
      padding: theme.spacing(2),
      textAlign: 'center',
      color: theme.palette.text.secondary,
    },
  }));
  const classes = useStyles();
  return (
    <div className={classes.root}>
      <Grid container spacing={3}>
        <Grid item xs={12} sm={6}>
          <Paper className={classes.paper}>
            <ZipInfo />
          </Paper>
        </Grid>
        <Grid item xs={12} sm={6}>
          <Paper className={classes.paper}></Paper>
        </Grid>
        <Grid item xs={12}>
          <Paper className={classes.paper}>
          </Paper>
        </Grid>
      </Grid>
    </div>
  );
}
