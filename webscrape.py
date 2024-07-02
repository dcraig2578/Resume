print("What are you searching for?")
search_this = input()
import time
import csv
from selenium import webdriver
driver=webdriver.Chrome()
driver.get("https://www.ebay.com/")
from selenium.webdriver.common.by import By
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.chrome.service import Service as ChromeService
from webdriver_manager.chrome import ChromeDriverManager
from selenium.common.exceptions import WebDriverException
from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
wait=WebDriverWait(driver,20)
search = wait.until(EC.element_to_be_clickable((By.ID, 'gh-ac')))
search.send_keys(search_this)
search.send_keys(Keys.ENTER)
listings = driver.find_elements(By.CLASS_NAME, "s-item__title")
prices = driver.find_elements(By.CLASS_NAME, "s-item__price")
listings_arr = [item.text for item in listings]
prices_arr = [item.text for item in prices]
time.sleep(3)
driver.quit()
with open("ebay_listings.csv", "w", newline="", encoding="utf-8") as file:
    writer=csv.writer(file)
    writer.writerow(["Title","Price"])
    for title,price in zip(listings_arr,prices_arr):
        writer.writerow([title,price])