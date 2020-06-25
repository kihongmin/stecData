import re

from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.common.by import By
from selenium.common.exceptions import NoSuchElementException
from bs4 import BeautifulSoup

from . import connect
from ..es.start_date import StartDate
from ..es.recruitment import Recruitment
from ..es.level import Level


base_url = 'https://career.nexon.com'


def run(is_load_all = False):   #이전 데이터 전부다 가져오나
    driver = connect()
    driver.get(base_url)
    WebDriverWait(driver, 10).until(
        EC.presence_of_element_located((By.CSS_SELECTOR, "#container > ul > li:nth-child(1) > a > img"))
    )
    #홈페이지 접속 후 채용 공고 페이지로 이동하기 위해 클릭
    driver.find_element_by_css_selector('#container > ul > li:nth-child(1) > a > img').click()
    while True:
        WebDriverWait(driver, 10).until(
            EC.presence_of_element_located((By.ID, "con_right"))
        )
        html = driver.page_source
        soup = BeautifulSoup(html,'html.parser')
        #채용공고 페이지
        posts = soup.select('#con_right > div.content > table > tbody > tr')
        #채용 공고 페이지의 채용 공고들
        for post in posts:
            post_date, is_posted_yesterday = StartDate.transform(post.select('td')[5].text)
            if not is_load_all and not is_posted_yesterday : #어제꺼만 가져오는데 어제꺼 아니면 continue
                continue
            post_title = post.select('td.tleft.fc_02 > a > span')[0].text
            post_url = base_url+post.select('td.tleft.fc_02 > a')[0].get('href')

            levels = []
            for level in post.select('td')[1].text.split('/'):
                levels.append(
                    Level.string2code(
                        text=level
                    )
                )
            levels = [level for level in levels if level is not None]
            levels = sorted(levels)
            if len(levels) < 1:
                levels = [Level.newbie, Level.unlimited]

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
                level=levels,
                job=None,
                contents=[post_contents])
            tmp_post.run()

        try:
            driver.find_element_by_css_selector('#con_right > div.content > div > a.next').click()
        except NoSuchElementException:
            break

    driver.quit()
