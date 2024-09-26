#!/bin/bash
cd samples
cd gups
nohup srun -J gups -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/gups.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/gups.err ./gups -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 3000000  &