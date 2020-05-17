import re

from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.common.by import By
from bs4 import BeautifulSoup

from . import connect
from ..es.recruitment import Recruitment
from ..es.level import Level
from ..es.start_date import StartDate


start_url = 'https://careers.kakao.com/jobs?page=1&company=ALL&keyword=&part=TECHNOLOGY&skilset='
def get_pageurl(i):
    return 'https://careers.kakao.com/jobs?page='+str(i)+'&company=ALL&keyword=&part=TECHNOLOGY&skilset='
def run():
    driver = connect()
    driver.get(start_url)
    next_list = driver.find_elements_by_css_selector('#mArticle > div > div.paging_list > span > a.change_page.btn_lst')
    total_page = re.sub('[^0-9]','',next_list[-1].get_attribute('href'))
    for i in range(2,int(total_page)+1):
        WebDriverWait(driver, 10).until(
            EC.presence_of_element_located((By.CSS_SELECTOR, "#mArticle > div > ul.list_notice > li"))
        )
        html = driver.page_source
        soup = BeautifulSoup(html,'html.parser')

        posts = soup.select("#mArticle > div > ul.list_notice > li")
        for post in posts:
            #kakao post_date 없음
            #post_date = StartDate.transform(
            #    date=post.select('div.cw_jopinfo > a > span.cw_info > span.cw_range')[0].text)
            post_date = None
            post_title = post.select('div > div > div > a > span')[0].text
            post_url = 'https://careers.kakao.com'+post.select('div > div > div > a')[0].get('href')
            post_newbie = Level.text2code(
                text_list=post_title)

            tmp_driver = connect()
            tmp_driver.get(post_url)
            tmp_html = tmp_driver.page_source
            tmp_driver.quit()

            soup = BeautifulSoup(tmp_html,'html.parser')
            post_contents = []
            txt = soup.select('#mArticle > div > div.board_view > div.cont_board > div')

            if txt:
                post_contents.append(re.sub('[\s]+|\\u200b', ' ', txt[0].text))
            tmp_post = Recruitment(
                title=post_title,
                url = post_url,
                company = 'kakao',
                start_date = post_date,
                level = post_newbie,
                job=None,
                contents=post_contents
            )            
            tmp_post.run()

        driver.get(get_pageurl(i))
