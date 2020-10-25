const { SM, Cascade, Feedback, Delay, Increment, Parallel, Parallel2 } = require ("../SM")

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
    return {state: undefined, output: this.fn(state, input)}
  }
}

// Problem Wk.3.1.4: Combining accounts
const FN1 = (state, input) => {
  let newState = 0
  if(input != 0) {
    newState = state*1.02 + input - 100
  } else {
    newState = state*1.02
  }
  return {state: newState, output: newState}
}

const FN2 = (state, input) => {
  let newState = state*1.01 + input
  return {state: newState, output: newState}
}
const BA1 = new PureFunction(FN1)
const BA2 = new PureFunction(FN2)

const maxAccount = new Parallel(BA1, BA2)

const switchAccount = new Parallel2(BA1, BA2)

console.log(switchAccount)
// abandon
// Problem Wk.3.1.5: Sequential combinations
// Part 1: Sum machine 
class SumTM extends SM {
  constructor() {
    super()
    this.SM = new Accumulator(0)
  }
  done(state){
    return state>100
  }
}

// Part 3: Counting machine
class CountUpTo extends SM {
  constructor(initialValue) {
    super(initialValue)
    this.state = 0
  }
  start() {
    this.state = 0
  }
  done(){
    return this.state >= this.startState
  }
  getNextState(state, input) {
    return state + 1
  }
}

const m  = new CountUpTo(3)
console.log(m.run(20))

// Part 4: Multiple Counting machine 
function makeSequenceCounter(nums) {

}

// Problem Wk.3.1.6: Feedback SM 
const negate = new PureFunction((state) => {
  let newState = !state
  return {state: newState, output: newState}
})

// new Feedback(new Cascade(new negate(step), new Delay(true)))