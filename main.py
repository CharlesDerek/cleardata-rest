from flask import Flask, jsonify, request, send_from_directory

app = Flask(__name__)
@app.route('/v1/get/<string>/', methods = ["GET"])
def get(string):
    alphabet = "abcdefghijklmnopqrstuvwxyz"
    if string == None:
        return jsonify({"result": False})
    for char in alphabet:
        if char not in string:
            return jsonify({"result": False})
    return jsonify({"result": True})

@app.route('/', methods = ["GET"])
def index():
    return send_from_directory("../public", "index.html")

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=3000)
