exports.up = function(knex) {
  return knex.raw(`

    CREATE TABLE users (
      id INT GENERATED ALWAYS AS IDENTITY,
      first_name VARCHAR(255) NOT NULL,
      last_name VARCHAR(255) NOT NULL,
      email varchar(255) UNIQUE NOT NULL,
      password VARCHAR(255) UNIQUE NOT NULL,
      PRIMARY KEY(id)
    );

    CREATE TABLE medical_appointments (
      id INT GENERATED ALWAYS AS IDENTITY,
      id_user INT,
      doctor_name VARCHAR(255) NOT NULL,
      appointment_type VARCHAR(255) NOT NULL,
      note TEXT NOT NULL,
      date TIMESTAMP NULL,
      PRIMARY KEY(id),
      FOREIGN KEY(id_user)
      REFERENCES users(id)
    );

    CREATE TABLE reports (
      id INT NOT NULL,
      id_appointments INT,
      path VARCHAR(255) NOT NULL,
      PRIMARY KEY(id),
      FOREIGN KEY(id_appointments)
      REFERENCES medical_appointments(id)
    );

  `)
};

exports.down = function(knex) {
  return knex.raw(`
    drop table reports
    drop table medical_appointments
    drop table user
    `)
};
