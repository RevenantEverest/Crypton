const pgp = require('pg-promise')()
const dbConfig = require('../db/dbConfig')
const db = pgp(dbConfig)

module.exports = {
  createUser(user){
    return db.one(`INSERT INTO users (username, password) VALUES($[username], $[password]) RETURNING *`,user);
  }
}
