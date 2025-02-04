from flask import Flask, request, jsonify
from flask_cors import CORS
from datetime import datetime
# Habilitar CORS en toda la aplicación


app = Flask(__name__)
CORS(app)

# Ruta para el endpoint de simulación
@app.route('/api/simulacion', methods=['POST'])
def simulacion():
    # Obtenemos los datos que se envían desde el frontend
    data = request.get_json()

    # Simulamos una respuesta (puedes adaptarlo a lo que necesites)
    response = {
        "full_name": f"{data.get('nombre')}  {data.get('apellido')}",
        "age": f"{calcular_edad(data.get('fechaNacimiento'))}",  # Simulamos la edadcalcular_edad(data.get("fechaNacimiento"))
        "contributed_weeks": f"{data.get('semanasCotizadas')}", 
        "salary": data.get('salario'),
        "affiliation_entity": data.get('entidadAfiliado'),
        "estimated_pension": 2500  # Pensiones estimadas simuladas
    }

    return jsonify(response), 200  # Devolvemos una respuesta JSON

def calcular_edad(fecha_nacimiento):
    fecha_nacimiento = datetime.strptime(fecha_nacimiento, "%Y-%m-%d")  # Formato de fecha esperado
    hoy = datetime.today()
    edad = hoy.year - fecha_nacimiento.year - ((hoy.month, hoy.day) < (fecha_nacimiento.month, fecha_nacimiento.day))
    return edad

# Ejecutar el servidor
if __name__ == '__main__':
    app.run(debug=True)
