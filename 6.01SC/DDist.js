class DDist {
  constructor(dictionary) {
    // map 对象
    this.d = new Map()
    for (var [k, v] of dictionary) {
      this.d.set(this.handleKey(k), v)
    }
  }
  prob(elt) {
    elt = this.handleKey(elt)
    if (this.d.get(elt) != undefined) {
      return this.d.get(elt)
    }
    return 0
  }
  support() {
    let res = []
    for (var [k, v] of this.d) {
      if (v > 0) {
        res.push(k)
      }
    }
    return res
  }
  draw() {
    const r = Math.random()
    let sum = 0
    this.support().forEach(val => {
      sum += this.prob(val)
      if (r < sum) {
        return val
      }
    })
  }
  handleKey(k) {
    return typeof k == "string" ? JSON.stringify(k) : k
  }
  marginalizeOut(index) {
    const map = new Map()
    const d = this.d
    d.forEach((v1, k1) => {
      //  暂时不知什么原因 key上多了加了双引号
      console.log(typeof k1)
      let key = ""
      if (Array.isArray(k1)) {
        key = k1[index]
      } else {
        key = k1
      }
      // let db = PB(k1).d

      // if (db) {
      //   db.forEach((v2, k2) => {
      //     k2 = k2.replace(/\"/g, "")
      //     console.log(v1, v2)
      //     map.set([k1, k2], (v1 * v2).toFixed(2))
      //   })
      // }
    })
    return new DDist(map)
  }

  conditionOnVar() { }
}

const dict = new DDist([[[0, 0], 0.5], [[0, 1], 0.2], [[1, 0], 0.1], [[1, 1], 0.2]])

// console.log(dict.prob([0, 0]))

function TgivenD(D) {
  if (D == 'disease') {
    return new DDist([["positive", 0.99], ["negative", 0.01]])
  }

  else if (D == "noDisease") {
    return new DDist([["positive", 0.001], ["negative", 0.999]])
  }

  else { return new Error("invalid value for D") }

}

// console.log(TgivenD("disease").prob("negative"))

// PA = dist.DDist({’a1’ : 0.9, ’a2’ : 0.1}) def PBgA(a): if a == ’a1’: return dist.DDist({’b1’ : 0.7, ’b2’ : 0.3}) else:return dist.DDist({’b1’ : 0.2, ’b2’ : 0.8})
// >>> PAB = JDist(PA, PBgA)
// >>> PAB
// DDist((a1, b2): 0.270000, (a1, b1): 0.630000, (a2, b2): 0.080000, (a2, b1): 0.020000)

function removeElt(items, i) {
  const result = items.slice().splice(i, 1)
  if (result.length == 1) return result[0]
  return result
}

function incrDictEntry(d, k, v) {
  k = JSON.stringify(k)
  v = d.has(k) ? d.get(k) + v : v
  d.set(k, v)
}

function JDist(PA, PB) {
  const map = new Map()
  const dPA = PA.d
  dPA.forEach((v1, k1) => {
    //  暂时不知什么原因 key上多了加了双引号
    k1 = k1.replace(/\"/g, "")
    let db = PB(k1).d

    if (db) {
      db.forEach((v2, k2) => {
        k2 = k2.replace(/\"/g, "")
        console.log(v1, v2)
        map.set([k1, k2], (v1 * v2).toFixed(2))
      })
    }
  })

  return new DDist(map)
}

module.exports = {
  DDist,
  JDist
}

