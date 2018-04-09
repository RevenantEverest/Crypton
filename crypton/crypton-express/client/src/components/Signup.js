import React from 'react'
import Home from './Home'
import {Redirect} from 'react-router'
import services from '../services/services'

export default class Signup extends React.Component {
  constructor(props){
    super(props)
    this.state = {
      username: '',
      password: '',
      userCreated: false
    }
    this.handleSubmit = this.handleSubmit.bind(this)
    this.handleInputChange = this.handleInputChange.bind(this)
  }

  handleSubmit(e){
    e.preventDefault()
    console.log('my username --> ', this.state.username);
    console.log('my password --> ', this.state.password);
    services.createUser(this.state)
    .then(results => {
      this.setState({
        userCreated: true,
      })
    })
    .catch(err => console.log(err))
    //add axios call here
    //setState that userCreated is true
  }

  handleInputChange(e){
    let name = e.target.name
    let value = e.target.value
    this.setState({
      [name]:value
    })
  }
  render(){
    return (
      <div className='form'>
        <h1>Welcome To Crypton</h1>
        <form onSubmit={this.handleSubmit}>
          <input type='text' placeholder='Username' name='username' onChange={this.handleInputChange} /><br />
          <input type='password' placeholder='Password' name='password' onChange={this.handleInputChange} /><br />
          <input type='submit' value='Login' />
        </form>
        {this.state.userCreated ? <Redirect to='/home' /> :  '' }
        </div>
    )
  }
}
