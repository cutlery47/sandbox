from flask import Flask
from flask import request, make_response
import psycopg2
app = Flask(__name__)

def connect_to_postgres():
    db = psycopg2.connect(dbname="Flask",
                         user="ORtiz",
                         password="12345",
                         host="localhost",
                         port="5432")
    return db

@app.route("/people/", methods=['GET'])
def get_people():
    db = connect_to_postgres()

    cur = db.cursor()
    cur.execute(f"SELECT * FROM \"People\"")
    data = cur.fetchall()

    response = []
    for el in data:
        dat = {'first_name': el[0], 'id': el[1]}
        response.append(dat)

    cur.close()
    db.close()

    return make_response(response)

@app.route("/people/", methods=['POST'])
def add_people():
    first_name = request.form['first_name']

    db = connect_to_postgres()

    cur = db.cursor()
    cur.execute(f"INSERT INTO \"People\" (first_name) VALUES (\'{first_name}\')")

    db.commit()
    cur.close()
    db.close()

    return "Person added sucessfully"

@app.route("/about/")
def about():
    return "fucking nigger"

if __name__ == "__main__":
    app.run(debug=True)
