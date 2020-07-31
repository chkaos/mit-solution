// Problem 6
// Consider the following hierarchy of classes:
// 
// class Person(object):     
//     def __init__(self, name):         
//         self.name = name     
//     def say(self, stuff):         
//         return self.name + ' says: ' + stuff     
//     def __str__(self):         
//         return self.name  
// 
// class Lecturer(Person):     
//     def lecture(self, stuff):         
//         return 'I believe that ' + Person.say(self, stuff)  
// 
// class Professor(Lecturer): 
//     def say(self, stuff): 
//         return self.name + ' says: ' + self.lecture(stuff)
// 
// class ArrogantProfessor(Professor): 
//     def say(self, stuff): 
//         return 'It is obvious that ' + self.say(stuff)
// As written, this code leads to an infinite loop when using the Arrogant 
// Professor class.
// 
// Change the definition of ArrogantProfessor so that the following behavior is 
// achieved:
// 
// e = Person('eric') 
// le = Lecturer('eric') 
// pe = Professor('eric') 
// ae = ArrogantProfessor('eric')
// 
// >>> e.say('the sky is blue')
// eric says: the sky is blue
// 
// >>> le.lecture('the sky is blue')
// I believe that eric says: the sky is blue
// 
// >>> pe.say('the sky is blue')
// eric says: I believe that eric says: the sky is blue
// 
// >>> pe.lecture('the sky is blue')
// I believe that eric says: the sky is blue
// 
// >>> ae.say('the sky is blue')
// eric says: It is obvious that eric says: the sky is blue
// 
// >>> ae.lecture('the sky is blue')
// It is obvious that eric says: the sky is blue
// =============================================================================

class Person {
  constructor(name) {
    this.name = name
  }
  Say(stuff) {
    return this.name + ' says: ' + stuff
  }
}

class Lecturer extends Person {
  Lecture(stuff) {
    return 'I believe that ' + super.Say(stuff)
  }
}

class Professor extends Lecturer {
  Say(stuff) {
    return this.name + ' says: ' + super.Lecture(stuff)
  }
}

class ArrogantProfessor extends Lecturer {
  Say(stuff) {
    return this.name + ' says: ' + this.Lecture(stuff)
  }
  Lecture(stuff) {
    return 'It is obvious that ' + super.Say(stuff)
  }
}

const e = new Person('eric')
const le = new Lecturer('eric')
const pe = new Professor('eric')
const ae = new ArrogantProfessor('eric')

console.log(e.Say('the sky is blue'))
console.log(le.Lecture('the sky is blue'))

console.log(pe.Say('the sky is blue'))
console.log(pe.Lecture('the sky is blue'))

console.log(ae.Say('the sky is blue'))
console.log(ae.Lecture('the sky is blue'))

//  =============================================================================
//  Problem 7
//  
//  You are given the following superclass.Do not modify this.
//    
//  class Container(object):
//      """ Holds hashable objects. Objects may occur 0 or more times """
//      def __init__(self):
//          """ Creates a new container with no objects in it. I.e., any object 
//              occurs 0 times in self. """
//          self.vals = {}
//      def insert(self, e):
//          """ assumes e is hashable
//              Increases the number times e occurs in self by 1. """
//          try:
//              self.vals[e] += 1
//          except:
//              self.vals[e] = 1
//      def __str__(self):
//          s = ""
//          for i in sorted(self.vals.keys()):
//              if self.vals[i] != 0:
//                  s += str(i) + ":" + str(self.vals[i]) + "\n"
//          return s
//  Write a class that implements the specifications below.Do not override any 
//  methods of Container.
//    
//  class Bag(Container):
//      def remove(self, e):
//          """ assumes e is hashable
//              If e occurs in self, reduces the number of 
//              times it occurs in self by 1. Otherwise does nothing. """
//           write code here
//  
//      def count(self, e):
//          """ assumes e is hashable
//              Returns the number of times e occurs in self. """
//           write code here
//  For example,
//  d1 = Bag()
//  d1.insert(4)
//  d1.insert(4)
//  print(d1)
//  d1.remove(2)
//  print(d1)
//  prints
//  4: 2
//  4: 2
//  For example,
//  d1 = Bag()
//  d1.insert(4)
//  d1.insert(4)
//  d1.insert(4)
//  print(d1.count(2))
//  print(d1.count(4))
//  prints
//  0
//  3
//  
//  =============================================================================

class Container {
  constructor() {
    this.vals = {}
  }
  insert(e) {
    if (this.vals[e]) {
      this.vals[e] += 1
    } else {
      this.vals[e] = 1
    }
    console.log(e, " was inserted")
  }
  show() {
    let res = ""
    Object.keys(this.vals).forEach(key => {
      res += `${key}|${this.vals[key]} `
    })
    return res
  }
}

class Bag extends Container {
  remove(e) {
    if (this.vals[e]) {
      this.vals[e] -= 1
      if (this.vals[e] == 0) {
        delete this.vals[e]
      }
    } else {
      console.log(e, " was not found")
    }
  }
  count(e) {
    if (this.vals[e]) {
      return this.vals[e]
    } else {
      console.log(e, " was not found")
    }
  }
}

const d1 = new Bag()
d1.insert(4)
d1.insert(4)

d1.remove(2)

d1.insert(2)
console.log(d1.count(2))
console.log(d1.count(4))

console.log(d1.show())
