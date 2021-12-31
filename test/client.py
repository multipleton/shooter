# IMPORTANT #
# RUN HOST: python3 client.py 1200 john.doe create
# RUN PLAYER: python3 client.py 1201 bob.simple join
# IMPORTANT #

# FLOW #
# 1. create user
# 2. create server and start server || join server
# 3. connect by udp to get data
# 4. sometimes send player position and bullet spawn
# FLOW #

import sys
import socket
import json
import random

from urllib import request

address = ("127.0.0.1", int(sys.argv[1]))
server = ("127.0.0.1", 9000)

username = sys.argv[2]
action = sys.argv[3]

# HTTP

def login():
  url = "http://127.0.0.1:8080/auth/login"
  body = { "username": username }
  encoded = json.dumps(body).encode('utf8')
  headers = {'content-type': 'application/json'}
  req = request.Request(url, data=encoded, headers=headers, method="POST")
  response = request.urlopen(req)
  return json.loads(response.read().decode('utf8'))['id']

def create(userId):
  url = "http://127.0.0.1:8080/servers/users/" + str(userId) + "/create"
  req = request.Request(url, method="POST")
  response = request.urlopen(req)
  return json.loads(response.read().decode('utf8'))['id']

def join(userId, serverId):
  url = "http://127.0.0.1:8080/servers/" + str(userId) + "/users/" + str(serverId) + "/join"
  req = request.Request(url, method="POST")

def start(serverId):
  url = "http://127.0.0.1:8080/servers/" + str(serverId) + "/start"
  req = request.Request(url, method="POST")

userId = login()
serverId = -1

if action == "create":
  serverId = create(userId)

if action == "join":
  join(userId, 1) # NOTE: server id hardcoded

# UDP

opened_socket = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
opened_socket.bind(address)

connect = {
  "type": "connect",
  "data": {
    "id": userId
  }
}

byte_message = bytes(json.dumps(connect), "utf-8")
opened_socket.sendto(byte_message, server)

def generate_random_position():
  return {
    "x": random.randint(0, 50),
    "y": random.randint(0, 50),
    "angle": random.randint(0, 360)
  }

def send_position(position):
  message = {
    "type": "position",
    "data": position
  }
  byte_message = bytes(json.dumps(message), "utf-8")
  opened_socket.sendto(byte_message, server)

def send_bullet(position):
  message = {
    "type": "bullet",
    "data": position
  } 
  byte_message = bytes(json.dumps(message), "utf-8")
  opened_socket.sendto(byte_message, server)

if action == "create":
  start(serverId)

while True:
    data, address = opened_socket.recvfrom(1024)
    print("Received:", data.decode('utf-8'))
    position = generate_random_position()
    if random.randint(0, 100) > 90:
      send_position()
    if random.randint(0, 100) > 95:
      send_bullet()

opened_socket.close()
