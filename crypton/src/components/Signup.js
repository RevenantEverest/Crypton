import React from 'react'

export default class Signup extends React.Component {
  constructor(props){
    super(props)
    this.state = {
      username: '',
      password: ''
    }
    this.handleSubmit = this.handleSubmit.bind(this)
    this.handleInputChange = this.handleInputChange.bind(this)
  }

  handleSubmit(e){
    e.preventDefault()
    console.log('my username --> ', this.state.username);
    console.log('my password --> ', this.state.password);

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
        <form onSubmit={this.handleSubmit}>
          <input type='text' placeholder='Username' name='username' onChange={this.handleInputChange} /><br />
          <input type='password' placeholder='Password' name='password' onChange={this.handleInputChange} /><br />
          <input type='submit' value='Login' />
        </form>
        </div>
    )
  }
}
