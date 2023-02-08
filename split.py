import csv
import pandas as pd

df = pd.read_csv('')

df_small=pd.read_csv('',nrows=10000)

i=1

for file in df:
    print(file.shape)
    file.to_csv(f' file_{i}.csv')
    i+=1