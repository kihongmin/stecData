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
    driver = connect()
    driver.get(start_url)

    final_data = []
    title_set = set()
    dict={'#entType001':'신입', '#entType002':'경력','#entType004':'인턴'}

    for button in ['#entType001','#entType002','#entType004']:
        driver.find_element_by_css_selector(button+' > a').click()
        newbie = dict[button]
        WebDriverWait(driver, 10).until(
            EC.presence_of_element_located((By.CSS_SELECTOR, "#jobListDiv > ul > li:nth-child(1) > a > span > strong"))
        )
        for i in range(10):
            # while문으로 click하면 랜덤하게 진행됨 -> 일단은 for문으로 진행
            if button == '#entType004':
                break
            try:
                driver.find_element_by_css_selector('#moreDiv > button').click()
                driver.implicitly_wait(5)
            except:
                break
        html = driver.page_source
        soup = BeautifulSoup(html,'html.parser')
        posts = soup.select('#jobListDiv > ul > li')
        print(len(posts))
        for post in posts:
            post_date = StartDate.transform(post.select('a > span > em')[0].text)
            post_url = 'https://recruit.navercorp.com'+post.select('a')[0].get('href')
            post_title = post.select('a > span > strong')[0].text
            post_newbie = Level.string2code(newbie)
            print(post_date,post_url,post_title,post_newbie)

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
                company = 'netmarble',
                start_date = post_date,
                level = post_newbie,
                job=None,
                contents=post_contents
            )
            tmp_post.run()


    return final_data, driver

def body_text(driver,json):
    driver.get(json['url'])
    WebDriverWait(driver, 10).until(
        EC.presence_of_element_located((By.CSS_SELECTOR, '#content > div > div.career_detail > div.dtl_context > div.context_area'))
    )
    html = driver.page_source
    soup = BeautifulSoup(html,'html.parser')
    txt = soup.select('#content > div > div.career_detail > div.dtl_context > div.context_area')
    if txt:
        json['contents'] = re.sub('\n|\xa0','',txt[0].text)
        print(json)
        return json
    else:
        return None



if __name__ == "__main__":
    run()
