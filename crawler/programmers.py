import re,os,sys,json

import selenium
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.common.by import By
from bs4 import BeautifulSoup

sys.path.append(os.path.dirname(os.path.abspath(os.path.dirname(__file__))))

from . import connect, transfrom_date, make_newbie, programmers_newbie
from recruitment import Recruitment
from selenium import webdriver

start_url = 'https://programmers.co.kr/job'
base_url = 'https://programmers.co.kr'

def run():
    driver = connect()
    driver.get(start_url)
    while True:
        crawled_data = []
        WebDriverWait(driver, 10).until(
            EC.presence_of_element_located((By.ID, "tab_position"))
        )
        html = driver.page_source
        soup = BeautifulSoup(html,'html.parser')

        posts = soup.select('#list-positions-wrapper > ul > li > div.item-body')
        for post in posts:

            post_main = post.select('h4 > a')[0]
            post_url = base_url + post_main.get('href')
            post_title = post_main.text

            company_name = post.select('h5')[0].text
            post_date = ''
            post_newbie = programmers_newbie(post.select('ul.company-info > li.experience')[0].text)

            tmp_driver = connect()
            tmp_driver.get(post_url)
            tmp_html = tmp_driver.page_source
            tmp_driver.quit()

            tmp_soup = BeautifulSoup(tmp_html,'html.parser')
            post_contents = []
            for section in tmp_soup.select('body > div.main > div.position-show > div > div > div.content-body.col-item.col-xs-12.col-sm-12.col-md-12.col-lg-8 > section'):
                if section.get('class')[0] in set(['section-stacks','section-position','section-requirements','section-preference','section-description']):
                    post_contents.append(re.sub('[\s]+', ' ', section.text))

            tmp_post = Recruitment(
                title=post_title,
                url = post_url,
                company = company_name,
                start_date = '',
                level = post_newbie,
                job='',
                contents=post_contents
            )
            #print(tmp_post.get(True))
            #tmp_post.run()

        _next = soup.select('#paginate > nav > ul > li.next.next_page.page-item > a')[-1].get('href')
        if _next == '#':
            break
        else:
            driver.get('https://programmers.co.kr'+_next)
'''
def body_text(driver,json):
    print(json)
    driver.get(json['url'])
    WebDriverWait(driver, 10).until(
        EC.presence_of_element_located((By.CSS_SELECTOR, "body > div.main > div.position-show > div > div"))
    )
    html = driver.page_source
    soup = BeautifulSoup(html,'html.parser')
    for section in soup.select('body > div.main > div.position-show > div > div > div.content-body.col-item.col-xs-12.col-sm-12.col-md-12.col-lg-8 > section'):
        if section.get('class')[0] in set(['section-stacks','section-position','section-requirements','section-preference','section-description']):
            json['contents'] += section.text
    return json

def run(driver_path=None):
    print('start crawling : programmers')
    json_list, driver = programmers(driver_path)
    print('The number of programmers post : %d'%(len(json_list)))
    for i, json in enumerate(json_list):
        if i%10 == 0:
            print('programmers post : %d'%(i))
        json = body_text(driver,json)

    driver.quit()
    print('finish crawling : programmers')
    return json_list
'''

if __name__ == "__main__":
    a = run('/Users/mingihong/chromedriver')
