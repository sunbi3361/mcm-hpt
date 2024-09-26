#!/bin/bash
cd samples
cd blackscholes
nohup srun -J blackscholes -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/blackscholes.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/blackscholes.err ./blackscholes -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -width=8192 -height=8192 &