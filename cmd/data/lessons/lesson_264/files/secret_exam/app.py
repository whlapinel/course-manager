from flask import Flask, render_template, request, redirect, url_for
# Bug(5) Missing import of Task
app = Flask(__name__)


@app.route("/")
def index():
    # Bug (6): Missing the call to get tasks
    # Bug (7): Need to pass tasks into render_template() below
    return render_template("index.html")


@app.route("/add", methods=["POST"])
def add_task():
    description = request.form.get("description")
    if description:
        pass  # remove this line after you've solved Bug(3) or the code won't work
        # Bug(8): Need to call add_task and pass in the description
    return redirect(url_for("index"))


@app.route("/complete/<int:task_id>", methods=["POST"])
def complete_task(task_id):
    # Bug(9): Call complete_task() with the id
    return redirect(url_for("index"))


if __name__ == "__main__":
    # Bug(10): Fix the indentation below
app.run(debug=True)


