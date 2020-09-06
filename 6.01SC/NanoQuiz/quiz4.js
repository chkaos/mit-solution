class Triangle {
  constructor(h) {
    this.height = h
  }
  sample(n) {
    if(Math.abs(n) < this.height){
      return this.height - Math.abs(n)
    }
    return 0
  }
}

const triangle = new Triangle(5)
triangle.sample(4)

// trapezoid = SummedSignal(SummedSignal(Rn(Triangle(2),2),
// Rn(Triangle(2),4)),
// Rn(Triangle(2),6))
// # Another way of doing it.
// #trapezoid = polyR(Triangle(2), poly.Polynomial([1, 0, 1, 0, 1, 0, 0]))