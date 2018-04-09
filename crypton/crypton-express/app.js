const express = require('express')
const app = express()
const logger = require('morgan')
const PORT = process.env.PORT || 3001
const NODE_ENV = process.env.NODE_ENV || 'dev'
const bodyParser = require('body-parser')
const methodOveride = require('method-override')
const usersRouter = require('./routes/usersRouter')

app.use(logger('dev'))
app.use(express.static('public'))
app.use(bodyParser.json())


app.get('/', (req, res) => {
    res.json('we good')
})

app.use('/signup', usersRouter);


app.listen(PORT, () => console.log(`listening on port ${PORT}`))
