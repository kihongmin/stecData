import re
import datetime
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC

class job:
    def __init__(self,title,url,company,start_date,newbie):
        self.data = {
        'title' : title,
        'url' : url,
        'company' : company,
        'start_date' : start_date,
        'newbie' : newbie,
        'contents' : ''
        }

    def set_contents(self,contents):
        self.data['contents'] += contents
def head(driver_path):
    return webdriver.Chrome(driver_path)

def headless(driver_path):
    options = webdriver.ChromeOptions()
    options.add_argument('--window-size=1920,1080')
    options.add_argument('--headless')
    options.add_argument("disable-gpu")
    driver = webdriver.Chrome(driver_path, chrome_options=options)
    return driver

def transfrom_date(date,is_rocket=False):
    if is_rocket:
        p = re.compile('등록')
        if not p.search(date):
            return None
    date = re.sub('[^0-9]+','',date)
    now = datetime.datetime.now()
    if len(date) == 6:
        date = date[2:]
    if (now - datetime.timedelta(days=1)).strftime('%m%d') == date:
        return str(now.year)+date
    else:
        return None


def make_newbie(newbie_list):
    newbie = {'인턴':10,'신입':50,'경력':90}
    newbie_list = list(filter(lambda x:x in newbie,newbie_list))
    return list(map(lambda x:newbie[x],newbie_list))

def programmers_newbie(newbie_list):
    p = re.compile('경력 무관')
    m = re.compile('경력')
    if p.search(newbie_list):
        return [50]
    elif m.search(newbie_list):
        return [90]
    else:
        return []
