const knex = require('../database/connection')

module.exports = {

  async createUser(req, res, next) {
    try {
      const { first_name } = req.body
      const { last_name } = req.body
      const { email } = req.body
      const { password } = req.body

      await knex('users')
      .insert({ first_name, last_name, email, password})

      res.send('ok')
    } 
    catch (error) {
      console.error(error)
    }
  },

  async credentials(req, res, next) {
    const { email } = req.body
    const { password } = req.body

    console.log({email, password})

    const credentials = await knex
      .select('first_name', 'last_name', 'email')
      .from('users')
      .where('email', '=', `${email}`)
      .andWhere('password', '=', `${password}`)

    res.send(credentials)
  },

}
