import re, os, sys
import selenium
sys.path.append(os.path.dirname(os.path.abspath(os.path.dirname(__file__))))
from util import *
from selenium import webdriver
from bs4 import BeautifulSoup
import warnings
warnings.filterwarnings(action='ignore')


def rocketpunch(driver_path=None):
    if not driver_path:
        driver_path='chromedriver'
    driver = headless(driver_path)
#    driver = webdriver.Chrome(driver_path)
    driver.get('https://www.rocketpunch.com/jobs?page=1')
    final_data = []
    while True:
        crawled_data = []
        WebDriverWait(driver, 10).until(
            EC.presence_of_element_located((By.ID, "company-list"))
        )
        html = driver.page_source
        soup = BeautifulSoup(html,'html.parser')

        company_list = soup.select('#company-list > div.company.item')
        backup_url = driver.current_url
        for company in company_list:
            company_name = company.select('div.content > div.company-name > a > h4 > strong')[0].text
            post_list = company.select('div.content > div.company-jobs-detail > div.job-detail')
            for post in post_list:
                post_date = transfrom_date(post.select('div.job-dates > span')[-1].text,True)
                if not post_date:
                    continue
                post_main = post.select('div > a.nowrap.job-title.primary.link')[0]
                post_title = post_main.text
                post_url = 'https://www.rocketpunch.com' + post_main.get('href')
                post_newbie = make_newbie(post.select('div > span.job-stat-info')[0].text.replace(',',' ').split())
                crawled_data.append(
                    job(post_title,post_url,company_name,post_date,post_newbie).data
                )
        final_data.extend(crawled_data)
        _next = driver.find_elements_by_css_selector('#search-results > div.ui.blank.right.floated.segment > div.ui.pagination.menu > a')[-1].get_attribute('href')
        if _next:
            driver.get(_next)
        else:
            break

    return final_data, driver

def body_text(driver,json):
    driver.get(json['url'])
    WebDriverWait(driver, 10).until(
        EC.presence_of_element_located((By.CSS_SELECTOR, "#wrap > div.eight.wide.job-content.column > section.row > h4"))
    )
    html = driver.page_source
    soup = BeautifulSoup(html,'html.parser')
    for section in soup.select('#wrap > div.eight.wide.job-content.column > section.row > h4'):
        if section.text in set(["주요 업무", "업무 관련 기술 / 활동 분야", "채용 상세"]):
            json['contents'] += section.parent.text.replace('\n',' ')

    return json

def run(driver_path=None):

    print('start crawling : rocketpunch')
    json_list, driver = rocketpunch(driver_path)
    print('The number of rocketpunch post : %d'%(len(json_list)))
    for i, json in enumerate(json_list):
        if i%10 == 0:
            print('rocketpunch post : %d'%(i))
        json = body_text(driver,json)
    driver.quit()
    print('finish crawling : rocketpunch')

    return json_list


if __name__ == "__main__":
    run()


    #rocketpunch()
    #driver = headless('/Users/mingihong/chromedriver')
    #t = job(0,'https://www.rocketpunch.com/jobs/43074/Trading-Business-Development-at-Tridge    ',0,0,0)
    #t = body_text(driver,t)
    #print(t.data)
