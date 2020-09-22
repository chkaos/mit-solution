// Problem Wk.2.1.1: State machines

class SM {
  constructor(initialValue) {
    if (initialValue == undefined)
      this.startState = initialValue // 记住初始值
    this.state = 0
    this.input = 0   // input output 不一定需要
    this.output = 0
    this.indent = 0  // log 部分
  }
  // 启动方法 startState 赋值给 state
  start() {
    this.state = this.startState
  }
  // 给一个输入值, 改变内部 state 并返回输出
  step(input, verbose = false) {
    // if (input === undefined || input === null) input = null
    const { state, output } = this.getNextValues(this.state, input)

    if (verbose) this.debug()

    this.state = state
    this.input = input
    this.output = output
    return output
  }
  transduce(inputs, verbose = false, compact = false) {
    this.start()
    if (verbose) {
      console.log(`StartState: ${this.state}`)
      let n = inputs.length
      for (let i = 0; i < n; i++) {
        console.log("Step: ", i)
        // console.log(" ",  this.constructor.name)
        this.step(inputs[i], verbose)
      }
    } else {
      let outputs = []
      for (let i in inputs) {
        outputs.push(this.step(inputs[i]))
      }
      return outputs
    }

  }
  // 无输入状态机调用run, n为次数
  run(n = 10, verbose = false) {
    return this.transduce(Array(n).fill(null), verbose)
  }
  getNextValues(state, input) {
    const nextState = this.getNextState(state, input)
    return { state: nextState, output: nextState }
  }
  debug(indent = 0) {
    this.indent = indent
    console.log(`${this.indent + " " + this.constructor.name} In: ${this.input} Out: ${this.output} Next State: ${this.state}`)
  }
  done() {
    return false
  }
}

// Problem Wk.2.1.2: Turnstile state machine
class Turnstile extends SM {
  constructor() {
    super("locked")
  }

  getNextValues(state, input) {

    if (state == "locked") {
      if (input == "coin") {
        return { state: "unlocked", output: "enter" }
      }
      else { return { state: "locked", output: "pay" } }
    } else {
      if (input == "turn") {
        return { state: "locked", output: "enter" }
      }
      else {
        return { state: "unlocked", output: "enter" }
      }
    }
  }
}

// Wk.2.1.3: Double Delay State Machine

class Delay2Machine extends SM {
  constructor(val0, val1) {
    super()
    this.startState = { val0, val1 }
  }
  getNextValues(state, input) {
    const { val0, val1 } = state
    const nextState = { val0: val1, val1: input }
    return { state: nextState, output: val0 }
  }
}

console.log("Test1", new Delay2Machine(100, 10).transduce([1, 0, 2, 0, 0, 3, 0, 0, 0, 4]))
console.log("Test2", new Delay2Machine(10, 100).transduce([0, 0, 0, 0, 0, 0, 1]))
console.log("Test3", new Delay2Machine(-1, 0).transduce([1, 2, -3, 1, 2, -3]))

const m = new Delay2Machine(100, 10)
m.start()
let arr1 = [-1, -2, -3, -4, -5, -6]
arr1.forEach(i => {
  m.getNextValues(m.state, i)
});

let arr2 = [1, 0, 2, 0, 0, 3, 0, 0, 0, 4]
const output = []
arr2.forEach(i => {
  output.push(m.step(i))
})

console.log("Test4", output)

// # execute runTestsDelay() to carry out the testing, you should get:
// #Test1: [100, 10, 1, 0, 2, 0, 0, 3, 0, 0]
// #Test2: [10, 100, 0, 0, 0, 0, 0]
// #Test3: [-1, 0, 1, 2, -3, 1]
// #Test4: [100, 10, 1, 0, 2, 0, 0, 3, 0, 0]

// Problem Wk.2.1.4: Comments Machine

// >> str = '''def f(x): # comment
// return 1'''
// >>> m = CommentsSM()
// >>> m.transduce(str)
// [None, None, None, None, None, None, None, None, None, None,
// '#', ' ', 'c', 'o', 'm', 'm', 'e', 'n', 't',
// None, None, None, None, None, None, None, None, None, None, None, None] 

class CommentsSM extends SM {
  constructor() {
    super()
  }
  getNextValues(state, input) {

    if (input == "#" || (state && state != "\n")) {
      return { state: input, output: state }
    } else {
      return { state: undefined, output: undefined }
    }
  }
}

const str = `def f(x): # comment
                 return 1`
const cm = new CommentsSM()
console.log(cm.transduce(str))