export const adaptSimulationResponse = (apiResponse) => {
  // Si la respuesta es un array en lugar de un objeto con "simulations", la adaptamos
  if (Array.isArray(apiResponse)) {
    apiResponse = { simulations: apiResponse };
  }

  // Validación de datos
  if (!apiResponse || !apiResponse.simulations || !Array.isArray(apiResponse.simulations)) {
    console.error("⚠️ No se recibieron datos válidos o la estructura es incorrecta:", apiResponse);
    return [];
  }

  //console.log("📥 Datos originales de la API:", JSON.stringify(apiResponse.simulations, null, 2));
  //console.log("📥 Daticos:", apiResponse); 
  return apiResponse.simulations.map(sim => ({
    regimen: sim.Regimen || "Desconocido",  
    opciones: Array.isArray(sim.Options) ?  
      sim.Options.map(opt => ({
        semanas: opt.NumberOfWeeks ?? 0,  
        prioridad: opt.Priority ?? 0,     
        monto: opt.Amount ?? 0           
      }))
      : []
  }));
};
