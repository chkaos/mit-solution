// Part C: Substitution Cipher

const findPermutations = (string) => {
  if (!string || typeof string !== "string") {
    return "Please enter a string"
  }

  if (!!string.length && string.length < 2) {
    return string
  }

  let permutationsArray = []

  for (let i = 0; i < string.length; i++) {
    let char = string[i]

    if (string.indexOf(char) != i)
      continue

    let remainder = string.slice(0, i) + string.slice(i + 1, string.length)

    for (let permutation of findPermutations(remainder)) {
      permutationsArray.push(char + permutation)
    }
  }
  return permutationsArray
}

class SubMessage {
  constructor(text) {
    this.text = text
    this.validWords = ["hello", "world"]   // 题目是读取文件里的单词写到这里, 这里为了简便写死几个单词
    this.letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    this.vowels = "aeiou"
  }
  getMessageText() {
    return this.text
  }
  isWord(word) {
    word = word.replace(/[^a-zA-Z0-9]/g, "").toLowerCase()

    return this.validWords.includes(word)
  }
  buildTransposeDict(permutation) {
    let dict = {}
    const vowels = this.vowels
    for (let i = 0; i < vowels.length; i++) {
      dict[vowels[i]] = permutation[i]
      dict[vowels[i].toUpperCase()] = permutation[i].toUpperCase()
    }

    return dict

  }
  applyTranspose(dict) {
    let res = ""
    for (let i = 0; i < this.text.length; i++) {
      res += dict[this.text[i]] || this.text[i]
    }
    return res
  }
}

class EncryptedSubMessage extends SubMessage {
  constructor(text) {
    super(text)
  }
  decryptMessage() {
    const res = []
    let highPoint = 0
    const permutationArray = findPermutations(this.vowels)
    console.log(permutationArray)
    for (let i = 0; i < permutationArray.length; i++) {
      let point = 0
      let dict = this.buildTransposeDict(permutationArray[i])
      let words = this.applyTranspose(dict)
      let wordArray = words.split(" ")
      wordArray.forEach(word => {
        if (this.isWord(word)) {
          point += 1
        }
      })
      if (point > highPoint) {
        highPoint = point
        res.push({ "text": words, point })
      }
    }
    return res
  }
}


const message = new SubMessage("Hello World!")
const permutation = "eaiuo"
const encDict = message.buildTransposeDict(permutation)
console.log("Original message:", message.getMessageText(), "Permutation:", permutation)
console.log("Expected encryption:", "Hallu Wurld!")
console.log("Actual encryption:", message.applyTranspose(encDict))

const encMessage = new EncryptedSubMessage(message.applyTranspose(encDict))
console.log("Decrypted message:", encMessage.decryptMessage())