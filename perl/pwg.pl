#!/usr/bin/env perl

use warnings;
use strict;
$|++;

srand(time());
my $dictfile = '/usr/share/dict/words';
my %words;
my @fragment_lengths = 3..5;
my $length = 12;
my $middle_length = 3;
my @specials = ('!', '@', '#', '$', '%', '^', '&', '*', '(', ')');


sub read_words {
    my @rawwords;
    my $f;
    open($f, "<$dictfile") or die "Cannot open $dictfile: $!\n";
    @rawwords = map {
        chomp;
        lc;
    } grep {
        /^[a-zA-Z]{$fragment_lengths[0],$fragment_lengths[-1]}?$/
    } <$f>;
    close($f);
    foreach my $fragment_length (@fragment_lengths) {
        @{$words{$fragment_length}} = grep { (length($_) == $fragment_length) } @rawwords;
    }
}

sub middle {
    my $length = shift || $middle_length;
    my $symbol = $specials[int(rand($#specials+1))];
    my $digits;
    foreach my $i (0..($length - 2)) {
        $digits .= int(rand(10));
    }
    return $digits . $symbol;
}

sub generate {
    my $count = shift || 1;
    my @passwords;
    foreach my $i (0..($count - 1)) {
        my $start_length = $fragment_lengths[int(rand($#fragment_lengths))];
        my $midlen = $middle_length;
        if ($start_length == 3) {
            $midlen ++;
        }
        my $end_length = $length - ($start_length + $midlen);
        my $start = @{$words{$start_length}}[int(rand($#{$words{$start_length}}))];
        $start = ucfirst($start) if (int(rand(10) % 2) == 0);
        my $middle = middle($midlen);
        my $end = @{$words{$end_length}}[int(rand($#{$words{$end_length}}))];
        $end = ucfirst($end) if (int(rand(10) % 2) == 0);
        push(@passwords, $start . $middle . $end);
    }
    return @passwords;
}


my $count = shift || 1;
read_words();
my @passwords = generate($count);
map {print $_, "\n"} @passwords;

# vim:filetype=perl
