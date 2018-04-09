import axios from 'axios'

const services = {}

services.createUser = (userInfo) => {
  console.log('this is user info --> ', userInfo);
    return axios({
      method: 'POST',
      url: 'http://localhost:3000/signup',
      data: {
        username: userInfo.username,
        password: userInfo.password
      }
    })
  }

export default services
