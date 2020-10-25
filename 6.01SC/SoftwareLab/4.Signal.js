class Signal {
  constructor() { }
  plot() {
    // 非python, 暂时不支持
  }
  __add__(other) {
    return new SummedSignal(this, other)
  }
  __rmul__(scalar) { }
  __mul__(scalar) { }
  period(n = null, z = null) { }
  crossings(n = null, z = null) { }
  mean(n = null) { }
  samplesInRange(lo, hi) {
    const outputs = []
    for (let i = lo; i < hi; i++) {
      outputs.push(this.sample(i))
    }
    return outputs
  }
}

class CosineSignal extends Signal {
  constructor(omega, phase) {
    super()
    this.omega = omega
    this.phase = phase
  }
  sample(n) {
    return Math.cos(this.omega * n + this.phase)
  }
}

// Problem Wk.4.1.1: Constructing Signals
class ConstantSignal extends Signal {
  constructor(c) {
    super()
    this.c = c
  }
  sample() {
    return this.c
  }
}

// Problem Wk.4.1.2: Step Signal
class StepSignal extends Signal {
  sample(n) {
    return n == 0 ? 1 : 0
  }
}

// Problem Wk.4.1.3: Subclasses of Signals
class SummedSignal extends Signal {
  constructor(s1, s2) {
    super()
    this.s1 = s1
    this.s2 = s2
  }
  sample(n) {
    return
  }
}

class ScaledSignal extends Signal {
  constructor(s, c) {
    super()
    this.s = s
    this.c = c
  }
  sample(n) {
    return
  }
}

// Problem Wk.4.1.4: Additional subclasses 

class R extends Signal {
  constructor(s) {
    super()
    this.s = s
  }
  sample(n) {
    return this.sig.sample(n-1)
  }
}

class RN extends Signal {
  constructor(s, n) {
    super()
    this.s = s
    this.n = n
  }
  sample(n) {
    return this.sig.sample(n-this.n)
  }
}

// Problem Wk.4.1.5: PolyR on Signals