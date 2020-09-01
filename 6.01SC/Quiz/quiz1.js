class Animal {
  constructor(name, age) {
    this.name = name
    this.age = age
    this.isDog = false
  }
  birthday() {
    this.age += 1
  }
  inDogYears() {
    if (this.isDog) return this.age * 7
    return this.age
  }
}

const snuffy = new Animal('Aloysius Snuffleupagus', 25)
const garfield = new Animal('Garfield', 32)
const george = new Animal('Curious George', 70)
const snoopy = new Animal('Snoopy', 60)

garfield.birthday()

console.log(garfield.age)

console.log(snoopy.inDogYears())
snoopy.birthday()
console.log(snoopy.inDogYears())

snoopy.isDog = true
console.log(snoopy.inDogYears())

snuffy.isDog = false
Animal.isDog = true
console.log(snuffy.inDogYears())

console.log(garfield.inDogYears())