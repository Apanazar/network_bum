from selenium.webdriver import FirefoxOptions
from selenium import webdriver

options = FirefoxOptions()
options.add_argument("--headless")

driver = webdriver.Firefox(options=options, executable_path=r'./geckodriver')
driver.set_page_load_timeout(15)

file1 = open("result.txt", "r")

num = 0
while True:
    line = file1.readline()
    num += 1

    if not line:
        break
    
    print("URL: \t", line.strip())
    screen_name = "./screen/" + str(line.strip()) + ".png"
    print(screen_name)
    
    try:
        driver.get("http://"+line.strip())
        driver.save_screenshot(screen_name)
    except:
        continue

file1.close()
driver.close()
driver.quit()