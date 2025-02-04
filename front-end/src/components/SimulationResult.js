import React, { useEffect, useState } from "react";
import { adaptSimulationResponse } from "../adapters/apiAdapter";

const SimulationResult = ({ simulationData }) => {
  const [formattedData, setFormattedData] = useState(null);

  useEffect(() => {
    if (simulationData) {
      const adaptedData = adaptSimulationResponse(simulationData);
      setFormattedData(adaptedData);
      
    }
  }, [simulationData]);

  if (!formattedData) return <p>Cargando resultados...</p>;
  console.log(simulationData); 
  return (
    simulationData && (
      <div>
        <h2>Resultados de la Simulación</h2>
        <p>Nombre: {simulationData.nombre} {simulationData.apellido}</p>
        <p>Edad: {simulationData.edad} años</p>
        <p>Semanas Cotizadas: {simulationData.semanasCotizadas}</p>
        <p>Salario: ${simulationData.salario}</p>
        <p>Entidad Afiliada: {simulationData.afiliado}</p>
        <p>Pensión Estimada: ${simulationData.pensionEstimado}</p>
      </div>
    )
  );
};

export default SimulationResult;
