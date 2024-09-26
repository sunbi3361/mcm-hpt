#!/bin/bash
cd samples
cd mvt
nohup srun -J mvt -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/mvt.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/mvt.err ./mvt -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 3000000 -n=32768 &