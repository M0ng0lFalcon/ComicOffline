from funs import *

# Global Variables
NAME = 'ZhouShuHuiZhan'
BASE_URL = 'https://www.cartoonmad.com/comic/7654.html'

# https://www.cartoonmad.com/5e568/4975/001/001.jpg
IMG_URL_BASE = 'https://www.cartoonmad.com/'
FIRST_PARAM = '75652/'
SECOND_PARAM = '7654/'
IMG_URL_UNIT = (IMG_URL_BASE, FIRST_PARAM, SECOND_PARAM)


startTime = time.time()


tableRes = GetTableRes(BASE_URL)

chapterNum = GetChaptersNum(tableRes)
pageNums = GetEachPageNum(tableRes)
urls = GetUrls(chapterNum, pageNums, IMG_URL_UNIT)

linkUnits = MakeLinkUnits(urls, NAME)


MakeCommicDirs(NAME, chapterNum, pageNums)


loop = asyncio.get_event_loop()
loop.run_until_complete(DownloadMain(loop, linkUnits))
endTime = time.time()

print(f'[!] Running time: {endTime - startTime} s')
