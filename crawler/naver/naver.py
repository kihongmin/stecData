import re, os, sys
import selenium
sys.path.append(os.path.dirname(os.path.abspath(os.path.dirname(__file__))))
from util import *
from selenium import webdriver
from bs4 import BeautifulSoup
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
import warnings
warnings.filterwarnings(action='ignore')


def naver(driver_path=None):
    if not driver_path:
        driver_path='/Users/mingihong/chromedriver'
#    driver = headless(driver_path)
    driver = webdriver.Chrome(driver_path)
    driver.get('https://recruit.navercorp.com/naver/job/list/developer')

    final_data = []
    title_set = set()
    dict={'#entType001':'신입', '#entType002':'경력','#entType004':'인턴'}

    for button in ['#entType001','#entType002','#entType004']:
        driver.find_element_by_css_selector(button+' > a').click()
        newbie = dict[button]
        WebDriverWait(driver, 10).until(
            EC.presence_of_element_located((By.CSS_SELECTOR, "#jobListDiv > ul > li:nth-child(1) > a > span > strong"))
        )
        while True:
            if button == '#entType004':
                break
            WebDriverWait(driver, 10).until(
                EC.presence_of_element_located((By.CSS_SELECTOR, '#moreDiv > button'))
            )
            try:
                driver.find_element_by_css_selector('#moreDiv > button').click()
            except:
                break
        html = driver.page_source
        soup = BeautifulSoup(html,'html.parser')

        posts = soup.select('#jobListDiv > ul > li')
        for post in posts:
            post_date = transfrom_date(post.select('a > span > em')[0].text)
            post_url = 'https://recruit.navercorp.com'+post.select('a')[0].get('href')
            post_title = post.select('a > span > strong')[0].text
            post_newbie = make_newbie([newbie])
            
            final_data.append(
                job(post_title,post_url,'naver',post_date,post_newbie).data
            )
    return final_data, driver

def body_text(driver,json):
    driver.get(json['url'])
    WebDriverWait(driver, 10).until(
        EC.presence_of_element_located((By.CSS_SELECTOR, '#content > div > div.career_detail > div.dtl_context > div.context_area'))
    )
    html = driver.page_source
    soup = BeautifulSoup(html,'html.parser')
    txt = soup.select('#content > div > div.career_detail > div.dtl_context > div.context_area')
    if txt:
        json['contents'] = re.sub('\n|\xa0','',txt[0].text)
        print(json)
        return json
    else:
        return None


def run(driver_path=None):
    json_list, driver = naver(driver_path)
    for json in json_list:
        json = body_text(driver,json)
    driver.quit()

    return json_list


if __name__ == "__main__":
    run()
