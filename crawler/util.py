import re
import datetime
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC

def headless(driver_path):
    options = webdriver.ChromeOptions()
    options.add_argument('--window-size=1920,1080')
    options.add_argument('--headless')
    options.add_argument("disable-gpu")
    driver = webdriver.Chrome(driver_path, chrome_options=options)
    return driver

def transfrom_date(date):
    p = re.compile('등록')
    if p.search(date):
        date = re.sub('[^0-9]+','',date)
        now = datetime.datetime.now()
        if (now - datetime.timedelta(days=1)).strftime('%m%d') == date:
            return str(now.year)+date
        else:
            return None
    else:
        return None
