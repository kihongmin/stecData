import re

from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.common.by import By
from bs4 import BeautifulSoup

from . import connect
from ..es.recruitment import Recruitment
from ..es.level import Level
from ..es.start_date import StartDate


start_url = 'https://m.netmarble.com/rem/www/noticelist.jsp'


def run(is_load_all = False):   #이전 데이터 전부다 가져오나
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

            post_date, is_posted_yesterday = StartDate.transform(
                date=post.select('div.cw_jopinfo > a > span.cw_info > span.cw_range')[0].text)
            if not is_load_all and not is_posted_yesterday : #어제꺼만 가져오는데 어제꺼 아니면 continue
                continue
            post_title = post.select('div.cw_jopinfo > a > span.cw_title')[0].text
            post_url = 'https://m.netmarble.com/rem/www'+post.select('div.cw_jopinfo > a')[0].get('href')[1:]
            post_newbie = Level.string2code(
                text=post.select('div.cw_jopinfo > a > span.cw_info > span.cw_type')[0].text)

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
                start_date = post_date,
                level = post_newbie,
                job=None,
                contents=post_contents
            )
            tmp_post.run()

        _next = driver.find_element_by_css_selector('#contents > div > div > div > div.recruit_list_wrapper > div.recruit_pagination > button.page_next')
        if not _next.get_attribute('disabled'):
            _next.click()
        else:
            break
