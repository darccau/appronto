const express = require('express')
const router = express.Router()

const patientController = require('./controllers/patientController')

router.get('/patients', patientController.index)

router.post('/patients', patientController.create)

router.put('/patients/:id', patientController.update)

router.delete('/patients/:id', patientController.delete)

module.exports = router
