
roi=[ 0.10, 0.20]


for r in roi:
    total = 0
    yearly=4000000

    for i in range(1,21):
        total = total + yearly
        gain = total * r
        gain = total * r
        total = total + gain
        print "Year: %2d"%i," V: %10.2d"% total, " U: %10.2d"% (total/67.7)
