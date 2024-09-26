#!/bin/bash
cd samples
cd fastwalshtransform
nohup srun -J fastwalshtransform -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/fastwalshtransform.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/fastwalshtransform.err ./fastwalshtransform -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 100000000 -length=966367641 &