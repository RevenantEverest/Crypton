import React, { Component } from 'react';
import logo from './logo.svg';
import Home from './components/Home'
import './App.css';
import {BrowserRouter as Router, Link, Route} from 'react-router-dom'
import Signup from './components/Signup'

class App extends Component {
  render() {
    return (
      <Router>
        <div>
          <Route path='/home' component={Home} />
          <Route path='/signup' component={Signup} />
        </div>
      </Router>
    );
  }
}

export default App;
