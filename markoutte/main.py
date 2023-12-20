import sys
sys.path.append('algs')
from tests import Test

tst = Test()
tst.loremVarPatternTest()
tst.dnaVarPatternTest()
tst.aliceVarPatternTest()
#tst.dialoguesVarPatternTest()
tst.loremVarTextTest()
tst.dnaVarTextTest()
tst.aliceVarTextTest()
#tst.dialoguesVarTextTest()

print("All the tests have been run")