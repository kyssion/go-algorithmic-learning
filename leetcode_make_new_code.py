import argparse
from asyncio import base_tasks
import imp
import logging
import os
parser = argparse.ArgumentParser(description='manual to this script')
parser.add_argument("--code_num", type=int, default=0)
parser.add_argument("--desc", type=str, default="")
args = parser.parse_args()
print("code_num = "+str(args.code_num))

def get_num_code( codeNum ):
    if codeNum>=0 and codeNum <=900:
        return 3
    elif codeNum >=901 and codeNum<=9000:
        return 4
    elif codeNum>=9001 and codeNum<99999:
        return 5
    else :
        return 0

def fil_code_str(code_num_str, length):
    if len(code_num_str)>length:
        return code_num_str
    base_str = ""
    for x in range(length-len(code_num_str)):
        base_str= "0"+base_str
    return "code"+base_str+code_num_str

if args.code_num == 0:
    logging.error("code num is error")
    exit()

code_num = args.code_num
basePath= "leetcode"

if os.path.exists(basePath) is False:
    os.mkdir(basePath)

timesEnd = (int((code_num-1)/100)+1)*100
timesStart = timesEnd - 100 + 1
aggregation_path = ""+str((int(timesStart/100)+1))+"_code_"+str(timesStart)+"_"+str(timesEnd)

if os.path.exists(basePath+"/"+aggregation_path) is False:
    os.mkdir(basePath+"/"+aggregation_path)

num_length = get_num_code(code_num)
code_str = fil_code_str(str(code_num),num_length)
path_item = basePath+"/"+aggregation_path+"/"+code_str

logging.info("add leetcode code path is "+path_item)

if os.path.exists(basePath+"/"+aggregation_path+"/"+code_str) is False:
    os.mkdir(basePath+"/"+aggregation_path+"/"+code_str)


mainFile = open(path_item+"/"+"main.go",'w')
mdFile = open(path_item+"/"+"题目.md",'w')

mainFile.write(
'''
package main

func main() {

}
'''
)
mdFile.write("# "+args.desc)
mainFile.close()
mdFile.close()

logging.info("init leetcode file success !")