#!/bin/bash
cd samples
cd color
nohup srun -J color -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/color.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/color.err ./color -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 100000000 -numNodes=134217728 -numItems=268435456 &