import math

def trig():
    ch=input("Pick your choice\n1)Sin\n2)Cos\n3)Tan\n\n")
    ch=int(ch)
    val=input("Enter the number: ")
    val=int(val)

    match ch:
        case 1:
            print(math.sin(val))
        case 2:
            print(math.cos(val))
        case 3:
            print(math.tan(val))
        case default:
            print("Invalid input")


def main():

    val = input("Enter your choice\n1)Trignometry\n2)Exit\n\n")
    val=int(val)

    match val:
        case 1:
            trig()
        case default:
            print("Goodbye!")

main()