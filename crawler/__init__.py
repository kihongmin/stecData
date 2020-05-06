import re
import datetime

from selenium import webdriver


_chrome_options_ = ['--window-size=1920,1080', '--headless', 'disable-gpu']
_driver_path_ = './chromedriver'


def connect(driver_path=_driver_path_, args=_chrome_options_):
    options = webdriver.ChromeOptions()
    if args is not None:
        for arg in args:
            options.add_argument(arg)
    driver = webdriver.Chrome(driver_path, chrome_options=options)
    return driver


def transfrom_date(date,is_rocket=False):
    if is_rocket:
        p = re.compile('등록')
        if not p.search(date):
            return None
    date = re.sub('[^0-9]+','',date)
    now = datetime.datetime.now()
    len_date = len(date)
    if len_date == 4:
        date = '2020'+date
    elif len_date == 6:
        date = '20'+date
    elif len_date == 12:
        date = '20'+date[:6]
    elif len_date == 16:
        date = date[:8]
    else:
        return None
    if (now - datetime.timedelta(days=1)).strftime('%Y%m%d') == date:
        return date
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
