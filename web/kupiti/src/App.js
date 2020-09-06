import React from 'react';
import { ThemeProvider } from '@material-ui/core/styles';

//Local
import './App.css';
import Main from './containers/Main'
import Theme from './theme/Theme'

function App() {
  return (
    <ThemeProvider theme={Theme}>
      <Main/>
    </ThemeProvider>
  );
}

export default App;
