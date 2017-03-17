# File : l2.py
# Second ML script
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

import numpy as np
from sklearn import tree
from sklearn.metrics import accuracy_score
from sklearn.datasets import load_iris

# Feature & Targets(Labels)
ds = load_iris()
print "F: ", ds.feature_names
print "T: ", ds.target_names
print "L: ", len(ds.data)

# Trim out validation data
test_list = [49, 99, 149]
t = np.delete( ds.target, test_list)
d = np.delete( ds.data, test_list, axis=0)

clf = tree.DecisionTreeClassifier()
clf = clf.fit(d, t)

out = clf.predict()

for e in test_list:
    print "F: ", ds.data[e], " E: ", ds.target[e], " A: ", clf.predict(ds.data[e])
