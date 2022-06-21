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
     
      break;

      case 2:
     
      break;

      case 3:
      let a = parseInt(prompt('Enter the first number '));
      let b =parseInt(prompt('Enter the second number '));
      let c = parseInt(prompt('Enter the third number '));
      findRoots(a, b, c)

      break;
     
      case 4:
     
      break;
     
      case 5:
     
      break;
     
      case 6:
     
      break;
     
      default:
      hello("Choice...not found");
    }
   
    function hello(msg){
      alert(msg);
    }