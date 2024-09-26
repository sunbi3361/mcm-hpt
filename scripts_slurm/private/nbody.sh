#!/bin/bash
cd samples
cd nbody
nohup srun -J nbody -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/nbody.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/nbody.err ./nbody -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 1000000000 -iter=8 -particles=33554432 &