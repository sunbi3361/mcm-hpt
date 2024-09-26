#!/bin/bash
cd samples
cd pagerank
nohup srun -J pagerank -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/pagerank.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/pagerank.err ./pagerank -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 100000000 -node=16384 -sparsity=0.5 -iterations=1 &