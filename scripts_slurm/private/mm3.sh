#!/bin/bash
cd samples
cd mm3
nohup srun -J mm3 -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/mm3.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/mm3.err ./mm3 -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 100000000 -ni=16384 -nj=16384 -nk=8192 -nl=8192 -nm=8192 &