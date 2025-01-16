from flask import Flask, render_template, request, redirect, url_for
from task import Task

app = Flask(__name__)
task_manager = Task()
Task._create_table()


@app.route("/")
def index():
    tasks = task_manager.get_tasks()
    return render_template("index.html", tasks=tasks)


@app.route("/add", methods=["POST"])
def add_task():
    description = request.form.get("description")
    if description:
        task_manager.add_task(description)
    return redirect(url_for("index"))


@app.route("/complete/<int:task_id>", methods=["POST"])
def complete_task(task_id):
    task_manager.complete_task(task_id)
    return redirect(url_for("index"))


if __name__ == "__main__":
    app.run(debug=True)
