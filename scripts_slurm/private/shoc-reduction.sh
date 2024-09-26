#!/bin/bash
cd samples
cd shoc-reduction
nohup srun -J shoc-reduction -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/shoc-reduction.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/shoc-reduction.err ./shoc-reduction -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 1000000000 -Size=536870912 -Iterations=2 &