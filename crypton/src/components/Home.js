import React from 'react'
import Signup from './Signup'

export default class Home extends React.Component {
  render(){
    return (
      <div>
        <h1 className='Welcome'>Welcome to Crypton</h1>
        <Signup />
      </div>
    )
  }
}
