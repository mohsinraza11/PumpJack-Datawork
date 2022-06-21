
function distance(x1,y1,x2,y2)
{
    const distan=Math.sqrt(Math.pow(x2-x1,2)+Math.pow(y2-y1,2));
    console.log("The distance between the two points is :",distan.toFixed(4));
}

function findRoots(a, b, c)
{
if (a == 0) {
document.write("Invalid");
return;
}

let d = b * b - 4 * a * c;
let sqrt_val = Math.sqrt(Math.abs(d));

if (d > 0) {
document.write(
"Roots are real and different \n");

document.write(
(-b + sqrt_val) / (2 * a) +
+ (-b - sqrt_val) / (2 * a));
}
else if (d == 0) {
document.write(
"Roots are real and same \n" + "<br/>");

document.write(-b / (2 * a) +
+ -b / (2 * a)) ;
}
else
{
document.write("Roots are complex \n");

document.write(-b / (2 * a) + " + i"
+ sqrt_val / (2 * a) + "\n"
+ -b / (2 * a)
+ " - i" + sqrt_val) / (2 * a) ;
}
}


function trignometric()
{
  document.write("You are using trignometric identities");
  
  let value=parseInt(prompt('Enter your number:  '));
  
  document.write("The value for sin is: "+Math.sin(value)+"\n");
  
  document.write("The value for cos is: "+Math.cos(value)+"\n");

  document.write("The value for tan is: "+Math.tan(value));
}
function isPrime(num) {
  const num = parseInt(prompt("Enter a positive number: "));
  if (num === 2) {
    return true;
  } else if (num > 1) {
    for (var i = 2; i < num; i++) {

      if (num % i !== 0) {
        return true;
      } else if (num === i * i) {
        return false
      } else {
        return false;
      }
    }
  } else {
    return false;
  }

}


     document.write("Choose the formula. \n");
     document.write("1-Prime Counting Function \n");
     document.write("2-Distance Formula \n");
     document.write("3-Quadratic Formula \n");
     document.write("4-Slope formula \n");
     document.write("5-Trignometric Identity \n");
     document.write("6-Sum of numbers \n");
     
     let choice=parseInt(prompt('Enter your choice '));
    switch(choice) {
      case 1:
            console.log(isPrime());
      break;

      case 2:
        const prompt=require("prompt-sync")({fake_val:"OPTIONAL CONFIG VALUES HERE",});
        let x1=parseInt(prompt("Enter First X- Cordinate: "));
        let x2=parseInt(prompt("Enter Second X- Cordinate: "));
        let y1=parseInt(prompt("Enter First Y- Cordinate: "));
        let y2=parseInt(prompt("Enter Second Y- Cordinate: "));
        distance(x1,y1,x2,y2);
     
      break;

      case 3:
      let a = parseInt(prompt('Enter the first number '));
      let b =parseInt(prompt('Enter the second number '));
      let c = parseInt(prompt('Enter the third number '));
      findRoots(a, b, c)

      break;
     
      case 4:
              var x1 = prompt("x1: ");
              var y1 = prompt("y1: ");
              var x2 = prompt("x2: ");
              var y2 = prompt("y2: ");
              m = (y2-y1)/(x2-x1);
              alert("Gradient calculated = " + m);
      break;
     
      case 5:
        trignometric()
      break;
     
      case 6:
            
     
      break;
     
      default:
      hello("Choice...not found");
    }
   
    function hello(msg){
      alert(msg);
    }
