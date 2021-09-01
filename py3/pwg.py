import os
import random
import sys
import typing as t

from pathlib import Path


class PasswordGenerator:
    def __init__(self, length: int = 12):
        random.seed()
        self.dictfile = Path("/usr/share/dict/words")
        self.min_frag_len: int = 3
        self.max_frag_len: int = 5
        self.min_specials: int = 1
        self.max_specials: int = 2
        self.min_numbers: int = 1
        self.max_numbers: int = 3
        self.specials: str = '!@#$%^&*()'

    def get_words(self):
        words: t.List[str] = []
        filesize: int = self.dictfile.stat()[6]
        with self.dictfile.open('r') as f:
            while len(words) != 2:
                f.seek((f.tell() + random.randint(0, filesize - 1)) % filesize)
                # we may be mid-line so throw it away
                f.readline()
                line = f.readline().strip().lower()
                if line.isalpha() and len(line) in range(self.min_frag_len, self.max_frag_len + 1):
                    if (random.randint(0, 10) % 2) == 0:
                        words.append(line.capitalize())
                    else:
                        words.append(line)
        return words

    def get_specials(self) -> str:
        specials: str = ""
        num_specials: int = random.randint(self.min_specials, self.max_specials)
        for i in range(num_specials):
            specials = "".join((specials, self.specials[random.randint(0, len(self.specials) - 1)]))
        return specials

    def get_numbers(self) -> str:
        numbers: str = ""
        count_numbers: int = random.randint(self.min_numbers, self.max_numbers)
        for i in range(count_numbers):
            numbers = "".join((numbers, str(random.randint(0, 9))))
        return numbers

    def generate(self, count: int = 1) -> t.List[str]:
        passwords: t.List[str] = []
        for i in range(count):
            specials = self.get_specials()
            numbers = self.get_numbers()
            words = self.get_words()
            pw: str =  "".join((words[0], numbers, specials, words[1]))
            passwords.append(pw)
        return passwords


def main():
    count: int = 1
    args = sys.argv[1:]
    if len(args) == 1:
        count = int(args[0])
    gen = PasswordGenerator()
    print("\n".join(gen.generate(count)))

if __name__ == "__main__":
    main()
