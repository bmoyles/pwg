#!/usr/bin/env python2.7

import os
import random
import sys


class PasswordGenerator(object):
    def __init__(self, length=12):
        random.seed()
        self.dictfile = '/usr/share/dict/words'
        self.min_frag_len = 3
        self.max_frag_len = 6
        self.min_specials = 1
        self.max_specials = 2
        self.min_numbers = 1
        self.max_numbers = 3
        self.specials = '!@#$%^&*()'

    def _get_words(self):
        words = []
        filesize = os.stat(self.dictfile)[6]
        with open(self.dictfile, 'r') as f:
            while len(words) != 2:
                f.seek((f.tell() + random.randint(0, filesize - 1)) % filesize)
                f.readline()
                line = f.readline().strip().lower()
                if line.isalpha() and len(line) in xrange(self.min_frag_len, self.max_frag_len + 1):
                    if (random.randint(0, 10) % 2) == 0:
                        words.append(line.capitalize())
                    else:
                        words.append(line)
        return words

    def _get_specials(self):
        specials = ""
        num_specials = random.randint(self.min_specials, self.max_specials)
        for i in xrange(0, num_specials):
            specials = ''.join((specials, self.specials[random.randint(0, len(self.specials) - 1)]))
        return specials

    def _get_numbers(self):
        numbers = ""
        num_numbers = random.randint(self.min_numbers, self.max_numbers)
        for i in xrange(0, num_numbers):
            numbers = ''.join((numbers, str(random.randint(0, 9))))
        return numbers

    def generate(self, count=1):
        passwords = []
        for i in xrange(0, count):
            specials = self._get_specials()
            numbers = self._get_numbers()
            words = self._get_words()
            pw = ''.join((words[0], numbers, specials, words[1]))
            passwords.append(pw)
        return passwords


def main():
    args = sys.argv[1:]
    if (len(args) != 1):
        number = 1
    else:
        number = int(args[0])
    pwg = PasswordGenerator()
    print("\n".join(pwg.generate(number)))


if __name__ == '__main__':
    main()

# vim:filetype=python
