from crawler.nexon import nexon
from crawler.netmarble import netmarble
from crawler.programmers import programmers
from crawler.rocketpunch import rocketpunch

def main(chrome_path=None):
    json_list = []
    json_list.extend(nexon.run(chrome_path))
    json_list.extend(netmarble.run(chrome_path))
    json_list.extend(rocketpunch.run(chrome_path))
    json_list.extend(programmers.run(chrome_path))
    return json_list

main()

#if __name__ == "__main__":
