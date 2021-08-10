const multer = require('multer')
const path = require('path')

const uploadsFolder = path.resolve(__dirname, '..', '..', 'tmp', 'uploads')

module.exports = {
  dest: uploadsFolder,
  storage: multer.diskStorage({
    destination: uploadsFolder,
    filename: (req, file, cb) => {
      cb(null, `${Date.now()}-${file.originalname}`)
    }
  }),

  limits: {
    fileSize: 4 * 1024 * 1024
  },

  fileFilter: (req, file, cb) => {
    const mimeTypes = [
      'image/jpeg',
      'image/png'
    ]

    if (!mimeTypes.includes(file.mimetype)) {
      return cb(null, false)
    }

    cb(null, true)
  }
} 

