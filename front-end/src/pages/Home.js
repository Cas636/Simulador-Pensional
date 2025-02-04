import React, { useState, useEffect } from "react";
import SimulationForm from "../components/SimulationForm.js";
import SimulationResult from "../components/SimulationResult.js";
import { adaptSimulationResponse } from "../adapters/apiAdapter";
import '../assets/Home.css';

const Home = () => {
  const [formData, setFormData] = useState(null); // 🔹 Observer → Observamos cambios en formData
  const [simulationData, setSimulationData] = useState(null);

  const handleSimulationSubmit = (data) => {
    setFormData(data); // Se activa el useEffect cuando formData cambia
  };

  useEffect(() => {
    if (formData) {
      console.log("📡 Enviando datos a la API...", formData);
      fetch("http://127.0.0.1:5000/api/simulacion", {
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
        {simulationData && <SimulationResult simulationData={simulationData} />}
      </section>
    </div>
  );
};

export default Home;
