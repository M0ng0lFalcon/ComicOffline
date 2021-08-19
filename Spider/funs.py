import requests
from bs4 import BeautifulSoup
import re
import os
import aiohttp
import asyncio
import time


def GetTableRes(BASE_URL):
    r = requests.get(BASE_URL)

    soup = BeautifulSoup(r.text, 'lxml')

    tableRes = soup.find('table', attrs={'width': '800', 'align': 'center'})

    return tableRes


def GetEachPageNum(tableRes):
    res = []
    numTag = tableRes.find_all(
        'font', attrs={'style': 'font-size:8pt;color: #888888;'})
    for i in numTag:
        numString = str(i.string)
        pageNum = ParsePageNum(numString)
        res.append(int(pageNum))

    return res


def GetChaptersNum(tableRes):
    chapters = tableRes.find_all('a', attrs={'target': '_blank'})

    return len(chapters)


def GetUrls(ChapterNum, PageNums, ImgUrlUnit):
    res = []
    for chapter in range(1, ChapterNum+1):
        li = []
        print('[*] Parsing Chapter', chapter)
        for page in range(1, PageNums[chapter-1]+1):
            url = MakeUrl(ImgUrlUnit[0], ImgUrlUnit[1],
                          ImgUrlUnit[2], chapter, page)
            print('[!] Adding URL:', url)

            li.append(url)

        res.append(li)
    return res


async def DownloadJob(session, linkUnit):
    print('[&] Downloading ', linkUnit[1])
    img = await session.get(linkUnit[1])

    imgCode = await img.read()

    fs = open(f'{linkUnit[0]}', 'wb')
    fs.write(imgCode)
    fs.close()

    print(f'[!] Download Complete {linkUnit[0]}')
    return f'[!] Downlad Successfully {linkUnit[0]}'


async def DownloadMain(loop, linkUnits):
    async with aiohttp.ClientSession() as session:
        tasks = [loop.create_task(DownloadJob(session, linkUnits[_]))
                 for _ in range(len(linkUnits))]
        finished, unfinished = await asyncio.wait(tasks)
        all_result = [r.result() for r in finished]
        for res in all_result:
            print(res)

# ------------ Utils ---------------


def ParsePageNum(numString):
    pattern = re.compile(r'\d+')
    res = pattern.findall(numString)

    return res[0]


def MakeUrl(IMG_URL_BASE, FIRST_PARAM, SECOND_PARAM, chapter, page):
    prefix = IMG_URL_BASE + FIRST_PARAM + SECOND_PARAM

    chapterStr = '%03d' % chapter
    pageStr = '%03d' % int(page)

    suffix = f'{chapterStr}/{pageStr}.jpg'

    url = prefix + suffix

    return url


def MakeCommicDirs(name, chapterNum, pageNums):
    os.mkdir(name)
    for i in range(chapterNum):
        os.mkdir(f'{name}/{i+1}-{pageNums[i]}')


def MakeLinkUnits(urls, name):
    linkUnits = []
    chapterNum = 1
    for chapter in urls:
        pageNum = 1
        for page in chapter:
            filePath = f'./{name}/{chapterNum}-{len(chapter)}/{pageNum}.jpg'
            linkUnit = (filePath, page)
            linkUnits.append(linkUnit)
            pageNum += 1
        chapterNum += 1

    return linkUnits
