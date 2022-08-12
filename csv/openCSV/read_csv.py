#!/bin/python3
#
import pandas as dp

#Define column names
cols = ['index','data']

# Read in the CSV with pandas
data = dp.read_csv('data.csv', names=cols)

# Print out the maximum value in the integer column
print(data['index'].max())

