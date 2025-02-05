import React, { useState, useEffect } from "react";
import '../assets/Home.css';

const SimulationForm = ({ onSubmit }) => {
  const [formData, setFormData] = useState({
    UserName: "",
    Salary: "",
    DateOfBirth: "",
    WeeksContributed: "",
    Gender: "",
    AccumulatedFunds: "",
    HasBeneficiary: "no",
    TypeBeneficiary: "",
  });

  const [isChanged, setIsChanged] = useState(false);

  useEffect(() => {
    // console.log("📢 Formulario actualizado:", formData);
  }, [formData]);

  const handleChange = (e) => {
    const { name, value } = e.target;

    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));

    setIsChanged(true); // Se activa cuando el usuario cambia un valor
  };

  const handleSubmit = (e) => {
    e.preventDefault();

    if (!isChanged) return;

    const formattedDate = new Date(formData.DateOfBirth).toISOString();

    const formattedData = {
      Salary: parseFloat(formData.Salary),  // Convertir a número
      DateOfBirth: formattedDate,
      WeeksContributed: parseInt(formData.WeeksContributed, 10), // Convertir a entero
      Gender: formData.Gender === "hombre" ? "M" : "F",
      AccumulatedFunds: parseFloat(formData.AccumulatedFunds), // Convertir a número
      HasBeneficiary: formData.HasBeneficiary === "si",
      TypeBeneficiary: formData.TypeBeneficiary || null,
      UserName: formData.UserName, // Aseguramos que se incluya el nombre del usuario
    };

    onSubmit(formattedData);  // Enviamos los datos al componente superior

    setIsChanged(false);
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <label>Nombre de usuario:</label>
        <input type="text" name="UserName" value={formData.UserName} onChange={handleChange} />
      </div>
      <div>
        <label>Salario:</label>
        <input type="number" name="Salary" value={formData.Salary} onChange={handleChange} required />
      </div>
      <div>
        <label>Fecha de nacimiento:</label>
        <input type="date" name="DateOfBirth" value={formData.DateOfBirth} onChange={handleChange} required />
      </div>
      <div>
        <label>Semanas Cotizadas:</label>
        <input type="number" name="WeeksContributed" value={formData.WeeksContributed} onChange={handleChange} required />
      </div>
      <div>
        <label>Género:</label>
        <select name="Gender" value={formData.Gender} onChange={handleChange} required>
          <option value="">Selecciona</option>
          <option value="hombre">Hombre</option>
          <option value="mujer">Mujer</option>
        </select>
      </div>
      <div>
        <label>Fondos Acumulados:</label>
        <input type="number" name="AccumulatedFunds" value={formData.AccumulatedFunds} onChange={handleChange} required />
      </div>
      <div>
        <label>¿Tiene beneficiario?</label>
        <select name="HasBeneficiary" value={formData.HasBeneficiary} onChange={handleChange} required>
          <option value="no">No</option>
          <option value="si">Sí</option>
        </select>
      </div>
      {formData.HasBeneficiary === "si" && (
        <div>
          <label>Tipo de beneficiario:</label>
          <select name="TypeBeneficiary" value={formData.TypeBeneficiary} onChange={handleChange} required>
            <option value="">Selecciona</option>
            <option value="Spouse">Conyugue</option>
            <option value="Parent">Padre</option>
            <option value="Child">Hijo</option>
            <option value="Sibling">Hermano</option>
          </select>
        </div>
      )}
      <button type="submit" disabled={!isChanged}>Simular</button>
    </form>
  );
};

export default SimulationForm;
