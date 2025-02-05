import React from "react";
import '../assets/SimulationResult.css';  // Asegúrate de importar el archivo CSS

const SimulationResult = ({ userName, simulationData }) => {
  if (!simulationData) return <p>Cargando resultados...</p>;

  return (
    <div className="simulation-result">
      <h2 className="result-title">Resultados de la Simulación para <span className="user-name">{userName}</span></h2>
      {simulationData.map((regimenData, index) => (
        <div key={index} className="regimen-section">
          <h3 className="regimen-title">Régimen: {regimenData.regimen}</h3>
          <ul className="options-list">
            {regimenData.opciones.map((opcion, idx) => (
              <li key={idx} className="option-item">
                <strong className="option-priority">Prioridad {opcion.prioridad}:</strong> 
                <span className="option-weeks">{opcion.semanas} semanas</span> - 
                <span className="option-amount">${opcion.monto.toFixed(2)}</span>
              </li>
            ))}
          </ul>
        </div>
      ))}
    </div>
  );
};

export default SimulationResult;
