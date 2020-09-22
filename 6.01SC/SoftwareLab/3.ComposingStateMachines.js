const { SM, Cascade, Delay, Increment } = require ("../SM")

// Problem Wk.3.1.1: Simulating Cascade
// 1.
let d1 = new Delay(1)
let d2 = new Delay(2)
let c =  new Cascade(d1, d2)
console.log(c.transduce([3,5,7,9]))

// 2.
let d3 = new Delay(1)
let d4 = new Increment(3)
let c2 =  new Cascade(d3, d4)
console.log(c2.transduce([3,5,7,9]))

// Problem Wk.3.1.2: Cascading Machines
// pass

// Problem Wk.3.1.3: Function Machines 
class PureFunction extends SM {
  constructor(fn) {
    super()
    this.fn = fn
  }
  getNextValues(state, input){
    return {state: undefined, output: this.fn(input)}
  }
}