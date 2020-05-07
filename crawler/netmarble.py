import re,os,sys,json

import selenium
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.common.by import By
from bs4 import BeautifulSoup

from . import connect, transfrom_date, make_newbie, programmers_newbie
from recruitment import Recruitment
from selenium import webdriver

start_url = 'https://m.netmarble.com/rem/www/noticelist.jsp'

def run():
    driver = connect()
    driver.get(start_url)

    while True:
        WebDriverWait(driver, 10).until(
            EC.presence_of_element_located((By.CSS_SELECTOR, "#contents > div > div > div > div.recruit_list_wrapper > ul"))
        )
        html = driver.page_source
        soup = BeautifulSoup(html,'html.parser')

        posts = soup.select('#contents > div > div > div > div.recruit_list_wrapper > ul > li')
        for post in posts:

            post_date = transfrom_date(post.select('div.cw_jopinfo > a > span.cw_info > span.cw_range')[0].text)
            #if not post_date:
            #    continue
            post_title = post.select('div.cw_jopinfo > a > span.cw_title')[0].text
            post_url = 'https://m.netmarble.com/rem/www'+post.select('div.cw_jopinfo > a')[0].get('href')[1:]
            post_newbie = make_newbie(post.select('div.cw_jopinfo > a > span.cw_info > span.cw_type')[0].text)

            tmp_driver = connect()
            tmp_driver.get(post_url)
            tmp_html = tmp_driver.page_source
            tmp_driver.quit()

            soup = BeautifulSoup(tmp_html,'html.parser')
            post_contents = []
            txt = soup.select('#tmpCapture > div > table')
            if txt:
                post_contents.append(re.sub('[\s]+', ' ', txt[0].text))

            tmp_post = Recruitment(
                title=post_title,
                url = post_url,
                company = 'netmarble',
                start_date = '',
                level = post_newbie,
                job='',
                contents=post_contents
            )
            tmp_post.run()
            #print(tmp_post.get(True))


        _next = driver.find_element_by_css_selector('#contents > div > div > div > div.recruit_list_wrapper > div.recruit_pagination > button.page_next')
        if not _next.get_attribute('disabled'):
            _next.click()
        else:
            break
