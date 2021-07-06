const express = require('express')
const router = require('./routes')


const app = express()
app.use(express.json())
app.use(router)


app.listen(1337, () => {
  console.log('Server listen on port 1337')
})

