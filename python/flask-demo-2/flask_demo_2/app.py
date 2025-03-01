from flask import Flask, jsonify, request
from middleware import MwTracker

mw_tracker = MwTracker()

app = Flask(__name__)

@app.route('/')
def home():
    return jsonify({"message": "Welcome to the Flask API!"})

@app.route('/hello', methods=['GET'])
def hello():
    name = request.args.get('name', 'World')
    return jsonify({"message": f"Hello, {name}!"})

def main():
    app.run(debug=True)

if __name__ == '__main__':
    main()
