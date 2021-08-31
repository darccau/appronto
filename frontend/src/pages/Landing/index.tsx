import React from "react";
import { FaArrowRight } from 'react-icons/fa';
import { Link } from 'react-router-dom'

// import logoImg from '../../assets/images/logo.svg';
// <img src={logoImg} alt="Happy" />

import './styles.css';

export default function Landing() {
  return (
    <div id="page-landing">
      <div className="content-wrapper">
      <h1>Appronto</h1>
 
        <main>
          <h1>Seus exames com voce!</h1>
          <p>Gerencie seus prontuarios e torne sua vida mais simples</p>
        </main>

        <Link to="/orphanages/create" className="enter-app">
          <FaArrowRight size={26} color="rgba(0, 0, 0, 0.6)" />
        </Link>
      </div>
    </div>
  );
}
