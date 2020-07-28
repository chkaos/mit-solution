// Part B: Cipher Like Caesar

class Message {
  constructor(text) {
    this.text = text
    this.validWords = ["apple", "banana", "pie"]   // 题目是读取文件里的单词写到这里, 这里为了简便写死几个单词
    this.letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
  }
  getMessageText() {
    return this.text
  }
  getValidWords() {
    return this.validWords.slice()
  }
  buildShiftDict(shift) {
    let dict = {}
    const letters = this.letters
    for (let i = 0; i < letters.length; i++) {
      let charcode = letters.charCodeAt(i)
      if (97 <= charcode && 122 >= charcode) {
        charcode = charcode + shift > 122 ? 97 + charcode + shift - 122 : charcode + shift
      } else if (65 <= charcode && 90 >= charcode) {
        charcode = charcode + shift > 90 ? 65 + charcode + shift - 90 : charcode + shift
      }
      dict[letters[i]] = String.fromCharCode(charcode)
    }

    return dict

  }
  applyShift(shift) {
    const dict = buildShiftDict(shift)
    let res = ""
    for (let i = 0; i < this.text.length; i++) {
      res += dict[this.text[i]] || this.text[i]
    }
    return res
  }
}

class PlaintextMessage extends Message {
  constructor(text, shift) {
    super(text)
    this.shift = shift
    this.encryptionDict = {}
    this.messageTextEncrypted = ""
  }
  getShift() {
    return this.shift
  }
  changeShift(shift) {
    this.shift = shift
  }
  getEncryptionDict() {
    return Object.assign({}, this.encryptionDict)
  }
  getMessageTextEncrypted() {
    return this.messageTextEncrypted
  }
}

class CiphertextMessage extends Message {
  constructor(text) {
    super(text)
  }
  decryptMessage() {
    // 配合合法单词数据解密 最多尝试25次暴力破解
    // todo
  }
}

const amessage = new Message()
console.log(amessage.buildShiftDict(5))