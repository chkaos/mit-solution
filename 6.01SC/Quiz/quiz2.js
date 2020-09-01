class Hammock {
  constructor() {
    this.inHammock = 0
    this.last = ""
  }
  sitDown(name) {
    if (this.inHammock == 0 || this.last == name) {
      this.inHammock += 1
      console.log("welcome")
      return
    }

    this.last = name
    console.log('sorry, no room')

  }
  leave() {
    if (this.inHammock > 0) {
      this.inHammock--
    }
    console.log(this.inHammock)
  }
}

const myHammock = new Hammock()
myHammock.sitDown('George')
// welcome!
myHammock.sitDown('Bobby')
// sorry, no room
myHammock.sitDown('Bobby')
// welcome!
myHammock.leave()
// 1
myHammock.leave()
// 0
myHammock.leave()
// 0
myHammock.sitDown('Martha')
// welcome!
myHammock.sitDown('Wilhelm')
// sorry, no room
myHammock.sitDown('Klaus')
// sorry, no room
myHammock.sitDown('Wilhelm')
// sorry, no room
myHammock.leave()
// 0