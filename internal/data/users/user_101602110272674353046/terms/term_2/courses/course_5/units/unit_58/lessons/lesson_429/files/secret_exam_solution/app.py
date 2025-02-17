from flask import Flask, render_template, request

app = Flask(__name__)

planets = {
    "Mercury": {"Type": "Terrestrial", "Day Length (Earth days)": 58.6, "Moons": 0},
    "Venus": {"Type": "Terrestrial", "Day Length (Earth days)": 243, "Moons": 0},
    "Earth": {"Type": "Terrestrial", "Day Length (Earth days)": 1, "Moons": 1},
    "Mars": {"Type": "Terrestrial", "Day Length (Earth days)": 1.03, "Moons": 2},
    "Jupiter": {"Type": "Gas Giant", "Day Length (Earth days)": 0.41, "Moons": 79},
    "Saturn": {"Type": "Gas Giant", "Day Length (Earth days)": 0.45, "Moons": 83},
    "Uranus": {"Type": "Ice Giant", "Day Length (Earth days)": 0.72, "Moons": 27},
    "Neptune": {"Type": "Ice Giant", "Day Length (Earth days)": 0.67, "Moons": 14},
}


@app.route("/")
def home():
    return render_template("index.html", title="Home", planets=planets)


@app.route("/search")
def search():
    return render_template("search.html", title="Search")


@app.route("/result")
def result():
    query = request.args.get("term", "").capitalize()
    result = planets.get(query, "No matching planet found.")
    return render_template("result.html", query=query, result=result)


if __name__ == "__main__":
    app.run(debug=True)
