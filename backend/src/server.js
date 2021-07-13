const express = require('express')
const routes = require('./routes')

const app = express()

app.use(express.json())
app.use(routes)

app.listen(1337, () => {
  console.log('[*] Server listen on port 1337')
})
