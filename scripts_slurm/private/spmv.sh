#!/bin/bash
cd samples
cd spmv
nohup srun -J spmv -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/spmv.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/spmv.err ./spmv -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 100000000 -dim=8388608 -sparsity=0.00001 &