#!/bin/bash
cd samples
cd mis
nohup srun -J mis -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/mis.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/mis.err ./mis -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 100000000 -numNodes=134217728 -numItems=2097152 &