from selenium import webdriver
from selenium.webdriver.common.by import By
from time import sleep
import os
import urllib
#获取浏览器对象
def getBrowser(url):
    #设置浏览器参数
    options=webdriver.ChromeOptions()
    #开发者模式
    options.add_experimental_option('excludeSwitches', ['enable-automation'])
    #去掉自动化测试提示
    options.add_experimental_option('useAutomationExtension', False)
    #装载参数
    browser = webdriver.Chrome(options=options)
    #修改识别标志，防检测
    browser.execute_cdp_cmd(
            'Page.addScriptToEvaluateOnNewDocument',
            {'source':'Object.defineProperty(navigator,"webdriver",{get:()=>undefined})'}
            )
    #启动浏览器
    browser.get(url)
    #等待加载
    sleep(3)
    #下移加载评论框
    browser.execute_script("var q=document.documentElement.scrollTop=500")    
    sleep(2)
    print("浏览器页面打开完成")
    return browser

def getEmoji(browser,addr=r'./hoyo_emoji'):
    itemList=browser.find_elements(By.CSS_SELECTOR,'[class="mhy-emoticon__item"]')
    sleep(1)
    #emoji地址
    emojiAddrList=[]
    #emoji标题
    emojiNameList=[]
    for item in itemList:
        emojiNameList.append(item.get_attribute('title'))
        temp=item.get_attribute("style")
        idxChar=temp.index('("')+1
        emojiAddrList.append(temp[idxChar+1:len(temp)-3])
    #创建子目录
    addr=addr+'/'+emojiNameList[0]
    try:
        if addr not in os.listdir():
            os.mkdir(addr)    # 创建文件夹
    except:
        pass
    download(emojiAddrList,emojiNameList,addr)

def download(emojiAddrList=[],emojiNameList=[],addr=''):
    for idx in range(len(emojiAddrList)):
        #读取emoji数据
        img=urllib.request.urlopen(emojiAddrList[idx]).read()
        #创建或打开文件
        f=open(addr+'/'+emojiNameList[idx]+'.png','wb')
        #写入
        f.write(img)
        #关闭
        f.close()
        
def main():
    #目的地址
    browser=getBrowser('https://www.miyoushe.com/bh3/article/10465564')
    #存储目录地址
    addr=r'./hoyo_emoji'
    #点击表情图标
    browser.find_element(By.XPATH,'//*[@id="reply"]/div[1]/div[2]/div[2]/div[1]/i[1]').click()
    
    #创建目录
    try:
        if addr not in os.listdir():
            os.mkdir(addr)    # 创建文件夹
    except:
        pass
    
    while(1):
        root=browser.find_elements(By.CSS_SELECTOR,'[class~="mhy-emoticon__set"]')
        for item in root:
            item.click()
            getEmoji(browser,addr)
        
        #检测是否还有表情包
        try:
            #没有则退出
            browser.find_element(By.CSS_SELECTOR,'[class="mhy-emoticon__next mhy-emoticon__pagerbtn--disabled"]')
            break
        except:
            #有则进行下一轮
            browser.find_element(By.CSS_SELECTOR,'[class="mhy-icon iconfont icon-xiaojiantou"]').click()
            
    print('完成')
        
if __name__=='__main__':
    main()