#!/usr/bin/env python
import sys
import pandas as pd
import pymongo
import json
import os

class MongoBase:
    
    def __init__(self,collection):
        self.collection=collection
        self.OpenDB()
    
    def OpenDB(self):
        user=''
        passwd=''
        host='localhost'
        port='27017'
        auth_db=''
        db='semaforo'
        #uri = "mongodb://"+user+":"+passwd+"@"+host+":"+port+"/"+auth_db+"?authMechanism=SCRAM-SHA-1"
        #uri = "mongodb://"+host+":"+port+"/"+auth_db+"?authMechanism=SCRAM-SHA-1"
        uri = "mongodb://"+host+":"+port+"/"+db
        
        self.con = pymongo.MongoClient(uri, connect=False)
        
        self.db=self.con['semaforo']
        
        self.collection=self.db[self.collection]
    
    def closeDB(self):
        self.con.close()


def import_content(filepath):
    mng_client = pymongo.MongoClient('localhost', 27017)
    mng_db = mng_client['semaforo'] # Replace mongo db name
    collection_name = 'estados' # Replace mongo db collection name
    db_cm = mng_db[collection_name]
    cdir = os.path.dirname(__file__)
    file_res = os.path.join(cdir, filepath)

    data = pd.read_csv(file_res)
    data_json = json.loads(data.to_json(orient='records'))
    db_cm.remove()
    db_cm.insert(data_json)

if __name__ == "__main__":
  #filepath = '/path/to/csv/path'  # pass csv file path
  #import_content(filepath)
  
  mongo=MongoBase('teste')
  dicts = {'one': [1, 2, 3], 'two': [2, 3, 4], 'three': [3, 4, 5]}
  df = pd.DataFrame(dicts)
  mongo.collection.insert(json.loads(df.T.to_json()).values())
  mongo.closeDB()
  print(json.loads(df.T.to_json()).values())