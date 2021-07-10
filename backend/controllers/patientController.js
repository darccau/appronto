const knex = require('../database')

module.exports = {
  async index(req, res) {
    const results = await knex('report')
    return res.json(results)
  },

  async create(req, res, next) {
    try {
      console.log(req.body)
      const { first_name } = req.body
      const { last_name } = req.body
      const { report } = req.body

      await knex('report')
        .insert({ first_name, last_name, report })
    } catch (error) {
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
      .whe
    }
    catch (error) {
      console.error(error)
    }
  },

  async delete(req, res, next) {
    try {
      const { id } = req.params
      console.log(id)
      await knex('report')
      .where({ id })
      .del()
      // console.log('all right')
    }
    catch(error) {
      console.error(error)
    }
  }

}


