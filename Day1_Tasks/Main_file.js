const prompt = require('prompt-sync')({fake_val: "OPTIONAL CONFIG VALUES HERE",});
function distance(x1,y1,x2,y2)
{
    const distan=Math.sqrt(Math.pow(x2-x1,2)+Math.pow(y2-y1,2));
    console.log("The distance between the two points is :",distan.toFixed(4));
}

function findRoots(a, b, c)
{
if (a == 0) {
console.log("Invalid");
return;
}

let d = b * b - 4 * a * c;
let sqrt_val = Math.sqrt(Math.abs(d));

if (d > 0) {
console.log(
"Roots are real and different \n");

console.log(
(-b + sqrt_val) / (2 * a) +
+ (-b - sqrt_val) / (2 * a));
}
else if (d == 0) {
console.log(
"Roots are real and same \n" + "<br/>");

console.log(-b / (2 * a) +
+ -b / (2 * a)) ;
}
else
{
console.log("Roots are complex \n");

console.log(-b / (2 * a) + " + i"
+ sqrt_val / (2 * a) + "\n"
+ -b / (2 * a)
+ " - i" + sqrt_val) / (2 * a) ;
}
}


function trignometric()
{
  console.log("You are using trignometric identities");
  
  let value=parseInt(prompt('Enter your number:  '));
  
  console.log("The value for sin is: "+Math.sin(value)+"\n");
  
  console.log("The value for cos is: "+Math.cos(value)+"\n");

  console.log("The value for tan is: "+Math.tan(value));
}
function isPrime() {
  var num = parseInt(prompt("Enter a positive number: "));
  if (num === 2) {
        console.log("Prime number")
  } else if (num > 1) {
    for (var i = 2; i < num; i++) {

      if (num % i !== 0) {
        console.log("Prime number");
        break;
      } else if (num === i * i) {
        console.log("Not a Prime number");
        break;
      } else {
        console.log("Not a Prime number");
        break;
      }
    }
}
  else {
    console.log("Not a Prime number");
  }
}
    
 
function sum1(){
const num1 =  parseInt(prompt('Enter the first number '));
const num2 = parseInt(prompt('Enter the second number '));
const sum = num1 + num2;
return (console.log(`The sum of ${num1} and ${num2} is ${sum}`));
}




     console.log("Choose the formula. \n");
     console.log("1-Prime Function \n");
     console.log("2-Distance Formula \n");
     console.log("3-Quadratic Formula \n");
     console.log("4-Slope formula \n");
     console.log("5-Trignometric Identity \n");
     console.log("6-Sum of numbers \n");
     
     let choice=parseInt(prompt('Enter your choice '));
    switch(choice) {
      case 1:
        isPrime()
      break;

      case 2:
        var x1=parseInt(prompt("Enter First X- Cordinate: "));
        var x2=parseInt(prompt("Enter Second X- Cordinate: "));
        var y1=parseInt(prompt("Enter First Y- Cordinate: "));
        var y2=parseInt(prompt("Enter Second Y- Cordinate: "));
        distance(x1,y1,x2,y2);
     
      break;

      case 3:
      let a = parseInt(prompt('Enter the first number '));
      let b =parseInt(prompt('Enter the second number '));
      let c = parseInt(prompt('Enter the third number '));
      findRoots(a, b, c)

      break;
     
      case 4:
              var x1 = parseInt(prompt("x1: "));
              var y1 = parseInt(prompt("y1: "));
              var x2 = parseInt(prompt("x2: "));
              var y2 = parseInt(prompt("y2: "));
              m = (y2-y1)/(x2-x1);
              console.log("Gradient calculated = " + m);
      break;
     
      case 5:
        trignometric()
      break;
     
      case 6:
            sum1();
     
      break;
     
      default:
      hello("Choice...not found");
    }
   
    function hello(msg){
      console.log(msg);
    }
