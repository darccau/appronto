import React from "react";

import PrimaryButton from "../../components/PrimaryButton";
import Sidebar from "../../components/Sidebar";

import './styles.css';
import { FiPlus } from "react-icons/fi";

export default function OrphanagesMap() {
  return (
    <div id="page-create-orphanage">
      <Sidebar />

      <main>
        <form className="create-orphanage-form">
          <fieldset>
            <legend>Create Appointment</legend>

            <div className="input-block">
              <label htmlFor="name">Doctor name</label>
              <input id="name" />
            </div>

            <div className="input-block">
              <label htmlFor="name">Appointment type</label>
              <input id="name" />
            </div>

            <div className="input-block">
              <label htmlFor="about">Note</label>
              <textarea id="name" maxLength={300} />
            </div>

            <div className="input-block">
              <label htmlFor="images">Fotos</label>

              <div className="uploaded-image">

              </div>

              <button className="new-image">
                <FiPlus size={24} color="#15b6d6" />
              </button>
            </div>
          </fieldset>

          <PrimaryButton type="submit">Submit</PrimaryButton>
        </form>
      </main>
    </div>
  );
}

