
function distance(x1,y1,x2,y2)
{
    const distan=Math.sqrt(Math.pow(x2-x1,2)+Math.pow(y2-y1,2));
    console.log("The distance between the two points is :",distan.toFixed(4));
}

const prompt=require("prompt-sync")({fake_val:"OPTIONAL CONFIG VALUES HERE",});
let x1=parseInt(prompt("Enter First X- Cordinate: "));
let x2=parseInt(prompt("Enter Second X- Cordinate: "));
let y1=parseInt(prompt("Enter First Y- Cordinate: "));
let y2=parseInt(prompt("Enter Second Y- Cordinate: "));
distance(x1,y1,x2,y2);