import os

from selenium import webdriver


_chrome_options_ = ['--window-size=1920,1080', '--headless', 'disable-gpu']
_driver_path_ = os.path.abspath('./assets/chromedriver')


def connect(driver_path=_driver_path_, args=_chrome_options_):
    options = webdriver.ChromeOptions()
    if args is not None:
        for arg in args:
            options.add_argument(arg)
    driver = webdriver.Chrome(driver_path, chrome_options=options)
    return driver