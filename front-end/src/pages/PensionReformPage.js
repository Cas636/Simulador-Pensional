import React from "react";
import "../assets/PensionReformPage.css";


  const PensionReformPage = () => {
   

  return (
    <div className="PensionReformPage">

      <main className="Content">
        <section className="Intro">
        <h1> Todo lo que necesitas saber sobre el nuevo modelo pensional</h1>
          <h3>Lo más importante</h3>
          <ul>
            <li>
              El nuevo sistema combina esfuerzos entre Colpensiones y las Administradoras del Régimen de Ahorro Individual.
            </li>
            <li>
              Los primeros 2.3 salarios mínimos de cotización irán a Colpensiones, el excedente a una Administradora.
            </li>
            <li>
              Los requisitos de edad y semanas cotizadas se mantienen, con una reducción gradual para mujeres a partir de 2025.
            </li>
          </ul>
        </section>

        <section className="Categories">
          <h2>Clasificación de afiliados</h2>
          <div className="Category">
            <h3>Régimen de Transición</h3>
            <p>
              Afiliados que cumplen con los requisitos previos al 30 de junio de 2025 seguirán bajo la Ley 100 de 1993.
            </p>
          </div>
          <div className="Category">
            <h3>Oportunidad de Traslado</h3>
            <p>
              Afiliados con ciertas condiciones podrán elegir entre permanecer en su Administradora actual o trasladarse a Colpensiones.
            </p>
          </div>
          <div className="Category">
            <h3>Nueva Ley</h3>
            <p>
              Aplica para afiliados con menos de 750 (mujeres) o 900 (hombres) semanas cotizadas al 1° de julio de 2025.
            </p>
          </div>
        </section>

        <section className="FAQ">
          <h2>Preguntas frecuentes</h2>
          <ul>
            <li>¿Cuándo entra en vigencia la nueva ley? <span>1 de julio de 2025.</span></li>
            <li>¿Qué pasará con mis recursos actuales? <span>Se mantendrán según el régimen aplicable.</span></li>
            <li>¿Cómo conocer mis semanas cotizadas? <span>Consulta en tu Administradora o Colpensiones.</span></li>
          </ul>
        </section>
      </main>
    </div>
  );
};

export default PensionReformPage;
