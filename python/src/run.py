import requests


def hello():
    return "Hello, World!"


if __name__ == "__main__":
    print(hello())
    requests.get("https://ipinfo.io/ip")
