
import React, { useState, useEffect } from "react";
import SimulationForm from "../components/SimulationForm.js";
import SimulationResult from "../components/SimulationResult.js";
import { adaptSimulationResponse } from "../adapters/apiAdapter";
import '../assets/Home.css';

const Home = () => {
  const [formData, setFormData] = useState(null); // 🔹 Observer → Observamos cambios en formData
  const [simulationData, setSimulationData] = useState(null);
  const [userName, setUserName] = useState(""); // Estado para almacenar el nombre del usuario
  const [showReloadButton, setShowReloadButton] = useState(false);

  const handleSimulationSubmit = (data) => {
    setFormData(data);
    setUserName(data.UserName);
    setShowReloadButton(true); // 🔹 Muestra el botón después del envío
  };

  useEffect(() => {
    if (formData) {
      console.log("📡 Enviando datos a la API...", formData);
      fetch("http://localhost:8080/api/pension/v1/simulate", {
        method: "POST",
          body: JSON.stringify(formData),
          headers: { "Content-Type": "application/json" },
        })
          .then((res) => res.json())
          .then((data) => {
            console.log("Datos recibidos de la API:", data); // Verifica lo que está recibiendo
            const adaptedData = adaptSimulationResponse(data);
            setSimulationData(adaptedData);
          })
          .catch((err) => console.error("❌ Error al enviar:", err));
    }
  }, [formData]);

  
  return (
    <div className="Home">
      <section id="simulador">
        <p>Completa el siguiente formulario para conocer tu proyección pensional.</p>
        <SimulationForm onSubmit={handleSimulationSubmit} />
        {simulationData && <SimulationResult userName={userName} simulationData={simulationData} />}
        
        {/* 🔹 Botón que se muestra después de enviar el formulario */}
        <button 
          type="button" 
          onClick={() => window.location.reload()} 
          hidden={!showReloadButton} 
        >
          Hacer otra Simulación
        </button>
      </section>
    </div>
  );
};

export default Home;
