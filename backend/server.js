const express = require('express')
const router = require('./routes')


const app = express()

app.use(express.json())
app.use(router)

app.use((error, req, res, next) => {
  res.status(error.status || 500)
  res.json({ error: error.message })
})

app.listen(1337, () => {
  console.log('Server is running')
})

