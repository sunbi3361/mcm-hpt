#!/bin/bash
./pagerank -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv \
-max-inst 100000000 \
-node=8192 -sparsity=0.5 -iterations=1 1> pagerank_8192.out &
