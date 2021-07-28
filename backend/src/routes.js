const express = require('express')
// const multer = require('multer')
// const multerConfig = require('./config/multer')
const router = express.Router()

const patientController = require('./controllers/patientController')

router.get('/patients', patientController.index)

router.post('/patients', patientController.create)

router.put('/patients/:id', patientController.update)

router.delete('/patients/:id', patientController.delete)

// router.post('/reports', multer(multerConfig).single('file'), () => {
//   console.log({"fuck": "that shit"})
// })

module.exports = router
