import os
from flask import Flask, render_template, request

app = Flask(__name__)


@app.route("/")
def index():
    return render_template("index.html", title="Home", name="Harold")


@app.route("/about")
def about():
    return render_template("about.html", title="About")

@app.get('/contact')
def contact():
    return render_template('contact.html', name="Billy Bob")

@app.post('/contact')
def contact_post():
    return render_template('contact.html', msg=request.form.get('message'))

if __name__ == "__main__":
    app.run(debug=True)
