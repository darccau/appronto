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

}


//   async index(req, res) {
//     const results = await knex('report')
//     return res.json(results)
//   },
// 
//   async consultation(req, res, next) {
//     try {
//       const { id } = req.params
// 
//       const results = await knex('report')
//         .where({ id })
//       console.log(results)
//       res.sendStatus(200)
//     } 
//     catch (error) {
//       console.error(error)
//     }
//   },
  //
//   async update(req, res, next) {
//     try {
//       const { id } = req.params
//       const { first_name } = req.body
//       const { last_name } = req.body
//       const { report } = req.body
// 
//       await knex('report')
//         .update({ first_name, last_name, report })
//         .where({ id })
//       res.send('ok')
//     }
//     catch (error) {
//       console.error(error)
//     }
//   },
// 
//   async delete(req, res, next) {
//     try {
//       const { id } = req.params
// 
//       await knex('report')
//       .where({ id })
//       .del()
// 
//       res.sendStatus(200)
//     }
//     catch(error) {
//       console.error(error)
//     }
//   },
 
