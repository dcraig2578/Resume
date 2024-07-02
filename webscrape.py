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
# from bs4 import BeauitfulSoup 
wait=WebDriverWait(driver,20)
# # //div[@class='s-item__image-wrapper'])[2] //*[@id="gh-ac"]
# <input type="text" class="gh-tb ui-autocomplete-input" aria-autocomplete="list" aria-expanded="false" size="50" maxlength="300" 
# aria-label="Search for anything" placeholder="Search for anything" id="gh-ac" name="_nkw" autocapitalize="off" autocorrect="off"
#  spellcheck="false" autocomplete="off" aria-haspopup="true" role="combobox" aria-owns="ui-id-1">
search = wait.until(EC.element_to_be_clickable((By.ID, 'gh-ac')))
search.send_keys(search_this)
search.send_keys(Keys.ENTER)
# listings_arr = []
# prices_arr = []
listings = driver.find_elements(By.CLASS_NAME, "s-item__title")
prices = driver.find_elements(By.CLASS_NAME, "s-item__price")
listings_arr = [item.text for item in listings]
prices_arr = [item.text for item in prices]
print(', '.join(listings_arr))
print(', '.join(prices_arr))
# with open('scrapedata.txt', 'w') as f:
#     for item in listings_arr:
#         f.write(f"{listings_arr}\n")
# with open('scrapedata.txt', 'w') as f:
#     for item in prices_arr:
#         f.write(f"{prices_arr}\n")

# for item in listings:
#     print(item.text)
#     # listings_arr.append(listings)
# print(listings_arr)
# prices = driver.find_elements(By.CLASS_NAME, "s-item__price")
# for item in prices:
#     print(item.text)
# send_keys(Keys.ARROW_DOWN * 2)
# test = wait.until(EC.element_to_be_clickable((By.XPATH,"(//div[@class='s-item__title']")))
# test = wait.until(EC.element_to_be_clickable((By.ID, 'rtm_html_278')))
# poopsy.append(test)
# print(poopsy)
# test = driver.find_element(By.XPATH, "(//div[@class='s-item__image-wrapper'])")
# test.click()
time.sleep(3)
driver.quit()

with open("ebay_listings.csv", "w", newline="", encoding="utf-8") as file:
    writer=csv.writer(file)
    writer.writerow(["Title","Price"])
    for title,price in zip(listings_arr,prices_arr):
        writer.writerow([title,price])

# print("Titles and Prices saved to ebay_listings.csv")