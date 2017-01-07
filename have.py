#-*-coding:utf8-*-

import re
import string
import sys
import os
import urllib
import urllib2
from bs4 import BeautifulSoup
import requests
from lxml import etree

reload(sys) 
sys.setdefaultencoding('utf-8')

f = open('have.swift', 'wr')
wd = '''
//
//  Created by ZhangLiangZhi on 2017/1/7.
//  Copyright © 2017年 xigk. All rights reserved.
//


func initJSword1() {
    var arrMeaning = [String]()
    var arrProp = [JSprop]()
'''
wd = wd + '\n    '

soup=BeautifulSoup(open('have.html'), "lxml")


getfy = ""
allul= soup.find_all('ul')
for ul in allul:
	if ul:
		if ul['class'][0] == 'base-list':
			# print ul
			getfy = ul
			break
			pass

# print getfy
allli = getfy.find_all('li')
for li in allli:
	# print li
	prop = li.span.string
	allspan = li.p.find_all('span')

	wd = wd + 'arrMeaning = []\n    '
	for span in allspan:
		meaning = span.string
		wd = wd + 'arrMeaning.append("' + meaning + '")\n    '
	pass
	wd = wd + 'arrProp.append(JSprop(prop: "' + prop + '", meaning: arrMeaning))\n    '

wd = wd + 'garrJSWord.append(JSWord(id: 1, word: "' + 'have' + '", arrProp: arrProp))\n}'

print wd
f.write(wd)
f.close()





