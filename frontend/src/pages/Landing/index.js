import React from "react";
// import { Link } from 'react-router-dom'

import './styles.css';

export default function Landing() {
  return (
    <div id="page-landing">
      <div className="content-wrapper">
 
        <main>
          <h1>Seus exames com voce!</h1>
          <p>Gerencie seus prontuarios e torne sua vida mais simples</p>
        </main>

          <div id="login">
          <form>
          <label for="user">Usuario</label>
          <input type="text" name="user"/>
          <label for="password">Senha</label>
          <input type="password" name="password"/>
          </form>
          </div>

      </div>
    </div>
  );
}
