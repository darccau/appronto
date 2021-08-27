const express = require('express')
const multer = require('multer')

const multerConfig = require('./config/multer.js')

const appointment = require('./controllers/appointment')
const user = require('./controllers/user')

const router = express.Router()
const upload = multer(multerConfig)

router.post('/user/create', user.createUser)

router.post('/appointment/create', appointment.createAppointment)

router.post('/appointment/upload/reports', upload.array('reports'),
  appointment.uploadReports)

router.post('/user/credentials', user.credentials) 

router.post('/appointment/appointments', appointment.appointmentsByUsers)

router.post('/appointments/reports', appointment.reports)

module.exports = router
