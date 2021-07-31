const multer = require('multer')
const path = require('path')

export default {
  storage: multer.diskStorage({
    destination: path.join(__dirname, '..', '..', 'upload'),
    filename: (req, res, cb) => {
      cb(null, `${Date.now()}-${file.originalname}`)
    },
  }),
  limits: {
    filesize: 4 * 1024 * 1024
  },
  fileFilter: (req, res, cb) => {
    const mimeTypes = [
      'images/jpg',
      'images/png'
    ]

    if (!mimeTypes.includes(file.mimetype)) {
      return cb(null, false)
    }
    cb(null, true)
  },
} 

