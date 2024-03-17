from flask import Flask, request, jsonify, send_file
from sayaBot import model

app = Flask(__name__)

@app.route('/process_json', methods=['POST'])
def process_json():
    try:
        # Запускаем модель для обработки JSON-файла
        response_data = model('example.json')

        # Возвращаем созданный JSON-файл
        return send_file(response_data), 200
    except Exception as e:
        return jsonify({"error": "asd" + str(e)}), 400

if __name__ == '__main__':
    app.run(debug=True)
