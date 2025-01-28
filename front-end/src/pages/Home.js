import React, { useState } from 'react';
import '../assets/Home.css'; 

// Componente para el formulario de simulación
const SimulationForm = ({ onSubmit }) => {
  const [edad, setEdad] = useState('');
  const [semanasCotizadas, setSemanasCotizadas] = useState('');
  const [salario, setSalario] = useState('');
  const [genero, setGenero] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    onSubmit({ edad, semanasCotizadas, salario, genero });
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <label>Edad:</label>
        <input
          type="number"
          value={edad}
          onChange={(e) => setEdad(e.target.value)}
          required
        />
      </div>
      <div>
        <label>Semanas Cotizadas:</label>
        <input
          type="number"
          value={semanasCotizadas}
          onChange={(e) => setSemanasCotizadas(e.target.value)}
          required
        />
      </div>
      <div>
        <label>Salario:</label>
        <input
          type="number"
          value={salario}
          onChange={(e) => setSalario(e.target.value)}
          required
        />
      </div>
      <div>
        <label>Género:</label>
        <select
          value={genero}
          onChange={(e) => setGenero(e.target.value)}
          required
        >
          <option value="mujer">Mujer</option>
          <option value="hombre">Hombre</option>
        </select>
      </div>
      <button type="submit">Simular</button>
    </form>
  );
};

const Home = () => {
  const handleSimulationSubmit = (formData) => {
    // Lógica para manejar la simulación, enviar datos a la API o realizar cálculos
    console.log('Datos de simulación:', formData);
  };

  return (
    <div className="Home">

      <section id="simulador">
        <p>Completa el siguiente formulario para conocer tu proyección pensional.</p>
        <SimulationForm onSubmit={handleSimulationSubmit} />
      </section>
    </div>
  );
};

export default Home;
