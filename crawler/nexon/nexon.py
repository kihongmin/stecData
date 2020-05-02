import re, os, sys
import selenium
sys.path.append(os.path.dirname(os.path.abspath(os.path.dirname(__file__))))
from util import *
from selenium import webdriver
from bs4 import BeautifulSoup
import warnings
warnings.filterwarnings(action='ignore')


def nexon(driver_path=None):
    if not driver_path:
        driver_path='/Users/mingihong/chromedriver'
    driver = headless(driver_path)
#    driver = webdriver.Chrome(driver_path)
    driver.get('https://career.nexon.com/user/recruit/notice/noticeList')
    WebDriverWait(driver, 10).until(
        EC.presence_of_element_located((By.CSS_SELECTOR, "#container > ul > li:nth-child(1) > a > img"))
    )
    driver.find_element_by_css_selector('#container > ul > li:nth-child(1) > a > img').click()
    final_data = []
    while True:
        crawled_data = []
        WebDriverWait(driver, 10).until(
            EC.presence_of_element_located((By.ID, "con_right"))
        )
        html = driver.page_source
        soup = BeautifulSoup(html,'html.parser')

        posts = soup.select('#con_right > div.content > table > tbody > tr')
        for post in posts:
            post_date = transfrom_date(post.select('td')[5].text)
            if not post_date:
                continue
            post_title = post.select('td.tleft.fc_02 > a > span')[0].text
            post_url = 'https://career.nexon.com'+post.select('td.tleft.fc_02 > a')[0].get('href')
            post_newbie = post.select('td')[1].text

            crawled_data.append(
                job(post_title,post_url,'nexon',post_date,post_newbie).data
            )

        final_data.extend(crawled_data)

        try:
            driver.find_element_by_css_selector('#con_right > div.content > div > a.next').click()
        except:
            break


    return final_data, driver

def body_text(driver,json):
    driver.get(json['url'])

    html = driver.page_source
    soup = BeautifulSoup(html,'html.parser')
    txt = re.sub('[\s]+',' ',soup.select('#con_right > div.content > div.list_txt01')[0].text)
    json['contents'] += txt
    return json

def run(driver_path=None):
    print('start crawling : nexon')
    json_list, driver = nexon(driver_path)
    print('The number of nexon post : %d'%(len(json_list)))
    for i,json in enumerate(json_list):
        if i%10 == 0:
            print('nexon post : %d'%(i))
        jojsonb = body_text(driver,json)
    driver.quit()
    print('finish crawling : nexon')

    return json_list


if __name__ == "__main__":
    run()
