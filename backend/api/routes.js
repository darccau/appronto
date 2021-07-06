const express = require('express')
const router = express.Router()
// const queries = require('./pacient')

router.get('/', (req, res) => {
  console.log('check')
  res.send('{fuck: that shit}')
})

// router.get('/create', () => {
//   console.log('create')
// })

// router.get('/delete', () => {
//   console.log('delete')
// })

module.exports = router
