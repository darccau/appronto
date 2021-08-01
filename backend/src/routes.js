const express = require('express')
const multer = require('multer')
const multerConfig = require('../config/multer.js')

const patientController = require('./controllers/patientController')

const router = express.Router()
const upload = multer(multerConfig)

router.get('/patients', patientController.index)

router.post('/patients', patientController.create)

router.put('/patients/:id', patientController.update)

router.delete('/patients/:id', patientController.delete)

router.post('/patients/reports', upload.array('reports'), patientController.upload)

module.exports = router
