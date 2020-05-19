
import re,time

from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.common.by import By
from bs4 import BeautifulSoup

from . import connect
from ..es.recruitment import Recruitment
from ..es.level import Level
from ..es.start_date import StartDate


main_url = 'https://recruit.navercorp.com/naver/recruitMain'
start_url = 'https://recruit.navercorp.com/naver/job/list/developer'

def run():
    driver = connect()
    driver.get(main_url)

    n_posts = int(driver.find_element_by_css_selector('#content > div.spot > div.recruit_post > ul > li.fst > a > span > strong').text)

    driver.get(start_url)
    WebDriverWait(driver, 10).until(
        EC.presence_of_element_located((By.CSS_SELECTOR, '#moreDiv > button'))
    )
    for i in range(round(n_posts/10)):
        driver.find_element_by_css_selector('#moreDiv > button').click()
        time.sleep(5)
    html = driver.page_source
    soup = BeautifulSoup(html,'html.parser')
    posts = soup.select('#jobListDiv > ul > li')

    print("The number of crawled naver data : ",len(posts))
    for post in posts:
        post_date = StartDate.transform(post.select('a > span > em')[0].text)
        post_url = 'https://recruit.navercorp.com'+post.select('a')[0].get('href')
        post_title = post.select('a > span > strong')[0].text
        post_newbie = Level.text2code(
            text_list = post_title)

        tmp_driver = connect()
        tmp_driver.get(post_url)
        html = tmp_driver.page_source

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
