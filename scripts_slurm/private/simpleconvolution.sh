#!/bin/bash
cd samples
cd simpleconvolution
nohup srun -J simpleconvolution -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/simpleconvolution.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/simpleconvolution.err ./simpleconvolution -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 100000000 -width=16384 -height=32760 &