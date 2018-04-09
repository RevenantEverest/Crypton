const usersDB = require('../models/usersDB')

module.exports = {

create(req, res){
  console.log('this is req.body --> ',req.body);
    usersDB.createUser(req.body)
    .then(results => {
      res.json({
        message: 'ok',
        data: results
      })
    })
    .catch(err => {
      console.log(err)
      res.status(400).json({message: '400', err})
    })
  }
}
