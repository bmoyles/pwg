#!/usr/bin/env ruby

class PasswordGenerator
  def initialize()
    srand Time.now.to_i
    @words = Hash.new
    @lengths = (4..5)
    @specials = ['!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '[', ']', '{', '}']
    read_words
  end

  def generate()
    length = @lengths.to_a[rand(@lengths.to_a.length)]
    word1 = @words[length][rand(@words[length].length)]
    word2 = @words[length][rand(@words[length].length)]
    mid = ''
    (word1.length + word2.length).to_i.upto(10) do |count|
      mid = mid + rand(9).to_s
    end
    mid = mid + @specials[rand(@specials.length)]
    if (rand(10) % 2 == 0)
      word1.capitalize!
    end
    if (rand(10) % 2 == 0)
      word2.capitalize!
    end
    word = word1 + mid + word2
  end

  private

  def read_words()
    File.open('/usr/share/dict/words') do |file|
      file.each do |word|
        word.chomp!
        word.downcase!
        if @lengths.include? word.length
          if not @words.has_key? word.length
            @words[word.length] = Array.new
          end
          @words[word.length].push(word)
        end
      end
    end
  end

end

if ARGV.length < 1
  count = 1
else
  count = ARGV[0].to_i
end
pwg = PasswordGenerator.new
1.upto(count) do
  puts pwg.generate
end
