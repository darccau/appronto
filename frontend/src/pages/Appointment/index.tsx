import React from "react";
import Sidebar from "../../components/Sidebar";
import axios from "axios";

import './styles.css';

const retriveAppointments = ()  => {
  axios.post('http://localhost:1337/appointment/appointments', {
      "id_user":"1"
      }).then(response => {console.log(response.data)})
}

retriveAppointments()

export default function Orphanage() {
  return (
    <div id="page-orphanage">
      <Sidebar />

      <main>
        <div className="orphanage-details">

          <div className="images">
            <button className="active" type="button">
              <img src="https://www.gcd.com.br/wp-content/uploads/2020/08/safe_image.jpg" alt="Lar das meninas" />
            </button>
            <button type="button">
              <img src="https://www.gcd.com.br/wp-content/uploads/2020/08/safe_image.jpg" alt="Lar das meninas" />
            </button>
            <button type="button">
              <img src="https://www.gcd.com.br/wp-content/uploads/2020/08/safe_image.jpg" alt="Lar das meninas" />
            </button>
            <button type="button">
              <img src="https://www.gcd.com.br/wp-content/uploads/2020/08/safe_image.jpg" alt="Lar das meninas" />
            </button>
            <button type="button">
              <img src="https://www.gcd.com.br/wp-content/uploads/2020/08/safe_image.jpg" alt="Lar das meninas" />
            </button>
            <button type="button">
              <img src="https://www.gcd.com.br/wp-content/uploads/2020/08/safe_image.jpg" alt="Lar das meninas" />
            </button>
          </div>
          
          <div className="orphanage-details-content">
            <h1>Lar das meninas</h1>
            <p>Presta assistência a crianças de 06 a 15 anos que se encontre em
            situação de risco e/ou vulnerabilidade social.</p>

            <hr />

            <h2>Instruções para visita</h2>
            <p>Venha como se sentir mais à vontade e traga muito amor para dar.</p>

            <div className="open-details">
              <div className="hour">
              bla blabla bla  
              bla blabla bla  
              bla blabla bla  
              bla blabla bla  
              </div>
            </div>

          </div>
        </div>
      </main>
    </div>
  );
}
