#!/bin/bash
cd samples
cd gemver
nohup srun -J gemver -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/gemver.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/gemver.err ./gemver -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 100000000 -n=32768 &