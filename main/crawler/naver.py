
import re

from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.common.by import By
from bs4 import BeautifulSoup

from . import connect
from ..es.recruitment import Recruitment
from ..es.level import Level
from ..es.start_date import StartDate


start_url = 'https://recruit.navercorp.com/naver/job/list/developer'


def run(driver_path=None):
    #가끔 아예 못 가져오는 경우 있음->총 3번 시도
    for t in range(3):
        driver = connect()

        for i in range(15):
            driver.get(start_url)
            WebDriverWait(driver, 10).until(
                EC.presence_of_element_located((By.CSS_SELECTOR, "#jobListDiv"))
            )
            try:
                driver.find_element_by_xpath('//*[@id="moreDiv"]/button').click()
                #왜인지 모르겠는데 아래 요청 보내면 post 목록 어느정도 정확하게 가져옴
                driver.current_url
                driver.implicitly_wait(5)
            except:
                driver.current_url
                driver.implicitly_wait(5)
                continue
        html = driver.page_source
        soup = BeautifulSoup(html,'html.parser')
        posts = soup.select('#jobListDiv > ul > li')
        if len(posts) != 0:
            break
        print('fail to load naver... try again')
        driver.quit()        

    for post in posts:
        post_date = StartDate.transform(post.select('a > span > em')[0].text)
        post_url = 'https://recruit.navercorp.com'+post.select('a')[0].get('href')
        post_title = post.select('a > span > strong')[0].text
        post_newbie = Level.text2code(
            text_list = post_title)

        tmp_driver = connect()
        tmp_driver.get(post_url)
        html = driver.page_source
        soup = BeautifulSoup(html,'html.parser')
        txt = soup.select('#content > div > div.career_detail > div.dtl_context > div.context_area')
        if txt:
            post_contents = [re.sub('\n|\xa0','',txt[0].text)]
        else:
            post_contents = []
        tmp_post = Recruitment(
            title=post_title,
            url = post_url,
            company = 'naver',
            start_date = post_date,
            level = post_newbie,
            job=None,
            contents=post_contents
        )

        tmp_post.run()
