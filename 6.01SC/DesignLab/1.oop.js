// PCAP Exercises

// 2.1 Fibonacci 
function fib(n) {
    if(n <= 1) return n

    return fib(n-1) + fib(n-2)
}

console.log(fib(8))

// 2.3 Two-dimensional vectors
class V2 {
  constructor(x, y){
    this.x = x 
    this.y = y
  }
  name (){
    return `Vector [${this.x},${this.y}]`
  }
  getX(){
    return this.x
  }
  getY(){
    return this.y
  }
  add(V) {
    return new V2(this.x+V.x, this.y+V.y)
  }
  mul(scalar){
    return new V2(this.x * scalar, this.y * scalar)
  }
}

console.log(new V2(1.1, 2.2).name())

const v = new V2(1.0, 2.0)
v.getX()

v.getY()


const a = new V2(1.0, 2.0)
const b = new V2(2.2, 3.3)

console.log( a.add(b))

console.log(a.mul(2))
console.log(a.add(b).mul(-1))

function PolyCalcHelper(arr1, arr2){
  const len1 = arr1.length
  const len2 = arr2.length
  let newArr1 = arr1.slice()
  let newArr2 = arr2.slice()
  if(len1 != len2) {
    let diff = Math.abs(len1 - len2)
    let padList = Array(diff).fill(0)
    if (len1 > len2){
      newArr2.splice(0, 0, ...padList)
    } else {
      newArr1.splice(0, 0, ...padList)
    }
  }

  return {
    arr1: newArr1,
    arr2: newArr2
  }
}

function Complex(R , I){
  if( isNaN( R ) || isNaN( I )) { throw new TypeError('Complex params require Number'); }
  this.r = R;                          
  this.i = I;                         
}

class Polynomial{
  constructor(coeffs){
    this.coeffs = coeffs
    this.name = this.getName(coeffs)
  }
  getX(base, bit){
    let res = ""
    if(bit == 0) return base
    if(bit > 1) {
      for(let i=0; i< bit; i ++) {
        res+="*"
      }
      return Math.abs(base) + " z" + res + bit
    } else {
      return Math.abs(base) + " z"
    }  
    
  }
  getName(coeffs){
    let name = []
    const len = coeffs.length;
    for(let i=0; i < len; i ++){
      const item = coeffs[i]
      if(item != 0) {
        name.push((item > 0? "+" : "-"))
        name.push(this.getX(item, len-1-i))
      }
    }
    name.shift()
    return name.join(" ")
  }
  coeff(i){
    return this.coeffs[this.coeffs.length - 1 - i] || null
  }
  add(other){
    let { arr1, arr2} = PolyCalcHelper(this.coeffs, other.coeffs)
    let coeffs = []

    arr1.forEach((item, index) => {
      let coeff = item + arr2[index]
      coeffs.push(coeff)
    })

    return new Polynomial(coeffs)
  }
  mul(other){
    let arr1 = this.coeffs.slice()
    let arr2 = other.coeffs.slice()
    let len = arr1.length + arr2.length - 1
    let coeffs = Array(len).fill(0)

    arr1.forEach((item, index) => {
      arr2.forEach((item2, index2) => {
        let product = item * item2
        coeffs[index + index2] += product    
      })
    })

    console.log(coeffs)

    return new Polynomial(coeffs)
  }
  val(x){
    let num = 0
    let len = this.coeffs.length
    for (let index = 0; index < len; index++) {
      let element = this.coeffs[index];
      if(element!=0) {
        num += element * Math.pow(x, len - 1 - index)
      }
    }
    return num
  }
  roots() {
    const len = this.coeffs.length
    switch(len) {
      case 1 :
        console.error("Order too low to solve for roots.")
        break
      case 2 :
        return [ - (this.coeffs[1] / this.coeffs[0]) ]
      case 3 :
        // x1=（-b+（b^2-4ac)^1/2）/2a ,x2=（-b-（b^2-4ac)^1/2）/2a 初中数学怀旧 
        let a = this.coeffs[0]
        let b = this.coeffs[1]
        let c = this.coeffs[2]
        // 韦达定理
        let x1, x2
        if((Math.pow(b, 2) - 4 * a * c) < 0 ) {  // 此处应有共轭复根
          // x1 = new Complex()
          // x2 = new Complex()
          x1 = -b/2*a + "-" + Math.pow(4*a*c-b*b,1/2)/2*a+"i";
          x2 = -b/2*a + "+" + Math.pow(4*a*c-b*b,1/2)/2*a+"i";
        } else {
          x1 = (-b + Math.sqrt(Math.pow(b, 2) - 4 * a * c)) / (2 * a)
          x2 = (-b - Math.sqrt(Math.pow(b, 2) - 4 * a * c)) / (2 * a)
        }
        return [ x1, x2 ]
      default: 
      console.error("Order too high to solve for roots.")
    }
  }
}
// x2 + 2x + 3 = 0
let p1 = new Polynomial([1, -7, 10, -4, 6])
console.log(p1.name)

let p2 = new Polynomial([3, 0, 0, 0])
console.log(p2.name)

let p3 = new Polynomial([8])
console.log(p3.name)

console.log(p1.coeff(3))

let p4 = new Polynomial([1, 2, 3])

console.log(p4.val(1))
console.log(p4.val(-1))

let p5 = new Polynomial([100, 200])
console.log(p4.add(p5).name)

console.log(p4.mul(p4).name)

console.log(p5.roots())

console.log(p4.roots())

let p6 = new Polynomial([3, 2, -1])

console.log(p6.roots())
// >>> (p1 + p2)(10)
// 1323.0
// >>> p1.mul(p1)
// 1.000 z**4 + 4.000 z**3 + 10.000 z**2 + 12.000 z + 9.000
// >>> p1 * p1
// 1.000 z**4 + 4.000 z**3 + 10.000 z**2 + 12.000 z + 9.000
// >>> p1 * p2 + p1
// 100.000 z**3 + 401.000 z**2 + 702.000 z + 603.000
// >>> p1.roots()
// [(-1+1.4142135623730947j), (-1-1.4142135623730947j)]
// >>> p2.roots()
// [-2.0]
// >>> p3 = Polynomial([3, 2, -1])
// >>> p3.roots()
// [-1.0, 0.33333333333333331]
// >>> (p1 * p1).roots()
// Order too high to solve for roots.


// Operator overloading  需要python特有语法 跳过