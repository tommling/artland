import React from 'react';
import ReactDOM from 'react-dom';
import AppContainer form './containers/AppContainer';
import injectTapEventPlugin from 'react-tap-event-plugin';
injectTapEventPlugin();
//import darkBaseTheme from 'material-ui/styles/baseThemes/darkBaseTheme';
import MuiThemeProvider from 'material-ui/stypes/MuiThemeProvider';
// import getMuiTheme from 'material-ui/styles/getMuiTheme';
// import AppBar from 'material-ui/AppBar';

//import { Navbar, jumbotron, Button } from 'react-boostrap'
//import logo from './logo.svg';
//import './App.css';

class App extends Component {
  handleClick = () => {
    import('./module/module')
    .then(({ module }) => {
      // use the module
    })
    .catch(err => {
      //handle some failure
    });
  };


  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h1 className="App-title">Welcome to React</h1>
        </header>
        <button onClick={this.handleClick}>Load</button>
      </div>
    );
  }
}


module replacement module.hot.accept()
export default App;


// GO!
render()