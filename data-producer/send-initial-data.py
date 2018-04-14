import socket
import json

soc = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
soc.connect(("127.0.0.1", 514))

clients_input = {
    "tid": 1,
    "body": "hey",
    "timestamp": 12,
    "params": {
        "a": True,
        "version": 2.21
    }
}
clients_input = json.dumps(clients_input)
soc.send(clients_input.encode("utf8"))  # we must encode the string to bytes
result_bytes = soc.recv(4096)  # the number means how the response can be in bytes
result_string = result_bytes.decode("utf8")  # the return will be in bytes, so decode

print("Result from server is {}".format(result_string))
