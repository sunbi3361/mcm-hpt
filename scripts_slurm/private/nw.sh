#!/bin/bash
cd samples
cd nw
nohup srun -J nw -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/nw.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/nw.err ./nw -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 100000000 -length=24576 &