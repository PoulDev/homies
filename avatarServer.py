from flask import Flask
from avatars import randAvatar, genAvatar

app = Flask(__name__)

@app.route('/')
def index():
    return randAvatar()

app.run('127.0.0.1', 8000, debug=True)