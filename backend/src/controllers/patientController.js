const knex = require('../database/connection')

module.exports = {
  async index(req, res) {
    const results = await knex('report')
    return res.json(results)
  },

  async create(req, res, next) {
    try {
      const { first_name } = req.body
      const { last_name } = req.body
      const { report } = req.body

      await knex('report')
        .insert({ first_name, last_name, report })
    } 
    catch (error) {
      console.error(error)
    }
  },

  async update(req, res, next) {
    try {
      const { first_name } = req.body
      const { last_name } = req.body
      const { report } = req.body

      await knex('report')
      .update({ first_name, last_name, report })
      .where(id)
    }
    catch (error) {
      console.error(error)
    }
  },

  async update(req, res, next) {
    try {
      const { id } = req.params
      const { first_name } = req.body
      const { last_name } = req.body
      const { report } = req.body

      await knex('report')
        .update({ first_name, last_name, report })
        .where({ id })

    } catch (error) {
      console.error(error)
    }
  },

  async delete(req, res, next) {
    try {
      const { id } = req.params

      await knex('report')
      .where({ id })
      .del()

      res.send(`${id} was deleted`)
    }
    catch(error) {
      console.error(error)
    }
  },

  async upload(req, res, next) {
    res.send('File was uploaded')
  }
}

