// github.com/dengsgo/math-engine damn!未完成, 卡太久了

class BinaryOp {
  constructor(left, right) {
    this.left = left
    this.right = right
  }
  __str__() {
    return this.opStr + '(' + this.left + ', ' + this.right + ')'
  }
}

class Sum extends BinaryOp {
  constructor() {
    super()
    this.opStr = 'Sum'
  }
}
class Prod extends BinaryOp {
  constructor() {
    super()
    this.opStr = 'Prod'
  }
}

class Quot extends BinaryOp {
  constructor() {
    super()
    this.opStr = 'Quot'
  }
}

class Diff extends BinaryOp {
  constructor() {
    super()
    this.opStr = 'Diff'
  }
}

class Assign extends BinaryOp {
  constructor() {
    super()
    this.opStr = 'Assign'
  }
}

class CNumber {
  constructor(val) {
    this.value = val
  }
  __str__() {
    return 'Num(' + this.value + ')'
  }
}

class Variable {
  constructor(name) {
    this.name = name
  }
  __str__() {
    return 'Variable(' + this.name + ')'
  }
}

const seps = ['(', ')', '+', '-', '*', '/', '=']
const precedence = { "+": 1, "-": 1, "*": 2, "/": 2 }

const tokenize = (input) => {
  const tokens = []
  let tem = ""
  for (let i = 0; i < input.length; i++) {
    let char = input[i]
    if (seps.includes(char)) {
      if (tem) {
        tokens.push(tem)
        tem = ""
      }
      tokens.push(char)
    } else if (char != " ") {
      tem += char
    }

  }
  return tokens
}

// console.log(tokenize("((fred + george) / (voldemort + 666))"))

const tokens = [
  '(', '(',
  'fred', '+',
  'george', ')',
  '/', '(',
  'voldemort', '+',
  '666', ')',
  ')'
]

class AST {
  constructor(source) {
    this.tokens = []
    this.source = source
    this.curTok = null
    this.curIndex = 0
    this.err = null
    this.init()
  }
  init() {
    let tokens = tokenize(this.source)
    if (tokens.length == 0 || !tokens) {
      throw new Error("INVALID TOKENS")
    } else {
      this.tokens = tokens
      this.curTok = tokens[0]
    }
  }
  parse() {
    const lhs = this.parsePrimary()
    console.log(lhs)
    return this.parseBinOpRHS(0, lhs)
  }
  binOp(left, right) {
    switch (this.curTok) {
      case "+":
        return new Sum(left, right)
      case "-":
        return new Diff(left, right)
      case "*":
        return new Prod(left, right)
      case "/":
        return new Quot(left, right)
    }
  }
  parseBinOpRHS(execPrec, lhs) {

    while (true) {
      let tokPrec = this.getTokPrecedence()
      // console.log(tokPrec, execPrec)
      if (tokPrec < execPrec) {
        return lhs
      }
      if (!this.getNextToken()) {
        return lhs
      }
      let rhs = this.parsePrimary()
      if (!rhs) {
        return null
      }
      let nextPrec = this.getTokPrecedence()
      if (tokPrec < nextPrec) {
        // 递归，将当前优先级+1
        rhs = this.parseBinOpRHS(tokPrec + 1, rhs)
        if (!rhs) {
          return null
        }
      }
      lhs = this.binOp(lhs, rhs)
    }
  }
  getTokPrecedence() {
    return precedence[this.curTok] || -1
  }
  parsePrimary() {
    // 数字
    if (numberTok(this.curTok)) {
      return this.parseNumber()
    }
    // 变量
    if (variableTok(this.curTok)) {
      return this.parseVariable()
    }
    // 操作符
    if (this.curTok == "(") {
      this.getNextToken()
      let err = this.parse()
      if (!err) return null
      if (this.curTok.Tok != ")") {
        return null
      }
      this.getNextToken()
    } else {
      if (numberTok(this.curTok)) {
        return this.parseNumber()
      }
      // 变量
      if (variableTok(this.curTok)) {
        return this.parseVariable()
      }
    }

    return null
  }
  getNextToken() {
    this.curIndex++
    if (this.curIndex < this.tokens.length) {
      this.curTok = this.tokens[this.curIndex]
      return this.curTok
    }
    return null
  }
  parseNumber() {
    const numberAST = new CNumber(Number(this.curTok))
    this.getNextToken()
    return numberAST
  }
  parseVariable() {
    const variableAST = new Variable(this.curTok)
    this.getNextToken()
    return variableAST
  }
}

const numberTok = (input) => {
  const value = Number(input)
  return typeof value === 'number' && isFinite(value)
}

const variableTok = (input) => {
  return !numberTok(input) && !seps.includes(input)
}

const forma = new AST("((fred + george) / (voldemort + 666))")
console.log(forma.parse())