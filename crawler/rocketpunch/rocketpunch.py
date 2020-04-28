import re, os, sys
import selenium
sys.path.append(os.path.dirname(os.path.abspath(os.path.dirname(__file__))))
from util import *
from selenium import webdriver
from bs4 import BeautifulSoup
import warnings
warnings.filterwarnings(action='ignore')

def make_newbie(newbie_list):
    newbie = {'인턴':10,'신입':50,'경력':90}
    newbie_list = list(filter(lambda x:x in newbie,newbie_list))
    return list(map(lambda x:newbie[x],newbie_list))

def rocketpunch():
    del_ws = re.compile(r'\s+')
    driver = headless('/Users/mingihong/chromedriver')
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
        for company in company_list:
            company_name = company.select('div.content > div.company-name > a > h4 > strong')[0].text
            post_list = company.select('div.content > div.company-jobs-detail > div.job-detail')
            for post in post_list:
                post_date = transfrom_date(post.select('div.job-dates > span')[-1].text)
                if not post_date:
                    continue
                post_main = post.select('div > a.nowrap.job-title.primary.link')[0]
                post_title = post_main.text
                post_url = post_main.get('href')
                post_newbie = make_newbie(post.select('div > span.job-stat-info')[0].text.replace(',',' ').split())
                crawled_data.append({
                    'title':post_title,
                    'url' : 'https://www.rocketpunch.com'+post_url,
                    'company' : company_name,
                    'start_date' : post_date,
                    'newbie' : post_newbie
                })
        final_data.extend(crawled_data)

        _next = driver.find_elements_by_css_selector('#search-results > div.ui.blank.right.floated.segment > div.ui.pagination.menu > a')[-1].get_attribute('href')
        if _next:
            driver.get(_next)
        else:
            break
        print(final_data)
    return final_data

if __name__ == "__main__":
    rocketpunch()
