import math

def trig():
    
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




def main():

    while True:
        val = input("Please enter your choice\n1)Trignometry\n2)Exit\n ")
        while True:
          try:
            val = float(val)
            break
          except ValueError:
            print("No valid option! Please try again ...\n")
            val = input("")

        if(val==1):
            trig()
            break
        elif (val==2):
            print("E X I T")
            break
        else:
            print("enter a valid option")
        
main()