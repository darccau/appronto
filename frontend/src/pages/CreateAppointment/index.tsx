import React, { useState } from "react";

import axios from "axios";

import PrimaryButton from "../../components/PrimaryButton";
import Sidebar from "../../components/Sidebar";
import { FiPlus } from "react-icons/fi";

import "./styles.css";

// const fuck = axios.get('http://localhost:1337/test')
// fuck.then(response => {
//   console.log(response.data)
//   })

export default function CreateAppointment() {
  return (
    <div id="page-create-orphanage">
      <Sidebar />
      <main>
        <form className="create-orphanage-form">
          <fieldset>
            <legend>Relatar consulta</legend>
            <div className="input-block">
              <label htmlFor="name">Nome do médico</label>
              <input id="name" />
            </div>

            <div className="input-block">
              <label htmlFor="name">Tipo de consulta</label>
              <input id="name" />
            </div>

            <div className="input-block">
              <label htmlFor="about">Notas</label>
              <textarea id="name" maxLength={300} />
            </div>

            <div className="input-block">
              <label htmlFor="images">Fotos</label>

              <div className="uploaded-image"></div>
              <button className="new-image">
                <FiPlus size={24} color="#15b6d6" />
              </button>
            </div>
          </fieldset>

          <PrimaryButton type="submit">Enviar</PrimaryButton>
        </form>
      </main>
    </div>
  );
}
