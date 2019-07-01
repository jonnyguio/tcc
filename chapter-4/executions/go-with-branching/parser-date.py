import pandas
import sys
import re

if __name__ == "__main__":
    if len(sys.argv) < 2:
        sys.exit("Usage: ./" + sys.argv[0] + " <file name>")
    filename = sys.argv[1]
    print("Formatting file {}".format(filename))
    df = pandas.read_csv(filename)
    time_regex = re.compile("((\d*)m)*(\d*.\d*)s")
    time_regex2 = re.compile("(\d*\.\d*)ms")
    for index, row in df.iterrows():

        match = time_regex.match(row['time'])
        if match is None:
            match = time_regex2.match(row['time'])
            if match is None:
                break
            else:
                miliseconds = float(match.group(1))
                total = miliseconds / 1000
        else:
            total = 0
            if match.group(2) is not None:
                minutes = float(match.group(2))
                total += minutes * 60
            if match.group(3) is not None:
                seconds = float(match.group(3))
                total += float(seconds)
        df.at[index, 'time'] = total
    df.to_csv(filename + ".fix")
