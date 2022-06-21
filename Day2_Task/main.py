import math


def PrimeCountingFunction():
    return "monday"
def DistanceFormula():
    x1=float(input("Enter the value of x1: "))
    y1=float(input("Enter the value  of y1: "))
    x2=float(input("Enter the value of x2: "))
    y2=float(input("Enter the value of y2: "))
    distance=math.sqrt(pow(x2-x1,2)+pow(y2-y1,2))
    print(f"The distacne between the two points is {distance}")
    
def QuadraticFormula():
    a = int(input('Enter a:'))  
    b = int(input('Enter b:'))  
    c = int(input('Enter c:')) 
    dis_form = b * b - 4 * a * c  
    sqrt_val = math.sqrt(abs(dis_form))  
  
  
    if dis_form > 0:  
        print(" real and different roots ")  
        print((-b + sqrt_val) / (2 * a))  
        print((-b - sqrt_val) / (2 * a))  
  
    elif dis_form == 0:  
        print(" real and same roots")  
        print(-b / (2 * a))  
  
  
    else:  
        print("Complex Roots")  
        print(- b / (2 * a), " + i", sqrt_val)  
        print(- b / (2 * a), " - i", sqrt_val)  
   
  
    if a == 0:  
        print("Input correct quadratic equation")  
    
def Slopeformula():
    return "thursday"
def TrignometricIdentity():
    return "friday"
def sum():
    x = int(input("Type First number: "))
    y = int(input("Type Second number: "))
    sum1 = int(x) + int(y)
    print("The sum is: ", sum1)
    
def default():
    return "Invalid choice"

switcher = {
    1: PrimeCountingFunction,
    2: DistanceFormula,
    3: QuadraticFormula,
    4: Slopeformula ,
    5: TrignometricIdentity,
    6: sum
    }
  
def switch(formula):
    return switcher.get(formula, default)()

if __name__ == "__main__":
    formula = float(input('Enter choice: '))
    print(switch(formula))
