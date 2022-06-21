def PrimeFunction():
    #def PrimeChecker(num):
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

    #number = int(input("Enter a number: "))
    #PrimeChecker(number)


def DistanceFormula():
    return "tuesday"


def QuadraticFormula():
    return "tuesday"


def Slopeformula():
    return "thursday"


def TrignometricIdentity():
    return "friday"


def sum():
    x = int(input("Type First number: "))
    y = int(input("Type Second number: "))
    sum1 = x + y
    print("The sum is: ", sum1)
    #break


def default():
    return "Invalid choice"


switcher = {
    1: PrimeFunction,
    2: DistanceFormula,
    3: QuadraticFormula,
    4: Slopeformula,
    5: TrignometricIdentity,
    6: sum
}


def switch(formula):
    return switcher.get(formula, default)()


if __name__ == "__main__":
    formula = float(input('Enter choice: '))
    switch(formula)
