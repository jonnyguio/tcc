import pandas
import sys
import re

if __name__ == "__main__":
    if len(sys.argv) < 2:
        sys.exit("Usage: ./" + sys.argv[0] + " <file name>")
    filename = sys.argv[1]
    df = pandas.read_csv(filename)
    time_regex = re.compile("(\d+)m(\d*.\d*)s")
    for index, row in df.iterrows():
        match = time_regex.match(row['time'])
        minutes = float(match.group(1))
        seconds = float(match.group(2))
        total = minutes * 60 + seconds
        df.at[index, 'time'] = total
    df.to_csv(filename + ".fix")
