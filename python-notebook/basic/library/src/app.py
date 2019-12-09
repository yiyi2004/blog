from time import sleep
import requests

def down_load_video(url, name):
    resp = requests.get(url)
    resp.close()

def main():
    resp = requests.get('https://www.baidu.com')
    print(resp.content)
    resp.close()
    sleep(1)


if __name__ == "__main__":
    main()

