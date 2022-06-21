import math
import numpy as np

import warnings
warnings.filterwarnings("ignore")

def input_coordinates():
    x1 = float(input("x1: "))
    y1 = float(input("y1: "))
    x2 = float(input("x2: "))
    y2 = float(input("y2: "))
    return x1,x2,y1,y2
 
def calculate_slope():
    x1,x2,y1,y2 = input_coordinates()
    m = (y2-y1)/(x2-x1)
    print("Gradient calculated = ", round(m,2))
    
 calculate_slope()
