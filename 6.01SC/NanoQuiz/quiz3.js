// Part 1: Constructing a State Machine

// Define a state machine for a vending machine.
// The vending machine has an unlimited number of sodas that it sells for 75 cents apiece.
// The user can deposit quarters in the machine.
// If the user presses the cancel button, all the coins she's put in so far are returned.
// If the user presses the dispense button,
// If she had not deposited at least 75 cents, she gets no soda and no change.
// If she had deposited 75 cents or more, she gets a soda and any amount over 75
// cents that she has deposited.
// More concretely,
// The input for the machine is one of:
// 'quarter' (25 cents) indicates that a coin has been deposited
// 'cancel' indicates that the user wants her money back
// 'dispense' indicates that the user wants a soda
// The output of the machine is:
//  (change, soda)
// where
// change is an integer that indicates (in cents) the amount of change (possibly 0)
// that is returned
// soda is a boolean that indicates whether or not the machine will dispense a soda
// For example:
// Vending().transduce(['dispense', 'quarter', 'quarter', 'quarter', 'quarter',
// 'dispense', 'quarter', 'cancel', 'dispense'])
// would return:
// [(0, False), (0, False), (0, False), (0, False), (0, False),
// (25, True), (0, False), (25, False), (0, False)] 

// class Vender {
//   constructor() {
//     this.pending = 0     // 0 unsellable 1 sellable
//     this.change = 0
//   }

//   transduce(actions) {
//     const result = []
//     actions.forEach(action => {
//       console.log(action)
//       result.push(this.doAction(action))
//     });
//     return result
//   }

//   doAction(actionName) {
//     return this[`${actionName}`].call(this)
//   }

//   quarter() {
//     this.change = this.change + 25
//     console.log(this.change)
//     return { amount: 0, state: false }
//   }

//   dispense() {
//     let state = false
//     if (this.change >= 75) {
//       state = true
//     }
//     let amount = this.getChange()
//     this.change = 0
//     return { amount, state }
//   }

//   cancel() {
//     let amount = this.change
//     this.change = 0
//     return { amount, state: false }
//   }

//   getChange() {
//     return this.change - 75 >= 0 ? this.change - 75 : 0
//   }
// }

class Vender {
  constructor() {
    this.state = 0
  }
  execute(action) {
    let result = { amount: 0, state: false }
    switch (action) {
      case "cancel":
        result.amount = this.state
        this.state = 0
        break
      case "dispense":
        if (this.state >= 75) {
          result.amount = this.state - 75
          result.state = true
          this.state = 0
        }
        break
      case "quarter":
        this.state += 25
        break
    }
    return result
  }
  transduce(actions) {
    const result = []
    actions.forEach(action => {
      result.push(this.execute(action))
    });
    return result
  }
}

let vendor = new Vender()
console.log(vendor.transduce(['dispense', 'quarter', 'quarter', 'quarter', 'quarter',
  'dispense', 'quarter', 'cancel', 'dispense']))
