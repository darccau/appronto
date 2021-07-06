const knex = require('knex')
const knexfile = require('knexfile')

const db = knex(knex.development)
module.exports = db

