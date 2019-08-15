import pandas
import sys
import re

if __name__ == "__main__":
    if len(sys.argv) < 2:
        sys.exit("Usage: ./" + sys.argv[0] + " <file name>")
    filename = sys.argv[1]
    df = pandas.read_csv(filename)
    time_regex = re.compile("((\d+)m)*((\d*.\d*)s)*((\d*\.\d*)ms)*")
    for index, row in df.iterrows():
        match = time_regex.match(row['time'])
        minutes = float(match.group(2) if match.group(2) is not None else 0)
        seconds = float(match.group(4) if match.group(4) is not None else 0)
        miliseconds = float(match.group(6) if match.group(6) is not None else 0)
        total = minutes * 60 + seconds + miliseconds / 1000.0
        df.at[index, 'time'] = total
    df.to_csv(filename + ".fix", mode='w', index=False)
