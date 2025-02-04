export const adaptSimulationResponse = (apiResponse) => {
    // Adaptación de la respuesta de la API
    const adaptedResponse = {
      nombre: apiResponse.full_name?.split(" ")[0] || "N/A",
      apellido: apiResponse.full_name?.split(" ")[3] || "N/A",
      edad: apiResponse.age || "No disponible",
      semanasCotizadas: apiResponse.contributed_weeks || 0,
      salario: apiResponse.salary || 0,
      afiliado: apiResponse.affiliation_entity || "No afiliado",
      pensionEstimado: apiResponse.estimated_pension || "Calculando...",
    };
  
    return adaptedResponse;
  };
  