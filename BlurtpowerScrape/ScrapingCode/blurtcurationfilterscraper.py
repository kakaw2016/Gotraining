import pandas as pd
import json
import re
from bs4 import BeautifulSoup
from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.by import By
import sys
import time


stdoutOrigin=sys.stdout 
sys.stdout = open("/home/ybf/Downloads/AUTOMATION/Beautifulsoupscraper/scrapeblurtTTTT", "w")

op = webdriver.ChromeOptions()
op.add_argument('--headless')
ser = Service("/home/ybf/Downloads/AUTOMATION/BlurtconnectCuration/BlurtScrapper/PowerdownScrape/chromedriver")
driver = webdriver.Chrome(service=ser, options=op)

urls = ['https://blurt.blog/instablurt/@panoramax/very-much-on-the-nails', 'https://blurt.blog/r2cornell/@nazirhussain/happy-independence-day-to-you-all', 'https://blurt.blog/r2cornell/@tanweeralam/75th-independence-day', 'https://blurt.blog/riddleschallenge/@nazirhussain/weekly-riddles-challenge-18', 'https://blurt.blog/techclub/@techclub/daily-curation-report-59-for-techclub-community']
for url in urls:
    driver.get(url)
    driver.implicitly_wait(4) #maximum time to load the link
    driver.execute_script("window.scrollTo(0, document.body.scrollHeight,)")

    time.sleep(5)

    names = []
    trs = driver.find_elements(By.XPATH, "/html/body/div[1]")
    for elements in trs:
        names.append(elements.text)
        print(names)
driver.close()
driver.quit()
sys.stdout.close()
sys.stdout=stdoutOrigin