import re

import selenium
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.common.by import By
from bs4 import BeautifulSoup

from . import connect
from . import transfrom_date
from ..recruitment import Recruitment
from ..codes import Level


base_url = 'https://career.nexon.com'
start_url = base_url + '/user/recruit/notice/noticeList'


def run():
    driver = connect()
    driver.get(start_url)
    WebDriverWait(driver, 10).until(
        EC.presence_of_element_located((By.CSS_SELECTOR, "#container > ul > li:nth-child(1) > a > img"))
    )
    driver.find_element_by_css_selector('#container > ul > li:nth-child(1) > a > img').click()
    while True: # 굳이 무한 루프 위험이 있는 while문을 쓰는 이유는?
        WebDriverWait(driver, 10).until(
            EC.presence_of_element_located((By.ID, "con_right"))
        )
        html = driver.page_source
        soup = BeautifulSoup(html,'html.parser')

        posts = soup.select('#con_right > div.content > table > tbody > tr')
        for post in posts:
            post_date = transfrom_date(post.select('td')[5].text)
            if not post_date:
                continue
            post_title = post.select('td.tleft.fc_02 > a > span')[0].text
            post_url = base_url+post.select('td.tleft.fc_02 > a')[0].get('href')
            # 이거 형태가 어케되는지, code로 변환하는 코드를 원래 코드에서 따로 못 찾겠던데
            post_newbie = post.select('td')[1].text

            tmp_driver = connect()
            tmp_driver.get(post_url)
            tmp_html = tmp_driver.page_source
            tmp_driver.quit()

            tmp_soup = BeautifulSoup(tmp_html,'html.parser')
            post_contents = tmp_soup.select('#con_right > div.content > div.list_txt01')[0].text
            post_contents = re.sub('[\s]+', ' ', post_contents)

            tmp_post = Recruitment(
                title=post_title,
                url=post_url,
                company='nexon',
                start_date=post_date,
                level=post_newbie,
                job='',
                contents=post_contents)
            ##########################################
            # 여기까지 정상작동하는지 확인해줄 것!     #
            # 작동 안하면 크롤링이 잘되도록 수정 부탁! #
            ##########################################
            tmp_post.run()

        try:
            driver.find_element_by_css_selector('#con_right > div.content > div > a.next').click()
        except: # 무슨 종류인지 분간 필요.
            break

    driver.quit()