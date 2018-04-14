import socket
import json


def readFromFile(filename):
    with open(filename, 'r') as f:
        return f.readlines()


def send_packages(soc, packages):
    def send_package(package):
        soc.send(package.encode("utf8"))  # we must encode the string to bytes
        result_bytes = soc.recv(4096)  # the number means how the response can be in bytes
        result_string = result_bytes.decode("utf8")  # the return will be in bytes, so decode
        return result_string

    for package in packages:
        print(package)
        # clients_input = json.dumps(clients_input)
        send_package(package)


DIR = 'test-sets/'


def get_socket(server, port):
    soc = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    soc.connect((server, port))
    return soc


data = readFromFile(DIR + "initial-graph.txt")
soc = get_socket("127.0.0.1", 514)
send_packages(soc, data)
