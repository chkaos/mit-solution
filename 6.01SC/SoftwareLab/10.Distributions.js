// Problem Wk.10.1.1: Probability distributions: DDist

const { DDist, JDist } = require("../DDist")

const foo = new DDist([["hi", 0.6], ["med", 0.1], ["lo", 0.3]])
console.log(foo)

const Disease = new DDist([["disease", 0.1], ["noDisease", 0.9]])

// Problem Wk.10.1.2: Conditional distributions
function PTgD(D) {
  if (D == "disease") {
    return new DDist([["posTest", 0.9], ["negTest", 0.1]])
  }

  else if (D == "noDisease") {
    return new DDist([["posTest", 0.5], ["negTest", 0.5]])
  }

  else { return new Error("invalid value for D") }
}

console.log(PTgD('disease').prob('posTest'))

// Problem Wk.10.1.3: Joint distributions
// １．
new DDist([[["noDisease", "posTest"], 0.05],
[["disease", "posTest"], 0.98],
[["noDisease", "negTest"], 0.95],
[["disease", "negTest"], 0.05]])
// 2. JDist marginalization example

console.log(JDist(Disease, PTgD))

// Problem Wk.10.1.4: Operations on Conditional Distributions

// Problem Wk.10.1.6: Implementing Joint distributions
console.log(JDist(Disease, PTgD).marginalizeOut(0))