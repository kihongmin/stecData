import re, os, sys
import selenium
sys.path.append(os.path.dirname(os.path.abspath(os.path.dirname(__file__))))
from util import *
from selenium import webdriver
from bs4 import BeautifulSoup
import warnings
warnings.filterwarnings(action='ignore')


def netmarble(driver_path=None):
    if not driver_path:
        driver_path='/Users/mingihong/chromedriver'
    driver = headless(driver_path)
#    driver = webdriver.Chrome(driver_path)
    driver.get('https://m.netmarble.com/rem/www/noticelist.jsp')

    final_data = []
    while True:
        crawled_data = []
        WebDriverWait(driver, 10).until(
            EC.presence_of_element_located((By.CSS_SELECTOR, "#contents > div > div > div > div.recruit_list_wrapper > ul"))
        )
        html = driver.page_source
        soup = BeautifulSoup(html,'html.parser')

        posts = soup.select('#contents > div > div > div > div.recruit_list_wrapper > ul > li')
        for post in posts:

            post_date = transfrom_date(post.select('div.cw_jopinfo > a > span.cw_info > span.cw_range')[0].text)
            if not post_date:
                continue
            post_title = post.select('div.cw_jopinfo > a > span.cw_title')[0].text
            post_url = 'https://m.netmarble.com/rem/www'+post.select('div.cw_jopinfo > a')[0].get('href')[1:]
            post_newbie = make_newbie(post.select('div.cw_jopinfo > a > span.cw_info > span.cw_type')[0].text)

            crawled_data.append(
                job(post_title,post_url,'netmarble',post_date,post_newbie).data
            )

        final_data.extend(crawled_data)
        _next = driver.find_element_by_css_selector('#contents > div > div > div > div.recruit_list_wrapper > div.recruit_pagination > button.page_next')
        if not _next.get_attribute('disabled'):
            _next.click()
        else:
            break


    return final_data, driver

def body_text(driver,json):
    driver.get(json['url'])

    html = driver.page_source
    soup = BeautifulSoup(html,'html.parser')
    txt = soup.select('#tmpCapture > div > table')
    if txt:
        json['contents'] = txt[0].text
        return json
    else:
        return None


def run(driver_path=None):

    print('start crawling : netmarble')
    json_list, driver = netmarble(driver_path)
    print('The number of netmarble post : %d'%(len(json_list)))
    for i,json in enumerate(json_list):
        if i%10 == 0:
            print('netmarble post : %d'%(i))
        json = body_text(driver,json)
    driver.quit()
    print('finish crawling : netmarble')

    return json_list


if __name__ == "__main__":
    run()
