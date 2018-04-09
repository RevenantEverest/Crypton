import React from 'react'
import Signup from './Signup'
import {Link} from 'react-router-dom'

export default class Home extends React.Component {
  constructor(props){
    super(props)
    this.state = {
      isLoggedIn: false
    }
  }
  render(){
    return (
      <div>
        <h1 className='Welcome'>Welcome to Crypton</h1>
        <Link to='/signup'>Sign up!</Link>
      </div>
    )
  }
}
