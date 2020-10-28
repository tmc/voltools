import React, { useState, useEffect, useRef } from 'react';

import Button from "@material-ui/core/Button";
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import Typography from "@material-ui/core/Typography";
import TextField from "@material-ui/core/TextField";
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import CardActions from '@material-ui/core/CardActions';

import { makeStyles } from "@material-ui/core/styles";

import { useDebounce } from '../hooks/debounce.js';

export default function ZipInfo(props) {
  const useStyles = makeStyles((theme) => ({
    offset: theme.mixins.toolbar,
  }));
  const classes = useStyles();

  const [searchTerm, setSearchTerm] = useState('');
  const [results, setResults] = useState([]);

  return (
    <div>
      <Card>
        <CardContent>
          <Typography variant="body2" component="p">
            Enter your zipcode here
          </Typography>
          <TextField label="Zip Code" />
        </CardContent>
      </Card>
    </div>
  );
}
