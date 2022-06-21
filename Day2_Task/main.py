import math


def PrimeCountingFunction():
    return "monday"
def DistanceFormula(x1,y1,x2,y2):
    distance=math.sqrt(pow(x2-x1,2)+pow(y2-y1,2))
    print(f"The distacne between the two points is {distance}")
def QuadraticFormula():
    return "tuesday"
def Slopeformula():
    return "thursday"
def TrignometricIdentity():
    return "friday"
def Sumofnumbers():
    return "Saturday"
def default():
    return "Invalid choice"

switcher = {
    1: PrimeCountingFunction,
    2: DistanceFormula,
    3: QuadraticFormula,
    4: Slopeformula ,
    5: TrignometricIdentity,
    6: Sumofnumbers
    }
  
def switch(formula):
    return switcher.get(formula, default)()

if __name__ == "__main__":
    formula = float(input('Enter choice: '))
    print(switch(formula))
