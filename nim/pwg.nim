import os
import random
import strutils


const
  dictFilePath = "/usr/share/dict/words"
  specialChars = "!@#$%^&*()"
  specialCharCount = len(specialChars)
  minWordLen = 3
  maxWordLen = 6
  minSpecialChars = 1
  maxSpecialChars = 2
  minNumbers = 1
  maxNumbers = 3


proc getSpecialChars(min=minSpecialChars, max=maxSpecialChars): string =
  result = ""
  let numSpecials = rand([min, max])
  for i in 1..numSpecials:
    result.add(specialChars[rand(specialCharCount)])


proc getNumbers(min=minNumbers, max=maxNumbers): string =
  result = ""
  let numNumbers = rand([min, max])
  for i in 1..numNumbers:
    result.add($rand(9))


proc getWords(minLen=minWordLen, maxLen=maxWordLen): seq[string] =
  result = @[]
  var
    dict: File

  if not open(dict, dictFilePath):
    raise newException(IOError, "Unable to open file " & dictFilePath)
  defer: dict.close()

  let dictSize = int(dict.getFileSize())

  while len(result) != 2:
    dict.setFilePos((dict.getFilePos() + rand(dictSize - 1)) %% dictSize)
    discard dict.readLine()
    var line = dict.readLine().strip().toLowerAscii()
    if line.isAlphaAscii() and len(line) in minLen..maxLen:
      if rand(10) %% 2 == 0:
        result.add(line.capitalizeAscii())
      else:
        result.add(line)


proc makePasswords(num: int = 1): seq[string] =
  result = @[]
  for i in 1..num:
    let specials = getSpecialChars()
    let numbers = getNumbers()
    let words = getWords()
    result.add(words[0] & numbers & specials & words[1])


proc main() =
  randomize()

  let num =
    if paramCount() == 1: paramStr(1).parseInt
    else: 1

  let passwords = makePasswords(num)
  for password in passwords:
    echo password


main()
