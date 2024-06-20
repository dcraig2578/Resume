import time
from selenium import webdriver
driver=webdriver.Chrome()
driver.get("https://www.ebay.com/")
from selenium.webdriver.common.by import By
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.chrome.service import Service as ChromeService
from webdriver_manager.chrome import ChromeDriverManager
from selenium.common.exceptions import WebDriverException
time.sleep(1)
motors_link = driver.find_element(By.LINK_TEXT, "Motors")
motors_link.click()
driver.execute_script("window.scrollTo(0, 900)")
time.sleep(1)
cars_link = driver.find_element(By.LINK_TEXT, "Engines & Engine Parts")
cars_link.click()
driver.execute_script("window.scrollTo(0, 450)")
time.sleep(1)
engines_link = driver.find_element(By.LINK_TEXT, "Engines")
engines_link.click()
driver.execute_script("window.scrollTo(0, 900)")
select = driver.find_element(By.XPATH, "(//div[@class='s-item__image-wrapper'])[2]")
select.click()
window_after = driver.window_handles[1]
driver.switch_to.window(window_after)
driver.execute_script("window.scrollTo(0, 900)")
add_cart = driver.find_element(By.LINK_TEXT, 'Add to cart')
add_cart.click()
proceed_link = driver.find_element(By.LINK_TEXT, 'Go to cart')
proceed_link.click()
time.sleep(10)
driver.quit()