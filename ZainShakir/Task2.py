import math

def distance_formula(x1,y1,x2,y2):
    distance=math.sqrt(pow(x2-x1,2)+pow(y2-y1,2))
    print(f"The distacne between the two points is {distance}")
    
    
    
x1=float(input("Enter the value of x1: "))
y1=float(input("Enter the value  of y1: "))
x2=float(input("Enter the value of x2: "))
y2=float(input("Enter the value of y2: "))

distance_formula(x1, y1, x2, y2)
    
    

