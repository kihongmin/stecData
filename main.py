from crawler.nexon import nexon
from crawler.netmarble import netmarble
from crawler.programmers import programmers
from crawler.rocketpunch import rocketpunch

if __name__ == "__main__":
    json_list = []
    json_list.extend(nexon.run())
    json_list.extend(netmarble.run())
    json_list.extend(rocketpunch.run())
    json_list.extend(programmers.run())
