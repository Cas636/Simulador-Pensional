import React from 'react';
import { Link } from 'react-router-dom';
import '../assets/Menu.css'; 


const Menu = () => {
  return (
    <div>
      <nav id="navigation">
        <ul>
          <li>
            <Link to="/">Inicio</Link>
          </li>
          <li>
            <Link to="/PensionReformPage">Información</Link>
          </li>
        </ul>
      </nav>
    </div>
  );
};

export default Menu;
