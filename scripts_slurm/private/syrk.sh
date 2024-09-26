#!/bin/bash
cd samples
cd syrk
nohup srun -J syrk -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/syrk.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/syrk.err ./syrk -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 3000000 -ni=32768 -nj=16384 &