import React, { useState, useEffect } from "react";
import '../assets/Home.css';

const SimulationForm = ({ onSubmit }) => {
  const [formData, setFormData] = useState({
    nombre: "",
    apellido: "",
    fechaNacimiento: "",
    genero: "",
    semanasCotizadas: "",
    salario: "",
    entidadAfiliado: "",
    tieneBeneficiario: "no",
    tipoBeneficiario: "",
  });

  // 🔹 Observer: Cada vez que formData cambia, lo mostramos en consola
  useEffect(() => {
    console.log("📢 Formulario actualizado:", formData);
  }, [formData]); 

  // Manejo de cambios en los campos del formulario
  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({ ...prevData, [name]: value }));
  };

  // Enviar los datos a `Home`
  const handleSubmit = (e) => {
    e.preventDefault();
    onSubmit(formData);
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <label>Nombre:</label>
        <input type="text" name="nombre" value={formData.nombre} onChange={handleChange} required />
      </div>
      <div>
        <label>Apellido:</label>
        <input type="text" name="apellido" value={formData.apellido} onChange={handleChange} required />
      </div>
      <div>
        <label>Fecha de nacimiento:</label>
        <input type="date" name="fechaNacimiento" value={formData.fechaNacimiento} onChange={handleChange} required />
      </div>
      <div>
        <label>Género:</label>
        <select name="genero" value={formData.genero} onChange={handleChange} required>
          <option value="">Selecciona</option>
          <option value="mujer">Mujer</option>
          <option value="hombre">Hombre</option>
        </select>
      </div>
      <div>
        <label>Semanas Cotizadas:</label>
        <input type="number" name="semanasCotizadas" value={formData.semanasCotizadas} onChange={handleChange} required />
      </div>
      <div>
        <label>Salario:</label>
        <input type="number" name="salario" value={formData.salario} onChange={handleChange} required />
      </div>
      <div>
        <label>Entidad donde está afiliado:</label>
        <input type="text" name="entidadAfiliado" value={formData.entidadAfiliado} onChange={handleChange} required />
      </div>
      <div>
        <label>¿Tiene beneficiario?</label>
        <select name="tieneBeneficiario" value={formData.tieneBeneficiario} onChange={handleChange} required>
          <option value="no">No</option>
          <option value="si">Sí</option>
        </select>
      </div>
      {formData.tieneBeneficiario === "si" && (
        <div>
          <label>Tipo de beneficiario:</label>
          <select name="tipoBeneficiario" value={formData.tipoBeneficiario} onChange={handleChange} required>
            <option value="">Selecciona</option>
            <option value="conyugue">Conyugue</option>
            <option value="padre">Padre</option>
            <option value="hijo">Hijo</option>
            <option value="hermano">Hermano</option>
          </select>
        </div>
      )}
      <button type="submit">Simular</button>
    </form>
  );
};

export default SimulationForm;
