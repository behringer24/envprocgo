#!/usr/bin/env python

import fileinput
import argparse
import sys
import re
import os

VERSION = '1.0'

parser = argparse.ArgumentParser(description='Preprocess configuration files and fille with environment variables.')

parser.add_argument('infile', nargs='?', type=argparse.FileType('r'), default=sys.stdin,
                                help='the input file to preprocess. Or use pipe for stdin')
parser.add_argument('outfile', nargs='?', type=argparse.FileType('w'), default=sys.stdout,
                                help='the output file to write the parsed result to. Or use pipe for stdout')
parser.add_argument('-c', '--char', action='store', default='$',
                                help='character that is used as an variable indicator')
parser.add_argument('-f', '--force', action='store_true', 
                                help='do not stop execution if environment variable is not found')
parser.add_argument('-v', '--version', action='version', version='%(prog)s ' + VERSION)

if len(sys.argv)==1:
    parser.print_help(sys.stderr)
    sys.exit(1)

args = parser.parse_args()
row = 0

for line in args.infile:
    row = row + 1
    matches = re.findall('\\' + args.char + r'\{env\:(.+?)\}', line)
    for match in matches:
        if os.environ.has_key(match):
            line = re.sub('\\' + args.char + r'\{env\:' + match + '\}', os.environ.get(match), line)
        elif not args.force:
            print >> sys.stderr, 'ERROR: No environment variable found for \'' + match + '\' in line ' + str(row)
            sys.exit(1)

    args.outfile.write(line)
    pass

args.outfile.close()
