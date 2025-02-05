from flask import Flask, request, jsonify
from flask_cors import CORS
from datetime import datetime
# Habilitar CORS en toda la aplicación


app = Flask(__name__)
CORS(app)

@app.route('/api/simulacion', methods=['POST'])
def simulacion():
    data = request.get_json()

    response = {
<<<<<<< HEAD
        "full_name": f"{data.get('nombre')}  {data.get('apellido')}",
        "age": f"{calcular_edad(data.get('fechaNacimiento'))}",  # Simulamos la edadcalcular_edad(data.get("fechaNacimiento"))
        "contributed_weeks": f"{data.get('semanasCotizadas')}", 
        "salary": data.get('salary'),
        "affiliation_entity": data.get('entidadAfiliado'),
        "estimated_pension": 2500  # Pensiones estimadas simuladas
=======
        "full_name": f"{data.get('UserName', '')}",
        "age": calcular_edad(data.get("DateOfBirth", "2000-01-01")),
        "contributed_weeks": data.get('WeeksContributed', 0),
        "salary": data.get('Salary', 0),
        "affiliation_entity": "Entidad Default",  # Puedes cambiarlo si lo tienes
        "estimated_pension": 2500  # Valor simulado
>>>>>>> d1bd61e (Guardando cambios locales antes de actualizar)
    }

    return jsonify(response), 200

from datetime import datetime

def calcular_edad(fecha_nacimiento):
    # Eliminar la parte de la hora y zona horaria (si existe)
    fecha_nacimiento = fecha_nacimiento.split("T")[0]  # Extraer solo YYYY-MM-DD
    fecha_nacimiento = datetime.strptime(fecha_nacimiento, "%Y-%m-%d")  # Convertir a datetime

    hoy = datetime.today()
    edad = hoy.year - fecha_nacimiento.year - ((hoy.month, hoy.day) < (fecha_nacimiento.month, fecha_nacimiento.day))
    return edad

# Ejecutar el servidor
if __name__ == '__main__':
    app.run(debug=True)
