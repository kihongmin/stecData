import re, os, sys
import selenium
sys.path.append(os.path.dirname(os.path.abspath(os.path.dirname(__file__))))
from util import *
from selenium import webdriver
from bs4 import BeautifulSoup
import warnings
warnings.filterwarnings(action='ignore')


def programmers(driver_path=None):
    if not driver_path:
        driver_path='/Users/mingihong/chromedriver'
    driver = headless(driver_path)
#    driver = webdriver.Chrome(driver_path)
    driver.get('https://programmers.co.kr/job')
    final_data = []
    while True:
        crawled_data = []
        WebDriverWait(driver, 10).until(
            EC.presence_of_element_located((By.ID, "tab_position"))
        )
        html = driver.page_source
        soup = BeautifulSoup(html,'html.parser')

        posts = soup.select('#list-positions-wrapper > ul > li > div.item-body')
        for post in posts:
            post_main = post.select('h4 > a')[0]
            post_url = 'https://programmers.co.kr'+post_main.get('href')
            post_title = post_main.text
            company_name = post.select('h5')[0].text
            post_date = ''
            post_newbie = programmers_newbie(post.select('ul.company-info > li.experience')[0].text)
            crawled_data.append(
                job(post_title,post_url,company_name,post_date,post_newbie)
            )
        final_data.extend(crawled_data)
        _next = soup.select('#paginate > nav > ul > li.next.next_page.page-item > a')[-1].get('href')
        if _next == '#':
            break
        else:
            driver.get('https://programmers.co.kr'+_next)

    return final_data, driver

def body_text(driver,job):
    driver.get(job.data['url'])
    WebDriverWait(driver, 10).until(
        EC.presence_of_element_located((By.CSS_SELECTOR, "body > div.main > div.position-show > div > div"))
    )
    html = driver.page_source
    soup = BeautifulSoup(html,'html.parser')
    for section in soup.select('body > div.main > div.position-show > div > div > div.content-body.col-item.col-xs-12.col-sm-12.col-md-12.col-lg-8 > section'):
        if section.get('class')[0] in set(['section-stacks','section-position','section-requirements','section-preference','section-description']):
            job.set_contents(section.text)
    return job

def run(driver_path=None):
    job_list, driver = programmers(driver_path)
    for job in job_list:
        job = body_text(driver,job)
    driver.quit()

    return job_list


if __name__ == "__main__":
    a = run()
