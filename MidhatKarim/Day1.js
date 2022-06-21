function calculate_slope(){
    var x1 = prompt("x1: ");
    var y1 = prompt("y1: ");
    var x2 = prompt("x2: ");
    var y2 = prompt("y2: ");
    m = (y2-y1)/(x2-x1);
    alert("Gradient calculated = " + m);
}

calculate_slope()
