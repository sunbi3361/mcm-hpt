#!/bin/bash    
./spmv_new -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 100000000 -dim=8388608 -sparsity=0.00001

# -dim=67108864
# -dim=33554432
# 16777216
# 8388608