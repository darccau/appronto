const express = require('express')

const app = express()
app.use(express.json())

app.get('/hello', (req, res) => {
  res.send('{id: fuck}')
})

app.listen(1337, () => console.log('Server listen on port 1337'))

