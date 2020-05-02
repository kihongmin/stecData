from crawler.nexon import nexon
from crawler.netmarble import netmarble
from crawler.programmers import programmers
from crawler.rocketpunch import rocketpunch

if __name__ == "__main__":
    chromedriver_path = '../geekermeter-data/chromedriver'
    json_list = []
    json_list.extend(nexon.run(chromedriver_path))
    json_list.extend(netmarble.run(chromedriver_path))
    json_list.extend(rocketpunch.run(chromedriver_path))
    json_list.extend(programmers.run(chromedriver_path))
