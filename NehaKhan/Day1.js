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


let a = parseInt(prompt('Enter the first number '));
let b =parseInt(prompt('Enter the second number '));
let c = parseInt(prompt('Enter the third number '));
findRoots(a, b, c)
