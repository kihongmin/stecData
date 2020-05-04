import selenium
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.common.by import By
from bs4 import BeautifulSoup

from .. import connect
from .. import transfrom_date
from ..recruitment import Recruitment


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
            post_newbie = post.select('td')[1].text

            tmp_post = Recruitment(
                title=post_title,
                url=post_url,
                company='nexon',
                start_date=post_date,
                level=post_newbie,
                job='')

            tmp_post.run()

        try:
            driver.find_element_by_css_selector('#con_right > div.content > div > a.next').click()
        except: # 무슨 종류인지 분간 필요.
            break