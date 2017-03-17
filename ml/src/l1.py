
# File : l1.py
# First ML script
# Supervised learning
# Problem type: Classification
# Problem definition
#   - There is set of targets T - {T1, T2,........}
#   - Each data point has a futures F1 - (Future 1, Future 2,......)
#                                   F2 - (Future 1, Future 2,......)
#   - Each data point has a target  F1->T1, F2->T2.........
#   - Object is to learn for the I/P dataset
#   - Build a decision tree
#   - Predict any future value F2 - (Future 1, Future 2,......) -> T?

# import pydotplus
from sklearn import tree

# Car Models
SMART, FORD, AUDI, BUGGATI = range(4)

# Training data
#  Futures - 3 Elements
#          - [ Price, Doors, HP]
futures = [[15000, 2, 100], [35000, 4, 255], [30250, 4, 200], [300000, 2, 1000]]
labels = [ SMART, AUDI, FORD, BUGGATI]

# futures = [[100, 0], [150, 1], [160, 1], [90, 0]]
# labels = [ 0, 1, 1, 0]

# Get a Classifier
clf = tree.DecisionTreeClassifier()
clf = clf.fit(futures, labels)

print clf.predict([40000, 4, 300])
print clf.predict([31000, 4, 250])
# print clf.predict([110, 0])
