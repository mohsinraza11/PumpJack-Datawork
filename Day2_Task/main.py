import math


def PrimeFunction():
    num = int(input('Enter any Number: '))
    if num > 1:
        for i in range(2, int(num / 2) + 1):
            if (num % i == 0):
                return print(num, 'is not a prime number')
                break

            else:
                return print(num, "is a prime number")
                break
        else:
            return print(num, "is not a prime number")
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
    try:
        x1 = float(input("x1: "))
        assert isinstance(x1, float)
    except:
        print("Not a number!")
        
    try:
        y1 = float(input("y1: "))
        assert isinstance(y1, float)
    except:
        print("Not a number!")
        
    try:
        x2 = float(input("x2: "))
        assert isinstance(x2, float)
    except:
        print("Not a number!")
        
    try:
        y2 = float(input("y2: "))
        assert isinstance(y2, float)
    except:
        print("Not a number!")
        
    m = (y2-y1)/(x2-x1)
    print("Gradient calculated = ", round(m,2))
    
    
def TrignometricIdentity():
    
    ch=input("Pick your choice\n1)Sin\n2)Cos\n3)Tan\n4)Exit\n\n")

    while True:
        while True:
          try:
            ch = int(ch)
            break
          except ValueError:
            print("No valid value! Please try again ...\n")
            ch=input("")

        if(ch==1):
            while True:
               try:
                 val = input("Please enter a number: ")
                 val = float(val)
                 break
               except ValueError:
                print("No valid value! Please try again ...")
            print("Sin: ",math.sin(val))
            break

        elif(ch==2):
            while True:
               try:
                 val = input("Please enter a number: ")
                 val = float(val)
                 break
               except ValueError:
                print("No valid value! Please try again ...")
            print("Cos: ",math.cos(val))
            break
        
        elif(ch==3):
            while True:
               try:
                 val = input("Please enter a number: ")
                 val = float(val)
                 break
               except ValueError:
                print("No valid value! Please try again ...")
            print("Tan: ",math.tan(val))
            break

        elif(ch==4):
            print("Good-bye!")
            break

        else:
            print("Invalid input. Try again...\n")
            ch=input("")


def sum():
    x = int(input("Type First number: "))
    y = int(input("Type Second number: "))
    sum1 = int(x) + int(y)
    print("The sum is: ", sum1)
    
def default():
    return "Invalid choice"

switcher = {
    1: PrimeFunction,
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
