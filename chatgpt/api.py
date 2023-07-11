import json

from flask import Flask, request , jsonify
from chatgpt import openai_function

app = Flask(__name__)

@app.route('/')
def get_data():

    return "<p>Hello world!</p>"



@app.route('/articles', methods=['GET'])
def articles_data():

    company = request.args.get("company")
    field = request.args.get("field")
    hobby = request.args.get("hobby")

    response = openai_function(company,field,hobby)
    print(response)
    data = {"text": response}
    json_string = jsonify(data)
    print("sucsess")

    return json_string

if __name__ == '__main__':
    # app.run(host='163.43.130.132', port=80)
    app.run(host='0.0.0.0', port=5000)

#http://10.200.0.234:5000/articles?company=yahoo&field=インフラ&hobby=IoT