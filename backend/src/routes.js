const express = require('express')
const multer = require('multer')

const multerConfig = require('./config/multer.js')
const patientController = require('./controllers/patientController')

const router = express.Router()
const upload = multer(multerConfig)

router.get('/patients', patientController.index)

router.get('/patient/:id', patientController.consultation)

router.post('/patient', patientController.create)

router.put('/patient/:id', patientController.update)

router.delete('/patient/:id', patientController.delete)

router.post('/patient/reports', upload.array('reports'), patientController.upload)

module.exports = router
