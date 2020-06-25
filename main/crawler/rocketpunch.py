import re

from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.common.by import By
from bs4 import BeautifulSoup

from . import connect
from ..es.recruitment import Recruitment
from ..es.level import Level
from ..es.start_date import StartDate


start_url = 'https://www.rocketpunch.com/jobs?page=1'


def run(is_load_all = False):   #이전 데이터 전부다 가져오나
    driver = connect()
    driver.get(start_url)

    while True:
        WebDriverWait(driver, 10).until(
            EC.presence_of_element_located((By.ID, "company-list"))
        )
        html = driver.page_source
        soup = BeautifulSoup(html,'html.parser')
        company_list = soup.select('#company-list > div.company.item')

        for company in company_list:
            company_name = company.select('div.content > div.company-name > a > h4 > strong')[0].text
            post_list = company.select('div.content > div.company-jobs-detail > div.job-detail')
            for post in post_list:
                post_date, is_posted_yesterday = StartDate.transform(
                    date=post.select('div.job-dates > span')[-1].text,
                    source='rocketpunch')
                if not is_load_all and not is_posted_yesterday : #어제꺼만 가져오는데 어제꺼 아니면 continue
                    continue
                
                post_main = post.select('div > a.nowrap.job-title.primary.link')[0]
                post_title = post_main.text

                post_url = 'https://www.rocketpunch.com' + post_main.get('href')
                levels = []
                for level in post.select('div > span.job-stat-info')[0].text.replace(',',' ').split():
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
                WebDriverWait(tmp_driver, 10).until(
                    EC.presence_of_element_located((By.CSS_SELECTOR, "#wrap > div.eight.wide.job-content.column > section.row > h4"))
                )
                tmp_html = tmp_driver.page_source
                tmp_driver.quit()

                tmp_soup = BeautifulSoup(tmp_html,'html.parser')
                post_contents = []
                for section in tmp_soup.select('#wrap > div.eight.wide.job-content.column > section.row > h4'):
                    if section.text in set(["주요 업무", "업무 관련 기술 / 활동 분야", "채용 상세"]):
                        post_contents.append(section.parent.text.replace('\n',' '))

                tmp_post = Recruitment(
                    title = post_title,
                    url = post_url,
                    company = company_name,
                    start_date = post_date,
                    level = levels,
                    job = None,
                    contents = post_contents
                )
                tmp_post.run()

        _next = driver.find_elements_by_css_selector('#search-results > div.ui.blank.right.floated.segment > div.ui.pagination.menu > a')[-1].get_attribute('href')
        if _next:
            driver.get(_next)
        else:
            driver.quit()
            return
