#!/bin/bash
cd samples
cd correlation
nohup srun -J correlation -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/correlation.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/correlation.err ./correlation -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 100000000 -num_m=16384 -num_n=32768 &