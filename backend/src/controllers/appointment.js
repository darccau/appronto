const knex = require("../database/connection");

module.exports = {
  async createAppointment(req, res, next) {
    try {
      const { doctor_name } = req.body;
      const { appointment_type } = req.body;
      const { id_user } = req.body;
      const { note } = req.body;

      await knex("medical_appointments").insert({
        doctor_name,
        id_user,
        appointment_type,
        note,
      });

      res.send().status("200");
    } catch (error) {
      console.error(error);
    }
  },

  // TODO improve add user_id correctly with file
  async uploadReports(req, res, next) {
    try {
      const files = req.files;
      const paths = files.map((file) => file.path);
      const id_appointments = 1;

      for (tes in paths) {
        let path = paths[tes];
        await knex("reports").insert({ path, id_appointments });
      }

      res.send().status("200");
    } catch (error) {
      console.error(error);
    }
  },

  async appointmentsByUsers(req, res, next) {
    try {
      const { id_user } = req.body;

      console.log(id_user);

      const appointments = await knex
        .select("*")
        .from("medical_appointments")
        .where("id_user", "=", `${id_user}`);

      res.send(appointments);
    } catch (error) {
      console.error(error);
    }
  },

  async reports(req, res, next) {
    try {
      const { id_user } = req.body;
      console.log(id_user);

      const filePaths = await knex
        .select("path")
        .from("reports")
        .where("id_appointments", "=", `${id_user}`);

      res.send(filePaths);
    } catch (error) {
      console.error(error);
    }
  },
};
